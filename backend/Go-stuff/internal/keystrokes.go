package internal

import (
	hook "github.com/robotn/gohook"
	"strings"
	"sync"
)

const maxCharLimit = 500

type RollingBuffer struct {
	mu       sync.Mutex
	buf      strings.Builder
	disabled bool
}

func (r *RollingBuffer) append(ch rune) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.disabled {
		return
	}

	r.buf.WriteRune(ch)

	if r.buf.Len() > maxCharLimit {
		s := r.buf.String()
		r.buf.Reset()
		r.buf.WriteString(s[len(s)-maxCharLimit:])
	}
}

func (r *RollingBuffer) Reset(disabled bool) {}

func (r *RollingBuffer) SnapShot() string {
	return "temp"
}

func StartCapture(buffer *RollingBuffer) {
	evchan := hook.Start()
	defer hook.End()

	for ev := range evchan {
		if ev.Kind == hook.KeyDown && ev.Keychar != 0 { // not like nav stuff, tab arrows etc
			buffer.append(rune(ev.Keychar))
		}
	}
}

//ts pmo
