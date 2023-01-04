package linguini

import (
	"io"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"github.com/wrk-grp/errnie"
)

type Time struct {
	parser *when.Parser
}

func NewTime() *Time {
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	return &Time{w}
}

func (t *Time) Read(p []byte) (n int, err error) {
	var res *when.Result

	res, err = t.parser.Parse(string(p), time.Now())
	errnie.Handles(err)

	copy(p, []byte(
		res.Time.Format(time.RFC3339Nano),
	))

	return len(p), io.EOF
}
