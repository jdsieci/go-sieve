package interp

import (
	"fmt"
	"mime"

	"github.com/foxcpp/go-sieve/parser"
)

func loadBodyTest(s *Script, test parser.Test) (Test, error) {
	if !s.RequiresExtension("body") {
		return nil, fmt.Errorf("missing require 'body'")
	}
	loaded := BodyTest{
		matcherTest: newMatcherTest(),
	}
	var key []string
	var content string
	err := LoadSpec(s, loaded.addSpecTags(&Spec{
		Tags: map[string]SpecTag{
			"raw": {
				MatchBool: func() {
					loaded.Raw = true
				},
			},
			"content": {
				NeedsValue: true,
				MatchStr: func(val []string) {
					content = val[0]
				},
			},
			"text": {
				MatchBool: func() {
					loaded.Text = true
				},
			},
		},
		Pos: []SpecPosArg{
			{
				MatchStr: func(val []string) {
					key = val
				},
				MinStrCount: 1,
			},
		},
	}), test.Position, test.Args, test.Tests, nil)
	if err != nil {
		return nil, err
	}
	if content != "" {
		if mediatype, _, err := mime.ParseMediaType(content); err != nil {
			return nil, fmt.Errorf("loadBodyTest: wrong media type '%s': %w", content, err)
		} else {
			loaded.Content = mediatype
		}
	}
	loaded.setKey(s, key)
	return loaded, nil
}
