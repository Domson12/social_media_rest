package api

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	db "github.com/Domson12/social_media_rest/db/sqlc"
	"github.com/Domson12/social_media_rest/util"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type createUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (Server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
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
		Password: hashedPassword,
	}

	user, err := Server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createUserResponse{
		Username: user.Username.String,
		Email:    user.Email,
	}

	ctx.JSON(http.StatusOK, rsp)
}

type getUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type getUserResponse struct {
	ID             int32          `json:"id"`
	Username       sql.NullString `json:"username"`
	Email          string         `json:"email"`
	FollowingCount sql.NullInt32  `json:"following_count"`
	FollowedCount  sql.NullInt32  `json:"followed_count"`
	Bio            sql.NullString `json:"bio"`
	Role           string         `json:"role"`
	ProfilePicture sql.NullString `json:"profile_picture"`
	CreatedAt      time.Time      `json:"created_at"`
	LastActivityAt time.Time      `json:"last_activity_at"`
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

	rsp := getUserResponse{
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		FollowingCount: user.FollowingCount,
		FollowedCount:  user.FollowedCount,
		Bio:            user.Bio,
		Role:           user.Role,
		ProfilePicture: user.ProfilePicture,
		CreatedAt:      user.CreatedAt,
		LastActivityAt: user.LastActivityAt,
	}

	ctx.JSON(http.StatusOK, rsp)
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

	rsp := make([]getUserResponse, len(users))
	for i, user := range users {
		rsp[i] = getUserResponse{
			ID:             user.ID,
			Username:       user.Username,
			Email:          user.Email,
			FollowingCount: user.FollowingCount,
			FollowedCount:  user.FollowedCount,
			Bio:            user.Bio,
			Role:           user.Role,
			ProfilePicture: user.ProfilePicture,
			CreatedAt:      user.CreatedAt,
			LastActivityAt: user.LastActivityAt,
		}
	}

	ctx.JSON(http.StatusOK, rsp)
}

type deleteUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (Server *Server) deleteUser(ctx *gin.Context) {
	var req deleteUserRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = Server.store.DeleteUser(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

type updateUsernameRequest struct {
	ID       int32  `uri:"id" binding:"required,min=1"`
	Username string `json:"username" binding:"required"`
}

func (Server *Server) updateUsername(ctx *gin.Context) {
	var req updateUsernameRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateUserParams{
		ID:       req.ID,
		Username: sql.NullString{String: req.Username, Valid: true},
	}

	user, err := Server.store.UpdateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
