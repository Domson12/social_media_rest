package api

import (
	"fmt"
	"io"

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
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))


	router.POST("/users/register", server.createUser)
	router.POST("/users/login", server.login)
	authRoutes.PUT("/users/:id", server.updateUser)
	authRoutes.GET("/users/:id", server.getUser)
	authRoutes.GET("/users", server.getUsers)
	authRoutes.POST("/users/follow", server.followUser)
	authRoutes.DELETE("/users/unfollow", server.unfollowUser)
	authRoutes.DELETE("/users/:id", server.deleteUser)
	authRoutes.POST("/posts/addPost", server.createPost)
	authRoutes.GET("/posts/:id", server.getPost)
	authRoutes.GET("/posts", server.getPosts)
	authRoutes.PUT("/posts/:id", server.updatePost)
	authRoutes.DELETE("/posts/:id", server.deletePost)
	authRoutes.POST("/posts/like", server.likePost)
	authRoutes.DELETE("/posts/unlike", server.unlikePost)
	authRoutes.POST("/posts/comment", server.addComment)
	authRoutes.GET("/ws", server.webSocket)

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	if err == io.EOF {
		return gin.H{"error": "Body need to be fulfiled"}
	}
	return gin.H{"error": err.Error()}
}
