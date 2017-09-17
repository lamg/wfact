package wfact

import (
	"bytes"
	"fmt"
	"github.com/lamg/errors"
	"io"
	"os"
	"time"
)

const (
	// ErrorNWTrunc is the error set by Truncater.NextWriter
	ErrorNWTrunc = iota
	// ErrorNWDtA is the error set by DateArchiver.NextWriter
	ErrorNWDtA
)

// WriterFct is an interface for genariting io.Writers
// without worrying about what stands behing those io.Writers
type WriterFct interface {
	// Get the current io.Writer
	Current() io.Writer
	// Close current io.Writer and create one new
	NextWriter()
	// Returns the error of a call to NextWriter
	Err() *errors.Error
}

// Truncater is an implementation of WriterFct that uses
// an os.File as io.Writer and truncates at each call to
// NextWriter
type Truncater struct {
	filename string
	w        *os.File
	e        *errors.Error
}

// Init initializes Truncater
func (wf *Truncater) Init(fname string) {
	wf.filename = fname
}

// NextWriter truncates the *os.File with name fname
func (wf *Truncater) NextWriter() {
	if wf.w != nil {
		wf.w.Close()
	}
	os.Rename(wf.filename, wf.filename+"~")
	var e error
	wf.w, e = os.Create(wf.filename)
	if e != nil {
		wf.e = &errors.Error{
			Code: ErrorNWTrunc,
			Err:  e,
		}
	}
}

// Current returns the *os.File as a io.Writer
func (wf *Truncater) Current() (w io.Writer) {
	w = wf.w
	return
}

// Err returns any error during operation
func (wf *Truncater) Err() (e *errors.Error) {
	e = wf.e
	return
}

// DateArchiver is an WriterFct that creates new
// files at each call to NextWriter with the date
// as a number appended at the end of the file name
type DateArchiver struct {
	filename, cfn string
	w             *os.File
	e             *errors.Error
}

// Init initializes DateArchiver
func (d *DateArchiver) Init(fname string) {
	d.filename = fname
}

// NextWriter closes the current file and creates
// a new one
func (d *DateArchiver) NextWriter() {
	if d.w != nil {
		d.w.Close()
	}
	var nw time.Time
	nw = time.Now()
	d.cfn = fmt.Sprintf("%s.%d%d%d%d%d%d",
		d.filename, nw.Year(), nw.Month(), nw.Day(), nw.Hour(),
		nw.Minute(), nw.Second())
	var ec error
	d.w, ec = os.Create(d.cfn)
	if ec != nil {
		d.e = &errors.Error{
			Code: ErrorNWDtA,
			Err:  ec,
		}
	}
}

// Current returns the current *os.File as a io.Writer
func (d *DateArchiver) Current() (w io.Writer) {
	w = d.w
	return
}

// Err returns any error during operation
func (d *DateArchiver) Err() (e *errors.Error) {
	e = d.e
	return
}

// DWF is a dummy writer factory for testing
type DWF struct {
	bf *bytes.Buffer
}

// NewDWF creates a new DWF
func NewDWF() (d *DWF) {
	d = &DWF{bytes.NewBufferString("")}
	return
}

// Current returns the *bytes.Buffer as an io.Writer
func (d *DWF) Current() (w io.Writer) {
	w = d.bf
	return
}

// NextWriter resets the *bytes.Buffer
func (d *DWF) NextWriter() {
	d.bf.Reset()
}

// Err returns no error. Made for implementing WriterFct
func (d *DWF) Err() (e *errors.Error) {
	return
}

// Content returns the content of the current writer
func (d *DWF) Content() (bs []byte) {
	bs = d.bf.Bytes()
	return
}
