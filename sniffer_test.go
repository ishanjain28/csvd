package csvd

import (
	"bytes"
	"encoding/csv"
	"strings"
	"testing"
)

const (
	csv1 = `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	csv2 = `first_name;last_name;username
"Rob";"Pike";rob
Ken;Thompson;ken
"Robert";"Griesemer";"gri"
`
	csv3 = `first_name$last_name$username
"Rob"$"Pike"$rob
Ken$Thompson$ken
"Robert"$"Griesemer"$"gri"
`
)

func TestAnalyse(t *testing.T) {
	r1 := csv.NewReader(strings.NewReader(csv1))
	s1 := defaultSniffer()
	s1.analyse(r1)

	r2 := csv.NewReader(strings.NewReader(csv2))
	s2 := defaultSniffer()
	s2.analyse(r2)

	for k, v := range s1.frequencyMap[','] {
		if k != 3 && v != 4 {
			t.Fail()
		}
	}

	for k, v := range s2.frequencyMap[';'] {
		if k != 3 && v != 4 {
			t.Fail()
		}
	}
}

func TestSniff(t *testing.T) {
	r1 := csv.NewReader(strings.NewReader(csv1))
	s1 := defaultSniffer()
	if s1.analyse(r1) != ',' {
		t.Fail()
	}

	r2 := csv.NewReader(strings.NewReader(csv2))
	s2 := defaultSniffer()
	if s2.analyse(r2) != ';' {
		t.Fail()
	}

	delimiter := DetectDelimiter(bytes.NewReader([]byte(csv1)))
	r3 := csv.NewReader(strings.NewReader(csv1))
	r3.LazyQuotes = true
	r3.Comma = delimiter

	data, _ := r3.ReadAll()
	if len(data) != 4 {
		t.Error(len(data))
	}

	// Custom sniffer.
	s := NewSniffer(20, '$')
	delimiter = DetectDelimiter(bytes.NewReader([]byte(csv3)), s)
	r4 := csv.NewReader(strings.NewReader(csv3))
	r4.LazyQuotes = true
	r4.Comma = delimiter
	data, _ = r4.ReadAll()
	if len(data) != 4 {
		t.Error(len(data))
	}
}

func TestIncrement(t *testing.T) {
	s := defaultSniffer()

	s.increment(',', 1)

	val, ok := s.frequencyMap[','][1]
	if !ok {
		t.Fail()
	}
	if val != 1 {
		t.Fail()
	}
}
