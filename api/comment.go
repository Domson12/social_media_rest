package api

import (
	"net/http"

	db "github.com/Domson12/social_media_rest/db/sqlc"
	"github.com/gin-gonic/gin"
)

type addCommentParams struct {
	PostID int32  `json:"post_id" binding:"required"`
	UserID int32  `json:"user_id" binding:"required"`
	Text   string `json:"text" binding:"required"`
}

type addCommentResponse struct {
	ID        int32  `json:"id"`
	PostID    int32  `json:"post_id"`
	UserID    int32  `json:"user_id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

func (Server *Server) addComment(ctx *gin.Context) {
	var req addCommentParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AddCommentParams{
		PostID: req.PostID,
		UserID: req.UserID,
		Text:   req.Text,
	}

	comment, err := Server.store.AddCommentToPostTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	commentValues := comment.Comment

	rsp := addCommentResponse{
		ID:        commentValues.ID,
		PostID:    commentValues.PostID,
		UserID:    commentValues.UserID,
		Text:      commentValues.Text,
		CreatedAt: commentValues.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	ctx.JSON(http.StatusOK, rsp)
}
