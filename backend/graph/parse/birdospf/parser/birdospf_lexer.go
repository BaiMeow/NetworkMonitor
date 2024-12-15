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
		"", "'area'", "'distance'", "'router'", "'stubnet'", "'xnetwork'", "'external'",
		"'xrouter'", "'metric'", "'metric2'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "Metric", "Metric2", "Prefix", "VERSION",
		"IP", "STRING", "INT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "Metric", "Metric2",
		"Prefix", "VERSION", "IP", "STRING", "INT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 15, 188, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 4, 10, 112, 8, 10, 11,
		10, 12, 10, 113, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 121, 8, 10,
		1, 10, 4, 10, 124, 8, 10, 11, 10, 12, 10, 125, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 4, 11, 138, 8, 11, 11, 11,
		12, 11, 139, 1, 11, 1, 11, 4, 11, 144, 8, 11, 11, 11, 12, 11, 145, 1, 11,
		1, 11, 4, 11, 150, 8, 11, 11, 11, 12, 11, 151, 1, 11, 1, 11, 4, 11, 156,
		8, 11, 11, 11, 12, 11, 157, 1, 12, 1, 12, 5, 12, 162, 8, 12, 10, 12, 12,
		12, 165, 9, 12, 1, 12, 1, 12, 1, 12, 5, 12, 170, 8, 12, 10, 12, 12, 12,
		173, 9, 12, 3, 12, 175, 8, 12, 1, 13, 4, 13, 178, 8, 13, 11, 13, 12, 13,
		179, 1, 14, 4, 14, 183, 8, 14, 11, 14, 12, 14, 184, 1, 14, 1, 14, 1, 163,
		0, 15, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 1, 0, 4, 1, 0, 48, 57, 2, 0, 65,
		90, 97, 122, 5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10,
		13, 13, 32, 32, 199, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 1, 31, 1, 0, 0, 0, 3, 36, 1, 0, 0, 0, 5, 45, 1, 0, 0, 0, 7,
		52, 1, 0, 0, 0, 9, 60, 1, 0, 0, 0, 11, 69, 1, 0, 0, 0, 13, 78, 1, 0, 0,
		0, 15, 86, 1, 0, 0, 0, 17, 93, 1, 0, 0, 0, 19, 101, 1, 0, 0, 0, 21, 105,
		1, 0, 0, 0, 23, 137, 1, 0, 0, 0, 25, 174, 1, 0, 0, 0, 27, 177, 1, 0, 0,
		0, 29, 182, 1, 0, 0, 0, 31, 32, 5, 97, 0, 0, 32, 33, 5, 114, 0, 0, 33,
		34, 5, 101, 0, 0, 34, 35, 5, 97, 0, 0, 35, 2, 1, 0, 0, 0, 36, 37, 5, 100,
		0, 0, 37, 38, 5, 105, 0, 0, 38, 39, 5, 115, 0, 0, 39, 40, 5, 116, 0, 0,
		40, 41, 5, 97, 0, 0, 41, 42, 5, 110, 0, 0, 42, 43, 5, 99, 0, 0, 43, 44,
		5, 101, 0, 0, 44, 4, 1, 0, 0, 0, 45, 46, 5, 114, 0, 0, 46, 47, 5, 111,
		0, 0, 47, 48, 5, 117, 0, 0, 48, 49, 5, 116, 0, 0, 49, 50, 5, 101, 0, 0,
		50, 51, 5, 114, 0, 0, 51, 6, 1, 0, 0, 0, 52, 53, 5, 115, 0, 0, 53, 54,
		5, 116, 0, 0, 54, 55, 5, 117, 0, 0, 55, 56, 5, 98, 0, 0, 56, 57, 5, 110,
		0, 0, 57, 58, 5, 101, 0, 0, 58, 59, 5, 116, 0, 0, 59, 8, 1, 0, 0, 0, 60,
		61, 5, 120, 0, 0, 61, 62, 5, 110, 0, 0, 62, 63, 5, 101, 0, 0, 63, 64, 5,
		116, 0, 0, 64, 65, 5, 119, 0, 0, 65, 66, 5, 111, 0, 0, 66, 67, 5, 114,
		0, 0, 67, 68, 5, 107, 0, 0, 68, 10, 1, 0, 0, 0, 69, 70, 5, 101, 0, 0, 70,
		71, 5, 120, 0, 0, 71, 72, 5, 116, 0, 0, 72, 73, 5, 101, 0, 0, 73, 74, 5,
		114, 0, 0, 74, 75, 5, 110, 0, 0, 75, 76, 5, 97, 0, 0, 76, 77, 5, 108, 0,
		0, 77, 12, 1, 0, 0, 0, 78, 79, 5, 120, 0, 0, 79, 80, 5, 114, 0, 0, 80,
		81, 5, 111, 0, 0, 81, 82, 5, 117, 0, 0, 82, 83, 5, 116, 0, 0, 83, 84, 5,
		101, 0, 0, 84, 85, 5, 114, 0, 0, 85, 14, 1, 0, 0, 0, 86, 87, 5, 109, 0,
		0, 87, 88, 5, 101, 0, 0, 88, 89, 5, 116, 0, 0, 89, 90, 5, 114, 0, 0, 90,
		91, 5, 105, 0, 0, 91, 92, 5, 99, 0, 0, 92, 16, 1, 0, 0, 0, 93, 94, 5, 109,
		0, 0, 94, 95, 5, 101, 0, 0, 95, 96, 5, 116, 0, 0, 96, 97, 5, 114, 0, 0,
		97, 98, 5, 105, 0, 0, 98, 99, 5, 99, 0, 0, 99, 100, 5, 50, 0, 0, 100, 18,
		1, 0, 0, 0, 101, 102, 3, 23, 11, 0, 102, 103, 5, 47, 0, 0, 103, 104, 3,
		27, 13, 0, 104, 20, 1, 0, 0, 0, 105, 106, 5, 66, 0, 0, 106, 107, 5, 73,
		0, 0, 107, 108, 5, 82, 0, 0, 108, 109, 5, 68, 0, 0, 109, 111, 1, 0, 0,
		0, 110, 112, 3, 29, 14, 0, 111, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0,
		113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115,
		116, 3, 27, 13, 0, 116, 117, 5, 46, 0, 0, 117, 120, 3, 27, 13, 0, 118,
		119, 5, 46, 0, 0, 119, 121, 3, 27, 13, 0, 120, 118, 1, 0, 0, 0, 120, 121,
		1, 0, 0, 0, 121, 123, 1, 0, 0, 0, 122, 124, 3, 29, 14, 0, 123, 122, 1,
		0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0,
		0, 126, 127, 1, 0, 0, 0, 127, 128, 5, 114, 0, 0, 128, 129, 5, 101, 0, 0,
		129, 130, 5, 97, 0, 0, 130, 131, 5, 100, 0, 0, 131, 132, 5, 121, 0, 0,
		132, 133, 5, 46, 0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 6, 10, 0, 0, 135,
		22, 1, 0, 0, 0, 136, 138, 7, 0, 0, 0, 137, 136, 1, 0, 0, 0, 138, 139, 1,
		0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 141, 1, 0, 0,
		0, 141, 143, 5, 46, 0, 0, 142, 144, 7, 0, 0, 0, 143, 142, 1, 0, 0, 0, 144,
		145, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147,
		1, 0, 0, 0, 147, 149, 5, 46, 0, 0, 148, 150, 7, 0, 0, 0, 149, 148, 1, 0,
		0, 0, 150, 151, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0,
		152, 153, 1, 0, 0, 0, 153, 155, 5, 46, 0, 0, 154, 156, 7, 0, 0, 0, 155,
		154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158,
		1, 0, 0, 0, 158, 24, 1, 0, 0, 0, 159, 163, 5, 34, 0, 0, 160, 162, 9, 0,
		0, 0, 161, 160, 1, 0, 0, 0, 162, 165, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0,
		163, 161, 1, 0, 0, 0, 164, 166, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166,
		175, 5, 34, 0, 0, 167, 171, 7, 1, 0, 0, 168, 170, 7, 2, 0, 0, 169, 168,
		1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 171, 172, 1, 0,
		0, 0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 174, 159, 1, 0, 0, 0,
		174, 167, 1, 0, 0, 0, 175, 26, 1, 0, 0, 0, 176, 178, 7, 0, 0, 0, 177, 176,
		1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0,
		0, 0, 180, 28, 1, 0, 0, 0, 181, 183, 7, 3, 0, 0, 182, 181, 1, 0, 0, 0,
		183, 184, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185,
		186, 1, 0, 0, 0, 186, 187, 6, 14, 0, 0, 187, 30, 1, 0, 0, 0, 13, 0, 113,
		120, 125, 139, 145, 151, 157, 163, 171, 174, 179, 184, 1, 6, 0, 0,
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
	BirdOSPFLexerMetric  = 8
	BirdOSPFLexerMetric2 = 9
	BirdOSPFLexerPrefix  = 10
	BirdOSPFLexerVERSION = 11
	BirdOSPFLexerIP      = 12
	BirdOSPFLexerSTRING  = 13
	BirdOSPFLexerINT     = 14
	BirdOSPFLexerWS      = 15
)
