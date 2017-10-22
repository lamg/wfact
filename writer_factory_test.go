package wfact

import (
	fs "github.com/lamg/filesystem"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	cont  = []byte("Hola coco")
	fname = "coco.txt"
)

func TestWriterFact(t *testing.T) {
	fsm := fs.NewBufferFS()
	tss := []struct {
		name    string
		content string
	}{
		{"coco.txt", "HolaCoco"},
	}
	for _, j := range tss {
		tr := NewTruncater(j.name, fsm)
		tr.NextWriter()
		require.True(t, tr.Err() == nil)
		_, e := tr.Current().Write([]byte(j.content))
		require.NoError(t, e)
		bf, ok := fsm.GetBuffer(j.name)
		require.True(t, ok)
		require.True(t, bf.String() == j.content)
	}
}

func TestDateArchiver(t *testing.T) {
	fsm := fs.NewBufferFS()
	tss := []struct {
		name    string
		content string
	}{
		{"coco.txt", "hola coco"},
	}
	for _, j := range tss {
		dt := NewDateArchiver(j.name, fsm)
		dt.NextWriter()
		require.True(t, dt.Err() == nil)
		dt.Current().Write([]byte(j.content))
		bf, ok := fsm.GetBuffer(dt.cfn)
		require.True(t, ok)
		require.Equal(t, bf.String(), j.content)
	}
}
