// Code generated from Match.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Match

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 55, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 21, 10, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 7, 3, 29, 10, 3, 12, 3, 14, 3, 32, 11, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 45, 10, 4, 3,
	5, 3, 5, 3, 5, 7, 5, 50, 10, 5, 12, 5, 14, 5, 53, 11, 5, 3, 5, 2, 3, 4,
	6, 2, 4, 6, 8, 2, 2, 2, 57, 2, 10, 3, 2, 2, 2, 4, 20, 3, 2, 2, 2, 6, 44,
	3, 2, 2, 2, 8, 46, 3, 2, 2, 2, 10, 11, 5, 4, 3, 2, 11, 3, 3, 2, 2, 2, 12,
	13, 8, 3, 1, 2, 13, 21, 5, 6, 4, 2, 14, 15, 7, 8, 2, 2, 15, 21, 5, 4, 3,
	6, 16, 17, 7, 3, 2, 2, 17, 18, 5, 4, 3, 2, 18, 19, 7, 4, 2, 2, 19, 21,
	3, 2, 2, 2, 20, 12, 3, 2, 2, 2, 20, 14, 3, 2, 2, 2, 20, 16, 3, 2, 2, 2,
	21, 30, 3, 2, 2, 2, 22, 23, 12, 5, 2, 2, 23, 24, 7, 6, 2, 2, 24, 29, 5,
	4, 3, 6, 25, 26, 12, 4, 2, 2, 26, 27, 7, 7, 2, 2, 27, 29, 5, 4, 3, 5, 28,
	22, 3, 2, 2, 2, 28, 25, 3, 2, 2, 2, 29, 32, 3, 2, 2, 2, 30, 28, 3, 2, 2,
	2, 30, 31, 3, 2, 2, 2, 31, 5, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 33, 34, 5,
	8, 5, 2, 34, 35, 7, 9, 2, 2, 35, 45, 3, 2, 2, 2, 36, 37, 5, 8, 5, 2, 37,
	38, 7, 10, 2, 2, 38, 39, 7, 13, 2, 2, 39, 45, 3, 2, 2, 2, 40, 41, 5, 8,
	5, 2, 41, 42, 7, 11, 2, 2, 42, 43, 7, 13, 2, 2, 43, 45, 3, 2, 2, 2, 44,
	33, 3, 2, 2, 2, 44, 36, 3, 2, 2, 2, 44, 40, 3, 2, 2, 2, 45, 7, 3, 2, 2,
	2, 46, 51, 7, 12, 2, 2, 47, 48, 7, 5, 2, 2, 48, 50, 7, 12, 2, 2, 49, 47,
	3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2,
	52, 9, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 7, 20, 28, 30, 44, 51,
}
var literalNames = []string{
	"", "'('", "')'", "'.'", "'&'", "'|'", "'!'", "'+'", "'='", "'>'",
}
var symbolicNames = []string{
	"", "", "", "", "AND", "OR", "NOT", "EXIST", "EQUAL", "CONTAINS", "KI",
	"VALUE", "WHITESPACE",
}

var ruleNames = []string{
	"sentence", "clause", "condition", "key",
}

type MatchParser struct {
	*antlr.BaseParser
}

// NewMatchParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *MatchParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewMatchParser(input antlr.TokenStream) *MatchParser {
	this := new(MatchParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Match.g4"

	return this
}

// MatchParser tokens.
const (
	MatchParserEOF        = antlr.TokenEOF
	MatchParserT__0       = 1
	MatchParserT__1       = 2
	MatchParserT__2       = 3
	MatchParserAND        = 4
	MatchParserOR         = 5
	MatchParserNOT        = 6
	MatchParserEXIST      = 7
	MatchParserEQUAL      = 8
	MatchParserCONTAINS   = 9
	MatchParserKI         = 10
	MatchParserVALUE      = 11
	MatchParserWHITESPACE = 12
)

// MatchParser rules.
const (
	MatchParserRULE_sentence  = 0
	MatchParserRULE_clause    = 1
	MatchParserRULE_condition = 2
	MatchParserRULE_key       = 3
)

// ISentenceContext is an interface to support dynamic dispatch.
type ISentenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSentenceContext differentiates from other interfaces.
	IsSentenceContext()
}

type SentenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenceContext() *SentenceContext {
	var p = new(SentenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_sentence
	return p
}

func (*SentenceContext) IsSentenceContext() {}

func NewSentenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenceContext {
	var p = new(SentenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_sentence

	return p
}

func (s *SentenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenceContext) Clause() IClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClauseContext)
}

func (s *SentenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterSentence(s)
	}
}

func (s *SentenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitSentence(s)
	}
}

func (p *MatchParser) Sentence() (localctx ISentenceContext) {
	localctx = NewSentenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MatchParserRULE_sentence)

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
		p.SetState(8)
		p.clause(0)
	}

	return localctx
}

// IClauseContext is an interface to support dynamic dispatch.
type IClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClauseContext differentiates from other interfaces.
	IsClauseContext()
}

type ClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClauseContext() *ClauseContext {
	var p = new(ClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_clause
	return p
}

func (*ClauseContext) IsClauseContext() {}

func NewClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseContext {
	var p = new(ClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_clause

	return p
}

func (s *ClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ClauseContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ClauseContext) NOT() antlr.TerminalNode {
	return s.GetToken(MatchParserNOT, 0)
}

func (s *ClauseContext) AllClause() []IClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClauseContext)(nil)).Elem())
	var tst = make([]IClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClauseContext)
		}
	}

	return tst
}

func (s *ClauseContext) Clause(i int) IClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClauseContext)
}

func (s *ClauseContext) AND() antlr.TerminalNode {
	return s.GetToken(MatchParserAND, 0)
}

func (s *ClauseContext) OR() antlr.TerminalNode {
	return s.GetToken(MatchParserOR, 0)
}

func (s *ClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterClause(s)
	}
}

func (s *ClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitClause(s)
	}
}

func (p *MatchParser) Clause() (localctx IClauseContext) {
	return p.clause(0)
}

func (p *MatchParser) clause(_p int) (localctx IClauseContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewClauseContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IClauseContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, MatchParserRULE_clause, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(18)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MatchParserKI:
		{
			p.SetState(11)
			p.Condition()
		}

	case MatchParserNOT:
		{
			p.SetState(12)
			p.Match(MatchParserNOT)
		}
		{
			p.SetState(13)
			p.clause(4)
		}

	case MatchParserT__0:
		{
			p.SetState(14)
			p.Match(MatchParserT__0)
		}
		{
			p.SetState(15)
			p.clause(0)
		}
		{
			p.SetState(16)
			p.Match(MatchParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(26)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewClauseContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MatchParserRULE_clause)
				p.SetState(20)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(21)
					p.Match(MatchParserAND)
				}
				{
					p.SetState(22)
					p.clause(4)
				}

			case 2:
				localctx = NewClauseContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MatchParserRULE_clause)
				p.SetState(23)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(24)
					p.Match(MatchParserOR)
				}
				{
					p.SetState(25)
					p.clause(3)
				}

			}

		}
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *ConditionContext) EXIST() antlr.TerminalNode {
	return s.GetToken(MatchParserEXIST, 0)
}

func (s *ConditionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(MatchParserEQUAL, 0)
}

func (s *ConditionContext) VALUE() antlr.TerminalNode {
	return s.GetToken(MatchParserVALUE, 0)
}

func (s *ConditionContext) CONTAINS() antlr.TerminalNode {
	return s.GetToken(MatchParserCONTAINS, 0)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *MatchParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MatchParserRULE_condition)

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

	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(31)
			p.Key()
		}
		{
			p.SetState(32)
			p.Match(MatchParserEXIST)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.Key()
		}
		{
			p.SetState(35)
			p.Match(MatchParserEQUAL)
		}
		{
			p.SetState(36)
			p.Match(MatchParserVALUE)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(38)
			p.Key()
		}
		{
			p.SetState(39)
			p.Match(MatchParserCONTAINS)
		}
		{
			p.SetState(40)
			p.Match(MatchParserVALUE)
		}

	}

	return localctx
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MatchParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) AllKI() []antlr.TerminalNode {
	return s.GetTokens(MatchParserKI)
}

func (s *KeyContext) KI(i int) antlr.TerminalNode {
	return s.GetToken(MatchParserKI, i)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitKey(s)
	}
}

func (p *MatchParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MatchParserRULE_key)
	var _la int

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
		p.SetState(44)
		p.Match(MatchParserKI)
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MatchParserT__2 {
		{
			p.SetState(45)
			p.Match(MatchParserT__2)
		}
		{
			p.SetState(46)
			p.Match(MatchParserKI)
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

func (p *MatchParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ClauseContext = nil
		if localctx != nil {
			t = localctx.(*ClauseContext)
		}
		return p.Clause_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MatchParser) Clause_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
