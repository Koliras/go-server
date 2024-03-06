package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
    listenAdr   string
    store       *gorm.DB
}

func NewServer(listenAdr string, store *gorm.DB) *Server {
    return &Server{
        listenAdr: listenAdr,
        store: store,
    }
}

func (s *Server) Start() error {
    router := newRouter(s)
    return router.Run(s.listenAdr)
}

func newRouter(server *Server)*gin.Engine {
	router := gin.Default()
    userGroup := router.Group("users")
    userGroup.GET("", server.getAllUsers)
    userGroup.GET("/:id", server.getUserById)
    userGroup.POST("", server.createUser)
    userGroup.DELETE("/:id", server.deleteUser)
    return router
}

