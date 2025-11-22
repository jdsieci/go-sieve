package interp

import (
	"context"
	"fmt"
	"io"
	"mime"
	"strings"

	"github.com/emersion/go-message/mail"
)

type BodyTest struct {
	matcherTest
	Raw     bool
	Content string
}

func (b BodyTest) Check(_ context.Context, d *RuntimeData) (bool, error) {
	if _, ok := d.Msg.(ExtendedMessage); !ok {
		return false, fmt.Errorf("body: Message do not implement ExtendedMessage interface")
	}
	rawMsg, err := d.Msg.(ExtendedMessage).RawMessage()
	if err != nil {
		return false, err
	}
	msg, err := mail.CreateReader(rawMsg)
	if err != nil {
		return false, err
	}
	for {
		part, err := msg.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			return false, err
		}
		mediatype, _, err := mime.ParseMediaType(part.Header.Get("Content-Type"))
		if err != nil {
			return false, err
		}
		if !strings.HasPrefix(mediatype, b.Content) {
			continue
		}
		ret, err := b.tryMatchStream(d, part.Body)
		if err != nil {
			return false, err
		}
		if ret {
			return true, nil
		}
	}
	return false, nil
}
