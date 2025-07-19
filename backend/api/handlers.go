package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syd-perf/syd-perf/internal/ingest"
	"github.com/syd-perf/syd-perf/internal/perf"
)

func Register(r *gin.Engine, db *sql.DB) {
	r.GET("/v1/data/latest", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"metar": "YSSY"})
	})

	r.POST("/v1/perf/calc", func(c *gin.Context) {
		var cfg perf.Config
		if err := c.ShouldBindJSON(&cfg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := perf.CalcTakeoff(cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})

	r.GET("/v1/notams", func(c *gin.Context) {
		notes, _ := ingest.FetchNOTAM("https://example.com/rss")
		c.JSON(http.StatusOK, notes)
	})
}
