package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"{{Package}}/pkg/config"
	"{{Package}}/pkg/logger"
	"{{Package}}/pkg/service/repo"
	"{{Package}}/pkg/service/repo/pg"
)

// Handler for app
type Handler struct {
	log   logger.Log
	cfg   config.Config
	repo  repo.Repo
	store repo.Store
}

// New will return an instance of Auth struct
func New(cfg config.Config, l logger.Log, s repo.Store) (*Handler, error) {

	r := pg.NewRepo()

	return &Handler{
		log:   l,
		cfg:   cfg,
		store: s,
		repo:  r,
	}, nil
}

// Ping handler
// Return "Pong"
func (h *Handler) Ping(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte("pong"))
}
