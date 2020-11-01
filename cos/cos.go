package cos

import (
	"context"
	"io"
)

type Service interface {
	GetPreSignedURL(ctx context.Context, options GetPreSignedURLOptions) (string, error)
	ListObjects(ctx context.Context, options ListObjectsOptions) (ListObjectsResult, error)
	GetObject(ctx context.Context, name string) (io.ReadCloser, error)
	DownloadObject(ctx context.Context, key, saveTo string) error
	PutObject(ctx context.Context, name string, src io.Reader) error
	UploadObject(ctx context.Context, name string, filepath string) error
	DeleteObject(ctx context.Context, names ...string) (map[string]string, error)
	CopyObject(ctx context.Context, name, srcName string) error
	RenameObject(ctx context.Context, name, sourceName string) error
}
