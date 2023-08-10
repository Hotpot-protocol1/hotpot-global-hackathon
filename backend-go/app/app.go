package app

import (
	"fmt"
	"net/http"

	"github.com/Hotpot-protocol1/hotpot-global/config"
	"github.com/Hotpot-protocol1/hotpot-global/db"
	"github.com/Hotpot-protocol1/hotpot-global/server"
	"github.com/pkg/errors"
)

func Start(cfg config.Conf, db db.DBHandler) error {
	router, err := server.Router(cfg, db)
	if err != nil {
		return err
	}

	serverHost := fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)
	log := cfg.Log.New()
	log.WithField("api", "start").
		Info(fmt.Sprintf("Listening addr =  %s, tls = %v", serverHost, cfg.HTTP.SSL))

	// Maybe add different timeout values for upload/non-upload endpoints
	httpServer := http.Server{
		Addr:           serverHost,
		Handler:        router,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
		IdleTimeout:    cfg.IdleTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	switch cfg.HTTP.SSL {
	case true:
		if err := httpServer.ListenAndServeTLS(cfg.HTTP.ServerCertPath, cfg.HTTP.ServerKeyPath); err != nil {
			return errors.Wrap(err, "failed to start https server")
		}

	default:
		if err := httpServer.ListenAndServe(); err != nil {
			return errors.Wrap(err, "failed to start http server")
		}
	}

	return nil
}
