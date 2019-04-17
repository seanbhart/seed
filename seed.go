package seed

// Seed is a Thing object used as a viewscreen originator
type Seed struct {
	Thing       Thing
	ThingBranch ThingBranch
}

// NewSeed initializes a new Seed opject with the mandatory data
func NewSeed(thing Thing) *Seed {
	return &Seed{Thing: thing}
}

// ThingBranch is used to track a chain of Things from a specific Seed
// A branch necessarily uses nested branches to represent the entire chain
type ThingBranch struct {
	Thing    Thing
	Children []ThingBranch
}
