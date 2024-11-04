package config

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"sync"
)

type Config struct {
	once sync.Once
}

func initLogger() {
	once.Do(func() {
		if Env(EnvProd) {
			log.WithLevel(zerolog.WarnLevel)
		} else if Env(EnvTrace) {
			log.WithLevel(zerolog.TraceLevel)
		} else if Env(EnvDebug) {
			log.WithLevel(zerolog.DebugLevel)
		} else {
			log.WithLevel(zerolog.InfoLevel)
		}
	})

}
