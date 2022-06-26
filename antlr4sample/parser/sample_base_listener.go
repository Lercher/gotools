// Code generated from c:\git\src\github.com\lercher\gotools\antlr4sample\Sample.g by ANTLR 4.9.2. DO NOT EDIT.

package parser // Sample

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSampleListener is a complete listener for a parse tree produced by SampleParser.
type BaseSampleListener struct{}

var _ SampleListener = &BaseSampleListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSampleListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSampleListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSampleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSampleListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMain is called when production main is entered.
func (s *BaseSampleListener) EnterMain(ctx *MainContext) {}

// ExitMain is called when production main is exited.
func (s *BaseSampleListener) ExitMain(ctx *MainContext) {}
