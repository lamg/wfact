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
	tr := new(Truncater)
	tr.Init(fname)
	tr.NextWriter()
	require.NoError(t, tr.Err())
	tr.Current().Write(cont)
	var e error
	var bs []byte
	bs, e = ioutil.ReadFile(fname)
	require.NoError(t, e)
	require.Equal(t, bs, cont)
	tr.NextWriter()
	require.NoError(t, e)
	bs, e = ioutil.ReadFile(fname + "~")
	require.NoError(t, e)
	require.Equal(t, bs, cont)
}

func TestDateArchiver(t *testing.T) {
	dt := new(DateArchiver)
	dt.Init(fname)
	dt.NextWriter()
	require.NoError(t, dt.Err())
	dt.Current().Write(cont)
	var e error
	var bs []byte
	bs, e = ioutil.ReadFile(dt.cfn)
	require.NoError(t, e)
	require.Equal(t, bs, cont)
}
