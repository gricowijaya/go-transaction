package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/gricowijaya/go-transaction/db/sqlc"
)

// create the new account request the struct is similiar with Account Struct
type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`                  // it is required using the binding type in gin
	Currency string `json:"currency" binding:"required,oneof=USD EUR"` //  it is required and only USD and EUR binding type in gin
}

// it create the account by using the context, everything is context in gin
func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	// pass the request body from the http
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err)) // return error
		return
	}

	// argument for creating into db
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0, // create the balance to be 0
	}

	// store the created account into server
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err)) // check error if there's an error in the server code
	}

	// return the status Ok with the created account
	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// get the account from the server by passing the request params into the getAccount
	account, err := server.store.GetAccount(ctx, req.ID)
	// check if there's the row for the data that has been requested
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required",min=1`
	PageSize int32 `form:"page_size" binding:"required",min=5,max=10`
}

func (server Server) listAccount(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

  arg := db.ListAccountsParams{
    Limit: req.PageSize,
    Offset: (req.PageID - 1)* req.PageSize,
  }
	// get the account from the server by passing the request params into the getAccount
	account, err := server.store.ListAccounts(ctx, arg)
	// check if there's the row for the data that has been requested
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type updateAccountRequest struct {
  ID      int64 `json:"id" binding:"required"`
  Balance int64 `json:"balance" binding:"required"`
}

// create the update account balance
func (server Server) UpdateAccountBalance(ctx *gin.Context) {
  // create the request handler
  var req updateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

  arg := db.UpdateAccountParams{
    ID:  req.ID,
    Balance: req.Balance,
  }
   
  account, err := server.store.UpdateAccount(ctx, arg)
  if err != nil { 
    ctx.JSON(http.StatusInternalServerError, errorResponse(err))
    return
  }
  ctx.JSON(http.StatusOK, account)
  return
}

type deleteAccountRequest struct{
  ID int64 `json:"id" binding:"required"`
}

func (server Server) DeleteAccount(ctx *gin.Context) {
  var req deleteAccountRequest

  if err := ctx.ShouldBindJSON(&req); err!= nil {
    ctx.JSON(http.StatusInternalServerError, errorResponse(err))
    return
  }

  account := server.store.DeleteAccount(ctx, req.ID)

  ctx.JSON(http.StatusOK, account)
  return
}
