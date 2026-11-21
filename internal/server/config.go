package server

import "time"

// structure de configuration
type Config struct {
    Addr              string
    ReadHeaderTimeout time.Duration
    ReadTimeout       time.Duration
    WriteTimeout      time.Duration
    IdleTimeout       time.Duration
    MaxHeaderBytes    int
}

// configuration par defaut du serveur http
func DefaultConfig() Config {
    return Config{
        Addr:              ":8080",
        ReadHeaderTimeout: 5 * time.Second,
        ReadTimeout:       15 * time.Second,
        WriteTimeout:      30 * time.Second,
        IdleTimeout:       60 * time.Second,
        MaxHeaderBytes:    1 << 20,
    }
}