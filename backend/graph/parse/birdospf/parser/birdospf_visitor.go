// Code generated from BirdOSPF.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // BirdOSPF
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by BirdOSPFParser.
type BirdOSPFVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BirdOSPFParser#state.
	VisitState(ctx *StateContext) interface{}

	// Visit a parse tree produced by BirdOSPFParser#area.
	VisitArea(ctx *AreaContext) interface{}

	// Visit a parse tree produced by BirdOSPFParser#router.
	VisitRouter(ctx *RouterContext) interface{}

	// Visit a parse tree produced by BirdOSPFParser#routerEntry.
	VisitRouterEntry(ctx *RouterEntryContext) interface{}
}
