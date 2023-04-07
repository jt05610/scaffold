package scaffold

import "context"

type Streamer interface {
	Stream(ctx context.Context) error
}
