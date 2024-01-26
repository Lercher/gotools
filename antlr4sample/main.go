package main

// go build && antlr4sample.exe in.sample
// -> Samplenum 42
// -> Sample42.<EOF>

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/lercher/gotools/antlr4sample/parser"
)

type sampleListener struct {
	*parser.BaseSampleListener
}

func (it *sampleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewSampleLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSampleParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Main()
	fmt.Println("Samplenum", tree.GetSamplenum().GetText())
	antlr.ParseTreeWalkerDefault.Walk(new(sampleListener), tree)
}
