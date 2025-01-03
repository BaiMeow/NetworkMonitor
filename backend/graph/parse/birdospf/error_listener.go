package birdospf

import "github.com/antlr4-go/antlr/v4"

type errorMsg struct {
	line, col int
	msg       string
}

type errorListener struct {
	errs []errorMsg
	*antlr.DefaultErrorListener
}

var _ antlr.ErrorListener = (*errorListener)(nil)

func (l *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line int, col int, msg string, e antlr.RecognitionException) {
	l.errs = append(l.errs, errorMsg{
		line: line,
		col:  col,
		msg:  msg,
	})
}
