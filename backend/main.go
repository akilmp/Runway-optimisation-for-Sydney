package main

import (
	stdlog "log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/syd-perf/syd-perf/api"
	"github.com/syd-perf/syd-perf/internal/db"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	r := gin.Default()

	dsn := db.DSNFromEnv()
	database, err := db.Connect(dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("connect db")
	}
	defer database.Close()

	api.Register(r, database)

	addr := ":8080"
	if v := os.Getenv("ADDR"); v != "" {
		addr = v
	}
	if err := r.Run(addr); err != nil {
		log.Fatal().Err(err).Msg("run server")
	}
	stdlog.Print("server started")
}
