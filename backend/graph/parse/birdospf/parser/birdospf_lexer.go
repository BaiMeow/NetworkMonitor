// Code generated from BirdOSPF.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type BirdOSPFLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var BirdOSPFLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func birdospflexerLexerInit() {
	staticData := &BirdOSPFLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'area'", "'router'", "'network'", "'stubnet'", "'xnetwork'", "'external'",
		"'xrouter'", "'via'", "'dr'", "'distance'", "'unreachable'", "'metric'",
		"'metric2'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "DR", "Distance", "Unreachable",
		"Metric", "Metric2", "Prefix", "VERSION", "IP", "STRING", "INT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "DR",
		"Distance", "Unreachable", "Metric", "Metric2", "Prefix", "VERSION",
		"IP", "STRING", "INT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 223, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 4, 14, 147, 8, 14, 11, 14, 12, 14, 148, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 3, 14, 156, 8, 14, 1, 14, 4, 14, 159, 8, 14, 11, 14, 12, 14, 160,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 4,
		15, 173, 8, 15, 11, 15, 12, 15, 174, 1, 15, 1, 15, 4, 15, 179, 8, 15, 11,
		15, 12, 15, 180, 1, 15, 1, 15, 4, 15, 185, 8, 15, 11, 15, 12, 15, 186,
		1, 15, 1, 15, 4, 15, 191, 8, 15, 11, 15, 12, 15, 192, 1, 16, 1, 16, 5,
		16, 197, 8, 16, 10, 16, 12, 16, 200, 9, 16, 1, 16, 1, 16, 1, 16, 5, 16,
		205, 8, 16, 10, 16, 12, 16, 208, 9, 16, 3, 16, 210, 8, 16, 1, 17, 4, 17,
		213, 8, 17, 11, 17, 12, 17, 214, 1, 18, 4, 18, 218, 8, 18, 11, 18, 12,
		18, 219, 1, 18, 1, 18, 1, 198, 0, 19, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15,
		31, 16, 33, 17, 35, 18, 37, 19, 1, 0, 4, 1, 0, 48, 57, 2, 0, 65, 90, 97,
		122, 5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13,
		32, 32, 234, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 1, 39, 1, 0, 0, 0, 3, 44, 1, 0, 0, 0, 5, 51, 1, 0, 0, 0, 7, 59, 1,
		0, 0, 0, 9, 67, 1, 0, 0, 0, 11, 76, 1, 0, 0, 0, 13, 85, 1, 0, 0, 0, 15,
		93, 1, 0, 0, 0, 17, 97, 1, 0, 0, 0, 19, 100, 1, 0, 0, 0, 21, 109, 1, 0,
		0, 0, 23, 121, 1, 0, 0, 0, 25, 128, 1, 0, 0, 0, 27, 136, 1, 0, 0, 0, 29,
		140, 1, 0, 0, 0, 31, 172, 1, 0, 0, 0, 33, 209, 1, 0, 0, 0, 35, 212, 1,
		0, 0, 0, 37, 217, 1, 0, 0, 0, 39, 40, 5, 97, 0, 0, 40, 41, 5, 114, 0, 0,
		41, 42, 5, 101, 0, 0, 42, 43, 5, 97, 0, 0, 43, 2, 1, 0, 0, 0, 44, 45, 5,
		114, 0, 0, 45, 46, 5, 111, 0, 0, 46, 47, 5, 117, 0, 0, 47, 48, 5, 116,
		0, 0, 48, 49, 5, 101, 0, 0, 49, 50, 5, 114, 0, 0, 50, 4, 1, 0, 0, 0, 51,
		52, 5, 110, 0, 0, 52, 53, 5, 101, 0, 0, 53, 54, 5, 116, 0, 0, 54, 55, 5,
		119, 0, 0, 55, 56, 5, 111, 0, 0, 56, 57, 5, 114, 0, 0, 57, 58, 5, 107,
		0, 0, 58, 6, 1, 0, 0, 0, 59, 60, 5, 115, 0, 0, 60, 61, 5, 116, 0, 0, 61,
		62, 5, 117, 0, 0, 62, 63, 5, 98, 0, 0, 63, 64, 5, 110, 0, 0, 64, 65, 5,
		101, 0, 0, 65, 66, 5, 116, 0, 0, 66, 8, 1, 0, 0, 0, 67, 68, 5, 120, 0,
		0, 68, 69, 5, 110, 0, 0, 69, 70, 5, 101, 0, 0, 70, 71, 5, 116, 0, 0, 71,
		72, 5, 119, 0, 0, 72, 73, 5, 111, 0, 0, 73, 74, 5, 114, 0, 0, 74, 75, 5,
		107, 0, 0, 75, 10, 1, 0, 0, 0, 76, 77, 5, 101, 0, 0, 77, 78, 5, 120, 0,
		0, 78, 79, 5, 116, 0, 0, 79, 80, 5, 101, 0, 0, 80, 81, 5, 114, 0, 0, 81,
		82, 5, 110, 0, 0, 82, 83, 5, 97, 0, 0, 83, 84, 5, 108, 0, 0, 84, 12, 1,
		0, 0, 0, 85, 86, 5, 120, 0, 0, 86, 87, 5, 114, 0, 0, 87, 88, 5, 111, 0,
		0, 88, 89, 5, 117, 0, 0, 89, 90, 5, 116, 0, 0, 90, 91, 5, 101, 0, 0, 91,
		92, 5, 114, 0, 0, 92, 14, 1, 0, 0, 0, 93, 94, 5, 118, 0, 0, 94, 95, 5,
		105, 0, 0, 95, 96, 5, 97, 0, 0, 96, 16, 1, 0, 0, 0, 97, 98, 5, 100, 0,
		0, 98, 99, 5, 114, 0, 0, 99, 18, 1, 0, 0, 0, 100, 101, 5, 100, 0, 0, 101,
		102, 5, 105, 0, 0, 102, 103, 5, 115, 0, 0, 103, 104, 5, 116, 0, 0, 104,
		105, 5, 97, 0, 0, 105, 106, 5, 110, 0, 0, 106, 107, 5, 99, 0, 0, 107, 108,
		5, 101, 0, 0, 108, 20, 1, 0, 0, 0, 109, 110, 5, 117, 0, 0, 110, 111, 5,
		110, 0, 0, 111, 112, 5, 114, 0, 0, 112, 113, 5, 101, 0, 0, 113, 114, 5,
		97, 0, 0, 114, 115, 5, 99, 0, 0, 115, 116, 5, 104, 0, 0, 116, 117, 5, 97,
		0, 0, 117, 118, 5, 98, 0, 0, 118, 119, 5, 108, 0, 0, 119, 120, 5, 101,
		0, 0, 120, 22, 1, 0, 0, 0, 121, 122, 5, 109, 0, 0, 122, 123, 5, 101, 0,
		0, 123, 124, 5, 116, 0, 0, 124, 125, 5, 114, 0, 0, 125, 126, 5, 105, 0,
		0, 126, 127, 5, 99, 0, 0, 127, 24, 1, 0, 0, 0, 128, 129, 5, 109, 0, 0,
		129, 130, 5, 101, 0, 0, 130, 131, 5, 116, 0, 0, 131, 132, 5, 114, 0, 0,
		132, 133, 5, 105, 0, 0, 133, 134, 5, 99, 0, 0, 134, 135, 5, 50, 0, 0, 135,
		26, 1, 0, 0, 0, 136, 137, 3, 31, 15, 0, 137, 138, 5, 47, 0, 0, 138, 139,
		3, 35, 17, 0, 139, 28, 1, 0, 0, 0, 140, 141, 5, 66, 0, 0, 141, 142, 5,
		73, 0, 0, 142, 143, 5, 82, 0, 0, 143, 144, 5, 68, 0, 0, 144, 146, 1, 0,
		0, 0, 145, 147, 3, 37, 18, 0, 146, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0,
		0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150,
		151, 3, 35, 17, 0, 151, 152, 5, 46, 0, 0, 152, 155, 3, 35, 17, 0, 153,
		154, 5, 46, 0, 0, 154, 156, 3, 35, 17, 0, 155, 153, 1, 0, 0, 0, 155, 156,
		1, 0, 0, 0, 156, 158, 1, 0, 0, 0, 157, 159, 3, 37, 18, 0, 158, 157, 1,
		0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0,
		0, 161, 162, 1, 0, 0, 0, 162, 163, 5, 114, 0, 0, 163, 164, 5, 101, 0, 0,
		164, 165, 5, 97, 0, 0, 165, 166, 5, 100, 0, 0, 166, 167, 5, 121, 0, 0,
		167, 168, 5, 46, 0, 0, 168, 169, 1, 0, 0, 0, 169, 170, 6, 14, 0, 0, 170,
		30, 1, 0, 0, 0, 171, 173, 7, 0, 0, 0, 172, 171, 1, 0, 0, 0, 173, 174, 1,
		0, 0, 0, 174, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 176, 1, 0, 0,
		0, 176, 178, 5, 46, 0, 0, 177, 179, 7, 0, 0, 0, 178, 177, 1, 0, 0, 0, 179,
		180, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 182,
		1, 0, 0, 0, 182, 184, 5, 46, 0, 0, 183, 185, 7, 0, 0, 0, 184, 183, 1, 0,
		0, 0, 185, 186, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0,
		187, 188, 1, 0, 0, 0, 188, 190, 5, 46, 0, 0, 189, 191, 7, 0, 0, 0, 190,
		189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 193,
		1, 0, 0, 0, 193, 32, 1, 0, 0, 0, 194, 198, 5, 34, 0, 0, 195, 197, 9, 0,
		0, 0, 196, 195, 1, 0, 0, 0, 197, 200, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0,
		198, 196, 1, 0, 0, 0, 199, 201, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 201,
		210, 5, 34, 0, 0, 202, 206, 7, 1, 0, 0, 203, 205, 7, 2, 0, 0, 204, 203,
		1, 0, 0, 0, 205, 208, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206, 207, 1, 0,
		0, 0, 207, 210, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 209, 194, 1, 0, 0, 0,
		209, 202, 1, 0, 0, 0, 210, 34, 1, 0, 0, 0, 211, 213, 7, 0, 0, 0, 212, 211,
		1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0,
		0, 0, 215, 36, 1, 0, 0, 0, 216, 218, 7, 3, 0, 0, 217, 216, 1, 0, 0, 0,
		218, 219, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220,
		221, 1, 0, 0, 0, 221, 222, 6, 18, 0, 0, 222, 38, 1, 0, 0, 0, 13, 0, 148,
		155, 160, 174, 180, 186, 192, 198, 206, 209, 214, 219, 1, 6, 0, 0,
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

// BirdOSPFLexerInit initializes any static state used to implement BirdOSPFLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewBirdOSPFLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func BirdOSPFLexerInit() {
	staticData := &BirdOSPFLexerLexerStaticData
	staticData.once.Do(birdospflexerLexerInit)
}

// NewBirdOSPFLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewBirdOSPFLexer(input antlr.CharStream) *BirdOSPFLexer {
	BirdOSPFLexerInit()
	l := new(BirdOSPFLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &BirdOSPFLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "BirdOSPF.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BirdOSPFLexer tokens.
const (
	BirdOSPFLexerT__0        = 1
	BirdOSPFLexerT__1        = 2
	BirdOSPFLexerT__2        = 3
	BirdOSPFLexerT__3        = 4
	BirdOSPFLexerT__4        = 5
	BirdOSPFLexerT__5        = 6
	BirdOSPFLexerT__6        = 7
	BirdOSPFLexerT__7        = 8
	BirdOSPFLexerDR          = 9
	BirdOSPFLexerDistance    = 10
	BirdOSPFLexerUnreachable = 11
	BirdOSPFLexerMetric      = 12
	BirdOSPFLexerMetric2     = 13
	BirdOSPFLexerPrefix      = 14
	BirdOSPFLexerVERSION     = 15
	BirdOSPFLexerIP          = 16
	BirdOSPFLexerSTRING      = 17
	BirdOSPFLexerINT         = 18
	BirdOSPFLexerWS          = 19
)
