// Code generated from BirdOSPF.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // BirdOSPF
import "github.com/antlr4-go/antlr/v4"

// BaseBirdOSPFListener is a complete listener for a parse tree produced by BirdOSPFParser.
type BaseBirdOSPFListener struct{}

var _ BirdOSPFListener = &BaseBirdOSPFListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBirdOSPFListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBirdOSPFListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBirdOSPFListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBirdOSPFListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterState is called when production state is entered.
func (s *BaseBirdOSPFListener) EnterState(ctx *StateContext) {}

// ExitState is called when production state is exited.
func (s *BaseBirdOSPFListener) ExitState(ctx *StateContext) {}

// EnterArea is called when production area is entered.
func (s *BaseBirdOSPFListener) EnterArea(ctx *AreaContext) {}

// ExitArea is called when production area is exited.
func (s *BaseBirdOSPFListener) ExitArea(ctx *AreaContext) {}

// EnterRouter is called when production router is entered.
func (s *BaseBirdOSPFListener) EnterRouter(ctx *RouterContext) {}

// ExitRouter is called when production router is exited.
func (s *BaseBirdOSPFListener) ExitRouter(ctx *RouterContext) {}

// EnterRouterEntry is called when production routerEntry is entered.
func (s *BaseBirdOSPFListener) EnterRouterEntry(ctx *RouterEntryContext) {}

// ExitRouterEntry is called when production routerEntry is exited.
func (s *BaseBirdOSPFListener) ExitRouterEntry(ctx *RouterEntryContext) {}
