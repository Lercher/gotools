// Code generated from c://git//src//github.com//lercher//gotools//antlr4sample//Sample.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Sample

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SampleParser struct {
	*antlr.BaseParser
}

var SampleParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sampleParserInit() {
	staticData := &SampleParserStaticData
	staticData.LiteralNames = []string{
		"", "'Sample'", "", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "SAMPLE", "INT", "DOT", "WS",
	}
	staticData.RuleNames = []string{
		"main",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 4, 8, 2, 0, 7, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0,
		0, 0, 6, 0, 2, 1, 0, 0, 0, 2, 3, 5, 1, 0, 0, 3, 4, 5, 2, 0, 0, 4, 5, 5,
		3, 0, 0, 5, 6, 5, 0, 0, 1, 6, 1, 1, 0, 0, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SampleParserInit initializes any static state used to implement SampleParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSampleParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SampleParserInit() {
	staticData := &SampleParserStaticData
	staticData.once.Do(sampleParserInit)
}

// NewSampleParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSampleParser(input antlr.TokenStream) *SampleParser {
	SampleParserInit()
	this := new(SampleParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SampleParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
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

	// Getter signatures
	SAMPLE() antlr.TerminalNode
	DOT() antlr.TerminalNode
	EOF() antlr.TerminalNode
	INT() antlr.TerminalNode

	// IsMainContext differentiates from other interfaces.
	IsMainContext()
}

type MainContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	samplenum antlr.Token
}

func NewEmptyMainContext() *MainContext {
	var p = new(MainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SampleParserRULE_main
	return p
}

func InitEmptyMainContext(p *MainContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SampleParserRULE_main
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(2)
		p.Match(SampleParserSAMPLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(3)

		var _m = p.Match(SampleParserINT)

		localctx.(*MainContext).samplenum = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(4)
		p.Match(SampleParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(5)
		p.Match(SampleParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
