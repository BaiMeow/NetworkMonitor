// Code generated from BirdOSPF.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // BirdOSPF
import "github.com/antlr4-go/antlr/v4"

type BaseBirdOSPFVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseBirdOSPFVisitor) VisitState(ctx *StateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBirdOSPFVisitor) VisitArea(ctx *AreaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBirdOSPFVisitor) VisitRouter(ctx *RouterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBirdOSPFVisitor) VisitRouterEntry(ctx *RouterEntryContext) interface{} {
	return v.VisitChildren(ctx)
}
