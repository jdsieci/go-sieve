package interp

import (
	"bufio"
	"io"
)

func testStream(comparator Comparator, match Match, rel Relational, stream io.Reader, key string) (bool, []string, error) {
	switch comparator {
	case ComparatorOctet:
		switch match {
		case MatchContains:
			ret, err := streamContains(stream, []byte(key))
			return ret, nil, err
		case MatchIs:
		case MatchMatches:
		case MatchValue:
		case MatchCount:
		}
	case ComparatorASCIINumeric:
		switch match {
		case MatchContains:
		case MatchIs:
		case MatchMatches:
		case MatchValue:
		case MatchCount:
		}
	case ComparatorASCIICaseMap:
		switch match {
		case MatchContains:
		case MatchIs:
		case MatchMatches:
		case MatchValue:
		case MatchCount:
		}
	case ComparatorUnicodeCaseMap:
		switch match {
		case MatchContains:
		case MatchIs:
		case MatchMatches:
		case MatchValue:
		case MatchCount:
		}
	}
	return false, nil, nil
}

func streamContains(stream io.Reader, key []byte) (bool, error) {
	keyLenght := len(key)
	bufStream := bufio.NewReaderSize(stream, keyLenght)
	index := 0
	// Always contains empty key
	if keyLenght == 0 {
		return true, nil
	}
	for {
		buf, err := bufStream.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			return false, err
		}
		if buf == key[index] {
			index++
		} else {
			index = 0
		}
		if index == keyLenght {
			return true, nil
		}
	}
	return false, nil
}
