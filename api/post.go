package api

import (
	"database/sql"
	"net/http"

	db "github.com/Domson12/social_media_rest/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createPostRequest struct {
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body" binding:"required"`
	UserID int32  `json:"user_id" binding:"required"`
	Status string `json:"status" binding:"required"`
}

func (Server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePostParams{
		Title:  sql.NullString{String: req.Title, Valid: true},
		Body:   sql.NullString{String: req.Body, Valid: true},
		UserID: req.UserID,
		Status: req.Status,
	}

	post, err := Server.store.CreatePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, post)
}
