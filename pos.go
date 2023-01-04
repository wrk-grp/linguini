package linguini

import (
	"io"

	"github.com/jdkato/prose/v2"
	"github.com/wrk-grp/errnie"
)

/*
POS is a "part of speech" tagger to help structure text
content which ingresses mostly unstructured.
*/
type POS struct {
}

func NewPOS() *POS {
	return &POS{}
}

func (pos *POS) Read(p []byte) (n int, err error) {
	doc, _ := prose.NewDocument(string(p))

	for _, ent := range doc.Entities() {
		p = append(p, ent.Text...)
		errnie.Debugs(ent.Label, "->", ent.Text)
	}

	return len(p), io.EOF
}
