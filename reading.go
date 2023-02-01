package libx

import (
	"io"
	"io/ioutil"
	"net/http"
)

// reading read interface
type reading interface {
	IORead(io.Reader) *read // IORead: io mode read
	PathRead(string) *read  // PathRead: path read
	NetRead(string) *read   // NetRead: net url read
	Read() ([]byte, error)  // Read: start read func
}

// read
type read struct {
	path      string    // path: file path
	url       string    // url: file url
	rd        io.Reader // rd: io.Reader mode
	isNetRead bool      // isNetRead: is net file read default false
	isIoRead  bool      // isIoRead: is io.Reader mode read data default false
}

var _ reading = (*read)(nil)

// NewRead
func NewRead() *read {
	return &read{
		isNetRead: false,
		isIoRead:  false,
	}
}

// IORead
// io mode read
func (r *read) IORead(rd io.Reader) *read {
	r.rd = rd
	r.isIoRead = true
	r.isNetRead = false
	return r
}

// PathRead
// path or url mod read
func (r *read) PathRead(path string) *read {
	r.path = path
	r.isIoRead = false
	r.isNetRead = false
	return r
}

// NetRead
// net url read, http get method read file
func (r *read) NetRead(url string) *read {
	r.url = url
	r.isIoRead = false
	r.isNetRead = true
	return r
}

// Read
// start read func
func (r *read) Read() ([]byte, error) {
	var result []byte
	var err error
	if r.isIoRead {
		result, err = ioutil.ReadAll(r.rd)
	} else if r.isNetRead {
		resp, e := http.Get(r.url)
		if e != nil {
			err = e
		} else {
			defer resp.Body.Close()
			result, err = ioutil.ReadAll(resp.Body)
		}
	} else {
		result, err = ioutil.ReadFile(r.path)
	}
	return result, err
}
