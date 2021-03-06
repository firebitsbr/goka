// Code generated by go-bindata.
// sources:
// templates/common/base.go.html
// templates/common/head.go.html
// templates/common/menu.go.html
// templates/monitor/index.go.html
// templates/monitor/menu.go.html
// templates/monitor/processor.go.html
// templates/query/index.go.html
// DO NOT EDIT!

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesCommonBaseGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\x4a\x2c\x4e\x55\xaa\xad\xe5\xb2\x51\x74\xf1\x77\x0e\x89\x0c\x70\x55\xc8\x28\xc9\xcd\xb1\xe3\xaa\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\xca\x48\x4d\x4c\x51\x52\xd0\x03\xa9\x82\x48\xda\x24\xe5\xa7\x54\xa2\xaa\xc9\x4d\xcd\x2b\x85\xa8\x41\x16\x4d\xce\xcf\x2b\x49\xcd\x2b\x81\x6a\xd6\x87\x68\xb3\xd1\x87\x59\x91\x9a\x97\x52\x5b\xcb\x05\x08\x00\x00\xff\xff\xfd\x8f\xc0\x67\x8d\x00\x00\x00")

func templatesCommonBaseGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesCommonBaseGoHtml,
		"templates/common/base.go.html",
	)
}

func templatesCommonBaseGoHtml() (*asset, error) {
	bytes, err := templatesCommonBaseGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/common/base.go.html", size: 141, mode: os.FileMode(420), modTime: time.Unix(1490714385, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCommonHeadGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x54\x4b\x73\xe2\x38\x18\xbc\xe7\x57\xb8\x74\xe1\x90\xb5\x04\x98\x0d\x24\x85\xb3\x95\x25\xbb\x3c\xf2\x22\x90\x40\xd8\xcb\x96\x90\x3e\xdb\x72\x64\xc9\xb1\x64\xc0\xc3\xf0\xdf\xa7\x1c\x86\x84\xca\x64\x1e\x39\xcc\xcd\x5f\xbb\xd5\xee\x6e\x7f\xa5\xf5\x9a\x43\x20\x14\x38\x28\x02\xca\xd1\x66\x73\xd0\x2e\x1f\x4e\x0f\x1c\xc7\x71\xda\x09\x58\xea\xb0\x88\x66\x06\xac\x8f\x72\x1b\xb8\x2d\xb4\xff\x2a\xb2\x36\x75\xe1\x29\x17\x0b\x1f\xad\xdc\x9c\xba\x4c\x27\x29\xb5\x62\x2e\x01\x39\x4c\x2b\x0b\xca\xfa\x48\x80\x0f\x3c\x84\xdd\x49\x2b\xac\x84\xd3\xf5\x1a\xa7\x34\x84\xff\x9f\xa7\xcd\xa6\x4d\xb6\xf0\x9e\xb8\xa2\x09\xf8\x88\x83\x61\x99\x48\xad\xd0\x6a\x4f\x12\x7d\x4b\x5c\x08\x58\xa6\x3a\xb3\x7b\xac\xa5\xe0\x36\xf2\x39\x2c\x04\x03\xf7\x79\xf8\xc3\x11\x4a\x58\x41\xa5\x6b\x18\x95\xe0\xd7\x76\x42\x52\xa8\x47\x27\xca\x20\xf0\x2b\x65\x28\x73\x42\x48\xa0\x95\x35\x38\xd4\x3a\x94\x40\x53\x61\x30\xd3\x09\x61\xc6\xfc\x15\xd0\x44\xc8\xc2\xbf\x49\x41\x1d\x8e\xa9\x32\x87\x1d\xad\x38\x28\x03\xfc\xc4\xab\x56\x3f\xbf\xe0\x15\x27\x03\xe9\x57\x8c\x2d\x24\x98\x08\xc0\x56\x1c\x5b\xa4\xe0\x57\x2c\xac\x6c\xa9\x54\xd9\xff\x78\xc9\x45\xaf\x5c\xb4\x75\x83\x76\x6e\x12\xba\x62\x5c\xe1\xb9\xd6\xd6\xd8\x8c\xa6\xe5\x50\x1a\x7a\x01\x88\x87\x3d\xdc\x2c\x65\x5f\x31\x9c\x08\x85\x99\x31\xc8\x11\xca\x42\x98\x09\x5b\xf8\xc8\x44\xd4\x6b\x35\xdc\xbf\x27\x33\x21\xc6\xfd\x7f\xe1\xa2\xc6\xbb\xc9\x60\x74\xf6\x58\xb0\xbc\x77\xd6\x1b\x85\x5e\xfd\x26\xb9\x67\xcb\x65\x53\x2b\x6f\x34\xe3\x61\x63\x42\x0f\x87\xc9\xf8\xce\x7c\x22\x17\x47\xad\xc5\x9c\xff\x13\x47\x8d\x1c\x39\x2c\xd3\xc6\xe8\x4c\x84\x42\xf9\x88\x2a\xad\x8a\x44\xe7\x06\xfd\xee\x50\xae\x8d\x20\x81\x1f\x45\xcb\x7a\x85\xbe\xae\x89\x91\x99\x3c\x4c\x1a\xea\xbc\x3a\xc8\xad\x54\x5d\x6a\x64\x67\x90\x77\x9a\xf9\x32\xe6\xf9\xf4\x78\x3c\xc9\x2e\x17\xa3\x99\xd6\xc3\xb4\x3e\x9f\xce\xc2\x24\x1c\xdc\xf6\x1f\x96\x92\x8c\xd3\x9f\x45\xdb\x6e\xa4\x63\x32\xe6\x23\x42\x68\x4c\x57\x6f\xd7\xa4\xc4\x88\x14\x73\x43\xe2\xa7\x1c\xb2\x82\xd4\x70\xad\x86\xab\x5f\xa7\x67\xef\xb1\x41\xa7\x6d\xb2\x95\x7a\x47\xf7\xa3\x15\xc5\x6f\x7f\x7b\xfc\x6e\x35\x77\xec\xcf\xfe\xad\x98\x57\xeb\xcd\xa7\x45\x11\x8f\xaf\x82\x5e\x7c\x73\x45\x2f\x1f\x83\x7c\x3a\x59\xfd\xb7\xba\x1f\xaa\xce\xe0\xac\x29\xeb\x49\x67\x7a\xdd\x4f\xbb\xc7\x49\xb7\x73\xde\x5a\x76\xaf\xfb\x6c\x78\xde\xbc\x5b\xd1\xef\x57\xf3\x0b\x59\x18\x57\xb1\xc1\x4c\xea\x9c\x07\x92\x66\xf0\xa6\x2a\xa9\x39\x35\x11\x8e\x0d\x69\xe0\xda\x11\x6e\xec\x80\x0f\xb4\xc5\xbd\xd8\x60\x9d\x85\x84\x7b\x78\xd1\x78\xe7\x64\x9b\x6c\xaf\xb7\xf5\x1a\x14\xdf\x6c\x0e\xbe\x04\x00\x00\xff\xff\xee\x02\xd0\x16\x00\x05\x00\x00")

func templatesCommonHeadGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesCommonHeadGoHtml,
		"templates/common/head.go.html",
	)
}

func templatesCommonHeadGoHtml() (*asset, error) {
	bytes, err := templatesCommonHeadGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/common/head.go.html", size: 1280, mode: os.FileMode(420), modTime: time.Unix(1490714385, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCommonMenuGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\x41\x6e\x84\x30\x0c\x45\xf7\x9c\xc2\xca\x7e\xc8\x05\x32\x5c\x05\x19\x62\x44\x54\xd7\x8c\x42\x92\x8d\xe5\xbb\x57\xd3\x01\xda\x4e\x57\x89\xf5\x9f\xf2\x7e\xac\x1a\x69\x49\x42\xe0\x3e\x49\xaa\x33\xeb\x82\x60\x83\x99\x71\xdf\xef\x4e\xb0\x4d\x98\xe1\x75\xdc\x22\x2d\x58\xb9\xb8\xa1\x03\x08\x31\x5d\xd4\xbc\x49\xc1\x24\x94\x6f\x0b\xd7\x14\xbf\xf3\xbf\xc4\xf1\xc0\x4a\x18\x29\x1f\x39\x40\xc0\xb7\x7c\xca\x28\xd1\xc1\x9a\x69\xb9\x3b\xd5\x7e\xc2\x9d\xc6\x07\x96\xd5\xcc\xbb\x41\xb5\x7f\x76\x1c\x4b\x2a\x4c\x66\xc1\xe3\x21\xf2\x31\xb5\xff\xce\x79\x63\xc6\xc7\x4e\x67\xfb\x73\xfe\xd1\x57\xfe\xe5\x3f\x31\xc1\x76\x11\xaa\x13\x6f\xf3\xc7\x6b\x37\xe3\xf3\x9b\x24\xc5\x41\x6f\x76\x01\x24\xf1\x9a\x82\xaf\xfc\x56\xe9\xb8\x04\x2f\xd8\x86\x4e\x15\x48\x22\x98\x75\x5f\x01\x00\x00\xff\xff\x50\x2f\x43\xbe\x77\x01\x00\x00")

func templatesCommonMenuGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesCommonMenuGoHtml,
		"templates/common/menu.go.html",
	)
}

func templatesCommonMenuGoHtml() (*asset, error) {
	bytes, err := templatesCommonMenuGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/common/menu.go.html", size: 375, mode: os.FileMode(420), modTime: time.Unix(1490714385, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMonitorIndexGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x91\x4d\x6e\xc2\x30\x10\x85\xf7\x9c\x62\x64\xb1\x6c\xb0\x8a\x58\x55\xc1\x9b\xb2\x61\xd7\x1b\x54\x86\x19\xb0\x25\xd7\x46\xb6\x4b\x7f\x2c\xdf\xbd\x0a\x24\x0e\x49\x50\xd9\x24\x99\x79\xdf\xc4\xef\x79\x52\x42\x3a\x68\x4b\xc0\xf6\xce\x46\xb2\x91\xe5\x3c\xab\x51\x9f\x61\x6f\x64\x08\x6b\xe6\xdd\x17\x13\x33\x80\xdb\x5e\x83\x4a\x6d\xc9\x5f\x94\xb1\x66\xaa\x0f\xac\x9e\x97\xad\x06\x50\xab\xa5\x78\xf3\x6e\x4f\x21\x38\x1f\x6a\xae\x96\x9d\x92\xd2\x7c\x27\x03\xbd\x9f\x64\x54\xf0\xb2\x86\x45\xa9\x72\x2e\x88\x97\xf6\x48\x30\xd7\x16\xe9\xfb\x09\xe6\xa7\xee\x4f\x97\x81\x52\x85\x32\xf1\xc8\xce\x98\x38\x49\x4b\x06\x2e\xcf\x0a\xe9\x20\x3f\x4d\x1c\xb0\x77\xe8\x4a\x91\x44\x6d\x8f\x23\xae\x89\xba\x1a\x82\x51\x47\x43\x4c\xd4\x12\x94\xa7\xc3\x9a\xdd\x26\xce\x99\x17\xfb\x3c\xa5\x6b\xc2\x9c\x59\x7f\x59\xd0\x77\x6b\x2e\x27\x87\x71\xb5\x12\xb3\xa1\x53\x8e\xfa\xfc\xc8\xfc\xce\xe1\xcf\xd4\x39\x9a\x8e\x42\x53\x29\xe7\xf5\x6f\xb3\x64\x33\x01\x1b\x34\x8a\x57\xa3\xc9\x46\xd8\x6e\x6a\x8e\xf1\x1e\x82\x22\xa5\x7e\x57\x8b\x2b\xbf\xdd\x34\x41\x10\xa7\x49\xd0\x88\xff\x83\x8c\x1a\x83\x32\x25\xb2\xd8\xee\xbf\x08\xed\x47\xfb\xea\x90\xbf\x00\x00\x00\xff\xff\x68\xfb\xe0\xd6\xef\x02\x00\x00")

func templatesMonitorIndexGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesMonitorIndexGoHtml,
		"templates/monitor/index.go.html",
	)
}

func templatesMonitorIndexGoHtml() (*asset, error) {
	bytes, err := templatesMonitorIndexGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/monitor/index.go.html", size: 751, mode: os.FileMode(420), modTime: time.Unix(1490714385, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMonitorMenuGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\xce\x41\xaa\xc3\x30\x0c\x04\xd0\x7d\x4e\x21\x44\x96\x3f\xdf\xfb\xe2\x78\xd3\x83\x04\xc5\x56\x5a\x83\xa3\x80\x9d\x94\x82\xd0\xdd\x0b\x4d\x4b\x57\x03\x0f\x66\x18\xd5\xc4\x4b\x16\x06\x5c\x59\x8e\x29\x6e\xb2\xb3\xec\x68\xd6\xa9\x56\x92\x1b\x43\x9f\x25\xf1\xf3\x0f\xfa\x09\x2e\x23\xfc\xc7\x4d\xda\xb1\x72\x6d\x66\x1d\x80\x2f\x39\x78\x82\x58\xa8\xb5\x11\x85\x1e\x33\xd5\x61\xae\x24\x09\xe1\x5e\x79\x19\xd1\x7d\x0b\x4e\xf5\x9c\x32\xc3\x70\xfd\xe0\xf0\x43\xef\x28\x78\x57\x72\xe8\x54\x59\xd2\xfb\xc1\x99\xaf\x00\x00\x00\xff\xff\xf2\xc3\x81\x2e\xa4\x00\x00\x00")

func templatesMonitorMenuGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesMonitorMenuGoHtml,
		"templates/monitor/menu.go.html",
	)
}

func templatesMonitorMenuGoHtml() (*asset, error) {
	bytes, err := templatesMonitorMenuGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/monitor/menu.go.html", size: 164, mode: os.FileMode(420), modTime: time.Unix(1490714385, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMonitorProcessorGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x59\xdf\x6f\xdb\x38\x12\x7e\xf7\x5f\x31\xab\x43\xd7\xd2\xda\x96\xe4\xb4\xb7\x0f\x89\x9d\xc5\xfd\xc4\x15\xb8\xed\x16\xdd\xde\xde\x43\x10\x18\x8c\x34\xb6\xd9\xa5\x48\x81\xa4\xed\x04\xae\xfe\xf7\x03\xa9\x1f\xa6\x64\x29\x71\xef\xae\xed\x83\x2b\x73\x3e\x7e\x33\x9c\x19\x7d\x34\x99\xe3\x31\xc5\x35\xe5\x08\x5e\x22\xb8\x46\xae\xbd\xa2\x18\x2d\x52\xba\x87\x84\x11\xa5\x96\x9e\x14\x07\xef\x76\x04\xe0\x8e\x19\x28\xa1\x1c\xe5\x6c\xcd\x76\x34\xb5\xf6\x2e\x82\xcd\xb2\x74\x36\xbf\xea\xb1\xe5\x84\x23\x03\xfb\x39\x4b\x71\x4d\x76\x4c\x57\xa8\x61\xdc\x16\x49\x4a\xf9\xa6\xc1\x01\x2c\xb6\xf3\xdb\xe3\x31\xcc\xa5\x48\x50\x29\x21\xdf\x91\x0c\x8b\x62\x11\x6d\xe7\x0d\x57\x94\xd2\xfd\x20\xf1\xec\x41\xa4\x4f\x2e\x9f\x26\x0f\x0c\x6b\x48\xf9\xc5\x7e\xce\x94\x96\x34\xc7\xd4\xc1\x1a\xb4\x89\xc8\x1d\x31\x63\xb2\x3d\x60\x61\xb7\x1f\x0d\xc9\x22\xd2\xdb\x3e\xe3\x7b\x22\x35\xd5\x54\xf0\x21\xc0\xaf\x9a\xe8\x9d\x1a\xb2\x7e\x20\x1a\xc1\xcf\xd4\x26\x52\xc1\x33\x0c\x52\x0f\x19\xff\xb2\x93\x12\xf9\xa0\xf9\x1f\xff\xfe\x79\xc8\xf4\xb7\x8f\x1f\xc0\x7f\xc6\x2b\x27\xf9\x07\xa2\xa9\x00\x7f\x97\xa7\x44\xa3\x8a\x90\x6b\xf9\x34\x38\xe1\xbd\x14\x1b\x89\xaa\x67\xa9\x8b\xa8\x9d\x58\x83\xe8\x24\x7f\xa1\x4d\x35\x81\xa6\xa6\xb8\x55\x46\x7f\xa3\x78\xf0\x3a\xf3\x0c\xca\x29\x79\x64\x0b\xdc\xe9\x97\xb2\x5f\x55\x22\x69\xae\x41\x3f\xe5\xb8\xf4\x34\x3e\xea\xe8\x13\xd9\x93\x72\xb4\x61\xdd\x13\x09\x12\x79\x8a\xb2\xa9\xa3\x82\x25\xac\x77\x3c\x31\xcf\x7e\x13\x8b\x0a\x8e\xa3\xc6\xad\x99\x65\x2c\xff\xa4\x4a\xc3\x12\x56\xe1\x9a\x11\xad\x91\xfb\xab\x30\x23\xb9\x33\x69\x7a\x62\xda\x13\xb6\xc3\x29\x50\x9e\xe2\x63\x70\x74\x16\x95\x88\x2c\x17\x1c\xb9\x36\x8e\xad\x39\x54\x39\xa3\xda\xf7\x42\x2f\x08\x15\xa3\x09\xfa\x57\xc1\xc8\x99\x41\xd7\xfe\x77\x25\x90\xf2\x84\xed\x52\x54\xbe\xb7\x11\xbf\x13\x2f\x80\xcf\x9f\x1d\xbe\x90\x21\xdf\xe8\x2d\x7c\xb7\x84\xd7\x01\x1c\x5b\x25\x91\xa8\x77\x92\xc3\xdd\xfd\x8d\x33\x5c\x38\xcf\x36\xde\xbb\x71\x86\x5a\xd2\x84\x93\x0c\xc7\xf7\xb0\x74\xc9\x3f\x09\xca\x6d\x8c\x37\x6e\x6c\x35\xaf\x9d\xee\x90\x17\xc1\x14\xb4\xdc\xa1\x8b\x8e\x22\x48\x04\x57\x82\x61\xc8\xc4\xc6\xb7\x85\x07\x46\x95\x36\x9c\x27\x90\x8b\xa9\xb3\xee\xd2\xd4\xd5\xf8\x28\x11\x6d\x35\x4c\x0d\x56\xe1\x46\x8a\x5d\xfe\xe7\xa7\x66\xca\xb4\x5d\xd5\x56\x0d\xaa\xa8\xcd\x78\x78\x5a\x71\xa8\x76\x0f\x46\x3a\xf8\xc6\x8f\xa7\x67\x46\x46\x94\x7e\x6b\xaa\xf0\xcb\xda\xa6\x21\x68\x2f\xf6\xe5\xc2\x9b\xb8\x4b\x3e\x55\x87\xfd\x9b\xc1\x2a\x27\xf8\x6a\xf2\x89\x2c\x38\x76\xde\xbc\x2a\xf4\x7d\x7f\xdc\xfb\xa1\x88\xc7\xe1\x38\x98\xcc\x83\x9b\x16\x5b\x2b\xee\x72\xe2\x90\xbf\xd2\x7a\x17\xdf\x77\x19\xdc\xef\xd5\xea\xee\xc6\x4d\xff\xd8\x2c\xf4\x42\xac\xc5\xc5\x9c\x97\xa7\xc2\xba\x79\x7e\xa9\x9b\xb4\x44\x6c\x77\x13\x74\xdb\xc9\xb4\x4d\xb7\x9d\x4a\xc1\x6b\x24\xe1\xbd\xdd\xc7\xfa\x64\xc1\x49\x8f\x99\x57\x6e\x3a\x4b\x68\xec\x61\x59\x8f\xee\xfb\x1c\x4f\xe7\x41\xf5\x02\xb9\xd1\xd5\x9d\x7c\x01\xc3\x7c\x7a\xe5\x32\xb4\x28\x8c\xb6\xc2\x12\xbc\x85\x4e\x6f\xbd\x89\x8d\x69\xe2\x2d\x22\x9d\xde\x96\x23\x86\xbb\x1a\xf0\x46\x9d\x86\xb4\x2a\xec\xac\xb4\x53\x7f\x75\xa0\x3a\xd9\x9e\xd6\x1f\x2a\xbb\xbb\x85\xb6\x4f\x3b\xd8\x84\x28\x84\xf8\xda\xe4\x5c\xe4\xc8\xfb\xfb\x68\x6c\x42\x5a\xa8\x9c\xf0\x7a\xe3\x66\xe4\x01\x19\xd8\xcf\xd3\x8f\x8b\x5f\x72\xe4\x8b\xc8\xc0\x6e\x6d\xdc\xe3\x9b\x73\x57\x73\xeb\x4a\x62\x22\xf6\x68\x9a\x7f\xd4\xf1\x68\x96\xb7\x96\x22\x6b\x65\xb7\x82\xaf\x94\xd9\x63\x57\x62\xbd\x56\xa8\xe1\xfb\xef\x5f\x40\x94\xeb\x85\xdb\x65\x0c\x3f\x5d\x06\xbd\x86\xf8\xa6\x27\x1e\x69\xb6\xff\x25\x38\x09\x35\xed\xb9\xcb\x30\x5d\x65\xa8\x14\xd9\xa0\x5a\x59\x8c\xeb\xa6\x1f\x72\x37\x9e\x67\xa1\x79\x18\xdf\x1b\x6f\x41\x9f\x3b\x2d\x7a\x17\xbf\x3d\x64\xbd\xeb\xd8\x1e\xb2\xe7\xc3\x4f\xca\x1f\x1f\xbd\xa4\x95\xad\xce\x69\x1f\x7f\x1b\xe2\xba\xea\xf8\x32\xaf\x2d\x61\xc9\x8e\x99\x54\xe4\xd5\xef\x8c\x9e\x78\x6a\x93\x49\xa9\x5f\x07\x37\xb3\x55\x0f\x7e\x98\xc7\x31\x44\xe0\x6b\x31\xb3\xdf\x83\x50\x8b\xbf\xd3\x47\x4c\xfd\xab\xb3\x5c\x39\x3c\x54\xbd\x23\xef\xfc\x7a\x20\xf8\x29\xbe\xae\x9f\xfb\x12\x82\x5a\x5a\xdf\x5a\xc0\xac\xce\x4e\x00\x91\xad\xf3\xc9\xe1\x79\x71\xdc\xb7\xe1\x43\xd3\xc1\x65\xab\xc3\xa4\x03\x36\xff\x2c\x72\x3c\x31\xbc\xce\x3a\x26\xe3\x97\xa7\x98\xc5\x5f\x82\xab\xa2\xbf\x04\xaa\xc5\x25\x28\xd4\xb2\x86\x0d\xa2\x66\x2f\xd8\x5b\x87\x80\xaa\x0e\x1e\x28\xfd\xc4\x70\xe9\x65\x44\x6e\x28\x9f\x3d\x08\xad\x45\x76\x1d\xe7\x8f\x37\x5e\x3f\x95\x21\xeb\x21\x9a\x3d\x10\xe9\x81\x14\x86\xab\x1e\xb3\x43\x35\x3f\xe5\xb3\x03\x4d\xf5\xf6\x1a\xde\x60\x76\x03\xe5\xf3\x18\x26\xa7\x7e\x99\xc0\xf8\x95\xf1\xda\x9b\x87\x53\x5f\x4d\xc6\xaf\xca\xdf\xaa\xf5\x67\x29\x69\xe7\x9a\x76\xe5\x6a\x1a\x28\x4d\x18\xc3\xf4\x7f\x52\xd2\xaa\xbb\xe0\xd7\x92\xeb\x25\x51\x7d\xed\x06\xd0\x71\xfd\x4d\x05\xec\xab\xca\x4d\xd7\x91\x11\x44\xd7\x49\xf3\xb4\x62\x82\xa4\x3d\x8a\xd9\x07\x18\xd4\xb3\xf2\xd0\xb1\x46\x89\x3c\x31\xc9\xfb\x99\xe8\x6d\x98\x51\x5e\x2b\xd6\xb4\x1a\x21\x8f\xfe\xf6\x90\xcd\x4c\x4a\x7e\xf8\x31\x9e\x42\x1c\x74\x72\x42\xd7\xbe\x8d\x74\x09\xf1\xd9\xef\x34\xe3\x84\x91\x0d\x2c\x61\x1e\x77\x56\x58\x20\x53\x38\x8c\xf7\x1b\xef\x8e\x84\x36\xf1\x06\x26\x8e\x5a\x4b\x6d\x78\x8d\x65\x58\x51\x8b\x51\xdf\xe1\xa3\x25\x77\xa6\x15\x07\x14\xe4\x8b\xb5\xce\x95\x12\x98\x40\x14\xd9\x0d\x79\x88\xf7\x65\xa1\xab\x80\xdb\xc3\x73\xc2\x79\xe6\xd4\x6c\x06\x17\xa0\x54\x7d\xce\xee\xc7\xfe\x3f\x14\x6f\x50\xed\xc0\xfd\x32\x53\xbb\x24\xb1\xec\xff\x8d\x04\x9a\xde\x79\x46\xfd\x18\xd9\x4c\xc6\xaf\x40\x22\x23\x9a\xee\xd1\x7c\xdf\xd8\x1d\xee\x12\x15\x7c\x53\x8a\xd0\x8e\xf3\xb3\x9f\x75\xdf\x4c\x16\x5a\x3b\x74\x19\xc9\xcb\x5b\xde\xd7\x90\xc5\x2f\xda\xee\x4f\x9d\x66\x5b\xad\xf7\x35\xf8\xf2\x1d\xbf\xcd\xba\x3d\x64\x97\x01\xfb\xde\x87\xbe\x54\x99\x37\x42\x6d\x85\x5e\x95\xc7\x30\x69\x2f\xa1\xdc\x44\xf5\x01\x4e\x75\xeb\xcb\x50\xcf\xce\xf6\x47\xdb\x54\x09\x13\xca\x34\x55\x3b\xa6\xaa\xd8\x5e\x65\xf5\x7a\xa6\xff\xd8\x4c\x1f\xda\x90\xbd\xd2\xd8\x99\x5c\x6d\xc4\xd7\x03\x73\xfe\xc5\x7f\xe7\xe2\xc0\x4d\xa9\x34\x7a\x5d\x15\x75\x9e\xfd\xf6\xbd\x4b\x79\xfa\xa4\xe6\x94\xea\x79\x7d\x87\x67\x7b\x28\x9c\x94\xe7\xbb\x89\x85\x3a\xa8\xa2\x75\x8c\x2e\xb3\x3a\x05\xe4\x1a\x25\x10\x9e\x82\xc4\x4c\xec\x11\x52\xa2\x49\xeb\xa4\x69\x8e\x99\xe9\xeb\x50\x21\xc3\x44\xfb\xde\x1f\xda\x57\x77\x41\x65\xf8\x13\x63\xbe\x77\xda\x22\x1f\xc4\xa3\x17\x84\x86\xab\x39\x7f\x3b\xd7\x0e\x69\x70\xac\x43\x4e\xc3\xf2\x2a\xc0\xbd\x54\x48\xc3\xad\xce\x98\xdf\x77\x40\x77\x51\xdf\x20\xa6\xd0\xa6\xc7\x0f\x42\x92\xe7\xc8\x53\xdf\xd3\xd2\x0b\x42\x2b\xb2\x98\xfa\x5e\x8b\xba\xba\xfb\xba\x2c\xf6\x10\x1f\xa9\xf6\x83\xb0\x4c\xba\xdf\x5c\xfd\x15\x4d\xc1\x0e\x94\xa7\xe2\x10\x2a\xd4\x6f\x4d\x0c\x7b\xc2\xfc\xbe\x03\x7b\xfa\x3a\xfc\xa4\x04\xf7\xbd\xe3\x31\x7c\x20\x0a\x57\x39\xd1\xdb\xa2\x88\x9a\x1b\x77\xb3\xde\xc8\xbd\x82\x7f\xfb\xd7\xa2\xf0\xa6\x67\x37\xa2\x4d\x74\xc5\x14\xde\xc4\x71\x7c\x6a\xbd\xf2\x60\xc6\x80\x6a\xa0\x9c\x6a\x4a\x18\x7b\x1a\x7d\x1d\xef\x8b\xa8\xbc\xbb\xad\xfe\x1c\x51\xff\x69\xa0\x7a\xa8\xfe\x3b\x1e\x91\xa7\x45\x31\xfa\x4f\x00\x00\x00\xff\xff\x5c\xac\xb8\x72\x15\x19\x00\x00")

func templatesMonitorProcessorGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesMonitorProcessorGoHtml,
		"templates/monitor/processor.go.html",
	)
}

func templatesMonitorProcessorGoHtml() (*asset, error) {
	bytes, err := templatesMonitorProcessorGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/monitor/processor.go.html", size: 6421, mode: os.FileMode(420), modTime: time.Unix(1498560133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesQueryIndexGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\x4d\x6f\xe3\x36\x10\xbd\xe7\x57\x4c\x08\xa3\x39\xb4\xb2\xd0\xe4\xd6\x50\xea\xa5\x39\x15\x28\x50\xf4\xd0\x63\x40\x89\x23\x89\x28\x4d\xaa\xe4\x30\x8e\x21\xe8\xbf\x2f\xf4\x19\x5b\x96\xed\x2c\x16\x0b\xec\x45\xc9\x68\x3e\x38\xef\xcd\xe3\x58\x4d\x23\xb1\x50\x06\x81\xe5\xd6\x10\x1a\x62\x6d\x7b\xc7\xa5\x7a\x83\x5c\x0b\xef\x13\xe6\xec\x9e\xa5\x77\x00\x00\xc7\x6f\xbb\x60\xa1\x0c\xba\xd1\xb7\xf4\x7f\x64\xad\x79\x73\xab\xa3\x9d\x8c\x7e\x7d\x5c\xc4\x00\xf0\xea\x31\xfd\x3b\xa0\x3b\xc0\x9f\x78\xf0\x3c\xae\x1e\x97\x11\x4d\xa3\x0a\xd8\xa2\x73\xd6\xb5\xed\x32\xfb\xe8\x0c\xa1\xd1\x11\xf4\xcf\x48\x0a\x53\xa2\x9b\x0c\xe5\x77\xca\x7b\x95\x69\x64\xe0\xac\xc6\x31\xf6\xac\x17\x00\x9e\x05\x22\x6b\x80\x0e\x35\x26\x6c\x30\xd8\x0c\x42\x5b\x8f\x0c\xa4\x20\x31\xd5\x9c\x2b\x71\x5f\x0b\x93\xfe\x44\x6a\x87\xfe\x99\xc7\xbd\xc5\xe3\xa1\xc0\xca\x31\x9e\x9c\x35\x65\xfa\xd2\x81\xba\xe7\xf1\x68\x42\xd3\x0c\x38\xb7\x2f\xeb\x68\x63\xa9\xde\xce\xe9\x41\x23\xcf\x42\x07\xd2\xf6\xc2\x19\x65\xca\x4f\xd3\x36\xc6\xff\xf0\xbc\xfd\x3b\xf4\x79\xca\xdc\xd8\xfc\xb7\x73\x37\xb1\xe7\x6d\x70\x39\xfa\x15\x3f\x2f\xac\xdb\x81\x35\x3e\x64\x3b\x45\x09\xdb\x2b\x23\xed\x7e\xab\x6d\x2e\x48\x59\x03\x09\x3c\x34\xcd\x36\x13\x1e\x5f\x6b\x41\x55\xdb\xc6\x4d\xb3\xf5\xa8\x31\x27\x94\xaf\x43\xdd\xb6\x8d\x1f\xe0\x67\xf0\x28\x5c\x5e\x6d\xdf\x84\x0e\xf8\x0c\x0e\x29\x38\x03\x85\xd0\x1e\x9f\x57\x78\x5e\x0e\x4f\x99\x3a\x50\x54\x3a\x1b\x6a\x38\xfa\x3f\xd2\xe5\x85\xe4\x2b\x05\xa2\x8c\xcc\x95\xac\x9b\x63\xce\xc8\x40\x46\x26\x92\x58\x88\xa0\x09\xa4\xb3\xb5\xb4\x7b\x13\x91\x2d\x4b\x3d\x09\x60\x30\x12\x36\x79\x59\xba\x46\x4d\x2f\x8b\x59\x3f\xc2\x61\x27\x95\x5b\xea\x38\xe9\x34\xe8\x29\x7d\xee\x63\x87\x26\xdc\x00\x08\xfd\xf0\x37\xf3\xe4\xe0\xb7\x04\x8e\xe7\xf8\x89\x64\xd7\xad\x1e\xd8\x28\x23\xf1\xfd\x17\xd8\x0c\x88\xfa\x3a\x97\xf5\x74\xd6\xbd\x56\x29\x17\x50\x39\x2c\x12\x76\xdc\x4f\xaf\xa4\xcd\xc4\x52\xc7\xdd\x6c\xf0\x58\xa4\x3c\xd6\xea\x33\x00\xd7\x55\x7f\xd2\x41\x1c\xf4\x15\x05\xad\x5d\xa6\x23\x77\xaf\x2a\x50\x32\x61\x83\xbc\xd9\xa8\x18\xc2\x77\x9a\xf5\xd2\xdd\xa1\xa8\xfb\x59\x71\x56\x33\x70\xf8\x7f\x50\x0e\xe5\xf7\x97\xed\x70\x69\x2f\xc9\x96\xa5\xff\xf4\x2d\xdf\x96\xd9\x15\x12\x2e\xb8\x78\xdc\x61\x3e\x7f\xdf\x34\xa8\x3d\xae\xed\x99\xeb\x7b\x7a\xb1\x95\xff\xb2\x30\x4a\x0c\x0a\x1b\x8c\xbc\x87\x3f\x94\x84\x83\x0d\x50\x58\x57\x22\x01\x59\x10\x44\x22\xaf\x80\x2a\xdc\xfd\x7e\xa1\xcb\x35\x79\x2c\x42\x17\xe6\xb0\x2e\xfb\x15\x76\x94\xc7\x2b\x17\xaf\x7f\x2b\xd4\xc2\xa0\x86\xfe\x39\x6f\x8b\xc1\xf2\x21\xcf\xd1\xfb\x2b\x5f\x12\x43\x5c\x85\x42\x76\x04\xac\x70\x5c\x3d\x9d\x86\x92\x22\x8d\xfd\x96\xf9\x0f\x0f\xdd\x35\xa9\x9e\xd2\x6b\xd8\xd6\x0f\xcc\xac\x3c\xac\x9d\x56\x3b\xec\x4a\x8f\xd8\x79\xdc\xd9\x5f\xc5\xdc\x07\xd7\xa3\x6b\xfc\x33\x79\xbe\x04\x00\x00\xff\xff\x05\xda\xbe\x9a\xaf\x09\x00\x00")

func templatesQueryIndexGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesQueryIndexGoHtml,
		"templates/query/index.go.html",
	)
}

func templatesQueryIndexGoHtml() (*asset, error) {
	bytes, err := templatesQueryIndexGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/query/index.go.html", size: 2479, mode: os.FileMode(420), modTime: time.Unix(1490714385, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/common/base.go.html": templatesCommonBaseGoHtml,
	"templates/common/head.go.html": templatesCommonHeadGoHtml,
	"templates/common/menu.go.html": templatesCommonMenuGoHtml,
	"templates/monitor/index.go.html": templatesMonitorIndexGoHtml,
	"templates/monitor/menu.go.html": templatesMonitorMenuGoHtml,
	"templates/monitor/processor.go.html": templatesMonitorProcessorGoHtml,
	"templates/query/index.go.html": templatesQueryIndexGoHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"common": &bintree{nil, map[string]*bintree{
			"base.go.html": &bintree{templatesCommonBaseGoHtml, map[string]*bintree{}},
			"head.go.html": &bintree{templatesCommonHeadGoHtml, map[string]*bintree{}},
			"menu.go.html": &bintree{templatesCommonMenuGoHtml, map[string]*bintree{}},
		}},
		"monitor": &bintree{nil, map[string]*bintree{
			"index.go.html": &bintree{templatesMonitorIndexGoHtml, map[string]*bintree{}},
			"menu.go.html": &bintree{templatesMonitorMenuGoHtml, map[string]*bintree{}},
			"processor.go.html": &bintree{templatesMonitorProcessorGoHtml, map[string]*bintree{}},
		}},
		"query": &bintree{nil, map[string]*bintree{
			"index.go.html": &bintree{templatesQueryIndexGoHtml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

