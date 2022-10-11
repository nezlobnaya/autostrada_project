package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nezlobnaya/messing_with_autostrada/internal/server"
	"github.com/nezlobnaya/messing_with_autostrada/internal/version"
)

type config struct {
	addr    string
	baseURL string
	env     string
	cookie  struct {
		secretKey string
	}
	tls struct {
		certFile string
		keyFile  string
	}
	version bool
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.StringVar(&cfg.addr, "addr", "localhost:4444", "server address to listen on")
	flag.StringVar(&cfg.baseURL, "base-url", "https://localhost:4444", "base URL for the application")
	flag.StringVar(&cfg.env, "env", "development", "operating environment: development, testing, staging or production")
	flag.StringVar(&cfg.cookie.secretKey, "cookie-secret-key", "4XGwaJIi31s6IrX0l6U5KrLBHpQrnly2", "secret key for cookie authentication/encryption")
	flag.StringVar(&cfg.tls.certFile, "tls-cert-file", "./tls/cert.pem", "tls certificate file")
	flag.StringVar(&cfg.tls.keyFile, "tls-key-file", "./tls/key.pem", "tls key file")
	flag.BoolVar(&cfg.version, "version", false, "display version and exit")

	flag.Parse()

	if cfg.version {
		fmt.Printf("version: %s\n", version.Get())
		return
	}

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)

	app := &application{
		config: cfg,
		logger: logger,
	}

	logger.Printf("starting server on %s (version %s)", cfg.addr, version.Get())

	err := server.Run(cfg.addr, app.routes(), cfg.tls.certFile, cfg.tls.keyFile)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Print("server stopped")
}
