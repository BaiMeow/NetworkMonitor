// Code generated from BirdOSPF.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // BirdOSPF
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

type BirdOSPFParser struct {
	*antlr.BaseParser
}

var BirdOSPFParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func birdospfParserInit() {
	staticData := &BirdOSPFParserStaticData
	staticData.LiteralNames = []string{
		"", "'area'", "'distance'", "'router'", "'metric'", "'stubnet'", "'xnetwork'",
		"'external'", "'metric2'", "'xrouter'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "Prefix", "VERSION", "IP", "STRING",
		"INT", "WS",
	}
	staticData.RuleNames = []string{
		"state", "area", "router", "routerEntry",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 57, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 4,
		0, 10, 8, 0, 11, 0, 12, 0, 11, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 18, 8, 1,
		1, 1, 5, 1, 21, 8, 1, 10, 1, 12, 1, 24, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 4, 2, 31, 8, 2, 11, 2, 12, 2, 32, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 3, 3, 55, 8, 3, 1, 3, 0, 0, 4, 0, 2, 4, 6, 0, 0, 61, 0, 9,
		1, 0, 0, 0, 2, 13, 1, 0, 0, 0, 4, 25, 1, 0, 0, 0, 6, 54, 1, 0, 0, 0, 8,
		10, 3, 2, 1, 0, 9, 8, 1, 0, 0, 0, 10, 11, 1, 0, 0, 0, 11, 9, 1, 0, 0, 0,
		11, 12, 1, 0, 0, 0, 12, 1, 1, 0, 0, 0, 13, 14, 5, 1, 0, 0, 14, 17, 5, 12,
		0, 0, 15, 16, 5, 2, 0, 0, 16, 18, 5, 14, 0, 0, 17, 15, 1, 0, 0, 0, 17,
		18, 1, 0, 0, 0, 18, 22, 1, 0, 0, 0, 19, 21, 3, 4, 2, 0, 20, 19, 1, 0, 0,
		0, 21, 24, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 3, 1,
		0, 0, 0, 24, 22, 1, 0, 0, 0, 25, 26, 5, 3, 0, 0, 26, 30, 5, 12, 0, 0, 27,
		28, 5, 2, 0, 0, 28, 31, 5, 14, 0, 0, 29, 31, 3, 6, 3, 0, 30, 27, 1, 0,
		0, 0, 30, 29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 32, 33,
		1, 0, 0, 0, 33, 5, 1, 0, 0, 0, 34, 35, 5, 3, 0, 0, 35, 36, 5, 12, 0, 0,
		36, 37, 5, 4, 0, 0, 37, 55, 5, 14, 0, 0, 38, 39, 5, 5, 0, 0, 39, 40, 5,
		10, 0, 0, 40, 41, 5, 4, 0, 0, 41, 55, 5, 14, 0, 0, 42, 43, 5, 6, 0, 0,
		43, 44, 5, 10, 0, 0, 44, 45, 5, 4, 0, 0, 45, 55, 5, 14, 0, 0, 46, 47, 5,
		7, 0, 0, 47, 48, 5, 10, 0, 0, 48, 49, 5, 8, 0, 0, 49, 55, 5, 14, 0, 0,
		50, 51, 5, 9, 0, 0, 51, 52, 5, 12, 0, 0, 52, 53, 5, 4, 0, 0, 53, 55, 5,
		14, 0, 0, 54, 34, 1, 0, 0, 0, 54, 38, 1, 0, 0, 0, 54, 42, 1, 0, 0, 0, 54,
		46, 1, 0, 0, 0, 54, 50, 1, 0, 0, 0, 55, 7, 1, 0, 0, 0, 6, 11, 17, 22, 30,
		32, 54,
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

// BirdOSPFParserInit initializes any static state used to implement BirdOSPFParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBirdOSPFParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BirdOSPFParserInit() {
	staticData := &BirdOSPFParserStaticData
	staticData.once.Do(birdospfParserInit)
}

// NewBirdOSPFParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBirdOSPFParser(input antlr.TokenStream) *BirdOSPFParser {
	BirdOSPFParserInit()
	this := new(BirdOSPFParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &BirdOSPFParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "BirdOSPF.g4"

	return this
}

// BirdOSPFParser tokens.
const (
	BirdOSPFParserEOF     = antlr.TokenEOF
	BirdOSPFParserT__0    = 1
	BirdOSPFParserT__1    = 2
	BirdOSPFParserT__2    = 3
	BirdOSPFParserT__3    = 4
	BirdOSPFParserT__4    = 5
	BirdOSPFParserT__5    = 6
	BirdOSPFParserT__6    = 7
	BirdOSPFParserT__7    = 8
	BirdOSPFParserT__8    = 9
	BirdOSPFParserPrefix  = 10
	BirdOSPFParserVERSION = 11
	BirdOSPFParserIP      = 12
	BirdOSPFParserSTRING  = 13
	BirdOSPFParserINT     = 14
	BirdOSPFParserWS      = 15
)

// BirdOSPFParser rules.
const (
	BirdOSPFParserRULE_state       = 0
	BirdOSPFParserRULE_area        = 1
	BirdOSPFParserRULE_router      = 2
	BirdOSPFParserRULE_routerEntry = 3
)

// IStateContext is an interface to support dynamic dispatch.
type IStateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArea() []IAreaContext
	Area(i int) IAreaContext

	// IsStateContext differentiates from other interfaces.
	IsStateContext()
}

type StateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateContext() *StateContext {
	var p = new(StateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BirdOSPFParserRULE_state
	return p
}

func InitEmptyStateContext(p *StateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BirdOSPFParserRULE_state
}

func (*StateContext) IsStateContext() {}

func NewStateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateContext {
	var p = new(StateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BirdOSPFParserRULE_state

	return p
}

func (s *StateContext) GetParser() antlr.Parser { return s.parser }

func (s *StateContext) AllArea() []IAreaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAreaContext); ok {
			len++
		}
	}

	tst := make([]IAreaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAreaContext); ok {
			tst[i] = t.(IAreaContext)
			i++
		}
	}

	return tst
}

func (s *StateContext) Area(i int) IAreaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAreaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAreaContext)
}

func (s *StateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BirdOSPFListener); ok {
		listenerT.EnterState(s)
	}
}

func (s *StateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BirdOSPFListener); ok {
		listenerT.ExitState(s)
	}
}

func (s *StateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BirdOSPFVisitor:
		return t.VisitState(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BirdOSPFParser) State() (localctx IStateContext) {
	localctx = NewStateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BirdOSPFParserRULE_state)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(9)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == BirdOSPFParserT__0 {
		{
			p.SetState(8)
			p.Area()
		}

		p.SetState(11)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IAreaContext is an interface to support dynamic dispatch.
type IAreaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IP() antlr.TerminalNode
	INT() antlr.TerminalNode
	AllRouter() []IRouterContext
	Router(i int) IRouterContext

	// IsAreaContext differentiates from other interfaces.
	IsAreaContext()
}

type AreaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAreaContext() *AreaContext {
	var p = new(AreaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BirdOSPFParserRULE_area
	return p
}

func InitEmptyAreaContext(p *AreaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BirdOSPFParserRULE_area
}

func (*AreaContext) IsAreaContext() {}

func NewAreaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AreaContext {
	var p = new(AreaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BirdOSPFParserRULE_area

	return p
}

func (s *AreaContext) GetParser() antlr.Parser { return s.parser }

func (s *AreaContext) IP() antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserIP, 0)
}

func (s *AreaContext) INT() antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserINT, 0)
}

func (s *AreaContext) AllRouter() []IRouterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRouterContext); ok {
			len++
		}
	}

	tst := make([]IRouterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRouterContext); ok {
			tst[i] = t.(IRouterContext)
			i++
		}
	}

	return tst
}

func (s *AreaContext) Router(i int) IRouterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRouterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRouterContext)
}

func (s *AreaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AreaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AreaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BirdOSPFListener); ok {
		listenerT.EnterArea(s)
	}
}

func (s *AreaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BirdOSPFListener); ok {
		listenerT.ExitArea(s)
	}
}

func (s *AreaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BirdOSPFVisitor:
		return t.VisitArea(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BirdOSPFParser) Area() (localctx IAreaContext) {
	localctx = NewAreaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BirdOSPFParserRULE_area)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(13)
		p.Match(BirdOSPFParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(14)
		p.Match(BirdOSPFParserIP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BirdOSPFParserT__1 {
		{
			p.SetState(15)
			p.Match(BirdOSPFParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(16)
			p.Match(BirdOSPFParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BirdOSPFParserT__2 {
		{
			p.SetState(19)
			p.Router()
		}

		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IRouterContext is an interface to support dynamic dispatch.
type IRouterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IP() antlr.TerminalNode
	AllRouterEntry() []IRouterEntryContext
	RouterEntry(i int) IRouterEntryContext
	AllINT() []antlr.TerminalNode
	INT(i int) antlr.TerminalNode

	// IsRouterContext differentiates from other interfaces.
	IsRouterContext()
}

type RouterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRouterContext() *RouterContext {
	var p = new(RouterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BirdOSPFParserRULE_router
	return p
}

func InitEmptyRouterContext(p *RouterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BirdOSPFParserRULE_router
}

func (*RouterContext) IsRouterContext() {}

func NewRouterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RouterContext {
	var p = new(RouterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BirdOSPFParserRULE_router

	return p
}

func (s *RouterContext) GetParser() antlr.Parser { return s.parser }

func (s *RouterContext) IP() antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserIP, 0)
}

func (s *RouterContext) AllRouterEntry() []IRouterEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRouterEntryContext); ok {
			len++
		}
	}

	tst := make([]IRouterEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRouterEntryContext); ok {
			tst[i] = t.(IRouterEntryContext)
			i++
		}
	}

	return tst
}

func (s *RouterContext) RouterEntry(i int) IRouterEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRouterEntryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRouterEntryContext)
}

func (s *RouterContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(BirdOSPFParserINT)
}

func (s *RouterContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserINT, i)
}

func (s *RouterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RouterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RouterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BirdOSPFListener); ok {
		listenerT.EnterRouter(s)
	}
}

func (s *RouterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BirdOSPFListener); ok {
		listenerT.ExitRouter(s)
	}
}

func (s *RouterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BirdOSPFVisitor:
		return t.VisitRouter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BirdOSPFParser) Router() (localctx IRouterContext) {
	localctx = NewRouterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BirdOSPFParserRULE_router)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(BirdOSPFParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(26)
		p.Match(BirdOSPFParserIP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(30)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case BirdOSPFParserT__1:
				{
					p.SetState(27)
					p.Match(BirdOSPFParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(28)
					p.Match(BirdOSPFParserINT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case BirdOSPFParserT__2, BirdOSPFParserT__4, BirdOSPFParserT__5, BirdOSPFParserT__6, BirdOSPFParserT__8:
				{
					p.SetState(29)
					p.RouterEntry()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
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

// IRouterEntryContext is an interface to support dynamic dispatch.
type IRouterEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IP() antlr.TerminalNode
	INT() antlr.TerminalNode
	Prefix() antlr.TerminalNode

	// IsRouterEntryContext differentiates from other interfaces.
	IsRouterEntryContext()
}

type RouterEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRouterEntryContext() *RouterEntryContext {
	var p = new(RouterEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BirdOSPFParserRULE_routerEntry
	return p
}

func InitEmptyRouterEntryContext(p *RouterEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BirdOSPFParserRULE_routerEntry
}

func (*RouterEntryContext) IsRouterEntryContext() {}

func NewRouterEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RouterEntryContext {
	var p = new(RouterEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BirdOSPFParserRULE_routerEntry

	return p
}

func (s *RouterEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *RouterEntryContext) IP() antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserIP, 0)
}

func (s *RouterEntryContext) INT() antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserINT, 0)
}

func (s *RouterEntryContext) Prefix() antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserPrefix, 0)
}

func (s *RouterEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RouterEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RouterEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BirdOSPFListener); ok {
		listenerT.EnterRouterEntry(s)
	}
}

func (s *RouterEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BirdOSPFListener); ok {
		listenerT.ExitRouterEntry(s)
	}
}

func (s *RouterEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BirdOSPFVisitor:
		return t.VisitRouterEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BirdOSPFParser) RouterEntry() (localctx IRouterEntryContext) {
	localctx = NewRouterEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BirdOSPFParserRULE_routerEntry)
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BirdOSPFParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Match(BirdOSPFParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(35)
			p.Match(BirdOSPFParserIP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(36)
			p.Match(BirdOSPFParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.Match(BirdOSPFParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BirdOSPFParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.Match(BirdOSPFParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Match(BirdOSPFParserPrefix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(40)
			p.Match(BirdOSPFParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.Match(BirdOSPFParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BirdOSPFParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(42)
			p.Match(BirdOSPFParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(43)
			p.Match(BirdOSPFParserPrefix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(44)
			p.Match(BirdOSPFParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.Match(BirdOSPFParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BirdOSPFParserT__6:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(46)
			p.Match(BirdOSPFParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Match(BirdOSPFParserPrefix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(48)
			p.Match(BirdOSPFParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)
			p.Match(BirdOSPFParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BirdOSPFParserT__8:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(50)
			p.Match(BirdOSPFParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Match(BirdOSPFParserIP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(52)
			p.Match(BirdOSPFParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.Match(BirdOSPFParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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
