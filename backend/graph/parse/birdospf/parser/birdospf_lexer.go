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
		"'xrouter'", "'via'", "'dr'", "'distance'", "'metric'", "'metric2'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "DR", "Distance", "Metric", "Metric2",
		"Prefix", "VERSION", "IP", "STRING", "INT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "DR",
		"Distance", "Metric", "Metric2", "Prefix", "VERSION", "IP", "STRING",
		"INT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 18, 209, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 4, 13, 133, 8, 13,
		11, 13, 12, 13, 134, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 142, 8,
		13, 1, 13, 4, 13, 145, 8, 13, 11, 13, 12, 13, 146, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 4, 14, 159, 8, 14, 11,
		14, 12, 14, 160, 1, 14, 1, 14, 4, 14, 165, 8, 14, 11, 14, 12, 14, 166,
		1, 14, 1, 14, 4, 14, 171, 8, 14, 11, 14, 12, 14, 172, 1, 14, 1, 14, 4,
		14, 177, 8, 14, 11, 14, 12, 14, 178, 1, 15, 1, 15, 5, 15, 183, 8, 15, 10,
		15, 12, 15, 186, 9, 15, 1, 15, 1, 15, 1, 15, 5, 15, 191, 8, 15, 10, 15,
		12, 15, 194, 9, 15, 3, 15, 196, 8, 15, 1, 16, 4, 16, 199, 8, 16, 11, 16,
		12, 16, 200, 1, 17, 4, 17, 204, 8, 17, 11, 17, 12, 17, 205, 1, 17, 1, 17,
		1, 184, 0, 18, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18,
		1, 0, 4, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 5, 0, 45, 45, 48, 57, 65,
		90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 220, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 1, 37, 1, 0, 0, 0, 3, 42, 1, 0, 0, 0, 5, 49,
		1, 0, 0, 0, 7, 57, 1, 0, 0, 0, 9, 65, 1, 0, 0, 0, 11, 74, 1, 0, 0, 0, 13,
		83, 1, 0, 0, 0, 15, 91, 1, 0, 0, 0, 17, 95, 1, 0, 0, 0, 19, 98, 1, 0, 0,
		0, 21, 107, 1, 0, 0, 0, 23, 114, 1, 0, 0, 0, 25, 122, 1, 0, 0, 0, 27, 126,
		1, 0, 0, 0, 29, 158, 1, 0, 0, 0, 31, 195, 1, 0, 0, 0, 33, 198, 1, 0, 0,
		0, 35, 203, 1, 0, 0, 0, 37, 38, 5, 97, 0, 0, 38, 39, 5, 114, 0, 0, 39,
		40, 5, 101, 0, 0, 40, 41, 5, 97, 0, 0, 41, 2, 1, 0, 0, 0, 42, 43, 5, 114,
		0, 0, 43, 44, 5, 111, 0, 0, 44, 45, 5, 117, 0, 0, 45, 46, 5, 116, 0, 0,
		46, 47, 5, 101, 0, 0, 47, 48, 5, 114, 0, 0, 48, 4, 1, 0, 0, 0, 49, 50,
		5, 110, 0, 0, 50, 51, 5, 101, 0, 0, 51, 52, 5, 116, 0, 0, 52, 53, 5, 119,
		0, 0, 53, 54, 5, 111, 0, 0, 54, 55, 5, 114, 0, 0, 55, 56, 5, 107, 0, 0,
		56, 6, 1, 0, 0, 0, 57, 58, 5, 115, 0, 0, 58, 59, 5, 116, 0, 0, 59, 60,
		5, 117, 0, 0, 60, 61, 5, 98, 0, 0, 61, 62, 5, 110, 0, 0, 62, 63, 5, 101,
		0, 0, 63, 64, 5, 116, 0, 0, 64, 8, 1, 0, 0, 0, 65, 66, 5, 120, 0, 0, 66,
		67, 5, 110, 0, 0, 67, 68, 5, 101, 0, 0, 68, 69, 5, 116, 0, 0, 69, 70, 5,
		119, 0, 0, 70, 71, 5, 111, 0, 0, 71, 72, 5, 114, 0, 0, 72, 73, 5, 107,
		0, 0, 73, 10, 1, 0, 0, 0, 74, 75, 5, 101, 0, 0, 75, 76, 5, 120, 0, 0, 76,
		77, 5, 116, 0, 0, 77, 78, 5, 101, 0, 0, 78, 79, 5, 114, 0, 0, 79, 80, 5,
		110, 0, 0, 80, 81, 5, 97, 0, 0, 81, 82, 5, 108, 0, 0, 82, 12, 1, 0, 0,
		0, 83, 84, 5, 120, 0, 0, 84, 85, 5, 114, 0, 0, 85, 86, 5, 111, 0, 0, 86,
		87, 5, 117, 0, 0, 87, 88, 5, 116, 0, 0, 88, 89, 5, 101, 0, 0, 89, 90, 5,
		114, 0, 0, 90, 14, 1, 0, 0, 0, 91, 92, 5, 118, 0, 0, 92, 93, 5, 105, 0,
		0, 93, 94, 5, 97, 0, 0, 94, 16, 1, 0, 0, 0, 95, 96, 5, 100, 0, 0, 96, 97,
		5, 114, 0, 0, 97, 18, 1, 0, 0, 0, 98, 99, 5, 100, 0, 0, 99, 100, 5, 105,
		0, 0, 100, 101, 5, 115, 0, 0, 101, 102, 5, 116, 0, 0, 102, 103, 5, 97,
		0, 0, 103, 104, 5, 110, 0, 0, 104, 105, 5, 99, 0, 0, 105, 106, 5, 101,
		0, 0, 106, 20, 1, 0, 0, 0, 107, 108, 5, 109, 0, 0, 108, 109, 5, 101, 0,
		0, 109, 110, 5, 116, 0, 0, 110, 111, 5, 114, 0, 0, 111, 112, 5, 105, 0,
		0, 112, 113, 5, 99, 0, 0, 113, 22, 1, 0, 0, 0, 114, 115, 5, 109, 0, 0,
		115, 116, 5, 101, 0, 0, 116, 117, 5, 116, 0, 0, 117, 118, 5, 114, 0, 0,
		118, 119, 5, 105, 0, 0, 119, 120, 5, 99, 0, 0, 120, 121, 5, 50, 0, 0, 121,
		24, 1, 0, 0, 0, 122, 123, 3, 29, 14, 0, 123, 124, 5, 47, 0, 0, 124, 125,
		3, 33, 16, 0, 125, 26, 1, 0, 0, 0, 126, 127, 5, 66, 0, 0, 127, 128, 5,
		73, 0, 0, 128, 129, 5, 82, 0, 0, 129, 130, 5, 68, 0, 0, 130, 132, 1, 0,
		0, 0, 131, 133, 3, 35, 17, 0, 132, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0,
		0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136,
		137, 3, 33, 16, 0, 137, 138, 5, 46, 0, 0, 138, 141, 3, 33, 16, 0, 139,
		140, 5, 46, 0, 0, 140, 142, 3, 33, 16, 0, 141, 139, 1, 0, 0, 0, 141, 142,
		1, 0, 0, 0, 142, 144, 1, 0, 0, 0, 143, 145, 3, 35, 17, 0, 144, 143, 1,
		0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0,
		0, 147, 148, 1, 0, 0, 0, 148, 149, 5, 114, 0, 0, 149, 150, 5, 101, 0, 0,
		150, 151, 5, 97, 0, 0, 151, 152, 5, 100, 0, 0, 152, 153, 5, 121, 0, 0,
		153, 154, 5, 46, 0, 0, 154, 155, 1, 0, 0, 0, 155, 156, 6, 13, 0, 0, 156,
		28, 1, 0, 0, 0, 157, 159, 7, 0, 0, 0, 158, 157, 1, 0, 0, 0, 159, 160, 1,
		0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 1, 0, 0,
		0, 162, 164, 5, 46, 0, 0, 163, 165, 7, 0, 0, 0, 164, 163, 1, 0, 0, 0, 165,
		166, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168,
		1, 0, 0, 0, 168, 170, 5, 46, 0, 0, 169, 171, 7, 0, 0, 0, 170, 169, 1, 0,
		0, 0, 171, 172, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0,
		173, 174, 1, 0, 0, 0, 174, 176, 5, 46, 0, 0, 175, 177, 7, 0, 0, 0, 176,
		175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178, 179,
		1, 0, 0, 0, 179, 30, 1, 0, 0, 0, 180, 184, 5, 34, 0, 0, 181, 183, 9, 0,
		0, 0, 182, 181, 1, 0, 0, 0, 183, 186, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0,
		184, 182, 1, 0, 0, 0, 185, 187, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 187,
		196, 5, 34, 0, 0, 188, 192, 7, 1, 0, 0, 189, 191, 7, 2, 0, 0, 190, 189,
		1, 0, 0, 0, 191, 194, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 193, 1, 0,
		0, 0, 193, 196, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 195, 180, 1, 0, 0, 0,
		195, 188, 1, 0, 0, 0, 196, 32, 1, 0, 0, 0, 197, 199, 7, 0, 0, 0, 198, 197,
		1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0,
		0, 0, 201, 34, 1, 0, 0, 0, 202, 204, 7, 3, 0, 0, 203, 202, 1, 0, 0, 0,
		204, 205, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206,
		207, 1, 0, 0, 0, 207, 208, 6, 17, 0, 0, 208, 36, 1, 0, 0, 0, 13, 0, 134,
		141, 146, 160, 166, 172, 178, 184, 192, 195, 200, 205, 1, 6, 0, 0,
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
	BirdOSPFLexerT__0     = 1
	BirdOSPFLexerT__1     = 2
	BirdOSPFLexerT__2     = 3
	BirdOSPFLexerT__3     = 4
	BirdOSPFLexerT__4     = 5
	BirdOSPFLexerT__5     = 6
	BirdOSPFLexerT__6     = 7
	BirdOSPFLexerT__7     = 8
	BirdOSPFLexerDR       = 9
	BirdOSPFLexerDistance = 10
	BirdOSPFLexerMetric   = 11
	BirdOSPFLexerMetric2  = 12
	BirdOSPFLexerPrefix   = 13
	BirdOSPFLexerVERSION  = 14
	BirdOSPFLexerIP       = 15
	BirdOSPFLexerSTRING   = 16
	BirdOSPFLexerINT      = 17
	BirdOSPFLexerWS       = 18
)
