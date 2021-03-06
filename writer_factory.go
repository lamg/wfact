package wfact

import (
	"bytes"
	"fmt"
	fs "github.com/lamg/filesystem"
	"io"
	"time"
)

// WriterFct is an interface for genariting io.Writers
// without worrying about what stands behing those io.Writers
type WriterFct interface {
	// Get the current io.Writer
	Current() io.Writer
	// Close current io.Writer and create one new
	NextWriter()
	// Returns the error of a call to NextWriter
	Err() error
}

// Truncater is an implementation of WriterFct that uses
// an os.File as io.Writer and truncates at each call to
// NextWriter
type Truncater struct {
	filename string
	w        fs.File
	fsm      fs.FileSystem
	e        error
}

// NewTruncater creates a new Truncater
func NewTruncater(fname string,
	fsm fs.FileSystem) (wf *Truncater) {
	wf = &Truncater{filename: fname, fsm: fsm}
	return
}

// NextWriter truncates the *os.File with name fname
func (wf *Truncater) NextWriter() {
	if wf.w != nil {
		wf.w.Close()
	}
	wf.fsm.Rename(wf.filename, wf.filename+"~")
	wf.w, wf.e = wf.fsm.Create(wf.filename)
}

// Current returns the *os.File as a io.Writer
func (wf *Truncater) Current() (w io.Writer) {
	w = wf.w
	return
}

// Err returns any error during operation
func (wf *Truncater) Err() (e error) {
	e = wf.e
	return
}

// DateArchiver is an WriterFct that creates new
// files at each call to NextWriter with the date
// as a number appended at the end of the file name
type DateArchiver struct {
	filename, cfn string
	w             fs.File
	e             error
	fsm           fs.FileSystem
}

// NewDateArchiver creates a new DateArchiver
func NewDateArchiver(fname string,
	fsm fs.FileSystem) (d *DateArchiver) {
	d = &DateArchiver{filename: fname, fsm: fsm}
	return
}

// NextWriter closes the current file and creates
// a new one
func (d *DateArchiver) NextWriter() {
	if d.w != nil {
		d.w.Close()
	}
	var nw time.Time
	nw = time.Now()
	d.cfn = fmt.Sprintf("%s.%s",
		d.filename, nw.Format(time.RFC3339))
	d.w, d.e = d.fsm.Create(d.cfn)
}

// Current returns the current *os.File as a io.Writer
func (d *DateArchiver) Current() (w io.Writer) {
	w = d.w
	return
}

// Err returns any error during operation
func (d *DateArchiver) Err() (e error) {
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
func (d *DWF) Err() (e error) {
	return
}

// Content returns the content of the current writer
func (d *DWF) Content() (bs []byte) {
	bs = d.bf.Bytes()
	return
}
