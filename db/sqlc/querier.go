// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	AddChatRoomParticipant(ctx context.Context, arg AddChatRoomParticipantParams) (ChatRoomParticipant, error)
	AddComment(ctx context.Context, arg AddCommentParams) (Comment, error)
	AddCommentToPost(ctx context.Context, arg AddCommentToPostParams) error
	// Add a follow relationship between two users
	AddFollow(ctx context.Context, arg AddFollowParams) (Follow, error)
	AddLikeToPost(ctx context.Context, arg AddLikeToPostParams) error
	CreateChatRoom(ctx context.Context, name string) (ChatRoom, error)
	CreateMessage(ctx context.Context, arg CreateMessageParams) (Message, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	CreateReadReceipt(ctx context.Context, arg CreateReadReceiptParams) (ReadReceipt, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteChatRoom(ctx context.Context, id int32) error
	DeleteChatRoomParticipant(ctx context.Context, userID int32) error
	DeleteComment(ctx context.Context, id int32) error
	DeleteMessage(ctx context.Context, id int32) error
	DeletePost(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id int32) error
	DeleteUserPosts(ctx context.Context, userID int32) error
	GetChatRoom(ctx context.Context, id int32) (ChatRoom, error)
	GetChatRoomParticipant(ctx context.Context, userID int32) (ChatRoomParticipant, error)
	GetChatRoomParticipants(ctx context.Context, arg GetChatRoomParticipantsParams) ([]ChatRoomParticipant, error)
	GetChatRooms(ctx context.Context, arg GetChatRoomsParams) ([]ChatRoom, error)
	GetComment(ctx context.Context, id int32) (Comment, error)
	GetComments(ctx context.Context, arg GetCommentsParams) ([]Comment, error)
	GetLikes(ctx context.Context) ([]Like, error)
	GetLikesByPostId(ctx context.Context, postID int32) ([]Like, error)
	GetLikesByUserId(ctx context.Context, userID int32) ([]Like, error)
	GetMessage(ctx context.Context, id int32) (Message, error)
	GetMessages(ctx context.Context, arg GetMessagesParams) ([]Message, error)
	GetPost(ctx context.Context, id int32) (Post, error)
	GetPosts(ctx context.Context, arg GetPostsParams) ([]Post, error)
	GetPostsWithUsers(ctx context.Context, arg GetPostsWithUsersParams) ([]GetPostsWithUsersRow, error)
	GetReadReceipt(ctx context.Context, messageID int32) (ReadReceipt, error)
	GetUser(ctx context.Context, id int32) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByUsername(ctx context.Context, username sql.NullString) (User, error)
	GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error)
	LikePost(ctx context.Context, arg LikePostParams) error
	ListPostsByUserId(ctx context.Context, userID int32) ([]Post, error)
	RemoveCommentFromPost(ctx context.Context, arg RemoveCommentFromPostParams) error
	// Remove a follow relationship between two users
	RemoveFollow(ctx context.Context, arg RemoveFollowParams) (Follow, error)
	RemoveLikeFromPost(ctx context.Context, arg RemoveLikeFromPostParams) error
	UnlikePost(ctx context.Context, arg UnlikePostParams) error
	UpdateChatRoom(ctx context.Context, arg UpdateChatRoomParams) (ChatRoom, error)
	UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error)
	UpdateMessage(ctx context.Context, arg UpdateMessageParams) (Message, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error)
	UpdatePostBody(ctx context.Context, arg UpdatePostBodyParams) (Post, error)
	UpdatePostTitle(ctx context.Context, arg UpdatePostTitleParams) (Post, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateUserBio(ctx context.Context, arg UpdateUserBioParams) (User, error)
	UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (User, error)
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (User, error)
	UpdateUserProfilePicture(ctx context.Context, arg UpdateUserProfilePictureParams) (User, error)
	UpdateUserUsername(ctx context.Context, arg UpdateUserUsernameParams) (User, error)
}

var _ Querier = (*Queries)(nil)
