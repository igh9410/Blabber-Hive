// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	// File: internal/pkg/sqlc/queries.sql
	// CreateChatRoom.sql
	CreateChatRoom(ctx context.Context, arg CreateChatRoomParams) error
	// FindChatRoomByID.sql
	FindChatRoomByID(ctx context.Context, id pgtype.UUID) (FindChatRoomByIDRow, error)
	// FindChatRoomInfoByID.sql
	FindChatRoomInfoByID(ctx context.Context, id pgtype.UUID) ([]FindChatRoomInfoByIDRow, error)
	// FindChatRoomList.sql
	FindChatRoomList(ctx context.Context) ([]FindChatRoomListRow, error)
	// GetFirstPageMessages.sql
	GetFirstPageMessages(ctx context.Context, arg GetFirstPageMessagesParams) ([]GetFirstPageMessagesRow, error)
	// GetPaginatedMessages.sql
	GetPaginatedMessages(ctx context.Context, arg GetPaginatedMessagesParams) ([]GetPaginatedMessagesRow, error)
	// JoinChatRoomByID.sql
	JoinChatRoomByID(ctx context.Context, arg JoinChatRoomByIDParams) error
}

var _ Querier = (*Queries)(nil)
