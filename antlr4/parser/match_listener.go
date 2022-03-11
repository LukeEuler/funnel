// Code generated from Match.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Match

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MatchListener is a complete listener for a parse tree produced by MatchParser.
type MatchListener interface {
	antlr.ParseTreeListener

	// EnterSentence is called when entering the sentence production.
	EnterSentence(c *SentenceContext)

	// EnterClause is called when entering the clause production.
	EnterClause(c *ClauseContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// ExitSentence is called when exiting the sentence production.
	ExitSentence(c *SentenceContext)

	// ExitClause is called when exiting the clause production.
	ExitClause(c *ClauseContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)
}
