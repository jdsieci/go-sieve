package interp

import "context"

type BodyTest struct {
	matcherTest
	Raw     bool
	Text    bool
	Content string
}

func (b BodyTest) Check(_ context.Context, d *RuntimeData) (bool, error) {
	return false, nil
}
