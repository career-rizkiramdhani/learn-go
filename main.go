package main

import (
	"crud/config"
	"crud/migration"
	"crud/router"
	"crud/seeder"
	"crud/tlsutil"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitConfig()
	config.InitDB()
	migration.RunMigrations()
	seeder.Seed()

	e := echo.New()
	router.InitRoutes(e)
	// If TLS is enabled in config, start with TLS; otherwise start plain HTTP.
	if config.Cfg.UseTLS == "true" {
		// Ensure certificate files exist (generate self-signed if missing)
		if err := tlsutil.EnsureCert(config.Cfg.CertFile, config.Cfg.KeyFile); err != nil {
			e.Logger.Fatalf("failed to ensure TLS cert: %v", err)
		}
		e.Logger.Infof("starting server with TLS on :%s", config.Cfg.Port)
		e.Logger.Fatal(e.StartTLS(":"+config.Cfg.Port, config.Cfg.CertFile, config.Cfg.KeyFile))
	} else {
		e.Logger.Infof("starting server on :%s", config.Cfg.Port)
		e.Logger.Fatal(e.Start(":" + config.Cfg.Port))
	}
}
