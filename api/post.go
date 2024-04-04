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

type createPostResponse struct {
	ID      int32  `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	UserID  int32  `json:"user_id"`
	Status  string `json:"status"`
	Created string `json:"created"`
}

func (Server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePostParams{
		Title:       sql.NullString{String: req.Title, Valid: true},
		Body:        sql.NullString{String: req.Body, Valid: true},
		LikesIds:    []int32{},
		CommentsIds: []int32{},
		UserID:      req.UserID,
		Status:      req.Status,
	}

	post, err := Server.store.CreatePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createPostResponse{
		ID:      post.ID,
		Title:   post.Title.String,
		Body:    post.Body.String,
		UserID:  post.UserID,
		Status:  post.Status,
		Created: post.CreatedAt.String(),
	}

	ctx.JSON(http.StatusOK, rsp)
}

type getPostRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

type getPostResponse struct {
	ID      int32  `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	UserID  int32  `json:"user_id"`
	Status  string `json:"status"`
	Created string `json:"created"`
}

func (Server *Server) getPost(ctx *gin.Context) {
	var req getPostRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	post, err := Server.store.GetPost(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := getPostResponse{
		ID:      post.ID,
		Title:   post.Title.String,
		Body:    post.Body.String,
		UserID:  post.UserID,
		Status:  post.Status,
		Created: post.CreatedAt.String(),
	}

	ctx.JSON(http.StatusOK, rsp)
}

type getPostsRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSz int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type getPostsResponse struct {
	ID      int32  `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	UserID  int32  `json:"user_id"`
	Status  string `json:"status"`
	Created string `json:"created"`
}

func (Server *Server) getPosts(ctx *gin.Context) {
	var req getPostsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetPostsParams{
		Limit:  req.PageSz,
		Offset: (req.PageID - 1) * req.PageSz,
	}

	posts, err := Server.store.GetPosts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := make([]getPostsResponse, len(posts))
	for i, post := range posts {
		rsp[i] = getPostsResponse{
			ID:      post.ID,
			Title:   post.Title.String,
			Body:    post.Body.String,
			UserID:  post.UserID,
			Status:  post.Status,
			Created: post.CreatedAt.String(),
		}
	}

	ctx.JSON(http.StatusOK, posts)
}

type updatePostRequest struct {
	ID    int32  `uri:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type updatePostResponse struct {
	ID      int32  `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	UserID  int32  `json:"user_id"`
	Status  string `json:"status"`
	Created string `json:"created"`
}

func (Server *Server) updatePost(ctx *gin.Context) {
	var req updatePostRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdatePostParams{
		ID:    req.ID,
		Title: sql.NullString{String: req.Title, Valid: true},
		Body:  sql.NullString{String: req.Body, Valid: true},
	}

	post, err := Server.store.UpdatePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := updatePostResponse{
		ID:      post.ID,
		Title:   post.Title.String,
		Body:    post.Body.String,
		UserID:  post.UserID,
		Status:  post.Status,
		Created: post.CreatedAt.String(),
	}

	ctx.JSON(http.StatusOK, rsp)
}

type deletePostRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (Server *Server) deletePost(ctx *gin.Context) {
	var req deletePostRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := Server.store.DeletePost(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
