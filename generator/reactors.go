package generator

import (
	"fmt"
	"github.com/shono-io/shono/artifacts"
	"github.com/shono-io/shono/artifacts/benthos"
	"github.com/shono-io/shono/inventory"
	"github.com/shono-io/shono/local"
)

func generateReactors(inv inventory.Inventory) error {
	scopes, err := inv.ListScopes()
	if err != nil {
		return err
	}

	for _, s := range scopes {
		concepts, err := inv.ListConceptsForScope(s.Reference())
		if err != nil {
			return err
		}

		for _, c := range concepts {
			_, err := generateReactor(inv, c)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func generateReactor(inv inventory.Inventory, concept inventory.Concept) (*artifacts.Artifact, error) {
	// -- generate the artifacts for all the reactors in the registry
	appId := fmt.Sprintf("%s_%s", concept.Scope.Code(), concept.Code)
	artifact, err := benthos.NewConceptGenerator().Generate(appId, "reactors", inv, concept.Reference())
	if err != nil {
		return nil, err
	}

	if err := local.DumpArtifact(artifact); err != nil {
		return nil, err
	}

	return artifact, nil
}
