package agent

import (
	"crypto/tls"
	"github.com/benthosdev/benthos/v4/public/service"
	"github.com/shono-io/go-shono/backbone"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/plain"
	"net"
	"time"
)

func NewAgentFromConfig(conf *service.ParsedConfig) (*Agent, error) {
	id, sopts, err := parseShonoConfig(conf.Namespace("shono"))
	if err != nil {
		return nil, err
	}

	kopts, err := parseKafkaOpts(conf.Namespace("eventLog", "kafka"))
	if err != nil {
		return nil, err
	}

	return NewAgent(
		WithAgentId(id),
		WithShonoOpts(sopts...),
		WithKafkaBackboneOpts(kopts...),
	)
}

func parseShonoConfig(conf *service.ParsedConfig) (string, []kgo.Opt, error) {
	var opts []kgo.Opt

	tlsDialer := &tls.Dialer{NetDialer: &net.Dialer{Timeout: 10 * time.Second}}
	opts = append(opts, kgo.Dialer(tlsDialer.DialContext))

	broker, err := conf.FieldString("endpoint")
	if err != nil {
		return "", nil, err
	}
	opts = append(opts, kgo.SeedBrokers(broker))

	clientID, err := conf.FieldString("clientId")
	if err != nil {
		return "", nil, err
	}

	clientSecret, err := conf.FieldString("clientSecret")
	if err != nil {
		return "", nil, err
	}
	opts = append(opts, kgo.SASL(plain.Auth{User: clientID, Pass: clientSecret}.AsMechanism()))

	return clientID, opts, nil
}

func parseKafkaOpts(conf *service.ParsedConfig) ([]kgo.Opt, error) {
	var opts []kgo.Opt

	brokerList, err := conf.FieldStringList("seed_brokers")
	if err != nil {
		return nil, err
	}
	opts = append(opts, kgo.SeedBrokers(brokerList...))

	tlsConf, tlsEnabled, err := conf.FieldTLSToggled("tls")
	if err != nil {
		return nil, err
	}
	if tlsEnabled {
		opts = append(opts, kgo.DialTLSConfig(tlsConf))
	}

	saslConfs, err := saslMechanismsFromConfig(conf)
	if err != nil {
		return nil, err
	}
	if saslConfs != nil && len(saslConfs) > 0 {
		opts = append(opts, kgo.SASL(saslConfs...))
	}

	return opts, nil
}

type Opt func(*Agent)

func WithAgentId(id string) Opt {
	return func(a *Agent) {
		a.id = id
	}
}

func WithKafkaBackboneOpts(opts ...kgo.Opt) Opt {
	return func(a *Agent) {
		a.backboneOpts = append(a.backboneOpts, opts...)
	}
}

func WithShonoOpts(opts ...kgo.Opt) Opt {
	return func(a *Agent) {
		a.shonoOpts = append(a.shonoOpts, opts...)
	}
}

func NewAgent(opts ...Opt) (*Agent, error) {
	a := &Agent{}

	for _, opt := range opts {
		opt(a)
	}

	return a, nil
}

type Agent struct {
	id           string
	backboneOpts []kgo.Opt
	shonoOpts    []kgo.Opt

	bb backbone.Backbone
}

func (a *Agent) Connect() error {
	// -- initialize the backbone
	a.bb = backbone.NewBackbone(a.id, a.backboneOpts...)

	return nil
}

func (a *Agent) Close() error {
	if a.bb != nil {
		a.bb.Close()
	}

	return nil
}
