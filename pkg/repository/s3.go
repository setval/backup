package repository

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type s3 struct {
	client *minio.Client
}

func NewS3(endpoint, access, secret string, ssl bool) (*s3, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(access, secret, ""),
		Secure: ssl,
	})
	if err != nil {
		return nil, err
	}
	return &s3{
		client: client,
	}, err
}

func (s *s3) Upload(path string, reader io.Reader) error {
	_, err := s.client.PutObject(context.Background(), "backup", path, reader, -1, minio.PutObjectOptions{})
	return err
}
