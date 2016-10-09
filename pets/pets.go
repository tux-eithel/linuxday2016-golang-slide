// Package pets defines a basic animal e other types
package pets

// Pets is generic (cute) domestic animal
type Pets struct {
	Name string
}

// MyName returns the pet's name
func (p *Pets) MyName() string {
	return p.Name
}

// Duck is a stange pet, but in this modern times not too much
type Duck struct {
	Pets // embedded struct // HL
	Call string
}

// Speak returns the pet's call
func (d *Duck) Speak() string {
	return d.Call
}
