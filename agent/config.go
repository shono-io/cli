package agent

import (
	"github.com/benthosdev/benthos/v4/public/service"
)

type Config struct {
	Shono ShonoConfig `yaml:"shono"`
}

type ShonoConfig struct {
	ClientId     string `yaml:"clientId"`
	ClientSecret string `yaml:"clientSecret"`
	Endpoint     string `yaml:"endpoint,omitempty"`
}

type EventLogConfig struct {
	Kafka *KafkaConfig `yaml:"kafka,omitempty"`
}

type KafkaConfig struct {
}

func AgentConfig() *service.ConfigSpec {
	return service.NewConfigSpec().
		Beta().
		Categories("Agent").
		Version("0.1.0").
		Summary("Shono Agent configuration").
		Description("Shono Agent configuration").
		Field(shonoField()).
		Field(eventLogField()).
		Field(executorsField())
}

func shonoField() *service.ConfigField {
	return service.NewObjectField("shono",
		service.NewStringField("clientId").Description("The client id for this agent"),
		service.NewStringField("clientSecret").Description("The client secret for this agent"),
		service.NewStringField("endpoint").Description("The endpoint for the shono platform").Default("kafka.shono.io:9092"),
	).Description("The configuration to connect this agent to the shono platform")
}

func eventLogField() *service.ConfigField {
	return service.NewObjectField("eventLog",
		kafkaField())
}

func kafkaField() *service.ConfigField {
	return service.NewObjectField("kafka",
		service.NewStringListField("seed_brokers").
			Description("A list of broker addresses to connect to in order to establish connections. If an item of the list contains commas it will be expanded into multiple addresses.").
			Example([]string{"localhost:9092"}).
			Example([]string{"foo:9092", "bar:9092"}).
			Example([]string{"foo:9092,bar:9092"}),
		service.NewTLSToggledField("tls"),
		saslField(),
	).Optional()
}

func executorsField() *service.ConfigField {
	return service.NewObjectField("executor", dockerField())
}

func dockerField() *service.ConfigField {
	return service.NewObjectField("docker").Optional()
}
