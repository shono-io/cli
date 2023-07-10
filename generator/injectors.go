package generator

import (
	"github.com/shono-io/shono/artifacts/benthos"
	"github.com/shono-io/shono/inventory"
	"github.com/shono-io/shono/local"
	"github.com/sirupsen/logrus"
)

func generateInjectors(inv inventory.Inventory) {
	injectors, err := inv.ListInjectorsForScope(inventory.NewScopeReference("todo"))
	if err != nil {
		logrus.Panicf("failed to list injectors: %v", err)
	}

	for _, i := range injectors {
		artifact, err := benthos.NewInjectorGenerator().Generate("todo_task_injector", i.Code, inv, i.Reference())
		if err != nil {
			logrus.Panicf("failed to generate injector artifact: %v", err)
		}

		if err := local.DumpArtifact(artifact); err != nil {
			logrus.Panicf("failed to dump artifact: %v", err)
		}
	}
}
