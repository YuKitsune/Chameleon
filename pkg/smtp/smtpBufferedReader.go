package smtp

import (
	"bufio"
	"io"
)

// This is a bufio.Reader what will use our adjustable limit reader
// We 'extend' buffio to have the limited reader feature
type smtpBufferedReader struct {
	*bufio.Reader
	alr *adjustableLimitedReader
}

// Delegate to the adjustable limited reader
func (sbr *smtpBufferedReader) setLimit(n int64) {
	sbr.alr.setLimit(n)
}

// Set a new reader & use it to reset the underlying reader
func (sbr *smtpBufferedReader) Reset(r io.Reader) {
	sbr.alr = newAdjustableLimitedReader(r, CommandLineMaxLength)
	sbr.Reader.Reset(sbr.alr)
}

// Allocate a new SMTPBufferedReader
func newSMTPBufferedReader(rd io.Reader) *smtpBufferedReader {
	alr := newAdjustableLimitedReader(rd, CommandLineMaxLength)
	s := &smtpBufferedReader{bufio.NewReader(alr), alr}
	return s
}
