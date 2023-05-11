package shono

import (
	go_shono "github.com/shono-io/go-shono"
	"github.com/sirupsen/logrus"
	"github.com/twmb/franz-go/pkg/kgo"
)

func NewKafkaResource(id string, opts ...kgo.Opt) go_shono.Resource[any] {
	kc, err := kgo.NewClient(opts...)
	if err != nil {
		logrus.Panicf("failed to create kafka client: %v", err)
	}

	return go_shono.Resource[any]{
		Id: id,
		ClientFactory: func() any {
			return kc
		},
	}
}
