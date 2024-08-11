// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Block struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	BlockedUserID uuid.UUID `json:"blocked_user_id"`
}

type ChatRoom struct {
	ID        uuid.UUID        `json:"id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	Name      pgtype.Text      `json:"name"`
}

type Friend struct {
	UserID1 pgtype.UUID `json:"user_id_1"`
	UserID2 pgtype.UUID `json:"user_id_2"`
	ID      int32       `json:"id"`
}

type FriendRequest struct {
	SenderID    uuid.UUID        `json:"sender_id"`
	RecipientID uuid.UUID        `json:"recipient_id"`
	Status      string           `json:"status"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	ID          int32            `json:"id"`
}

type GooseDbVersion struct {
	ID        int32            `json:"id"`
	VersionID int64            `json:"version_id"`
	IsApplied bool             `json:"is_applied"`
	Tstamp    pgtype.Timestamp `json:"tstamp"`
}

type Message struct {
	ChatRoomID      pgtype.UUID      `json:"chat_room_id"`
	SenderID        pgtype.UUID      `json:"sender_id"`
	Content         pgtype.Text      `json:"content"`
	MediaUrl        pgtype.Text      `json:"media_url"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
	DeletedByUserID pgtype.UUID      `json:"deleted_by_user_id"`
	ReadAt          pgtype.Timestamp `json:"read_at"`
	ID              int32            `json:"id"`
}

type User struct {
	ID              uuid.UUID        `json:"id"`
	Username        string           `json:"username"`
	Email           string           `json:"email"`
	ProfileImageUrl pgtype.Text      `json:"profile_image_url"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
}

type UsersInChatRoom struct {
	ID         int32       `json:"id"`
	UserID     pgtype.UUID `json:"user_id"`
	ChatRoomID pgtype.UUID `json:"chat_room_id"`
}
