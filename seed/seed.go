// -----------------------------
// Copyright TangoJ Labs, LLC
// Apache 2.0 License
// -----------------------------

package seed

import (
	"log"

	"github.com/seanbhart/seed/thing"
)

// Seed is a Thing object used as a viewscreen originator
// The first Seed Branch will contain a duplicate Thing - this way the entire structure is contained
type Seed struct {
	Thing thing.Thing
	Tree  Branch
}

// TODO: Limit the depth of nested Things to recall
// NewSeed initializes a new Seed opject with the mandatory data
func NewSeed(th thing.Thing) *Seed {

	// Create a new Seed
	// The Tree will automatically be filled
	seed := Seed{
		Thing: th,
		Tree:  *NewBranch(th),
	}

	return &seed
}

// Branch is used to track a chain of Things from a specific Seed
// A Branch necessarily uses nested Branches to represent the entire chain
type Branch struct {
	Thing    thing.Thing
	Branches []Branch
}

// TODO: Save the Things locally and first check if Things are local before requesting remotely again
// NewBranch creates a branch from the passed Thing (must be a Default Thing with Features) and fills the child Branches
func NewBranch(th thing.Thing) *Branch {

	// Initialize the branches slice will nil value
	// var branches []Branch
	branches := []Branch{}

	// If the Thing is a default type (has nested Things as Features), create branches with those child Things
	if th.GetType() == 0 {
		// Cast the thing as a default Thing
		dThing := th.(*thing.DefaultThing)
		log.Printf("SEED: NewBranch Feature count: %d", len(dThing.GetFeatures()))
		// Loop through the Features for nested Things and create a Branch for each
		for _, f := range dThing.GetFeatures() {
			log.Printf("SEED: NewBranch Feature: %s", f.Address)
			// Retrieve the Thing for this feature
			childThing, err := f.Retrieve()
			if err != nil {
				log.Printf("SEED: NewBranch Feature: trouble retrieving")
				continue
			}
			// Create a new branch
			childBranch := *NewBranch(*childThing)

			// Add the branch to the parent Branches list
			branches = append(branches, childBranch)
		}
	}

	branch := Branch{
		Thing:    th,
		Branches: branches,
	}

	return &branch
}
