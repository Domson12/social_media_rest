package api

import (
	"fmt"

	db "github.com/Domson12/social_media_rest/db/sqlc"
	"github.com/Domson12/social_media_rest/token"
	"github.com/Domson12/social_media_rest/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %v", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

	router.POST("/users/register", server.createUser)
	router.POST("/users/login", server.login)
	router.PUT("/users/:id", server.updateUser)
	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.getUsers)
	router.POST("/users/follow", server.followUser)
	router.POST("/users/unfollow", server.unfollowUser)
	router.DELETE("/users/:id", server.deleteUser)
	router.POST("/posts/addPost", server.createPost)
	router.GET("/posts/:id", server.getPost)
	router.GET("/posts", server.getPosts)
	router.PUT("/posts/:id", server.updatePost)
	router.DELETE("/posts/:id", server.deletePost)

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
