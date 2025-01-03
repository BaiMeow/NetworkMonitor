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
		"", "'area'", "'distance'", "'router'", "'stubnet'", "'xnetwork'", "'network'",
		"'external'", "'xrouter'", "'via'", "'metric'", "'metric2'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "Metric", "Metric2", "Prefix",
		"VERSION", "IP", "STRING", "INT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"Metric", "Metric2", "Prefix", "VERSION", "IP", "STRING", "INT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 204, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 4, 12, 128, 8, 12, 11, 12, 12, 12, 129, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 3, 12, 137, 8, 12, 1, 12, 4, 12, 140, 8, 12, 11, 12,
		12, 12, 141, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 13, 4, 13, 154, 8, 13, 11, 13, 12, 13, 155, 1, 13, 1, 13, 4, 13,
		160, 8, 13, 11, 13, 12, 13, 161, 1, 13, 1, 13, 4, 13, 166, 8, 13, 11, 13,
		12, 13, 167, 1, 13, 1, 13, 4, 13, 172, 8, 13, 11, 13, 12, 13, 173, 1, 14,
		1, 14, 5, 14, 178, 8, 14, 10, 14, 12, 14, 181, 9, 14, 1, 14, 1, 14, 1,
		14, 5, 14, 186, 8, 14, 10, 14, 12, 14, 189, 9, 14, 3, 14, 191, 8, 14, 1,
		15, 4, 15, 194, 8, 15, 11, 15, 12, 15, 195, 1, 16, 4, 16, 199, 8, 16, 11,
		16, 12, 16, 200, 1, 16, 1, 16, 1, 179, 0, 17, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 1, 0, 4, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 5,
		0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32,
		215, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 1, 35, 1, 0, 0, 0, 3, 40, 1, 0, 0, 0,
		5, 49, 1, 0, 0, 0, 7, 56, 1, 0, 0, 0, 9, 64, 1, 0, 0, 0, 11, 73, 1, 0,
		0, 0, 13, 81, 1, 0, 0, 0, 15, 90, 1, 0, 0, 0, 17, 98, 1, 0, 0, 0, 19, 102,
		1, 0, 0, 0, 21, 109, 1, 0, 0, 0, 23, 117, 1, 0, 0, 0, 25, 121, 1, 0, 0,
		0, 27, 153, 1, 0, 0, 0, 29, 190, 1, 0, 0, 0, 31, 193, 1, 0, 0, 0, 33, 198,
		1, 0, 0, 0, 35, 36, 5, 97, 0, 0, 36, 37, 5, 114, 0, 0, 37, 38, 5, 101,
		0, 0, 38, 39, 5, 97, 0, 0, 39, 2, 1, 0, 0, 0, 40, 41, 5, 100, 0, 0, 41,
		42, 5, 105, 0, 0, 42, 43, 5, 115, 0, 0, 43, 44, 5, 116, 0, 0, 44, 45, 5,
		97, 0, 0, 45, 46, 5, 110, 0, 0, 46, 47, 5, 99, 0, 0, 47, 48, 5, 101, 0,
		0, 48, 4, 1, 0, 0, 0, 49, 50, 5, 114, 0, 0, 50, 51, 5, 111, 0, 0, 51, 52,
		5, 117, 0, 0, 52, 53, 5, 116, 0, 0, 53, 54, 5, 101, 0, 0, 54, 55, 5, 114,
		0, 0, 55, 6, 1, 0, 0, 0, 56, 57, 5, 115, 0, 0, 57, 58, 5, 116, 0, 0, 58,
		59, 5, 117, 0, 0, 59, 60, 5, 98, 0, 0, 60, 61, 5, 110, 0, 0, 61, 62, 5,
		101, 0, 0, 62, 63, 5, 116, 0, 0, 63, 8, 1, 0, 0, 0, 64, 65, 5, 120, 0,
		0, 65, 66, 5, 110, 0, 0, 66, 67, 5, 101, 0, 0, 67, 68, 5, 116, 0, 0, 68,
		69, 5, 119, 0, 0, 69, 70, 5, 111, 0, 0, 70, 71, 5, 114, 0, 0, 71, 72, 5,
		107, 0, 0, 72, 10, 1, 0, 0, 0, 73, 74, 5, 110, 0, 0, 74, 75, 5, 101, 0,
		0, 75, 76, 5, 116, 0, 0, 76, 77, 5, 119, 0, 0, 77, 78, 5, 111, 0, 0, 78,
		79, 5, 114, 0, 0, 79, 80, 5, 107, 0, 0, 80, 12, 1, 0, 0, 0, 81, 82, 5,
		101, 0, 0, 82, 83, 5, 120, 0, 0, 83, 84, 5, 116, 0, 0, 84, 85, 5, 101,
		0, 0, 85, 86, 5, 114, 0, 0, 86, 87, 5, 110, 0, 0, 87, 88, 5, 97, 0, 0,
		88, 89, 5, 108, 0, 0, 89, 14, 1, 0, 0, 0, 90, 91, 5, 120, 0, 0, 91, 92,
		5, 114, 0, 0, 92, 93, 5, 111, 0, 0, 93, 94, 5, 117, 0, 0, 94, 95, 5, 116,
		0, 0, 95, 96, 5, 101, 0, 0, 96, 97, 5, 114, 0, 0, 97, 16, 1, 0, 0, 0, 98,
		99, 5, 118, 0, 0, 99, 100, 5, 105, 0, 0, 100, 101, 5, 97, 0, 0, 101, 18,
		1, 0, 0, 0, 102, 103, 5, 109, 0, 0, 103, 104, 5, 101, 0, 0, 104, 105, 5,
		116, 0, 0, 105, 106, 5, 114, 0, 0, 106, 107, 5, 105, 0, 0, 107, 108, 5,
		99, 0, 0, 108, 20, 1, 0, 0, 0, 109, 110, 5, 109, 0, 0, 110, 111, 5, 101,
		0, 0, 111, 112, 5, 116, 0, 0, 112, 113, 5, 114, 0, 0, 113, 114, 5, 105,
		0, 0, 114, 115, 5, 99, 0, 0, 115, 116, 5, 50, 0, 0, 116, 22, 1, 0, 0, 0,
		117, 118, 3, 27, 13, 0, 118, 119, 5, 47, 0, 0, 119, 120, 3, 31, 15, 0,
		120, 24, 1, 0, 0, 0, 121, 122, 5, 66, 0, 0, 122, 123, 5, 73, 0, 0, 123,
		124, 5, 82, 0, 0, 124, 125, 5, 68, 0, 0, 125, 127, 1, 0, 0, 0, 126, 128,
		3, 33, 16, 0, 127, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 127, 1,
		0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 132, 3, 31, 15,
		0, 132, 133, 5, 46, 0, 0, 133, 136, 3, 31, 15, 0, 134, 135, 5, 46, 0, 0,
		135, 137, 3, 31, 15, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137,
		139, 1, 0, 0, 0, 138, 140, 3, 33, 16, 0, 139, 138, 1, 0, 0, 0, 140, 141,
		1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 1, 0,
		0, 0, 143, 144, 5, 114, 0, 0, 144, 145, 5, 101, 0, 0, 145, 146, 5, 97,
		0, 0, 146, 147, 5, 100, 0, 0, 147, 148, 5, 121, 0, 0, 148, 149, 5, 46,
		0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 6, 12, 0, 0, 151, 26, 1, 0, 0, 0,
		152, 154, 7, 0, 0, 0, 153, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155,
		153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 159,
		5, 46, 0, 0, 158, 160, 7, 0, 0, 0, 159, 158, 1, 0, 0, 0, 160, 161, 1, 0,
		0, 0, 161, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0,
		163, 165, 5, 46, 0, 0, 164, 166, 7, 0, 0, 0, 165, 164, 1, 0, 0, 0, 166,
		167, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 169,
		1, 0, 0, 0, 169, 171, 5, 46, 0, 0, 170, 172, 7, 0, 0, 0, 171, 170, 1, 0,
		0, 0, 172, 173, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0,
		174, 28, 1, 0, 0, 0, 175, 179, 5, 34, 0, 0, 176, 178, 9, 0, 0, 0, 177,
		176, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 179, 177,
		1, 0, 0, 0, 180, 182, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 182, 191, 5, 34,
		0, 0, 183, 187, 7, 1, 0, 0, 184, 186, 7, 2, 0, 0, 185, 184, 1, 0, 0, 0,
		186, 189, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188,
		191, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 190, 175, 1, 0, 0, 0, 190, 183,
		1, 0, 0, 0, 191, 30, 1, 0, 0, 0, 192, 194, 7, 0, 0, 0, 193, 192, 1, 0,
		0, 0, 194, 195, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0,
		196, 32, 1, 0, 0, 0, 197, 199, 7, 3, 0, 0, 198, 197, 1, 0, 0, 0, 199, 200,
		1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 202, 1, 0,
		0, 0, 202, 203, 6, 16, 0, 0, 203, 34, 1, 0, 0, 0, 13, 0, 129, 136, 141,
		155, 161, 167, 173, 179, 187, 190, 195, 200, 1, 6, 0, 0,
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
	BirdOSPFLexerT__0    = 1
	BirdOSPFLexerT__1    = 2
	BirdOSPFLexerT__2    = 3
	BirdOSPFLexerT__3    = 4
	BirdOSPFLexerT__4    = 5
	BirdOSPFLexerT__5    = 6
	BirdOSPFLexerT__6    = 7
	BirdOSPFLexerT__7    = 8
	BirdOSPFLexerT__8    = 9
	BirdOSPFLexerMetric  = 10
	BirdOSPFLexerMetric2 = 11
	BirdOSPFLexerPrefix  = 12
	BirdOSPFLexerVERSION = 13
	BirdOSPFLexerIP      = 14
	BirdOSPFLexerSTRING  = 15
	BirdOSPFLexerINT     = 16
	BirdOSPFLexerWS      = 17
)
