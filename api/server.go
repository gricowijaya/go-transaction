package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/gricowijaya/go-transaction/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// create the new server instance and setup the api route
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	// add route
	// method for the create account, still no middleware
	router.POST("/accounts", server.createAccount)
  router.GET("/accounts/:id", server.getAccount) // get by id
  router.GET("/accounts", server.listAccount) // get by params
  router.PUT("/accounts", server.UpdateAccountBalance) // get by params
  router.DELETE("/accounts", server.DeleteAccount) // get by params
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
