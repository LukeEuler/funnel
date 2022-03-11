package core

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/LukeEuler/funnel/antlr4/parser"
	"github.com/LukeEuler/funnel/model"
)

/*
grammar Match;

sentence    : clause;

clause      : condition
            | NOT clause
            | clause AND clause
            | clause OR clause
            | '(' clause ')'
            ;

condition   : key EXIST
            | key EQUAL VALUE
            | key CONTAINS VALUE
            ;

key         : KI ('.' KI)*;

// Tokens
AND         : '&';
OR          : '|';
NOT         : '!';
EXIST       : '+';
EQUAL       : '=';
CONTAINS    : '>';
KI          : ([a-zA-Z] | '0'..'9' | '_')+;
VALUE       : '\'' .*? '\'';
WHITESPACE  : [ \r\n\t]+ -> skip;
*/
type listener struct {
	rule model.Rule
	data model.Data

	// 一个中间态数据
	lastClause string
	record     map[string]bool
}

func (l *listener) set(rule model.Rule, data model.Data) {
	l.rule = rule
	l.data = data
	l.lastClause = ""
	l.record = make(map[string]bool)
}

// VisitTerminal is called when a terminal node is visited.
func (l *listener) VisitTerminal(node antlr.TerminalNode) {
}

// VisitErrorNode is called when an error node is visited.
func (l *listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (l *listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (l *listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSentence is called when production sentence is entered.
func (l *listener) EnterSentence(ctx *parser.SentenceContext) {}

// ExitSentence is called when production sentence is exited.
func (l *listener) ExitSentence(ctx *parser.SentenceContext) {
	match := l.record[l.lastClause]
	l.data.SetRule(l.rule, match)
}

// EnterClause is called when production clause is entered.
func (l *listener) EnterClause(ctx *parser.ClauseContext) {}

/*
clause      : condition
            | NOT clause
            | clause AND clause
            | clause OR clause
            | '(' clause ')'
            ;
*/

// ExitClause is called when production clause is exited.
func (l *listener) ExitClause(ctx *parser.ClauseContext) {
	mark := ctx.GetText()
	l.lastClause = mark
	_, ok := l.record[mark]
	if ok {
		// 避免重复检查
		return
	}

	switch ctx.GetChildCount() {
	case 1:
		l.record[mark] = l.record[ctx.Condition().GetText()]
	case 2:
		l.record[mark] = !l.record[ctx.Clause(0).GetText()]
	case 3:
		if ctx.AND() != nil || ctx.OR() != nil {
			left := ctx.Clause(0)
			right := ctx.Clause(1)

			if ctx.AND() != nil {
				l.record[mark] = l.record[left.GetText()] && l.record[right.GetText()]
			} else {
				l.record[mark] = l.record[left.GetText()] || l.record[right.GetText()]
			}
		} else {
			l.record[mark] = l.record[ctx.Clause(0).GetText()]
		}
	}
}

// EnterCondition is called when production condition is entered.
func (l *listener) EnterCondition(ctx *parser.ConditionContext) {}

/*
condition   : key EXIST
            | key EQUAL VALUE
            | key CONTAINS VALUE
            ;
*/

// ExitCondition is called when production condition is exited.
func (l *listener) ExitCondition(ctx *parser.ConditionContext) {
	mark := ctx.GetText()
	_, ok := l.record[mark]
	if ok {
		// 避免重复检查
		return
	}

	key := ctx.Key().GetText()

	switch ctx.GetChildCount() {
	case 2:
		l.record[mark] = l.data.KeyExist(key)
	case 3:
		value := ctx.VALUE().GetText()
		value = strings.TrimPrefix(value, "'")
		value = strings.TrimSuffix(value, "'")
		if ctx.EQUAL() != nil {
			l.record[mark] = l.data.ValueEqual(key, value)
		} else {
			l.record[mark] = l.data.ValueContains(key, value)
		}
	}
}

// EnterKey is called when production key is entered.
func (l *listener) EnterKey(ctx *parser.KeyContext) {}

// ExitKey is called when production key is exited.
func (l *listener) ExitKey(ctx *parser.KeyContext) {}
