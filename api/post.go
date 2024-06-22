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
	ID     int32   `uri:"id" binding:"required,min=1"`
	Title  *string `json:"title,omitempty"`
	Body   *string `json:"body,omitempty"`
	Status *string `json:"status,omitempty"`
}

type updatePostResponse struct {
	ID      int32  `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	UserID  int32  `json:"user_id"`
	Status  string `json:"status"`
	Created string `json:"created"`
}

func (server *Server) updatePost(ctx *gin.Context) {
	var req updatePostRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var post db.Post
	var err error

	switch {
	case req.Title != nil && req.Body != nil && req.Status != nil:
		arg := db.UpdatePostParams{
			ID:    req.ID,
			Title: sql.NullString{String: *req.Title, Valid: true},
			Body:  sql.NullString{String: *req.Body, Valid: true},
		}
		post, err = server.store.UpdatePost(ctx, arg)
	case req.Title != nil && req.Body != nil:
		arg := db.UpdatePostParams{
			ID:    req.ID,
			Title: sql.NullString{String: *req.Title, Valid: true},
			Body:  sql.NullString{String: *req.Body, Valid: true},
		}
		post, err = server.store.UpdatePost(ctx, arg)
	case req.Title != nil:
		arg := db.UpdatePostTitleParams{
			ID:    req.ID,
			Title: sql.NullString{String: *req.Title, Valid: true},
		}
		post, err = server.store.UpdatePostTitle(ctx, arg)
	case req.Body != nil:
		arg := db.UpdatePostBodyParams{
			ID:   req.ID,
			Body: sql.NullString{String: *req.Body, Valid: true},
		}
		post, err = server.store.UpdatePostBody(ctx, arg)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No update fields provided"})
		return
	}

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

type getPostsWithUsersRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSz int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type postWithUserResponse struct {
	PostID             int32  `json:"post_id"`
	PostTitle          string `json:"post_title"`
	PostBody           string `json:"post_body"`
	PostStatus         string `json:"post_status"`
	PostCreatedAt      string `json:"post_created_at"`
	UserID             int32  `json:"user_id"`
	UserUsername       string `json:"user_username"`
	UserEmail          string `json:"user_email"`
	UserBio            string `json:"user_bio"`
	UserRole           string `json:"user_role"`
	UserProfilePicture string `json:"user_profile_picture"`
	UserCreatedAt      string `json:"user_created_at"`
	UserLastActivityAt string `json:"user_last_activity_at"`
}

type getPostsWithUsersResponse struct {
	Posts []postWithUserResponse `json:"posts"`
}

func (server *Server) getPostsWithUsers(ctx *gin.Context) {
	var req getPostsWithUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetPostsWithUsersParams{
		Limit:  req.PageSz,
		Offset: (req.PageID - 1) * req.PageSz,
	}

	posts, err := server.store.GetPostsWithUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := make([]postWithUserResponse, len(posts))
	for i, post := range posts {
		rsp[i] = postWithUserResponse{
			PostID:             post.PostID,
			PostTitle:          NullStringToString(post.PostTitle),
			PostBody:           NullStringToString(post.PostBody),
			PostStatus:         post.PostStatus,
			PostCreatedAt:      post.PostCreatedAt.String(),
			UserID:             post.UserID,
			UserUsername:       post.UserUsername.String,
			UserEmail:          post.UserEmail,
			UserBio:            post.UserBio.String,
			UserRole:           post.UserRole,
			UserProfilePicture: post.UserProfilePicture.String,
			UserCreatedAt:      post.UserCreatedAt.String(),
			UserLastActivityAt: post.UserLastActivityAt.String(),
		}
	}

	ctx.JSON(http.StatusOK, getPostsWithUsersResponse{Posts: rsp})
}
