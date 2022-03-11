package core

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/LukeEuler/funnel/antlr4/parser"
	"github.com/LukeEuler/funnel/model"
)

var lis = &listener{}

type worker struct {
	data model.Data
	rule model.Rule
	err  error
}

func (w *worker) IsErr() bool {
	return w.err != nil
}

func newWorker(rule model.Rule, data model.Data) *worker {
	return &worker{
		data: data,
		rule: rule,
		err:  nil,
	}
}

func Match(rule model.Rule, data model.Data) (bool, error) {
	tree := getSentenceContext(rule.GetContent())
	w := newWorker(rule, data)
	lis.set(w)
	antlr.ParseTreeWalkerDefault.Walk(lis, tree)
	return data.Match(rule), w.err
}

func getSentenceContext(rule string) parser.ISentenceContext {
	// Setup the input
	input := antlr.NewInputStream(rule)
	// Create the Lexer
	lexer := parser.NewMatchLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	queryParser := parser.NewMatchParser(stream)
	queryParser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	queryParser.BuildParseTrees = true
	return queryParser.Sentence()
}
