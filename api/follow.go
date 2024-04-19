package api

import (
	"net/http"

	db "github.com/Domson12/social_media_rest/db/sqlc"
	"github.com/gin-gonic/gin"
)

type followUserRequest struct {
	FollowingUserID int32 `json:"following_user_id" binding:"required"`
	FollowedUserID  int32 `json:"followed_user_id" binding:"required"`
}

type followUserResponse struct {
	FollowingUserID int32 `json:"following_user_id"`
	FollowedUserID  int32 `json:"followed_user_id"`
}

func (Server *Server) followUser(ctx *gin.Context) {
	var req followUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AddFollowParams{
		FollowingUserID: req.FollowingUserID,
		FollowedUserID:  req.FollowedUserID,
	}

	follow, err := Server.store.FollowUserTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := followUserResponse{
		FollowingUserID: follow.FollowingUserID,
		FollowedUserID:  follow.FollowedUserID,
	}

	ctx.JSON(http.StatusOK, rsp)
}

type unfollowUserRequest struct {
	FollowingUserID int32 `json:"following_user_id" binding:"required"`
	FollowedUserID  int32 `json:"followed_user_id" binding:"required"`
}

type unfollowUserResponse struct {
	FollowingUserID int32 `json:"following_user_id"`
	FollowedUserID  int32 `json:"followed_user_id"`
}

func (Server *Server) unfollowUser(ctx *gin.Context) {
	var req unfollowUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.RemoveFollowParams{
		FollowingUserID: req.FollowingUserID,
		FollowedUserID:  req.FollowedUserID,
	}

	follow, err := Server.store.RemoveFollow(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := unfollowUserResponse{
		FollowingUserID: follow.FollowingUserID,
		FollowedUserID:  follow.FollowedUserID,
	}

	ctx.JSON(http.StatusOK, rsp)
}
