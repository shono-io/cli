package generator

import (
	"github.com/shono-io/shono/decl"
)

func Generate(path string) error {
	// -- get the inventory from the path
	inv, err := decl.NewInventory(path)
	if err != nil {
		return err
	}

	generateInjectors(inv)

	if err := generateReactors(inv); err != nil {
		return err
	}

	return nil
}
