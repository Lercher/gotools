// Code generated from c:\git\src\github.com\lercher\gotools\antlr4sample\Sample.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Sample

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 6, 10, 4,
	2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 2, 2, 3, 2, 2, 2, 2, 8, 2,
	4, 3, 2, 2, 2, 4, 5, 7, 3, 2, 2, 5, 6, 7, 4, 2, 2, 6, 7, 7, 5, 2, 2, 7,
	8, 7, 2, 2, 3, 8, 3, 3, 2, 2, 2, 2,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'Sample'", "", "'.'",
}
var symbolicNames = []string{
	"", "SAMPLE", "INT", "DOT", "WS",
}

var ruleNames = []string{
	"main",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SampleParser struct {
	*antlr.BaseParser
}

func NewSampleParser(input antlr.TokenStream) *SampleParser {
	this := new(SampleParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Sample.g4"

	return this
}

// SampleParser tokens.
const (
	SampleParserEOF    = antlr.TokenEOF
	SampleParserSAMPLE = 1
	SampleParserINT    = 2
	SampleParserDOT    = 3
	SampleParserWS     = 4
)

// SampleParserRULE_main is the SampleParser rule.
const SampleParserRULE_main = 0

// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSamplenum returns the samplenum token.
	GetSamplenum() antlr.Token

	// SetSamplenum sets the samplenum token.
	SetSamplenum(antlr.Token)

	// IsMainContext differentiates from other interfaces.
	IsMainContext()
}

type MainContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	samplenum antlr.Token
}

func NewEmptyMainContext() *MainContext {
	var p = new(MainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleParserRULE_main
	return p
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleParserRULE_main

	return p
}

func (s *MainContext) GetParser() antlr.Parser { return s.parser }

func (s *MainContext) GetSamplenum() antlr.Token { return s.samplenum }

func (s *MainContext) SetSamplenum(v antlr.Token) { s.samplenum = v }

func (s *MainContext) SAMPLE() antlr.TerminalNode {
	return s.GetToken(SampleParserSAMPLE, 0)
}

func (s *MainContext) DOT() antlr.TerminalNode {
	return s.GetToken(SampleParserDOT, 0)
}

func (s *MainContext) EOF() antlr.TerminalNode {
	return s.GetToken(SampleParserEOF, 0)
}

func (s *MainContext) INT() antlr.TerminalNode {
	return s.GetToken(SampleParserINT, 0)
}

func (s *MainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleListener); ok {
		listenerT.EnterMain(s)
	}
}

func (s *MainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleListener); ok {
		listenerT.ExitMain(s)
	}
}

func (p *SampleParser) Main() (localctx IMainContext) {
	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SampleParserRULE_main)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(2)
		p.Match(SampleParserSAMPLE)
	}
	{
		p.SetState(3)

		var _m = p.Match(SampleParserINT)

		localctx.(*MainContext).samplenum = _m
	}
	{
		p.SetState(4)
		p.Match(SampleParserDOT)
	}
	{
		p.SetState(5)
		p.Match(SampleParserEOF)
	}

	return localctx
}
