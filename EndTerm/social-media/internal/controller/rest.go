package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/olzhas-b/social-media/internal/clients/auth"
	core "github.com/olzhas-b/social-media/internal/core"
	"net/http"
)

type Server struct {
	ctx  context.Context
	core *core.Core
	auth auth.Interface
}

func NewServer(
	ctx context.Context,
	core *core.Core,
	auth auth.Interface,
) *Server {
	return &Server{
		ctx:  ctx,
		core: core,
		auth: auth,
	}
}

func (s *Server) Run(port string) error {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")

	posts := v1.Group("posts")
	{
		posts.GET("", s.AuthMiddleware(user, false), s.GetPosts)
		posts.GET("/:id", s.GetPost)
		posts.POST("", s.AuthMiddleware(user, true), s.AddPosts)
		posts.DELETE("/:id", s.AuthMiddleware(user, true), s.RemovePosts)
	}

	return http.ListenAndServe(":"+port, router)
}
