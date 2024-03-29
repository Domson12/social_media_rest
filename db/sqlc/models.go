// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
	"time"
)

type ChatRoom struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type ChatRoomParticipant struct {
	ChatRoomID int32 `json:"chat_room_id"`
	UserID     int32 `json:"user_id"`
}

type Comment struct {
	ID        int32     `json:"id"`
	PostID    int32     `json:"post_id"`
	UserID    int32     `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type Follow struct {
	FollowingUserID int32     `json:"following_user_id"`
	FollowedUserID  int32     `json:"followed_user_id"`
	CreatedAt       time.Time `json:"created_at"`
}

type Like struct {
	ID        int32     `json:"id"`
	PostID    int32     `json:"post_id"`
	UserID    int32     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Message struct {
	ID           int32          `json:"id"`
	ChatRoomID   int32          `json:"chat_room_id"`
	SenderUserID int32          `json:"sender_user_id"`
	Text         sql.NullString `json:"text"`
	CreatedAt    time.Time      `json:"created_at"`
	Status       string         `json:"status"`
}

type Post struct {
	ID    int32          `json:"id"`
	Title sql.NullString `json:"title"`
	// Content of the post
	Body          sql.NullString `json:"body"`
	LikesCount    int32          `json:"likes_count"`
	CommentsCount int32          `json:"comments_count"`
	UserID        int32          `json:"user_id"`
	Status        string         `json:"status"`
	CreatedAt     time.Time      `json:"created_at"`
}

type ReadReceipt struct {
	MessageID int32        `json:"message_id"`
	UserID    int32        `json:"user_id"`
	ReadAt    sql.NullTime `json:"read_at"`
}

type User struct {
	ID             int32          `json:"id"`
	Username       sql.NullString `json:"username"`
	Email          string         `json:"email"`
	Password       string         `json:"password"`
	FollowingCount sql.NullInt32  `json:"following_count"`
	FollowedCount  sql.NullInt32  `json:"followed_count"`
	Bio            sql.NullString `json:"bio"`
	Role           string         `json:"role"`
	ProfilePicture sql.NullString `json:"profile_picture"`
	CreatedAt      time.Time      `json:"created_at"`
	LastActivityAt time.Time      `json:"last_activity_at"`
}
