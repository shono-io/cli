package reaktors

import (
	go_shono "github.com/shono-io/go-shono"
	"github.com/shono-io/go-shono/events"
	"github.com/shono-io/go-shono/shono"
	"github.com/shono-io/shono-agent/logics"
)

func RegisterScopeRoutes(r *go_shono.Router, b shono.Backbone) {
	l := logics.NewScopeLogic(b)

	r.Register(go_shono.MustNewReaktor(string(events.ScopeCreated.EventId),
		go_shono.ListenFor(events.ScopeCreated),
		go_shono.WithHandler(l.OnCreated)))

	r.Register(go_shono.MustNewReaktor(string(events.ScopeDeleted.EventId),
		go_shono.ListenFor(events.ScopeDeleted),
		go_shono.WithHandler(l.OnDeleted)))
}
