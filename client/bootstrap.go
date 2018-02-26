// Code generated by go-bindata.
// sources:
// client/bootstrap/daemon-down.sh
// client/bootstrap/daemon-up.sh
// client/bootstrap/docker.sh
// client/bootstrap/keygen.sh
// client/bootstrap/token.sh
// DO NOT EDIT!

package client

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

var _clientBootstrapDaemonDownSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xce\xc1\x4a\xc3\x40\x10\xc6\xf1\xfb\x3c\xc5\x47\xdb\xeb\xe6\x0d\x72\x88\x36\x88\x60\x57\x28\x7a\xf0\xd4\xae\xd9\xd9\x74\x48\x33\x5b\x67\xb7\xa8\x6f\x2f\x22\x42\xe8\xfd\xe3\xff\xfb\xd6\xb8\x0b\x45\x06\x94\xc1\xe4\x52\x91\xb2\xe1\xdd\x44\x47\xd1\x11\x31\x7f\x2a\xea\x89\x11\x03\xcf\x59\x1b\xa2\xc2\x15\x8e\x89\xb6\x5d\xbf\x7b\xf6\x07\xdf\xed\xfa\x56\x94\xad\x4a\x70\x7f\x23\xa2\x35\xee\x4f\x3c\x4c\x90\x84\x70\x36\x0e\xf1\x1b\x76\x55\x15\x1d\x1b\xea\x9e\xf6\x7d\xb7\x7d\x3b\xec\x5f\xbd\x7f\xf4\x0f\xed\xb1\x5c\x63\x46\xcc\xc3\xc4\x86\x4b\x81\xfb\x80\x73\x49\xce\x95\x0d\x2b\x0d\x33\xb7\x9b\x05\xb5\x3a\xfe\xd6\x5f\xc2\xc4\xe0\x2f\x29\xf5\xff\x63\x43\xcb\x8c\xcd\x70\x09\x9b\x1b\x8a\x7e\x02\x00\x00\xff\xff\x8a\x49\xa7\x95\xe9\x00\x00\x00")

func clientBootstrapDaemonDownShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapDaemonDownSh,
		"client/bootstrap/daemon-down.sh",
	)
}

func clientBootstrapDaemonDownSh() (*asset, error) {
	bytes, err := clientBootstrapDaemonDownShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/daemon-down.sh", size: 233, mode: os.FileMode(420), modTime: time.Unix(1514021421, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientBootstrapDaemonUpSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x53\x5d\x6f\xdb\x46\x10\x7c\xbf\x5f\xb1\xa5\x54\xf8\x25\xc7\x73\x5a\x14\x48\x1d\xe8\x41\x4d\x84\x58\x48\x45\x19\x92\x8a\x22\x68\x0b\xe5\x74\x5c\x89\x17\x91\x77\xec\xed\x52\x8c\xfd\xeb\x8b\xa3\x3e\x28\xbb\x7e\x92\xb8\xdc\x9d\x99\x1d\xce\x0e\xe0\x37\x4d\xd6\x00\x99\x60\x6b\x86\xad\x0f\xb0\x09\xd6\xed\xac\xdb\x41\x53\x03\x17\x08\xb9\xc6\xca\xbb\x54\x08\x42\x06\x89\x42\x3c\xcc\x17\xab\xd1\x8f\x24\x3e\x8e\x27\xb3\x79\xb6\xce\xc6\xb3\xc9\xc8\x3a\x0c\x6c\xb5\x3c\x36\x8b\x0f\xf3\x6c\x35\x9e\x66\x93\xc5\xba\x6b\x7e\x77\xfb\xee\xad\x98\xce\xc6\x9f\x26\xeb\xc5\xe4\x61\xbe\x9c\xae\xe6\x8b\x2f\xa3\x66\x63\x4a\xdd\x38\x53\xd4\x3a\x57\x27\x00\x21\x06\xf0\xa1\x40\xb3\x07\xbb\x05\x5d\x06\xd4\xf9\x23\x84\xc6\x39\xeb\x76\xa9\x18\xff\xbe\x98\x8c\x3f\x7e\x59\x2f\xfe\xc8\xb2\x69\xf6\x69\xf4\x95\x9a\xdc\x43\xee\xcd\x1e\x03\xd4\x04\xf2\x5f\x90\x72\x6b\x4b\xc6\x00\x89\xd3\x15\x8e\x86\x57\x22\x93\xaf\x11\x7d\xa5\xf7\x08\xf8\xdd\x12\xc7\x1d\x73\xdf\xba\x54\xd8\x2d\xfc\x05\x3f\x80\x7c\x82\x64\xf8\x82\x23\x81\x7f\xde\x47\x1b\x9c\x00\x00\x40\x53\x78\x48\x3e\xdb\xb2\x8c\xc3\x17\x14\xe3\x1d\xeb\xb8\x41\xd2\x75\x5d\xab\x0a\x15\xc8\x2d\xbc\x44\x15\x5b\xfb\x3e\x8a\x79\x68\xca\xb2\x33\xb9\xd4\x8c\xc4\x70\x72\xe1\xe2\xf9\xb3\xfd\x62\xef\xf0\xa5\x8b\x11\x65\x16\x57\x7a\x08\xfe\x1b\x1a\x86\xdc\x06\x34\xec\xc3\xa3\xa8\xf6\xb9\x0d\x20\x6b\x18\xde\xcf\x67\x13\x55\x1f\x1b\xe2\xc0\xa2\x71\xbd\x66\x68\x2d\x17\xa0\x8d\x41\x22\x60\xdf\xc9\x29\x3c\xf1\x99\x97\xe2\x0f\x83\x76\x39\x04\x2c\xf1\xa0\x5d\x4f\x62\x91\x40\x8a\x01\x70\x61\x09\x2c\x81\xc3\x88\xa2\xc3\x23\x6c\xd0\xe8\x86\x10\x5a\x84\x36\x4e\xf4\x49\x8a\x1c\x1b\x04\xbd\x29\x11\x88\x75\x60\x31\xe8\xc0\x89\x7d\xdd\xab\x22\x88\x9d\x27\x29\x29\x4c\xf9\x86\x40\x97\xe4\xbb\x8e\xe0\x0f\x18\xc8\xea\xf2\x8d\x18\x40\xc1\x5c\xd3\x9d\x52\x6d\xdb\xa6\xe5\xa1\x48\xad\x57\xb5\x27\x26\x95\x7b\xc7\x12\xbf\xd7\x9e\x50\x72\x81\xf2\xb8\x8f\x3c\xee\x23\x9d\x67\x89\x07\x74\x92\xbd\xd4\xf2\xc2\x9b\x16\x5c\x95\x62\x70\x45\x18\xd0\xf8\xaa\x42\x97\x63\x7e\x4d\xf7\xad\x46\xd6\x4f\x4f\x3e\xdd\x59\x2e\x9a\x4d\xa4\xfd\xe9\xf6\xed\x2f\xea\xf6\x57\x75\xfb\xb3\xca\x7d\x47\xd0\xd0\x85\xd6\xba\xf3\xbf\xad\x0f\xd2\x58\x25\x06\x30\x26\xd0\x10\x90\x9a\x92\xdf\x1c\x3d\xec\xbf\x4a\xa1\x09\x82\xf7\x7c\xfe\x32\x27\x3b\x02\x56\x9e\x11\x0e\x35\x3d\x4f\x47\x68\x1c\xc8\x1c\xa4\x0c\x15\xfc\xdd\xa5\x50\xd6\x90\x0c\xe3\xf9\x25\x77\xf1\xfe\xce\xd5\x03\xa8\x83\x0e\x2a\x34\x4e\x1d\x47\xd3\x68\xc8\xdd\x6b\xc5\x7e\xa4\x4b\xd0\x9d\xd2\x75\xad\xba\x68\x9c\x5e\x20\x1c\xcf\x6b\x94\x70\x68\x30\xe9\xcb\xb1\x7d\xd4\x0d\xf5\xb5\xe5\xf2\x7e\xfd\x39\x9b\xff\x99\xad\xef\xe7\xcb\xd5\x72\x74\x73\x81\x53\x29\x51\xa1\xf6\xce\xb7\x6e\x1d\x9f\xe9\xe6\x3c\x25\xe3\x11\x43\xf2\xec\x8a\x4f\xef\x92\xff\x9d\x42\x12\xa3\xfd\x8a\xd8\xca\x37\x8e\xe9\x94\xe9\xf3\x75\x00\xf9\xeb\x4c\x1a\xed\xce\x46\x5b\x16\xff\x05\x00\x00\xff\xff\xe2\x8b\xf9\x8a\x16\x05\x00\x00")

func clientBootstrapDaemonUpShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapDaemonUpSh,
		"client/bootstrap/daemon-up.sh",
	)
}

func clientBootstrapDaemonUpSh() (*asset, error) {
	bytes, err := clientBootstrapDaemonUpShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/daemon-up.sh", size: 1302, mode: os.FileMode(493), modTime: time.Unix(1517734298, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientBootstrapDockerSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x5b\x6b\xdc\x3e\x10\xc5\xdf\xf5\x29\x4e\xb2\x4b\x2e\x10\xdb\x49\xf8\xf3\x7f\x68\x48\xa1\x4d\x42\x29\x0d\x04\xf6\xf2\x5c\xb4\xd6\x68\x2d\xe2\x95\x5c\xcd\x38\xdb\xa5\xf4\xbb\x17\x6b\x6d\xb7\x69\xd3\x8b\x9f\x64\xcd\x99\xd1\x99\xdf\xcc\xe4\xa0\x58\x39\x5f\x70\xa5\xd4\x04\x6f\x43\x10\x96\xa8\x1b\x86\xc6\x46\x97\x95\xf3\x04\x1b\x22\x4c\x28\x1f\x29\x42\x7b\xd3\x1f\xb3\x32\x6c\x9a\xc0\x94\xab\x09\x16\x95\x63\x38\x46\x13\x49\x64\x07\x5d\x37\x95\xce\xf1\xde\xb3\xe8\xba\x66\x94\x6d\xac\x11\x7c\xbd\x83\xb3\x70\x02\x13\x88\xfd\xb1\x80\x3e\x3b\x96\x94\x4f\x2c\x64\x10\x3c\x96\xab\xd6\x4b\x8b\x8b\xff\xf3\xf3\xff\xce\x10\xe9\x53\xeb\x22\x31\xb8\x35\x21\xd9\xd0\xb0\xb4\x05\x0b\x35\x9c\x2b\xc5\x24\xc8\x48\xa9\xdb\x87\x9b\x0f\x77\xb3\x8f\xf3\x87\xe5\xec\xe6\xee\x7a\x4d\x92\xef\x3d\xe6\x65\xd8\x0c\xc1\xdb\xbb\xf9\xe2\xfa\xb8\x90\x4d\x53\xac\x49\xb2\x5e\xc0\xd5\x71\xd7\xf6\x4d\x45\xe5\x63\xe7\x6e\x68\x33\x92\x36\x3b\xb8\x7d\x07\x64\x72\xe5\x2c\x2a\xcd\xd5\x20\xb8\x7c\x5d\x18\x7a\x2a\x7c\x5b\xd7\x57\x90\x8a\xbc\x02\xd0\x35\x24\x38\x57\xd6\x5d\x29\x65\x49\xca\xca\xba\x9a\x4e\x4e\xf1\x25\x45\x27\x78\x13\xd7\xfc\xaa\x3f\x03\xd3\x0b\x70\x68\x63\x49\x58\xce\xee\xbf\xdf\x5e\xc2\x10\x8b\xf3\x5a\x5c\xf0\xe8\x2a\xe4\x29\x38\x38\x48\x34\x5f\x7e\xbf\xfb\x12\xaa\xa4\xc9\x2c\xcf\xef\xbb\x57\xb2\x80\xc3\xe9\xe5\xe1\xde\x62\x3d\xd4\xd9\xae\x49\xfe\x56\x27\x69\xb2\x87\x94\x8e\xe9\x45\x5f\x81\x69\x54\x45\x92\x36\x7a\xec\x23\x5d\xdf\x5f\x3b\x9c\xef\x48\x06\x50\x2f\x4e\xdc\x59\x1c\x9c\xfc\x06\xe7\xe9\x0f\x3e\x26\x58\xc4\x1d\x24\xc0\x84\xad\xaf\x83\x36\x68\xd9\xf9\x75\xbf\x50\x31\xd9\x3b\xeb\x95\x2b\xb2\x21\x12\x22\x71\x88\xd2\x89\x24\x0c\xf3\x1b\x52\x46\x8e\xe3\x6c\x30\x7d\xb6\x3a\xe3\x6f\xb7\x2c\x3f\xf3\xa8\x9e\x45\x7f\x45\xa1\x1b\xc9\x3a\x5c\x6d\x63\xb4\x10\x8e\x8e\xc6\x9b\x6c\xdc\xa4\x64\x63\xcc\xf8\x27\x1b\x7f\x34\x60\x9d\xb2\x4e\xa9\x34\x2b\xa6\xf8\xe4\x4a\x1a\x90\xb2\xe8\x28\xea\x5b\x00\x00\x00\xff\xff\x81\xf7\xdc\x13\xdd\x03\x00\x00")

func clientBootstrapDockerShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapDockerSh,
		"client/bootstrap/docker.sh",
	)
}

func clientBootstrapDockerSh() (*asset, error) {
	bytes, err := clientBootstrapDockerShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/docker.sh", size: 989, mode: os.FileMode(493), modTime: time.Unix(1517974712, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientBootstrapKeygenSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\x41\x6b\xdc\x30\x10\x46\xef\xf3\x2b\xbe\xb0\x81\x5c\x2a\xef\xbd\xa1\x0b\x6d\x13\xda\x3d\x74\x13\x68\x7a\x2a\xc5\x68\xad\x71\x34\xd8\x91\x8c\x66\xdc\xad\x2f\xfd\xed\xc5\xeb\xa5\x24\x64\x2f\xd1\x55\x6f\x78\x4f\x9a\xd5\xc5\x7a\x2f\x69\xad\x91\x68\x85\xfb\x92\xc3\xd8\xb0\xc2\x63\x18\xf7\xbd\x34\x6e\x28\xf2\xdb\x1b\xa3\xe3\xc9\x0d\x5e\x0a\x7c\x0a\xc8\xa3\x0d\xa3\x29\x2c\xf2\x89\x9b\xef\x2b\x22\x65\x83\x63\xa2\xed\x4d\x7d\x73\xfb\xfd\x61\xbb\xfb\xf8\xb0\xbd\xdb\x7d\xb8\xfc\x7a\xf7\xed\x76\x5d\xa9\xc6\xb5\x84\xba\xa8\xaf\x25\x71\x31\xf1\x75\xe0\xa1\xcf\x13\xdd\xff\xf8\x54\xbf\x71\xa6\x1a\xc6\xfd\x9c\xfc\x39\x72\xd3\x41\x5a\x04\x56\x93\xe4\x4d\x72\x42\x2b\x3d\xc3\xf7\x85\x7d\x98\xc0\x7f\x44\x4d\xdf\x93\xb4\xf8\x09\xd7\xe2\xf2\xa5\x09\xbf\xae\xe7\x87\x24\x02\x80\x23\x73\x71\xa4\x5e\x37\xbd\x20\xe7\xb3\xc2\xb6\x7d\xf6\x01\x08\x99\x35\xd9\x22\x7c\x87\x27\xdf\x31\xc4\xaa\xff\xb8\x6a\x74\x1d\x4f\x8f\x9c\xe0\xa6\x73\x25\x9b\x73\xd6\xe3\x78\x2b\xd7\xc4\xbd\x32\x2d\xda\x2f\x9c\xb8\x9c\xd6\x82\x83\x58\x44\xca\x18\xbc\xea\x21\x97\xb0\x08\x9f\xcb\x5e\x9b\x9c\xa1\xa8\x87\xdb\xe1\xea\x8a\x5a\x21\x3a\xe1\xda\xf8\x84\x47\xb1\x38\xee\xab\x26\x3f\x61\xb3\xc1\xdf\x65\x09\x5d\xca\x87\x54\xc7\xac\xa6\x44\x8d\xb7\xb3\xa9\xff\x02\x00\x00\xff\xff\x37\x00\x91\x4b\x4d\x02\x00\x00")

func clientBootstrapKeygenShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapKeygenSh,
		"client/bootstrap/keygen.sh",
	)
}

func clientBootstrapKeygenSh() (*asset, error) {
	bytes, err := clientBootstrapKeygenShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/keygen.sh", size: 589, mode: os.FileMode(493), modTime: time.Unix(1515346036, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientBootstrapTokenSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8f\x41\x4b\xc3\x40\x10\x85\xef\xfb\x2b\x9e\x54\xe8\x69\xbb\x77\x21\x07\x11\x31\x45\x6d\x84\x08\x5e\x84\xb0\x4d\x46\x77\x89\x9d\x8d\x3b\xb3\x16\xff\xbd\x44\x1a\xed\x71\xe6\x9b\xef\x31\x6f\x75\xe1\xf6\x91\x9d\x04\x63\x84\x14\x96\x8c\x59\xe1\x8e\x98\xb2\x57\x12\x78\x46\xdb\xd6\xd0\x34\x12\xe3\x2d\x65\x14\x21\x1c\xa3\x06\x5c\x3f\x6d\x91\xe9\xb3\x90\xa8\x6c\xce\x1c\x78\x0c\x9e\x0e\x89\x4f\x52\x91\xc8\xef\xb8\x79\xd8\x6e\x8c\x94\x21\x61\x48\xfd\x48\x19\xb9\x30\xac\xcd\x07\xbc\x1a\x00\xb0\x5f\xb8\xac\x9b\xc7\xdb\x2b\xe7\xa7\xc9\x85\x24\xba\x00\x9a\x1f\xe8\xee\x77\xcd\xcb\xae\xab\x9b\xf6\xb9\xad\xd6\x7f\x37\x6e\x23\x12\xdc\xc8\xe9\xc8\xdd\x3c\xcb\xfa\xdf\x9a\xd3\xaa\xdf\xcc\x65\x67\x89\x35\x7f\x4f\x29\xb2\x56\x91\x29\x6b\xf4\x27\x54\xf6\xfd\x87\x2f\xdc\x87\xc9\x0f\x6e\x41\xe7\x2d\xcc\x4f\x00\x00\x00\xff\xff\x5f\x07\x1a\x29\x27\x01\x00\x00")

func clientBootstrapTokenShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapTokenSh,
		"client/bootstrap/token.sh",
	)
}

func clientBootstrapTokenSh() (*asset, error) {
	bytes, err := clientBootstrapTokenShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/token.sh", size: 295, mode: os.FileMode(493), modTime: time.Unix(1515346036, 0)}
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
	"client/bootstrap/daemon-down.sh": clientBootstrapDaemonDownSh,
	"client/bootstrap/daemon-up.sh": clientBootstrapDaemonUpSh,
	"client/bootstrap/docker.sh": clientBootstrapDockerSh,
	"client/bootstrap/keygen.sh": clientBootstrapKeygenSh,
	"client/bootstrap/token.sh": clientBootstrapTokenSh,
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
	"client": &bintree{nil, map[string]*bintree{
		"bootstrap": &bintree{nil, map[string]*bintree{
			"daemon-down.sh": &bintree{clientBootstrapDaemonDownSh, map[string]*bintree{}},
			"daemon-up.sh": &bintree{clientBootstrapDaemonUpSh, map[string]*bintree{}},
			"docker.sh": &bintree{clientBootstrapDockerSh, map[string]*bintree{}},
			"keygen.sh": &bintree{clientBootstrapKeygenSh, map[string]*bintree{}},
			"token.sh": &bintree{clientBootstrapTokenSh, map[string]*bintree{}},
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
