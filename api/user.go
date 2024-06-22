package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	db "github.com/Domson12/social_media_rest/db/sqlc"
	"github.com/Domson12/social_media_rest/util"
	"github.com/gin-gonic/gin"
)

// createUserRequest represents the request to create a new user
type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// createUserResponse represents the response of a successful user creation
type createUserResponse struct {
	User_id  int32  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

// createUser is a handler function that creates a new user
func (Server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
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

	token, err := Server.tokenMaker.CreateToken(user.ID, time.Minute*15)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createUserResponse{
		User_id:  user.ID,
		Username: user.Username.String,
		Email:    user.Email,
		Token:    token,
	}

	ctx.JSON(http.StatusOK, rsp)
}

// getUserRequest represents the request to get a user
type getUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

// getUserResponse represents the response of a successful user retrieval
type getUserResponse struct {
	ID             int32     `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Bio            string    `json:"bio"`
	Role           string    `json:"role"`
	ProfilePicture string    `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
	LastActivityAt time.Time `json:"last_activity_at"`
}

// NullStringToString converts a sql.NullString to a regular string
func NullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

// getUser is a handler function that retrieves a user
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
		Username:       NullStringToString(user.Username),
		Email:          user.Email,
		Bio:            NullStringToString(user.Bio),
		Role:           user.Role,
		ProfilePicture: NullStringToString(user.ProfilePicture),
		CreatedAt:      user.CreatedAt,
		LastActivityAt: user.LastActivityAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}

// getUsersRequest represents the request to get a list of users
type getUsersRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSz int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// getUsersResponse represents the response of a successful user retrieval
type getUsersResponse struct {
	Users []getUserResponse `json:"users"`
}

// getUsers is a handler function that retrieves a list of users
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
			Username:       NullStringToString(user.Username),
			Email:          user.Email,
			Bio:            NullStringToString(user.Bio),
			Role:           user.Role,
			ProfilePicture: NullStringToString(user.ProfilePicture),
			CreatedAt:      user.CreatedAt,
			LastActivityAt: user.LastActivityAt,
		}
	}

	ctx.JSON(http.StatusOK, getUsersResponse{Users: rsp})
}

// deleteUserRequest represents the request to delete a user
type deleteUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

// deleteUser is a handler function that deletes a user
func (server *Server) deleteUser(ctx *gin.Context) {
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

	err = server.store.DeleteUserTx(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

// updateUsernameRequest represents the request to update a user's username
type updateUserRequest struct {
	ID       int32  `uri:"id" binding:"required,min=1"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Bio      string `json:"bio,omitempty"`
}

func (Server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateUserTxParams{
		UserID:   req.ID,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Bio:      req.Bio,
	}

	user, err := Server.store.UpdateUserTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := getUserResponse{
		ID:             user.ID,
		Username:       NullStringToString(user.Username),
		Email:          user.Email,
		Bio:            NullStringToString(user.Bio),
		Role:           user.Role,
		ProfilePicture: NullStringToString(user.ProfilePicture),
		CreatedAt:      user.CreatedAt,
		LastActivityAt: user.LastActivityAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}

// loginRequest represents the request to login a user
type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// loginResponse represents the response of a successful user login
type loginResponse struct {
	Token string          `json:"token"`
	User  getUserResponse `json:"user"`
}

// login is a handler function that logs in a user
func (Server *Server) login(ctx *gin.Context) {
	var req loginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := Server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials or user does not exist"})
		return
	}

	err = util.CheckPasswordHash(req.Password, user.Password)
	fmt.Println(user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
		return
	}

	token, err := Server.tokenMaker.CreateToken(user.ID, time.Minute*15)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginResponse{
		Token: token,
		User: getUserResponse{
			ID:             user.ID,
			Username:       NullStringToString(user.Username),
			Email:          user.Email,
			Bio:            NullStringToString(user.Bio),
			Role:           user.Role,
			ProfilePicture: NullStringToString(user.ProfilePicture),
			CreatedAt:      user.CreatedAt,
			LastActivityAt: user.LastActivityAt,
		},
	}

	ctx.JSON(http.StatusOK, rsp)
}
