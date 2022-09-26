package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	db "mohamedElsonny/simple-bank/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

// Start running server on specified address (just to make the router private in this package)
func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	// Routes

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccountByID)
	router.GET("/accounts", server.listAccounts)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

func errorResponse(err error) *gin.H {
	return &gin.H{"errors": err.Error()}
}
