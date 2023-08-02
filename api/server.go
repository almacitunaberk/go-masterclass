package api

import (
	db "github.com/almacitunaberk/go_masterclass/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// Add routes to this Router
	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts/:id", server.GetAccount)
	router.GET("/accounts", server.ListAccount)

	router.POST("/transfers", server.CreateTransfer)

	router.POST("/users", server.CreateUser)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
