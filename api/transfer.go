package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/gricowijaya/go-transaction/db/sqlc"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required",min=1`
	ToAccountID   int64  `json:"to_account_id" binding:"required",min=1`
	Amount        int64  `json:"amount" binding:"required",gt=0`                // greater than
	Currency      string `json:"currency" binding:"required,currency""` // the currency of both 2 account should create have the same data type.
}

// this transfer function will validate the currency FROM and TO account
// it is using the implementation of Store interface on Server struct
func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err)) // return error in bad request
		return
	}

	// check the From Account
	if !server.validAccount(ctx, req.FromAccountID, req.Currency) {
		return
	}

	// check the To Account
	if !server.validAccount(ctx, req.ToAccountID, req.Currency) {
		return
	}

	// put the request body into the argument create transfer functions
	arg := db.TransferTXParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	// store the data and check the nil
	result, err := server.store.TransferTX(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountId int64, currency string) bool {
	account, err := server.store.GetAccount(ctx, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return false
	}
	if account.Currency != currency {
		err := fmt.Errorf("this account [%d] is not using the %s meanwhile the desired currency is %s", accountId, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}
	return true
}
