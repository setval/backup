package repository

import "io"

var _ Repository = (*s3)(nil)
var _ Repository = (*local)(nil)

type Repository interface {
	Upload(string, io.Reader) error
}
