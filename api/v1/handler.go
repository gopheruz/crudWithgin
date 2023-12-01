package v1

import (
	"ginApi/storage"
)

type handlerV1 struct {
	Storage storage.StorageI
}

type HandlerV1Options struct {
	Storage storage.StorageI
}

func New(options *HandlerV1Options) *handlerV1 {

	return &handlerV1{
		Storage: options.Storage,
	}
}
