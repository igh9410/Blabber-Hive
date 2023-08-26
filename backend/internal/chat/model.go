package chat

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ChatRoom struct {
	ID        uuid.UUID `json:"id"`
	UserID1   uuid.UUID `json:"user_id_1"`
	UserID2   uuid.UUID `json:"user_id_2"`
	CreatedAt time.Time `json:"created_at"`
}

type Message struct {
	ID              uuid.UUID `json:"id"`
	ChatRoomID      uuid.UUID `json:"chat_room_id"`
	SenderID        uuid.UUID `json:"sender_id"`
	Content         string    `json:"content"`
	MediaURL        string    `json:"media_url"`
	CreatedAt       time.Time `json:"created_at"`
	ReadAt          time.Time `json:"read_at"`
	DeletedByUserID uuid.UUID `json:"deleted_by_user_id"`
}

type CreateChatRoomReq struct {
}

type CreateChatRoomRes struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type Service interface {
	CreateChatRoom(ctx context.Context, req *CreateChatRoomReq) (*CreateChatRoomRes, error)
	GetChatRoomByID(ctx context.Context, chatRoomID uuid.UUID) (*ChatRoom, error)
}

type Repository interface {
	CreateChatRoom(ctx context.Context, chatRoom *ChatRoom) (*ChatRoom, error)
	FindChatRoomByID(ctx context.Context, chatRoomID uuid.UUID) (*ChatRoom, error)
	FetchRecentMessages(ctx context.Context, chatRoomID uuid.UUID, limit int) ([]Message, error)

	SaveMessage(ctx context.Context, message *Message) error
}