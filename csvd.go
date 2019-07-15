package csvd

import (
	"bytes"
	"encoding/csv"
)

// DetectDelimiter returns the delimiter. As a second argument you
// can pass in a *Sniffer instance to use instead of the defaults, this can provide a different
// set of delimiters to look for.
func DetectDelimiter(r *bytes.Reader, s ...*Sniffer) rune {
	var sniffer *Sniffer
	if len(s) != 0 {
		sniffer = s[0]
	} else {
		sniffer = defaultSniffer()
	}
	csvReader := csv.NewReader(r)
	csvReader.LazyQuotes = true
	// The delimiter to start with should be the one chosen by the user
	csvReader.Comma = sniffer.delimiter

	sniffer.analyse(csvReader)
	r.Seek(0, 0)

	return sniffer.delimiter
}
