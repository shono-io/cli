package logics

import (
	"context"
	go_shono "github.com/shono-io/go-shono"
	"github.com/shono-io/go-shono/events"
	"github.com/shono-io/go-shono/shono"
	"github.com/sirupsen/logrus"
)

func NewScopeLogic(b shono.Backbone) *ScopeLogic {
	return &ScopeLogic{b: b}
}

type ScopeLogic struct {
	b shono.Backbone
}

func (l *ScopeLogic) OnCreated(ctx context.Context, evt any, w go_shono.Writer) {
	if err := l.b.Apply(events.ScopeCreated.EventId, evt); err != nil {
		logrus.Errorf("failed to apply event: %v", err)
		return
	}
	logrus.Infof("Scope created: %v", evt)
}

func (l *ScopeLogic) OnDeleted(ctx context.Context, evt any, w go_shono.Writer) {
	if err := l.b.Apply(events.ScopeDeleted.EventId, evt); err != nil {
		logrus.Errorf("failed to apply event: %v", err)
		return
	}
	logrus.Infof("Scope deleted: %v", evt)
}
