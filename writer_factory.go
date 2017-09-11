package wfact

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

type WriterFct interface {
	Current() io.Writer
	NextWriter()
	Err() error
}

type Truncater struct {
	filename string
	w        *os.File
	e        error
}

func (wf *Truncater) Init(fname string) {
	wf.filename = fname
}

func (wf *Truncater) NextWriter() {
	if wf.w != nil {
		wf.w.Close()
	}
	os.Rename(wf.filename, wf.filename+"~")
	wf.w, wf.e = os.Create(wf.filename)
}

func (wf *Truncater) Current() (w io.Writer) {
	w = wf.w
	return
}

func (wf *Truncater) Err() (e error) {
	e = wf.e
	return
}

type DateArchiver struct {
	filename, cfn string
	w             *os.File
	e             error
}

func (d *DateArchiver) Init(fname string) {
	d.filename = fname
}

func (d *DateArchiver) NextWriter() {
	if d.w != nil {
		d.w.Close()
	}
	var nw time.Time
	nw = time.Now()
	d.cfn = fmt.Sprintf("%s.%d%d%d%d%d%d",
		d.filename, nw.Year(), nw.Month(), nw.Day(), nw.Hour(),
		nw.Minute(), nw.Second())
	d.w, d.e = os.Create(d.cfn)
}

func (d *DateArchiver) Current() (w io.Writer) {
	w = d.w
	return
}

func (d *DateArchiver) Err() (e error) {
	e = d.e
	return
}

type DWF struct {
	bf *bytes.Buffer
}

func NewDWF() (d *DWF) {
	d = &DWF{bytes.NewBufferString("")}
	return
}

func (d *DWF) Current() (w io.Writer) {
	w = d.bf
	return
}

func (d *DWF) NextWriter() {
	d.bf.Reset()

}

func (d *DWF) Err() (e error) {
	return
}
