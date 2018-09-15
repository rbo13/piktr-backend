package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

// Server serves as the server
// of the app
type Server struct {
	HTTPServer *http.Server
	Router     *chi.Mux
	address    string
}

// New returns the instance
// of our Server struct
func New(serverAddress string, router *chi.Mux) *Server {
	tlsConfig := &tls.Config{
		// Causes servers to use Go's default ciphersuite preferences,
		// which are tuned to avoid attacks. Does nothing on clients.
		PreferServerCipherSuites: true,
		// Only use curves which have assembly implementations
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519, // Go 1.8 only
		},
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	return &Server{
		HTTPServer: &http.Server{
			Addr:         serverAddress,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
			TLSConfig:    tlsConfig,
			Handler:      router,
		},
		Router:  router,
		address: serverAddress,
	}
}

// Start boots-up a server
// that runs on plain HTTP
func (srvr *Server) Start(cert, privKey string) {
	log.Printf("Server is running on http://localhost%s", srvr.address)
	err := srvr.HTTPServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Server is not starting due: %v", err)
	}
}

// StartTLS boots-up a server
// that runs on HTTPS
func (srvr *Server) StartTLS(cert, privKey string) {
	log.Printf("Server is running on https://localhost%s", srvr.address)
	err := srvr.HTTPServer.ListenAndServeTLS(cert, privKey)
	if err != nil {
		log.Fatalf("Server is not starting due: %v", err)
	}
}
