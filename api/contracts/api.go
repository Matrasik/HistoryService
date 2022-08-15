package contracts

import (
	"context"
)

type Api interface {
	GetPair(ctx context.Context, curr string)
}
