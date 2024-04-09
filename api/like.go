package api

import (
	"net/http"

	db "github.com/Domson12/social_media_rest/db/sqlc"
	"github.com/gin-gonic/gin"
)

type likePostRequest struct {
	UserID int32 `json:"user_id" binding:"required"`
	PostID int32 `json:"post_id" binding:"required"`
}

func (Server *Server) likePost(ctx *gin.Context) {
	var req likePostRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.LikePostParams{
		UserID: req.UserID,
		PostID: req.PostID,
	}

	err := Server.store.LikePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Post liked"})
}

type unlikePostRequest struct {
	UserID int32 `json:"user_id" binding:"required"`
	PostID int32 `json:"post_id" binding:"required"`
}

func (Server *Server) unlikePost(ctx *gin.Context) {
	var req unlikePostRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UnlikePostParams{
		UserID: req.UserID,
		PostID: req.PostID,
	}

	err := Server.store.UnlikePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Post unliked"})
}
