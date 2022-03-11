package core

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/LukeEuler/funnel/antlr4/parser"
	"github.com/LukeEuler/funnel/model"
)

var lis = &listener{}

func Match(rule model.Rule, data model.Data) bool {
	tree := getSentenceContext(rule.GetContent())
	lis.set(rule, data)
	antlr.ParseTreeWalkerDefault.Walk(lis, tree)
	return data.Match(rule)
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
