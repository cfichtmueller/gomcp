package gomcp

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
	"sync"
)

type HttpTransport struct {
	server             *Server
	corsAllowedOrigins string
	handleMutex        sync.Mutex
}

func NewHttpTransport(server *Server) *HttpTransport {
	return &HttpTransport{
		server:             server,
		corsAllowedOrigins: "*",
	}
}

func (t *HttpTransport) Handle(w http.ResponseWriter, r *http.Request) {
	t.handleMutex.Lock()
	defer t.handleMutex.Unlock()

	if r.Method == http.MethodOptions {
		t.addStandardHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		t.addStandardHeaders(w)
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	// Check if client accepts JSON
	acceptsJson := false
	for _, val := range r.Header.Values("Accept") {
		for _, accepted := range strings.Split(val, ",") {
			if accepted == "application/json" || accepted == "*/*" {
				acceptsJson = true
				break
			}
		}
	}
	if !acceptsJson {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	message, err := ReadJsonRpcRequest(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	res := t.server.handle(ctx, message)

	var body []byte

	if res.SendBody {
		bb, err := json.Marshal(res.Body)
		if err != nil {
			slog.Error("Failed to marshal JSON-RPC response", "error", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		} else {
			body = bb
		}
	}

	t.addStandardHeaders(w)
	w.WriteHeader(res.Status)
	if len(body) > 0 {
		w.Write(body)
	}
}

func (t *HttpTransport) addStandardHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Mcp-Protocol-Version", "2025-06-18")
	w.Header().Set("Access-Control-Allow-Origin", t.corsAllowedOrigins)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Mcp-Protocol-Version")
}
