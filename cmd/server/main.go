package main

import (
	"net/http"

	SolidE "github.com/horizon-technos/solide-core/internal/server"
)

func main() {
	cfg := SolidE.DefaultConfig()

	h := func(c *SolidE.Context) error {
		c.W.Write([]byte("Hello world"))
		return nil
	}

	serveur := SolidE.New(h)

	serveurHttp := &http.Server{
		Addr: cfg.Addr,
		Handler: serveur,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
        ReadTimeout:       cfg.ReadTimeout,
        WriteTimeout:      cfg.WriteTimeout,
        IdleTimeout:       cfg.IdleTimeout,
        MaxHeaderBytes:    cfg.MaxHeaderBytes,
	}
	
	serveurHttp.ListenAndServe()
}
