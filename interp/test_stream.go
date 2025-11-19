package interp

import (
	"bufio"
	"io"
	"strconv"
	"unicode"
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
			panic("testStream should not be used with MatchCount")
		}
	case ComparatorASCIINumeric:
		switch match {
		case MatchContains:
			return false, nil, ErrComparatorMatchUnsupported
		case MatchIs:
			lhsNum, err := streamNumericValue(stream)
			if err != nil {
				return false, nil, err
			}
			rhsNum := numericValue(key)
			return RelEqual.CompareNumericValue(lhsNum, rhsNum), nil, nil
		case MatchMatches:
			return false, nil, ErrComparatorMatchUnsupported
		case MatchValue:
			lhsNum, err := streamNumericValue(stream)
			if err != nil {
				return false, nil, err
			}
			rhsNum := numericValue(key)
			return RelEqual.CompareNumericValue(lhsNum, rhsNum), nil, nil
		case MatchCount:
			panic("testStream should not be used with MatchCount")
		}
	case ComparatorASCIICaseMap:
		switch match {
		case MatchContains:
		case MatchIs:
		case MatchMatches:
		case MatchValue:
		case MatchCount:
			panic("testStream should not be used with MatchCount")
		}
	case ComparatorUnicodeCaseMap:
		switch match {
		case MatchContains:
		case MatchIs:
		case MatchMatches:
		case MatchValue:
		case MatchCount:
			panic("testStream should not be used with MatchCount")
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

func streamNumericValue(stream io.Reader) (*uint64, error) {
	bufStream := bufio.NewReader(stream)
	runes := []rune{}
	for {
		r, _, err := bufStream.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if !unicode.IsDigit(r) {
			break
		}
		runes = append(runes, r)
	}
	digit, err := strconv.ParseUint(string(runes), 10, 64)
	if err != nil {
		return nil, nil
	}
	return &digit, nil
}
