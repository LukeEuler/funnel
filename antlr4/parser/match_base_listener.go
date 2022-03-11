// Code generated from Match.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Match

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMatchListener is a complete listener for a parse tree produced by MatchParser.
type BaseMatchListener struct{}

var _ MatchListener = &BaseMatchListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMatchListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMatchListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMatchListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMatchListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSentence is called when production sentence is entered.
func (s *BaseMatchListener) EnterSentence(ctx *SentenceContext) {}

// ExitSentence is called when production sentence is exited.
func (s *BaseMatchListener) ExitSentence(ctx *SentenceContext) {}

// EnterClause is called when production clause is entered.
func (s *BaseMatchListener) EnterClause(ctx *ClauseContext) {}

// ExitClause is called when production clause is exited.
func (s *BaseMatchListener) ExitClause(ctx *ClauseContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseMatchListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseMatchListener) ExitCondition(ctx *ConditionContext) {}

// EnterKey is called when production key is entered.
func (s *BaseMatchListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseMatchListener) ExitKey(ctx *KeyContext) {}
