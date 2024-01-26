package main

import (
	"fmt"
	"html/template"
)

// Environment combines data about a rendering and the preparation of a postback
type Environment struct {
	Language string
	Username string
	Namer
	BoundValue BoundValues
}

// NewEnvironment creates a new Environment
func NewEnvironment() *Environment {
	return &Environment{
		"DE",
		"SampleUser",
		createNamer(1000),
		make(BoundValues),
	}
}

// BoundValues provides a Put method for a Name
type BoundValues map[Name]*BoundValue

// Name is something we store a formatted value in an html form or index an action with
type Name string

// Namer is a function to provide a sequence of names
type Namer func() Name

func createNamer(i int) Namer {
	return func() Name {
		i++
		return Name(fmt.Sprintf("n%d", i))
	}
}

// FuncMap provides an Environment bound function map for templates
func (e *Environment) FuncMap() template.FuncMap {
	return template.FuncMap{
		"Username": func() string {
			return e.Username
		},
		"Language": func() string {
			return e.Language
		},
		"Name": func() Name {
			return e.Namer()
		},
		"bind": func(dot interface{}, prop, label string) *BoundValue {
			name := e.Namer()
			bv := &BoundValue{
				dot,
				prop,
				label,
				name,
			}
			e.BoundValue[name] = bv
			return bv
		},
	}
}

// BoundValue is context information for accessing a value of a JSON object
// under a Name with the ability to Put someting back on a roundtrip.
type BoundValue struct {
	dot      interface{}
	Property string
	Label    string
	Name     Name
}

// Value returns the Property of the dot
func (b *BoundValue) Value() interface{} {
	switch dot := b.dot.(type) {
	case map[string]interface{}:
		return dot[b.Property]
	default:
		return fmt.Errorf("%q is not a property of %v", b.Property, b.dot)
	}
}

// Put returns a value to its location in the generic JSON structure
func (b *BoundValue) Put(value interface{}) error {
	switch dot := b.dot.(type) {
	case map[string]interface{}:
		dot[b.Property] = value
		return nil
	default:
		return fmt.Errorf("put: %q is not a property of %v", b.Property, b.dot)
	}
}
