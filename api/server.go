package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/gricowijaya/go-transaction/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

// create the new server instance and setup the api route
// change the *db.Store into db.Store because it is an interface
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// register the custom validation from the validator engine
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)       // get by id
	router.GET("/accounts", server.listAccount)          // get by params
	router.PUT("/accounts", server.UpdateAccountBalance) // get by params
	router.DELETE("/accounts", server.DeleteAccount)     // get by params
	router.POST("/transfer", server.createTransfer)
	server.router = router

	return server
}

// function for starting and running the server
func (server *Server) Start(address string) error {
	// router field is private
	return server.router.Run(address)
}

// return the error response with gin
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
