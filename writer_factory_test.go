package wfact

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

var (
	cont  = []byte("Hola coco")
	fname = "coco.txt"
)

func TestWriterFact(t *testing.T) {
	tr := NewTruncater(fname)
	tr.NextWriter()
	require.True(t, tr.Err() == nil)
	tr.Current().Write(cont)
	var e error
	var bs []byte
	bs, e = ioutil.ReadFile(fname)
	require.NoError(t, e)
	require.Equal(t, bs, cont)
	tr.NextWriter()
	require.True(t, tr.Err() == nil)
	bs, e = ioutil.ReadFile(fname + "~")
	require.NoError(t, e)
	require.Equal(t, bs, cont)
}

func TestDateArchiver(t *testing.T) {
	dt := NewDateArchiver(fname)
	dt.NextWriter()
	require.True(t, dt.Err() == nil)
	dt.Current().Write(cont)
	var e error
	var bs []byte
	bs, e = ioutil.ReadFile(dt.cfn)
	require.NoError(t, e)
	require.Equal(t, bs, cont)
}
