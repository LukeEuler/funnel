package core

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/pkg/errors"

	"github.com/LukeEuler/funnel/antlr4/parser"
)

type listener struct {
	worker *worker

	// 一个中间态数据
	lastClause string
	record     map[string]bool
}

func (l *listener) set(worker *worker) {
	l.worker = worker
	l.lastClause = ""
	l.record = make(map[string]bool)
}

// VisitTerminal is called when a terminal node is visited.
func (l *listener) VisitTerminal(node antlr.TerminalNode) {}

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
	if l.worker.IsErr() {
		return
	}
	match := l.record[l.lastClause]
	l.worker.data.SetRule(l.worker.rule, match)
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
	if l.worker.IsErr() {
		return
	}
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
func (l *listener) EnterCondition(ctx *parser.ConditionContext) {
	if l.worker.IsErr() {
		return
	}
	if ctx.Key() == nil {
		l.worker.err = errors.Errorf("invalid condition: %s", ctx.GetText())
		return
	}
}

/*
condition   : key EXIST
            | key EQUAL VALUE
            | key CONTAINS VALUE
            ;
*/

// ExitCondition is called when production condition is exited.
func (l *listener) ExitCondition(ctx *parser.ConditionContext) {
	if l.worker.IsErr() {
		return
	}
	mark := ctx.GetText()
	_, ok := l.record[mark]
	if ok {
		// 避免重复检查
		return
	}

	key := ctx.Key().GetText()

	switch ctx.GetChildCount() {
	case 2:
		l.record[mark] = l.worker.data.KeyExist(key)
	case 3:
		value := ctx.VALUE().GetText()
		value = strings.TrimPrefix(value, "'")
		value = strings.TrimSuffix(value, "'")
		if ctx.EQUAL() != nil {
			l.record[mark] = l.worker.data.ValueEqual(key, value)
		} else {
			l.record[mark] = l.worker.data.ValueContains(key, value)
		}
	}
}

// EnterKey is called when production key is entered.
func (l *listener) EnterKey(ctx *parser.KeyContext) {}

// ExitKey is called when production key is exited.
func (l *listener) ExitKey(ctx *parser.KeyContext) {}
