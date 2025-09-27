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
		"", "'area'", "'router'", "'network'", "'stubnet'", "'xnetwork'", "'external'",
		"'xrouter'", "'via'", "'tag'", "'dr'", "'distance'", "'unreachable'",
		"'metric'", "'metric2'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "DR", "Distance", "Unreachable",
		"Metric", "Metric2", "Prefix", "VERSION", "IP", "STRING", "INT", "WS",
		"HEX",
	}
	staticData.RuleNames = []string{
		"state", "area", "router", "network", "routerEntry", "distance",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 21, 79, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 4, 0, 14, 8, 0, 11, 0, 12, 0, 15, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 5, 1, 24, 8, 1, 10, 1, 12, 1, 27, 9, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 5, 2, 33, 8, 2, 10, 2, 12, 2, 36, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 5, 3, 45, 8, 3, 10, 3, 12, 3, 48, 9, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 62, 8, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 68, 8, 4, 1, 4, 1, 4, 3, 4, 72, 8, 4, 1,
		5, 1, 5, 1, 5, 3, 5, 77, 8, 5, 1, 5, 0, 0, 6, 0, 2, 4, 6, 8, 10, 0, 2,
		1, 0, 13, 14, 2, 0, 19, 19, 21, 21, 85, 0, 13, 1, 0, 0, 0, 2, 19, 1, 0,
		0, 0, 4, 28, 1, 0, 0, 0, 6, 37, 1, 0, 0, 0, 8, 61, 1, 0, 0, 0, 10, 76,
		1, 0, 0, 0, 12, 14, 3, 2, 1, 0, 13, 12, 1, 0, 0, 0, 14, 15, 1, 0, 0, 0,
		15, 13, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 18, 5,
		0, 0, 1, 18, 1, 1, 0, 0, 0, 19, 20, 5, 1, 0, 0, 20, 25, 5, 17, 0, 0, 21,
		24, 3, 4, 2, 0, 22, 24, 3, 6, 3, 0, 23, 21, 1, 0, 0, 0, 23, 22, 1, 0, 0,
		0, 24, 27, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 3, 1,
		0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 29, 5, 2, 0, 0, 29, 30, 5, 17, 0, 0, 30,
		34, 3, 10, 5, 0, 31, 33, 3, 8, 4, 0, 32, 31, 1, 0, 0, 0, 33, 36, 1, 0,
		0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 5, 1, 0, 0, 0, 36, 34,
		1, 0, 0, 0, 37, 38, 5, 3, 0, 0, 38, 39, 5, 15, 0, 0, 39, 40, 5, 10, 0,
		0, 40, 41, 5, 17, 0, 0, 41, 46, 3, 10, 5, 0, 42, 43, 5, 2, 0, 0, 43, 45,
		5, 17, 0, 0, 44, 42, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0,
		46, 47, 1, 0, 0, 0, 47, 7, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 49, 50, 5, 2,
		0, 0, 50, 62, 5, 17, 0, 0, 51, 52, 5, 4, 0, 0, 52, 62, 5, 15, 0, 0, 53,
		54, 5, 5, 0, 0, 54, 62, 5, 15, 0, 0, 55, 56, 5, 3, 0, 0, 56, 62, 5, 15,
		0, 0, 57, 58, 5, 6, 0, 0, 58, 62, 5, 15, 0, 0, 59, 60, 5, 7, 0, 0, 60,
		62, 5, 17, 0, 0, 61, 49, 1, 0, 0, 0, 61, 51, 1, 0, 0, 0, 61, 53, 1, 0,
		0, 0, 61, 55, 1, 0, 0, 0, 61, 57, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 62, 63,
		1, 0, 0, 0, 63, 64, 7, 0, 0, 0, 64, 67, 5, 19, 0, 0, 65, 66, 5, 8, 0, 0,
		66, 68, 5, 17, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 71, 1,
		0, 0, 0, 69, 70, 5, 9, 0, 0, 70, 72, 7, 1, 0, 0, 71, 69, 1, 0, 0, 0, 71,
		72, 1, 0, 0, 0, 72, 9, 1, 0, 0, 0, 73, 74, 5, 11, 0, 0, 74, 77, 5, 19,
		0, 0, 75, 77, 5, 12, 0, 0, 76, 73, 1, 0, 0, 0, 76, 75, 1, 0, 0, 0, 77,
		11, 1, 0, 0, 0, 9, 15, 23, 25, 34, 46, 61, 67, 71, 76,
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
	BirdOSPFParserEOF         = antlr.TokenEOF
	BirdOSPFParserT__0        = 1
	BirdOSPFParserT__1        = 2
	BirdOSPFParserT__2        = 3
	BirdOSPFParserT__3        = 4
	BirdOSPFParserT__4        = 5
	BirdOSPFParserT__5        = 6
	BirdOSPFParserT__6        = 7
	BirdOSPFParserT__7        = 8
	BirdOSPFParserT__8        = 9
	BirdOSPFParserDR          = 10
	BirdOSPFParserDistance    = 11
	BirdOSPFParserUnreachable = 12
	BirdOSPFParserMetric      = 13
	BirdOSPFParserMetric2     = 14
	BirdOSPFParserPrefix      = 15
	BirdOSPFParserVERSION     = 16
	BirdOSPFParserIP          = 17
	BirdOSPFParserSTRING      = 18
	BirdOSPFParserINT         = 19
	BirdOSPFParserWS          = 20
	BirdOSPFParserHEX         = 21
)

// BirdOSPFParser rules.
const (
	BirdOSPFParserRULE_state       = 0
	BirdOSPFParserRULE_area        = 1
	BirdOSPFParserRULE_router      = 2
	BirdOSPFParserRULE_network     = 3
	BirdOSPFParserRULE_routerEntry = 4
	BirdOSPFParserRULE_distance    = 5
)

// IStateContext is an interface to support dynamic dispatch.
type IStateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
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

func (s *StateContext) EOF() antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserEOF, 0)
}

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
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == BirdOSPFParserT__0 {
		{
			p.SetState(12)
			p.Area()
		}

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(17)
		p.Match(BirdOSPFParserEOF)
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

// IAreaContext is an interface to support dynamic dispatch.
type IAreaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IP() antlr.TerminalNode
	AllRouter() []IRouterContext
	Router(i int) IRouterContext
	AllNetwork() []INetworkContext
	Network(i int) INetworkContext

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

func (s *AreaContext) AllNetwork() []INetworkContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INetworkContext); ok {
			len++
		}
	}

	tst := make([]INetworkContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INetworkContext); ok {
			tst[i] = t.(INetworkContext)
			i++
		}
	}

	return tst
}

func (s *AreaContext) Network(i int) INetworkContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INetworkContext); ok {
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

	return t.(INetworkContext)
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
		p.SetState(19)
		p.Match(BirdOSPFParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(20)
		p.Match(BirdOSPFParserIP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BirdOSPFParserT__1 || _la == BirdOSPFParserT__2 {
		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case BirdOSPFParserT__1:
			{
				p.SetState(21)
				p.Router()
			}

		case BirdOSPFParserT__2:
			{
				p.SetState(22)
				p.Network()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(27)
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
	Distance() IDistanceContext
	AllRouterEntry() []IRouterEntryContext
	RouterEntry(i int) IRouterEntryContext

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

func (s *RouterContext) Distance() IDistanceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDistanceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDistanceContext)
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
		p.SetState(28)
		p.Match(BirdOSPFParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(29)
		p.Match(BirdOSPFParserIP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(30)
		p.Distance()
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(31)
				p.RouterEntry()
			}

		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
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

// INetworkContext is an interface to support dynamic dispatch.
type INetworkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Prefix() antlr.TerminalNode
	DR() antlr.TerminalNode
	AllIP() []antlr.TerminalNode
	IP(i int) antlr.TerminalNode
	Distance() IDistanceContext

	// IsNetworkContext differentiates from other interfaces.
	IsNetworkContext()
}

type NetworkContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNetworkContext() *NetworkContext {
	var p = new(NetworkContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BirdOSPFParserRULE_network
	return p
}

func InitEmptyNetworkContext(p *NetworkContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BirdOSPFParserRULE_network
}

func (*NetworkContext) IsNetworkContext() {}

func NewNetworkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NetworkContext {
	var p = new(NetworkContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BirdOSPFParserRULE_network

	return p
}

func (s *NetworkContext) GetParser() antlr.Parser { return s.parser }

func (s *NetworkContext) Prefix() antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserPrefix, 0)
}

func (s *NetworkContext) DR() antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserDR, 0)
}

func (s *NetworkContext) AllIP() []antlr.TerminalNode {
	return s.GetTokens(BirdOSPFParserIP)
}

func (s *NetworkContext) IP(i int) antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserIP, i)
}

func (s *NetworkContext) Distance() IDistanceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDistanceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDistanceContext)
}

func (s *NetworkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NetworkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NetworkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BirdOSPFListener); ok {
		listenerT.EnterNetwork(s)
	}
}

func (s *NetworkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BirdOSPFListener); ok {
		listenerT.ExitNetwork(s)
	}
}

func (s *NetworkContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BirdOSPFVisitor:
		return t.VisitNetwork(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BirdOSPFParser) Network() (localctx INetworkContext) {
	localctx = NewNetworkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BirdOSPFParserRULE_network)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Match(BirdOSPFParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(38)
		p.Match(BirdOSPFParserPrefix)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Match(BirdOSPFParserDR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Match(BirdOSPFParserIP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.Distance()
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(42)
				p.Match(BirdOSPFParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(43)
				p.Match(BirdOSPFParserIP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
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
	AllINT() []antlr.TerminalNode
	INT(i int) antlr.TerminalNode
	Metric() antlr.TerminalNode
	Metric2() antlr.TerminalNode
	AllIP() []antlr.TerminalNode
	IP(i int) antlr.TerminalNode
	Prefix() antlr.TerminalNode
	HEX() antlr.TerminalNode

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

func (s *RouterEntryContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(BirdOSPFParserINT)
}

func (s *RouterEntryContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserINT, i)
}

func (s *RouterEntryContext) Metric() antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserMetric, 0)
}

func (s *RouterEntryContext) Metric2() antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserMetric2, 0)
}

func (s *RouterEntryContext) AllIP() []antlr.TerminalNode {
	return s.GetTokens(BirdOSPFParserIP)
}

func (s *RouterEntryContext) IP(i int) antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserIP, i)
}

func (s *RouterEntryContext) Prefix() antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserPrefix, 0)
}

func (s *RouterEntryContext) HEX() antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserHEX, 0)
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
	p.EnterRule(localctx, 8, BirdOSPFParserRULE_routerEntry)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BirdOSPFParserT__1:
		{
			p.SetState(49)
			p.Match(BirdOSPFParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(BirdOSPFParserIP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BirdOSPFParserT__3:
		{
			p.SetState(51)
			p.Match(BirdOSPFParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.Match(BirdOSPFParserPrefix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BirdOSPFParserT__4:
		{
			p.SetState(53)
			p.Match(BirdOSPFParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.Match(BirdOSPFParserPrefix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BirdOSPFParserT__2:
		{
			p.SetState(55)
			p.Match(BirdOSPFParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.Match(BirdOSPFParserPrefix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BirdOSPFParserT__5:
		{
			p.SetState(57)
			p.Match(BirdOSPFParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Match(BirdOSPFParserPrefix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BirdOSPFParserT__6:
		{
			p.SetState(59)
			p.Match(BirdOSPFParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
			p.Match(BirdOSPFParserIP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(63)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BirdOSPFParserMetric || _la == BirdOSPFParserMetric2) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(64)
		p.Match(BirdOSPFParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BirdOSPFParserT__7 {
		{
			p.SetState(65)
			p.Match(BirdOSPFParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Match(BirdOSPFParserIP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BirdOSPFParserT__8 {
		{
			p.SetState(69)
			p.Match(BirdOSPFParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			_la = p.GetTokenStream().LA(1)

			if !(_la == BirdOSPFParserINT || _la == BirdOSPFParserHEX) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
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

// IDistanceContext is an interface to support dynamic dispatch.
type IDistanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Distance() antlr.TerminalNode
	INT() antlr.TerminalNode
	Unreachable() antlr.TerminalNode

	// IsDistanceContext differentiates from other interfaces.
	IsDistanceContext()
}

type DistanceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDistanceContext() *DistanceContext {
	var p = new(DistanceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BirdOSPFParserRULE_distance
	return p
}

func InitEmptyDistanceContext(p *DistanceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BirdOSPFParserRULE_distance
}

func (*DistanceContext) IsDistanceContext() {}

func NewDistanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DistanceContext {
	var p = new(DistanceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BirdOSPFParserRULE_distance

	return p
}

func (s *DistanceContext) GetParser() antlr.Parser { return s.parser }

func (s *DistanceContext) Distance() antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserDistance, 0)
}

func (s *DistanceContext) INT() antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserINT, 0)
}

func (s *DistanceContext) Unreachable() antlr.TerminalNode {
	return s.GetToken(BirdOSPFParserUnreachable, 0)
}

func (s *DistanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DistanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BirdOSPFListener); ok {
		listenerT.EnterDistance(s)
	}
}

func (s *DistanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BirdOSPFListener); ok {
		listenerT.ExitDistance(s)
	}
}

func (s *DistanceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BirdOSPFVisitor:
		return t.VisitDistance(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BirdOSPFParser) Distance() (localctx IDistanceContext) {
	localctx = NewDistanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BirdOSPFParserRULE_distance)
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BirdOSPFParserDistance:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.Match(BirdOSPFParserDistance)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(BirdOSPFParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BirdOSPFParserUnreachable:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.Match(BirdOSPFParserUnreachable)
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
