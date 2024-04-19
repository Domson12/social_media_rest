package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Domson12/social_media_rest/util"
)

type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type CreateChatRoomTxParams struct {
	ChatRoomName string `json:"name"`
	UserID       int32  `json:"user_id"`
}

type CreateChatRoomTxResult struct {
	ChatRoom ChatRoom `json:"chat_room"`
}

type CreateMessageTxResult struct {
	Message Message `json:"message"`
}

type DeleteUserTxParams struct {
	UserID int32 `json:"user_id"`
}

type DeleteUserTxResult struct {
	DeletedUser User `json:"deleted_user"`
}

type UpdateUserTxParams struct {
	UserID   int32  `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Bio      string `json:"bio,omitempty"`
}

type UpdateUserTxResult struct {
	UserID   int32  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

type AddCommentToPostTxResult struct {
	Comment Comment `json:"comment"`
}

type FollowUserTxResult struct {
	FollowingUserID int32 `json:"following_user_id" binding:"required"`
	FollowedUserID  int32 `json:"followed_user_id" binding:"required"`
}

// create a chat room with a participant
func (store *Store) CreateChatRoomTx(ctx context.Context, arg CreateChatRoomTxParams) (CreateChatRoomTxResult, error) {

	var result CreateChatRoomTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		chatRoom, errTx := q.CreateChatRoom(ctx, arg.ChatRoomName)
		if errTx != nil {
			return errTx
		}

		_, errTx = q.AddChatRoomParticipant(ctx, AddChatRoomParticipantParams{
			ChatRoomID: chatRoom.ID,
			UserID:     arg.UserID,
		})
		if errTx != nil {
			return errTx
		}

		result.ChatRoom = chatRoom

		return nil
	})

	return result, err
}

// add like to a post and update the post's like_id array
func (store *Store) AddLikeToPostTx(ctx context.Context, arg LikePostParams) error {
	return store.execTx(ctx, func(q *Queries) error {
		err := q.LikePost(ctx, LikePostParams{
			PostID: arg.PostID,
			UserID: arg.UserID,
		})
		if err != nil {
			return err
		}

		err = q.AddLikeToPost(ctx, AddLikeToPostParams{
			ID: arg.PostID,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

// remove like from a post and update the post's like_id array
func (store *Store) RemoveLikeFromPostTx(ctx context.Context, arg UnlikePostParams) error {
	return store.execTx(ctx, func(q *Queries) error {
		err := q.UnlikePost(ctx, UnlikePostParams{
			PostID: arg.PostID,
			UserID: arg.UserID,
		})
		if err != nil {
			return err
		}

		err = q.RemoveLikeFromPost(ctx, RemoveLikeFromPostParams{
			ID: arg.PostID,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

// add comment to a post and update the post's comment_id array
func (store *Store) AddCommentToPostTx(ctx context.Context, arg AddCommentParams) (AddCommentToPostTxResult, error) {
	var result AddCommentToPostTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		comment, err := q.AddComment(ctx, AddCommentParams{
			PostID: arg.PostID,
			UserID: arg.UserID,
			Text:   arg.Text,
		})
		if err != nil {
			return err
		}

		err = q.AddCommentToPost(ctx, AddCommentToPostParams{
			ID: comment.ID,
		})
		if err != nil {
			return err
		}
		result.Comment = comment
		return nil
	})
	return result, err
}

// delete comment from a post and update the post's comment_id array
func (store *Store) DeleteCommentFromPostTx(ctx context.Context, commentID int32, postID int32) error {
	return store.execTx(ctx, func(q *Queries) error {
		err := q.DeleteComment(ctx, commentID)
		if err != nil {
			return err
		}

		err = q.RemoveCommentFromPost(ctx, RemoveCommentFromPostParams{
			ID: postID,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

// Delete user transaction (delete user and all related data from other tables)
func (store *Store) DeleteUserTx(ctx context.Context, userID int32) error {
	var result DeleteUserTxResult

	return store.execTx(ctx, func(q *Queries) error {
		user, errTx := q.GetUser(ctx, userID)
		if errTx != nil {
			return errTx
		}

		errTx = q.DeleteUserPosts(ctx, userID)
		if errTx != nil {
			return errTx
		}

		errTx = q.DeleteUser(ctx, userID)
		if errTx != nil {
			return errTx
		}

		result.DeletedUser = user

		return nil
	})
}

// Update user transaction (update user and all related data from other tables, handle cases when some fields are not provided)
func (store *Store) UpdateUserTx(ctx context.Context, arg UpdateUserTxParams) (User, error) {
	var result User

	return result, store.execTx(ctx, func(q *Queries) error {
		var err error

		if arg.Username != "" {
			result, err = q.UpdateUserUsername(ctx, UpdateUserUsernameParams{
				ID:       arg.UserID,
				Username: sql.NullString{String: arg.Username, Valid: true},
			})
			if err != nil {
				return err
			}
		}

		if arg.Email != "" {
			result, err = q.UpdateUserEmail(ctx, UpdateUserEmailParams{
				ID:    arg.UserID,
				Email: arg.Email,
			})
			if err != nil {
				return err
			}
		}

		if arg.Password != "" {
			hashedPassword, err1 := util.HashPassword(arg.Password)
			if err1 != nil {
				return err1
			}
			result, err = q.UpdateUserPassword(ctx, UpdateUserPasswordParams{
				ID:       arg.UserID,
				Password: hashedPassword,
			})
			if err != nil {
				return err
			}
		}

		if arg.Bio != "" {
			result, err = q.UpdateUserBio(ctx, UpdateUserBioParams{
				ID:  arg.UserID,
				Bio: sql.NullString{String: arg.Bio, Valid: true},
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// Follow user transaction (if following id is the same as followed user id return error)
func (store *Store) FollowUserTx(ctx context.Context, arg AddFollowParams) (FollowUserTxResult, error) {
	var result FollowUserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		user, errTx := q.GetUser(ctx, arg.FollowedUserID)
		if errTx != nil {
			return fmt.Errorf("user doesn't exist")
		}
		if user.ID == arg.FollowingUserID {
			return fmt.Errorf("can't follow yourself")
		}

		follow, errTx := q.AddFollow(ctx, arg)
		if errTx != nil {
			return errTx
		}
		result.FollowedUserID = follow.FollowedUserID
		result.FollowingUserID = follow.FollowingUserID

		return nil
	})

	return result, err
}
