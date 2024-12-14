// Code generated from BirdOSPF.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // BirdOSPF
import "github.com/antlr4-go/antlr/v4"

// BirdOSPFListener is a complete listener for a parse tree produced by BirdOSPFParser.
type BirdOSPFListener interface {
	antlr.ParseTreeListener

	// EnterState is called when entering the state production.
	EnterState(c *StateContext)

	// EnterArea is called when entering the area production.
	EnterArea(c *AreaContext)

	// EnterRouter is called when entering the router production.
	EnterRouter(c *RouterContext)

	// EnterRouterEntry is called when entering the routerEntry production.
	EnterRouterEntry(c *RouterEntryContext)

	// ExitState is called when exiting the state production.
	ExitState(c *StateContext)

	// ExitArea is called when exiting the area production.
	ExitArea(c *AreaContext)

	// ExitRouter is called when exiting the router production.
	ExitRouter(c *RouterContext)

	// ExitRouterEntry is called when exiting the routerEntry production.
	ExitRouterEntry(c *RouterEntryContext)
}
