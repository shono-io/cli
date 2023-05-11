package shono

import (
	go_shono "github.com/shono-io/go-shono"
	"github.com/sirupsen/logrus"
	"github.com/twmb/franz-go/pkg/sr"
)

func NewSchemaRegistryResource(id string, opts ...sr.Opt) go_shono.Resource[any] {
	src, err := sr.NewClient(opts...)
	if err != nil {
		logrus.Panicf("failed to create schema registry client: %v", err)
	}

	return go_shono.Resource[any]{
		Id: id,
		ClientFactory: func() any {
			return src
		},
	}
}
