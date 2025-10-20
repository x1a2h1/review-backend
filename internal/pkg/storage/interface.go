package storage

import "context"

type Uploader interface {
	Upload(ctx context.Context, input *Input) (*Output, error)
}