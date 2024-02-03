package repositories

import (
	"context"
	"mysql/entity"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comments)(entity.Comments, error)
	FindById(ctx context.Context, id int)(entity.Comments, error)
	FindAll(ctx context.Context)([]entity.Comments, error)
}