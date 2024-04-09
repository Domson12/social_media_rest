package db

import (
	"context"
	"database/sql"
	"fmt"
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
func (store *Store) AddLikeToPostTx(ctx context.Context, postID int32, userID int32) error {
	return store.execTx(ctx, func(q *Queries) error {
		err := q.LikePost(ctx, LikePostParams{
			PostID: postID,
			UserID: userID,
		})
		if err != nil {
			return err
		}

		err = q.AddLikeToPost(ctx, AddLikeToPostParams{
			ID: postID,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

// remove like from a post and update the post's like_id array
func (store *Store) RemoveLikeFromPostTx(ctx context.Context, postID int32, userID int32) error {
	return store.execTx(ctx, func(q *Queries) error {
		err := q.UnlikePost(ctx, UnlikePostParams{
			PostID: postID,
			UserID: userID,
		})
		if err != nil {
			return err
		}

		err = q.RemoveLikeFromPost(ctx, RemoveLikeFromPostParams{
			ID: postID,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

// add comment to a post and update the post's comment_id array
func (store *Store) AddCommentToPostTx(ctx context.Context, postID int32, userID int32, text string) error {
	return store.execTx(ctx, func(q *Queries) error {
		comment, err := q.AddComment(ctx, AddCommentParams{
			PostID: postID,
			UserID: userID,
			Text:   text,
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
		return nil
	})
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

//Delete user transaction (delete user and all related data from other tables)
