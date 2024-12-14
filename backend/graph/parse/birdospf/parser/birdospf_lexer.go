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
		"", "'area'", "'distance'", "'router'", "'metric'", "'stubnet'", "'xnetwork'",
		"'external'", "'metric2'", "'xrouter'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "Prefix", "VERSION", "IP", "STRING",
		"INT", "WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"Prefix", "VERSION", "IP", "STRING", "INT", "WS", "COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 16, 202, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 4, 10, 114,
		8, 10, 11, 10, 12, 10, 115, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 4,
		10, 124, 8, 10, 11, 10, 12, 10, 125, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 4, 11, 138, 8, 11, 11, 11, 12, 11, 139,
		1, 11, 1, 11, 4, 11, 144, 8, 11, 11, 11, 12, 11, 145, 1, 11, 1, 11, 4,
		11, 150, 8, 11, 11, 11, 12, 11, 151, 1, 11, 1, 11, 4, 11, 156, 8, 11, 11,
		11, 12, 11, 157, 1, 12, 1, 12, 5, 12, 162, 8, 12, 10, 12, 12, 12, 165,
		9, 12, 1, 12, 1, 12, 1, 12, 5, 12, 170, 8, 12, 10, 12, 12, 12, 173, 9,
		12, 3, 12, 175, 8, 12, 1, 13, 4, 13, 178, 8, 13, 11, 13, 12, 13, 179, 1,
		14, 4, 14, 183, 8, 14, 11, 14, 12, 14, 184, 1, 14, 1, 14, 1, 15, 1, 15,
		5, 15, 191, 8, 15, 10, 15, 12, 15, 194, 9, 15, 1, 15, 3, 15, 197, 8, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 2, 163, 192, 0, 16, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 1, 0, 4, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 5, 0, 45,
		45, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 214,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 1, 33, 1, 0, 0, 0, 3, 38, 1, 0, 0, 0, 5, 47, 1, 0, 0, 0, 7, 54,
		1, 0, 0, 0, 9, 61, 1, 0, 0, 0, 11, 69, 1, 0, 0, 0, 13, 78, 1, 0, 0, 0,
		15, 87, 1, 0, 0, 0, 17, 95, 1, 0, 0, 0, 19, 103, 1, 0, 0, 0, 21, 107, 1,
		0, 0, 0, 23, 137, 1, 0, 0, 0, 25, 174, 1, 0, 0, 0, 27, 177, 1, 0, 0, 0,
		29, 182, 1, 0, 0, 0, 31, 188, 1, 0, 0, 0, 33, 34, 5, 97, 0, 0, 34, 35,
		5, 114, 0, 0, 35, 36, 5, 101, 0, 0, 36, 37, 5, 97, 0, 0, 37, 2, 1, 0, 0,
		0, 38, 39, 5, 100, 0, 0, 39, 40, 5, 105, 0, 0, 40, 41, 5, 115, 0, 0, 41,
		42, 5, 116, 0, 0, 42, 43, 5, 97, 0, 0, 43, 44, 5, 110, 0, 0, 44, 45, 5,
		99, 0, 0, 45, 46, 5, 101, 0, 0, 46, 4, 1, 0, 0, 0, 47, 48, 5, 114, 0, 0,
		48, 49, 5, 111, 0, 0, 49, 50, 5, 117, 0, 0, 50, 51, 5, 116, 0, 0, 51, 52,
		5, 101, 0, 0, 52, 53, 5, 114, 0, 0, 53, 6, 1, 0, 0, 0, 54, 55, 5, 109,
		0, 0, 55, 56, 5, 101, 0, 0, 56, 57, 5, 116, 0, 0, 57, 58, 5, 114, 0, 0,
		58, 59, 5, 105, 0, 0, 59, 60, 5, 99, 0, 0, 60, 8, 1, 0, 0, 0, 61, 62, 5,
		115, 0, 0, 62, 63, 5, 116, 0, 0, 63, 64, 5, 117, 0, 0, 64, 65, 5, 98, 0,
		0, 65, 66, 5, 110, 0, 0, 66, 67, 5, 101, 0, 0, 67, 68, 5, 116, 0, 0, 68,
		10, 1, 0, 0, 0, 69, 70, 5, 120, 0, 0, 70, 71, 5, 110, 0, 0, 71, 72, 5,
		101, 0, 0, 72, 73, 5, 116, 0, 0, 73, 74, 5, 119, 0, 0, 74, 75, 5, 111,
		0, 0, 75, 76, 5, 114, 0, 0, 76, 77, 5, 107, 0, 0, 77, 12, 1, 0, 0, 0, 78,
		79, 5, 101, 0, 0, 79, 80, 5, 120, 0, 0, 80, 81, 5, 116, 0, 0, 81, 82, 5,
		101, 0, 0, 82, 83, 5, 114, 0, 0, 83, 84, 5, 110, 0, 0, 84, 85, 5, 97, 0,
		0, 85, 86, 5, 108, 0, 0, 86, 14, 1, 0, 0, 0, 87, 88, 5, 109, 0, 0, 88,
		89, 5, 101, 0, 0, 89, 90, 5, 116, 0, 0, 90, 91, 5, 114, 0, 0, 91, 92, 5,
		105, 0, 0, 92, 93, 5, 99, 0, 0, 93, 94, 5, 50, 0, 0, 94, 16, 1, 0, 0, 0,
		95, 96, 5, 120, 0, 0, 96, 97, 5, 114, 0, 0, 97, 98, 5, 111, 0, 0, 98, 99,
		5, 117, 0, 0, 99, 100, 5, 116, 0, 0, 100, 101, 5, 101, 0, 0, 101, 102,
		5, 114, 0, 0, 102, 18, 1, 0, 0, 0, 103, 104, 3, 23, 11, 0, 104, 105, 5,
		47, 0, 0, 105, 106, 3, 27, 13, 0, 106, 20, 1, 0, 0, 0, 107, 108, 5, 66,
		0, 0, 108, 109, 5, 73, 0, 0, 109, 110, 5, 82, 0, 0, 110, 111, 5, 68, 0,
		0, 111, 113, 1, 0, 0, 0, 112, 114, 3, 29, 14, 0, 113, 112, 1, 0, 0, 0,
		114, 115, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116,
		117, 1, 0, 0, 0, 117, 118, 3, 27, 13, 0, 118, 119, 5, 46, 0, 0, 119, 120,
		3, 27, 13, 0, 120, 121, 5, 46, 0, 0, 121, 123, 3, 27, 13, 0, 122, 124,
		3, 29, 14, 0, 123, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 123, 1,
		0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128, 5, 114,
		0, 0, 128, 129, 5, 101, 0, 0, 129, 130, 5, 97, 0, 0, 130, 131, 5, 100,
		0, 0, 131, 132, 5, 121, 0, 0, 132, 133, 5, 46, 0, 0, 133, 134, 1, 0, 0,
		0, 134, 135, 6, 10, 0, 0, 135, 22, 1, 0, 0, 0, 136, 138, 7, 0, 0, 0, 137,
		136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140,
		1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 143, 5, 46, 0, 0, 142, 144, 7, 0,
		0, 0, 143, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0,
		145, 146, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 149, 5, 46, 0, 0, 148,
		150, 7, 0, 0, 0, 149, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 149,
		1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 155, 5, 46,
		0, 0, 154, 156, 7, 0, 0, 0, 155, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0,
		157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 24, 1, 0, 0, 0, 159, 163,
		5, 34, 0, 0, 160, 162, 9, 0, 0, 0, 161, 160, 1, 0, 0, 0, 162, 165, 1, 0,
		0, 0, 163, 164, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 164, 166, 1, 0, 0, 0,
		165, 163, 1, 0, 0, 0, 166, 175, 5, 34, 0, 0, 167, 171, 7, 1, 0, 0, 168,
		170, 7, 2, 0, 0, 169, 168, 1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 169,
		1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0,
		0, 0, 174, 159, 1, 0, 0, 0, 174, 167, 1, 0, 0, 0, 175, 26, 1, 0, 0, 0,
		176, 178, 7, 0, 0, 0, 177, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179,
		177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 28, 1, 0, 0, 0, 181, 183, 7,
		3, 0, 0, 182, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 182, 1, 0, 0,
		0, 184, 185, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 187, 6, 14, 0, 0, 187,
		30, 1, 0, 0, 0, 188, 192, 5, 35, 0, 0, 189, 191, 9, 0, 0, 0, 190, 189,
		1, 0, 0, 0, 191, 194, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 192, 190, 1, 0,
		0, 0, 193, 196, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 195, 197, 5, 13, 0, 0,
		196, 195, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198,
		199, 5, 10, 0, 0, 199, 200, 1, 0, 0, 0, 200, 201, 6, 15, 0, 0, 201, 32,
		1, 0, 0, 0, 14, 0, 115, 125, 139, 145, 151, 157, 163, 171, 174, 179, 184,
		192, 196, 1, 6, 0, 0,
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
	BirdOSPFLexerPrefix  = 10
	BirdOSPFLexerVERSION = 11
	BirdOSPFLexerIP      = 12
	BirdOSPFLexerSTRING  = 13
	BirdOSPFLexerINT     = 14
	BirdOSPFLexerWS      = 15
	BirdOSPFLexerCOMMENT = 16
)
