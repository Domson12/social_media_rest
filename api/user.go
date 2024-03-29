package api

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"strconv"

	db "github.com/Domson12/social_media_rest/db/sqlc"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (Server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	// Enhanced error handling for JSON parsing
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// Log the error for debugging purposes
		log.Printf("Error parsing JSON: %v", err)

		// Check if the error is due to EOF
		if err == io.EOF {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Request body cannot be empty"})
		} else {
			// Handle other types of JSON parsing errors
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
		}
		return
	}

	arg := db.CreateUserParams{
		Username: sql.NullString{String: req.Username, Valid: true},
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	user, err := Server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type getUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (Server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := Server.store.GetUser(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type getUsersRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSz int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (Server *Server) getUsers(ctx *gin.Context) {
	var req getUsersRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetUsersParams{
		Limit:  req.PageSz,
		Offset: (req.PageID - 1) * req.PageSz,
	}

	users, err := Server.store.GetUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}
