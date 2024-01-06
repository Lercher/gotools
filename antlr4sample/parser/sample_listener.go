// Code generated from c://git//src//github.com//lercher//gotools//antlr4sample//Sample.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Sample

import "github.com/antlr4-go/antlr/v4"

// SampleListener is a complete listener for a parse tree produced by SampleParser.
type SampleListener interface {
	antlr.ParseTreeListener

	// EnterMain is called when entering the main production.
	EnterMain(c *MainContext)

	// ExitMain is called when exiting the main production.
	ExitMain(c *MainContext)
}
