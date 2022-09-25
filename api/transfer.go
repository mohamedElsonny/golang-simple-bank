package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	db "mohamedElsonny/simple-bank/db/sqlc"
)

type createTransferRequestBody struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,oneof=USD EUR CAD"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	// parse and validate request body
	var body createTransferRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	isValidCurrency := server.validateAccount(ctx, body.FromAccountID, body.Currency)

	if !isValidCurrency {
		return
	}

	isValidCurrency = server.validateAccount(ctx, body.ToAccountID, body.Currency)

	if !isValidCurrency {
		return
	}
	// map transfer request to required arguments for transfer function
	arg := db.TransferTxParams{
		FromAccountID: body.FromAccountID,
		ToAccountID:   body.ToAccountID,
		Amount:        body.Amount,
	}

	// execute transfer request and check errors
	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// response with 201 status code with the transfer data
	ctx.JSON(http.StatusCreated, result)
}

func (server *Server) validateAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	return true
}
