package server

import (
	"fmt"
	"log"

	authRoutes "github.com/LuanTruongPTIT/tutor-be/modules/auth/port/router"
	"github.com/LuanTruongPTIT/tutor-be/pkg/dbs"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	db     dbs.IDatabase
}

func NewServer(db dbs.IDatabase) *Server {
	return &Server{
		engine: gin.Default(),

		db: db,
	}
}
func (s Server) Run() error {
	fmt.Print("vao day")
	_ = s.engine.SetTrustedProxies(nil)
	if err := s.MapRoutes(); err != nil {
		log.Fatalf("MapRoutes error: %v", err)
	}
	if err := s.engine.Run(fmt.Sprintf(":%d", 8080)); err != nil {
		log.Fatalf("Running HTTP server: %v", err)

	}
	return nil
}
func (s Server) MapRoutes() error {
	v1 := s.engine.Group("/api/v1")
	authRoutes.Routes(v1, s.db)
	return nil
}
