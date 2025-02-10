package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jpoz/esbuild-example/pkg/assets"
)

// Server defines a basic HTTP server with a configurable address and internal router.
type Server struct {
	Addr string
	mux  *http.ServeMux
}

// New creates a new Server instance, configures routes using the built-in ServeMux,
// and leverages the enhanced routing semantics (e.g. proper prefix matching).
func New(addr string) *Server {
	mux := http.NewServeMux()
	s := &Server{
		Addr: addr,
		mux:  mux,
	}

	mux.HandleFunc("/", s.homeHandler)
	mux.HandleFunc("/api/quote", s.quoteHandler)
	mux.HandleFunc("/src/", assets.SrcHandler("/src/"))

	mux.HandleFunc("/health", s.healthHandler)

	return s
}

// homeHandler handles the "/" route.
// We check for an exact match to avoid catching other routes.
func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data, err := assets.Public.ReadFile("public/index.html")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

// healthHandler handles the "/health" route.
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/health" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintln(w, "OK")
}

// Listen starts the HTTP server on the configured address.
func (s *Server) Listen() {
	log.Printf("Server is starting on %s\n", s.Addr)
	if err := http.ListenAndServe(s.Addr, s.mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
