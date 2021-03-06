// Code generated by go-bindata.
// sources:
// .gitignore
// Makefile
// NAME
// README.md
// app.go.tmpl
// config/config.yaml
// config.go.tmpl
// controller/controller.go.tmpl
// controller/gen_controller.go.tmpl
// controller/gen_route.go.tmpl
// controller/route.go.tmpl
// fresh.conf
// gen_types.go.tmpl
// main.go.tmpl
// model/all/all.go.tmpl
// model/database/db.go.tmpl
// model/database/table.go.tmpl
// router.go.tmpl
// DO NOT EDIT!

package genorm

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

var _Gitignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8c\xc1\x4a\xc3\x40\x10\x86\xef\xf3\x14\x03\xbd\x85\x76\xf7\x20\x28\xf4\x28\xf6\x20\x28\x8a\xf8\x00\xdd\x6c\xa6\x9b\x91\x49\x26\xec\xcc\x5a\xf3\xf6\x92\xea\xed\xe3\xff\x3f\xbe\x1d\x3e\xf2\x9c\x2a\x93\xe1\x45\x2b\x2e\x55\x4b\x4d\x93\x61\x9a\x07\x5c\xa4\x15\x9e\x0d\xba\x40\x3f\x04\x5d\x18\x44\xa0\x0b\xa6\x1b\xae\xc2\x3d\xc0\x0e\x3f\xc9\x1c\xfb\x2d\xb1\xee\xb1\x6f\x2c\x03\x5e\xd9\x47\x3c\x17\x45\xdf\xbe\x43\x3e\x43\x17\x36\xdc\xf4\xb7\xe6\x4b\x73\xd4\x0b\xfa\x48\x58\x14\xb3\x7e\x53\x4d\x85\xd0\x55\x65\x8f\xb6\x50\xe6\x0b\xe7\x24\xb2\xe2\x75\xa4\x19\x9b\xd1\x7f\xf2\x85\x9d\x9e\x9f\x4e\xd0\x05\x6d\xb7\xd8\x7b\xd5\x2f\xca\x7e\x10\xcd\x49\xb0\x08\x0f\x84\x39\xe5\x91\xf6\xf8\x71\x3a\xe2\xe8\xbe\xd8\x31\xc6\xc2\x3e\xb6\x3e\x64\x9d\xe2\x6b\x32\xa7\x3a\xf1\x3c\x58\xbc\xf9\x91\xcd\x1a\x59\x7c\xb8\xbb\x87\xf0\xb7\x00\x88\x96\x08\x3e\x2d\x11\x7e\x03\x00\x00\xff\xff\x51\x45\x48\x4d\x1e\x01\x00\x00")

func GitignoreBytes() ([]byte, error) {
	return bindataRead(
		_Gitignore,
		".gitignore",
	)
}

func Gitignore() (*asset, error) {
	bytes, err := GitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".gitignore", size: 286, mode: os.FileMode(420), modTime: time.Unix(1498529781, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _makefile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xf2\x73\xf4\x75\xb5\x4d\x48\x4e\x2c\x51\x00\xb1\x12\xb8\xb8\x92\x4a\x33\x73\x52\xac\xb8\x38\x1d\xd2\xf3\x15\xc0\x6c\x05\xdd\x7c\x85\xa4\xcc\x3c\x7d\x15\x0d\x90\x0a\x4d\x2e\x2e\xbd\x00\x0f\x7f\xbf\x48\x2b\x88\x2c\x17\x20\x00\x00\xff\xff\x34\xa1\xfe\xdc\x41\x00\x00\x00")

func makefileBytes() ([]byte, error) {
	return bindataRead(
		_makefile,
		"Makefile",
	)
}

func makefile() (*asset, error) {
	bytes, err := makefileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Makefile", size: 65, mode: os.FileMode(420), modTime: time.Unix(1497174442, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _name = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\xd6\xf3\x4b\xcc\x4d\xad\xad\xe5\x02\x04\x00\x00\xff\xff\xc5\x48\x82\xef\x0a\x00\x00\x00")

func nameBytes() ([]byte, error) {
	return bindataRead(
		_name,
		"NAME",
	)
}

func name() (*asset, error) {
	bytes, err := nameBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME", size: 10, mode: os.FileMode(420), modTime: time.Unix(1497174547, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _readmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xa8\xae\xd6\xf3\x4b\xcc\x4d\xad\xad\xe5\xe2\x32\xd4\x53\x78\x3a\x7f\xd7\xb3\x39\x6b\x5e\x2e\xdc\xf9\x7c\xf6\xba\xe7\xbb\x27\x3f\x9b\xd7\x62\xcd\x65\xa4\xa7\xe0\x1f\xe4\xab\xa0\xad\xe0\xe2\xe4\x9c\x98\x9c\x91\xaa\x11\x94\x9a\x92\x59\xac\xef\x9b\x9a\x9b\x0c\xe2\x6a\x5a\x73\x19\xeb\x29\x38\x06\x78\x3a\x16\xa5\x17\x5b\x73\x99\xe8\x29\x84\x54\x16\xa4\x16\x5b\x73\x01\x02\x00\x00\xff\xff\x0f\x48\xaa\x33\x5c\x00\x00\x00")

func readmeMdBytes() ([]byte, error) {
	return bindataRead(
		_readmeMd,
		"README.md",
	)
}

func readmeMd() (*asset, error) {
	bytes, err := readmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "README.md", size: 92, mode: os.FileMode(420), modTime: time.Unix(1497174569, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _appGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\xbd\x6a\xdc\x4e\x14\xc5\x6b\xcd\x53\x5c\x0b\x0c\x12\xe8\x2f\xf1\x6f\x0d\x5b\x2c\xc9\x62\x2f\x6c\x16\x23\xd9\x95\x31\x61\x56\x73\x35\x1e\x56\x9a\x2b\x66\xb4\x6b\x4c\xf0\x13\xe4\xab\x4a\x95\x2a\x5d\xde\xc0\x4d\xc0\x2f\x93\x0f\x77\x79\x85\x70\x47\xbb\xc4\x31\xa9\xd2\x89\x39\x67\xce\x39\xfa\x4d\x2f\xeb\xb5\xd4\x08\x9d\x34\x56\x08\xd3\xf5\xe4\x06\x48\x44\x14\xfb\x1b\x5f\xcb\xb6\x8d\x85\x00\x88\xb5\x19\xae\x36\xab\xbc\xa6\xae\x68\xb6\x2b\xaa\xd7\x05\x5a\xd5\xa2\xf7\xf1\x13\x55\x1b\xfb\x9f\x26\x6b\x6a\xfe\x7a\x2a\x2a\xea\x50\x13\x2b\x6b\x33\xf8\xa2\x33\x4a\xb5\x78\x2d\x1d\xc6\x22\xfa\x9b\x8f\x82\xad\x25\x1d\x8b\x54\x88\x66\x63\x6b\xb8\x92\xdc\x3b\xed\xfb\x24\x85\x57\x22\xd2\xc6\xe6\x15\x0e\x2f\x48\x61\x52\x37\x3a\x9f\xd9\x6d\x2a\x22\x47\x9b\x01\xe1\x68\x02\x2c\x2f\xf1\x3a\xd9\x9f\xe5\xe7\x1e\x93\xdf\xb5\xf9\x82\xb4\x46\x97\xa4\x7f\xe8\x7c\xa9\xc4\x9a\xb6\xe8\x6e\x46\x09\xb5\xf1\x43\xc9\x86\x24\xd8\x52\x11\x79\x74\x5b\x74\xdc\xb1\x03\xc1\x3d\x55\x38\x0c\x43\xa6\x4a\x39\xf4\x3e\x83\xfd\x85\xa2\x80\xef\x6f\x3f\x3c\xbc\xbb\x7b\xb8\xff\xf8\xe3\xf3\xeb\xf3\xaa\xfc\xff\xeb\xfd\xa7\x6f\xef\xef\x7e\x7e\x79\xe3\xb0\x25\xa9\x9e\x91\x6d\x8c\xde\x47\xe7\x95\xd1\x56\xb6\x27\x44\x6b\x7f\xb1\xef\x38\x2d\x67\x2f\xab\xf9\xf1\x72\xba\xb8\xbc\xd8\xbd\x4f\x5e\xcd\x8f\x39\xec\x12\x26\x20\xfb\x1e\xad\x4a\xfe\x31\x20\x03\x26\x3c\x72\x8d\x46\xce\x65\x18\x36\x06\x31\xc4\xdb\x54\x88\xc8\x34\xb0\x43\x0d\x07\x23\xe2\xe7\xb8\xda\x68\x7e\x83\x70\xb5\x25\x9d\xcf\x6d\x43\x4d\x12\x1f\xfa\x23\x58\x18\x3f\xa0\x35\x56\x83\xb4\x0a\x78\x1b\x7f\x9f\x9c\x9d\x9d\x02\x59\x38\xf4\x71\xc6\xbb\x97\xb2\xc3\x0c\x1e\x91\xe3\x36\x11\xa1\x0b\x8c\x77\x7f\x34\x46\x4d\xad\x0a\xa0\x79\x90\x69\x80\x2d\x07\x13\xb0\xa6\x7d\x34\x7c\x66\x15\xcb\x61\xcb\xcc\x39\x72\x61\x4c\x9c\xb1\x3b\x24\xdf\x8a\x5f\x01\x00\x00\xff\xff\x6b\xb2\x36\x5a\xf9\x02\x00\x00")

func appGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_appGoTmpl,
		"app.go.tmpl",
	)
}

func appGoTmpl() (*asset, error) {
	bytes, err := appGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "app.go.tmpl", size: 761, mode: os.FileMode(420), modTime: time.Unix(1498528928, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x4c\x49\x29\x4a\x2d\x2e\xb6\x52\x50\x32\x34\x32\xd7\x33\xd0\x33\xd0\x33\xb4\xaa\xae\xd6\x0b\xc8\x2f\x2a\xa9\xad\x55\xe2\x4a\xcd\x2b\xb3\x52\x50\x4a\x49\x4d\x2a\x4d\x57\xe2\xca\xc9\x4f\x8f\xcf\x49\x2d\x4b\xcd\xb1\x52\x50\x2a\x29\x4a\x4c\x4e\x55\xe2\x02\x04\x00\x00\xff\xff\x8b\x1d\xeb\xd8\x3f\x00\x00\x00")

func configConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_configConfigYaml,
		"config/config.yaml",
	)
}

func configConfigYaml() (*asset, error) {
	bytes, err := configConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/config.yaml", size: 63, mode: os.FileMode(420), modTime: time.Unix(1497174631, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xcd\xae\xd3\x30\x10\x85\xd7\x9e\xa7\x18\x2c\x81\x12\x54\x25\x82\x65\xa5\x2c\x10\x94\x55\xc5\x02\x89\x35\x75\x93\x89\x6b\xd5\xf1\x44\xb6\x53\xee\x55\xd5\x77\x47\x93\x1f\xe0\xb2\x81\xaa\x52\xe2\x39\x33\xc7\x5f\xce\x8c\xa6\xbd\x1a\x4b\x38\x18\x17\x00\xdc\x30\x72\xcc\x58\x80\xd2\x8e\x6b\xc7\x53\x76\x5e\x83\xd2\x9c\x34\x80\xd2\xd6\xe5\xcb\x74\xae\x5a\x1e\xea\x8e\x07\xb2\x5c\x5b\xbe\xba\x9c\x6a\xcf\x56\x83\x7a\x36\x83\x47\x6d\x79\xbc\xda\xca\x85\x5a\x8e\xd5\xed\xbd\x86\x12\xa0\xae\xf1\x23\x87\xde\x59\x6c\x97\x47\xca\x71\x6a\x33\xe4\xe7\x91\x36\x65\x29\xe1\x1d\xd4\x87\xae\x8b\x94\x12\x4a\xc9\x05\x8b\x27\xb1\xda\x6b\xb3\x94\xf5\x09\xd4\x21\xdc\x70\xfe\xbd\xec\xa0\x70\x13\xf5\xc8\xf6\x48\x37\xf2\x7f\xa9\x9e\xed\x77\x2f\x75\x7d\x82\xc7\xcc\x14\xc9\x74\xeb\xed\xf2\x8a\x26\x74\x38\x9a\x98\x68\xc5\xfc\xec\x3c\x41\x3f\x85\xf6\x8f\xce\x82\x9e\x5c\xc6\x33\xb3\x2f\xf1\xed\x3a\x7c\x07\x75\x4e\x3b\xa4\x18\x71\xdf\xe0\x92\x5b\xf5\x95\x4c\x27\x06\xc5\x6f\xaf\x12\x94\xeb\xe7\xb6\x57\x0d\x06\xe7\x65\x50\x79\xb6\xd5\x21\x46\x8e\x7d\xa1\x3d\x9b\x6e\x8b\xa8\x77\x9e\xb0\x37\xce\x53\xb7\xc3\x7c\x21\x1c\x4d\xbe\xa0\x4b\xf8\x3a\xe9\x1d\xbe\x30\x9d\x5d\x05\x4b\xfc\x14\xa7\xea\xf0\xe4\x72\xf1\x4e\x94\x07\xc8\x7f\xb5\xdc\x37\xf8\x66\x41\xbe\x3f\x40\x09\x47\x83\xf3\x9a\xbe\x85\xc1\xc4\x74\x31\xbe\x90\xef\x58\xba\xff\x09\x3b\x6d\x43\xbf\x88\x57\xd8\x19\x90\x62\xfc\x1f\xb2\x48\x79\x8a\x61\x75\x90\xb5\xac\x71\x4b\x12\x6b\xe0\xa5\x0c\xcb\xc5\x9f\xa8\x37\x93\xcf\x47\xb6\x96\x62\xb5\xac\xb8\x41\x51\xbe\xd0\x8f\xf9\x58\xb4\xbd\xad\xb6\xf5\x97\xf0\x80\x9f\x01\x00\x00\xff\xff\x9e\x07\xee\x86\xe0\x02\x00\x00")

func configGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_configGoTmpl,
		"config.go.tmpl",
	)
}

func configGoTmpl() (*asset, error) {
	bytes, err := configGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.go.tmpl", size: 736, mode: os.FileMode(420), modTime: time.Unix(1498528942, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _controllerControllerGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xc1\x4a\x03\x31\x10\x86\xef\xf3\x14\xc3\x9e\xda\xa2\xd9\x27\xf0\xa0\xde\xa5\x54\xef\x32\xcd\x4e\xc7\xd0\xee\x4c\x48\xb2\x62\x09\x79\x77\x59\x57\xf6\x20\xd8\xdb\xcf\xff\x7d\xcc\xcf\xf4\x3d\x0a\x2b\x27\x2a\x8c\xc7\x2b\x4a\xd0\x58\x20\x92\x3f\x93\x30\xd6\xea\xf6\x4b\x7c\xa1\x91\x5b\x03\x08\x63\xb4\x54\x70\x03\x88\x9d\x84\xf2\x31\x1d\x9d\xb7\xb1\x97\xa0\xf7\x62\x1a\xfc\x9c\xba\x3f\x70\xb0\x91\xc5\x66\x72\x0e\x25\xdf\xa6\x7d\xa4\x44\xe3\xbf\x92\xfd\x38\x17\x93\x0e\xb6\x00\x9f\x94\xf0\x1d\x1f\xf0\x62\xe2\xde\x12\x79\x3e\xad\xd5\xec\xb9\x03\xeb\xc0\xe9\x75\xf2\x9e\x73\x86\x72\x8d\x8c\x4f\x94\x79\x3f\x4f\xe0\x32\xe4\xd6\x02\xa0\xd6\x44\x2a\x8c\xee\x60\x53\xe1\xdc\x1a\x9c\x26\xf5\x38\x58\xad\x6e\xf9\x7f\xe3\x71\x27\x41\xdd\xb3\x69\xe1\xaf\x72\xf7\x7b\x04\x77\xab\xf1\x98\x24\x6f\xb1\x42\x83\x5a\x59\x87\xd6\xe0\x3b\x00\x00\xff\xff\x11\x3d\x20\x6c\x61\x01\x00\x00")

func controllerControllerGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_controllerControllerGoTmpl,
		"controller/controller.go.tmpl",
	)
}

func controllerControllerGoTmpl() (*asset, error) {
	bytes, err := controllerControllerGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controller/controller.go.tmpl", size: 353, mode: os.FileMode(420), modTime: time.Unix(1498535245, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _controllerGen_controllerGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\xc9\x6e\xdb\x30\x10\xbd\xf3\x2b\x1e\x04\x34\x90\x5b\x97\xbe\x1b\xc8\x21\xdb\x21\x40\x9b\x1a\x69\xee\x0d\x43\x8d\x59\xd6\x12\xa9\x72\x09\x6a\x10\xfc\xf7\x42\x8b\x65\xd9\x4d\xd1\xea\x34\x7c\xf3\x66\x79\x33\xa3\xd5\x0a\x8a\x0c\x39\x11\x08\x2f\x7b\x28\x6d\xda\xc0\x5a\x21\x77\x42\x11\x52\xe2\x9b\xc1\x7c\x10\x0d\xe5\xcc\x98\x6e\x5a\xeb\x02\x4a\x06\x14\x4a\x87\xef\xf1\x85\x4b\xdb\xac\x94\x36\x1f\x95\x35\x5a\x76\x56\x71\xe6\xac\x6c\x43\xca\x76\x9e\x9d\x0e\xfe\x6f\x5e\xfb\x4f\xe7\xaa\xb6\xaa\x60\x0b\xc6\x5e\x85\xc3\x37\x5c\xa2\xb6\x8a\x3f\x39\x21\x69\x3b\x41\x1d\x8f\x3f\x92\xa9\xc8\x7d\x8d\x52\x92\xf7\x8c\xa5\xe4\x84\x51\x04\xfe\x68\x63\x20\x9f\x33\x0b\xfb\xb6\x17\x37\xa8\xba\x72\xca\xc3\x07\x17\x65\x40\x62\xc0\xb5\xf0\xb4\x11\x4e\x34\x0c\x48\x09\x7a\x0b\x7e\xe5\x54\x6c\xc8\x04\x8f\x9c\x7b\x74\xc8\xd8\x0a\xe7\xe9\x6a\x73\x7f\x74\x1f\x99\x3d\x11\xc7\x2a\x9d\xf5\xb4\x6f\x3b\xeb\x79\x6b\x5d\xb3\x2e\x52\x8a\x5d\x9f\xb5\x36\x84\x91\x54\xe0\x87\xb7\xe6\x6d\xd7\x33\x52\x22\x53\x8d\x0d\x0c\x56\x66\x6c\x1b\x8d\xc4\xad\x9d\xea\x94\x12\xef\x95\x36\xfc\xc6\x9a\x40\xbf\xc2\xa2\x97\xd4\x76\x72\x3c\xd6\x97\xa7\xaa\x53\x97\x4b\xf2\x6b\x6d\xaa\xf2\x62\xe0\xf0\x49\xfd\xa2\xaf\x73\xa2\x3e\xe7\x33\xf6\x62\xd4\xf8\xdf\xe3\x18\x53\x6a\x7f\x1d\x75\x5d\xdd\x1b\x8c\x33\xe9\x51\x63\x03\xf8\x97\x36\x68\x6b\x44\x3d\x05\xa0\xdb\x40\x0c\xba\xf6\xfc\xde\xdf\x35\x6d\xd8\x97\x63\xab\x93\x96\x41\xe4\xe1\x9b\xdd\xc0\x9d\x73\xd6\x95\x72\x89\x8b\x11\xf4\x6d\x0f\x7d\x26\xef\x85\xa2\x79\x14\x6e\x6c\x45\xeb\x21\xb8\xe7\x74\xef\x43\xfb\x9f\x84\xdc\x2d\xe7\xe4\x31\xc1\x1a\x6f\xad\x0a\xda\xc3\xd1\xcf\xa8\x1d\x55\xc5\x3c\x2c\x2f\x66\x0f\x47\x21\x3a\x33\x01\xf9\x30\x99\x07\x1b\x7a\x95\xe0\xb7\xb4\x15\xb1\x0e\x39\x83\x6a\x4f\x27\x1a\xcf\x27\x80\x7e\xb3\x53\xc0\x49\xd6\xc3\xd1\x0c\xd3\xef\x5f\x73\xf0\xcf\xab\xc2\xec\xcf\x2a\x8b\xa1\xd4\x1a\xef\x3e\xbc\x16\x4b\x1c\xf7\x5e\xcd\x6f\x6e\x89\xe9\x22\x32\x3b\xe4\xf9\x1d\x00\x00\xff\xff\x2c\x06\xb4\x4f\x5c\x04\x00\x00")

func controllerGen_controllerGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_controllerGen_controllerGoTmpl,
		"controller/gen_controller.go.tmpl",
	)
}

func controllerGen_controllerGoTmpl() (*asset, error) {
	bytes, err := controllerGen_controllerGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controller/gen_controller.go.tmpl", size: 1116, mode: os.FileMode(420), modTime: time.Unix(1498535237, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _controllerGen_routeGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x51\xcb\x6a\xe4\x30\x10\xbc\xeb\x2b\x1a\xb1\x07\x7b\xd9\x95\xef\x81\x39\x84\x49\x48\x2e\x99\x98\x3c\x3e\x40\x63\xb7\x35\x62\x2c\xc9\xc8\x72\x1e\x34\xfd\xef\xc1\xf2\x24\x13\x02\xf1\x45\xea\xea\xae\xaa\x2e\xa4\xaa\x02\x83\x1e\xa3\x4e\x08\xfb\x77\x30\xd6\x0f\x49\x0c\xba\x39\x6a\x83\x40\xa4\xea\xa5\xdc\x69\x87\xcc\x82\x08\xfe\x98\x18\xa6\x61\xc6\x70\xb1\x81\x69\x18\x30\x6e\xb5\xc3\x1e\x54\xee\x9d\x48\xba\xef\xeb\x30\xa6\x99\xa2\x2e\x4f\x35\xb3\xb0\x6e\x08\x31\x41\x21\x00\xa4\xb1\xe9\x30\xed\x55\x13\x5c\x65\xac\xff\x6f\x82\xb7\xcd\x5c\xc9\x1f\xc3\x36\x38\x34\x61\x9e\x1c\x6d\x1a\x2b\x67\xdb\xb6\xc7\x57\x1d\x51\x8a\x52\x88\x17\x1d\x81\xe8\x1c\x8a\xf9\x21\x4c\x09\x6f\x66\x0c\x7f\x8d\xf5\x2a\xe3\x98\x1b\xa2\x9b\x7c\x03\x11\x8d\x1d\x53\x6e\x17\x71\x3e\x17\xde\xb5\x37\xd6\x63\x09\x24\xe0\x77\xc7\x0d\x64\x85\xca\xa8\x90\x44\x8b\x7d\x1d\xb1\xb3\x6f\xcc\xb2\x5c\x11\xab\xe7\x11\x8b\x73\x7c\xb5\x0d\xce\x05\x7f\xab\x7d\xdb\x63\x2c\xca\x2c\x05\xa2\xa8\xbd\x41\x58\x7c\x47\xe6\x35\x43\xa2\xfc\xfe\xa0\xee\x30\x1d\x42\xcb\x9c\x13\xd5\x3a\x1d\x98\xe5\x3f\xb8\x0a\x44\x6a\x11\x95\x40\x64\xbb\xaf\x7f\x61\x26\x02\xdb\x81\x1d\x77\x21\x7d\xca\x41\x0e\x61\x4c\x72\x7d\x65\x7d\xff\xf8\xb4\xba\x05\x7d\x3b\xbb\x7f\xbf\x04\x00\x0b\x16\x1f\x01\x00\x00\xff\xff\x84\x87\x93\x7b\x6b\x02\x00\x00")

func controllerGen_routeGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_controllerGen_routeGoTmpl,
		"controller/gen_route.go.tmpl",
	)
}

func controllerGen_routeGoTmpl() (*asset, error) {
	bytes, err := controllerGen_routeGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controller/gen_route.go.tmpl", size: 619, mode: os.FileMode(420), modTime: time.Unix(1498529013, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _controllerRouteGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xca\x31\x0a\x02\x31\x10\x46\xe1\x7e\x4e\xf1\xb3\x95\x2b\xb8\x7b\x0a\x5b\x91\xbd\x41\x36\x8c\xe3\x20\x99\x84\x38\x29\x24\xe4\xee\xb2\xd8\xd9\x3c\xbe\xe2\xad\x2b\x84\x8d\x6b\x70\xc6\xfe\x81\xa8\x15\xa7\x12\xe2\x2b\x08\xa3\xf7\xe5\xfe\xe3\x2d\x24\x1e\x83\x34\x95\x5c\x1d\x93\xa8\x3f\xdb\xbe\xc4\x9c\x56\x51\xbb\x48\x36\x8d\x87\x26\xa2\x47\xb3\x88\x8d\x45\xdf\xbe\xe5\xe6\x7c\xaa\x47\x71\x16\xb5\xe5\x6a\xa2\xc6\x33\x3a\x01\xf5\x7f\x99\x69\xd0\x37\x00\x00\xff\xff\x96\x6d\xa6\xc5\x8e\x00\x00\x00")

func controllerRouteGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_controllerRouteGoTmpl,
		"controller/route.go.tmpl",
	)
}

func controllerRouteGoTmpl() (*asset, error) {
	bytes, err := controllerRouteGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controller/route.go.tmpl", size: 142, mode: os.FileMode(420), modTime: time.Unix(1498465950, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _freshConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4b\xcc\xc9\x4c\x89\x4f\xad\x28\xb1\x52\x80\x01\xbd\xf4\x7c\x1d\x05\xbd\x92\x82\x1c\x10\x99\x0b\xa6\x32\x4a\x72\x41\x54\x65\x62\x6e\x0e\x57\x5e\x7e\x7c\x51\x6a\x52\x69\x66\x0e\x42\x1b\x6e\xc5\x99\xe9\x79\xf9\x45\xa9\x29\x08\xc3\x15\x14\x12\x8b\x8b\x53\x4b\x8a\x75\x14\x4a\x72\x0b\x74\x14\x72\xf2\xd3\xb9\x00\x01\x00\x00\xff\xff\xc8\x0e\x9b\x57\x83\x00\x00\x00")

func freshConfBytes() ([]byte, error) {
	return bindataRead(
		_freshConf,
		"fresh.conf",
	)
}

func freshConf() (*asset, error) {
	bytes, err := freshConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fresh.conf", size: 131, mode: os.FileMode(420), modTime: time.Unix(1498529775, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gen_typesGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8d\xc1\x4a\x03\x31\x10\x86\xef\xf3\x14\x3f\xcb\x1e\xf4\x92\xd2\x6b\xa1\x4f\x20\x8a\xa0\x78\xb5\x69\x33\x2c\xd5\x34\x09\xd9\xd9\xc8\x32\xcc\xbb\x4b\xba\x07\xc1\xdb\x30\xf3\x7d\xdf\xec\x76\x13\x27\xae\x5e\x38\xe0\xbc\x62\x2a\x42\xc5\x5f\xbe\xfd\xc4\x50\x75\xaf\xdb\x68\x46\x74\xbd\x95\x5c\x05\x0f\x04\x0c\x72\xbd\xf1\x40\x8f\x44\xcd\x57\x7c\xe2\x88\xbe\x70\x2f\xf9\x87\x48\xb5\xfa\x34\x31\xc6\x86\xc3\x11\x92\xdf\x72\x15\x0e\xcf\xbe\xc0\xbd\xaf\x85\x67\x33\x92\xb5\xf4\xf8\xd8\xdc\x13\xaf\x30\xc3\x2c\x75\xb9\x08\x94\x80\x3f\x7d\xff\xdf\x1f\x9b\xfb\xf0\x71\x61\x98\xdd\xc1\xb1\xed\x7b\xc0\x0c\xaa\x69\x89\xd1\x9f\x23\xf7\x17\xdd\xdd\x48\x33\x9c\xbe\xe6\x9c\x0e\x83\xea\x92\x02\xd7\x78\x4d\xdb\xf9\xee\x0d\x27\xa8\x72\x0a\x66\x64\xa4\x0a\x4e\xa1\xb7\x7f\x03\x00\x00\xff\xff\x19\x31\x59\xf6\x12\x01\x00\x00")

func gen_typesGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_gen_typesGoTmpl,
		"gen_types.go.tmpl",
	)
}

func gen_typesGoTmpl() (*asset, error) {
	bytes, err := gen_typesGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gen_types.go.tmpl", size: 274, mode: os.FileMode(420), modTime: time.Unix(1497164150, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mainGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x93\xcf\x6a\x1b\x31\x10\xc6\xcf\xd2\x53\x4c\x05\x86\xdd\xb2\xc8\x3d\x84\x1e\x02\x7b\x08\x6d\x5a\x02\x4d\x30\x49\xa1\x87\x34\x14\xd9\xfa\xe3\x21\x5a\x49\x48\x72\x52\x53\xf2\xee\x45\xd2\x26\x6b\xe2\x83\xe5\x91\x3e\xcd\x37\xf3\xf3\x28\x88\xdd\xa3\x30\x0a\x26\x81\x8e\x52\x9c\x82\x8f\x19\x3a\x4a\x98\xb6\xc2\xb0\xb2\x4e\xb9\x2c\xe8\xd7\xe8\x0f\x19\x6d\x09\x7c\x2a\xdf\x29\xc7\x9d\x77\x4f\xf3\x4f\x74\xa6\xed\x1e\xd3\x4e\x58\xcb\x28\x25\xcc\x60\xde\x1f\xb6\x7c\xe7\xa7\xb5\xf4\x93\x32\x7e\x6d\xfc\x23\xe6\xb4\xb6\xde\x30\xda\x53\xfa\x24\x62\x31\xdb\x69\x03\xe5\xf3\xf1\x8b\x77\x1a\x0d\x25\x22\x84\x1b\x31\x29\x68\x79\x29\x09\x28\xab\x00\x5d\x3e\xbd\x56\xd5\xdf\xd0\x2a\x80\x11\x58\x0b\xd7\x6d\xe1\x47\x31\x95\x5a\x03\xca\x26\x80\xaa\xb1\xde\xac\x45\x08\x3c\xa0\x64\x6f\x36\x55\x30\x02\xbb\xb9\xb8\xbe\xac\x65\xe9\x83\xdb\x01\x3a\xcc\x5d\x0f\xff\x28\x29\x76\x7f\x60\x04\x9f\xf8\xf5\xa3\xc4\xd8\x95\x34\x6c\xa8\xb1\x97\x6a\xa3\xe2\xd4\x53\xb2\x4d\x03\xa8\x18\xe1\x7c\x84\x46\x8a\xdf\x2a\x51\xcd\xbb\x13\x9f\x9e\x12\xd4\x55\xf7\x61\x04\x87\xb6\xe4\x27\xd6\x1b\x7e\x19\xa3\x8f\xba\x63\x51\x09\x09\xaf\xfd\xab\xb2\x39\x40\xde\x2b\x08\x22\xef\x01\x13\xac\x9a\x8d\x8f\xe7\xb0\x4a\x6c\x80\x93\xdc\xf5\xa0\xa7\xe4\x65\x01\x38\xce\x08\x13\xff\x19\x71\xea\x5a\xd0\x6d\x53\x3f\x00\xfb\xed\x58\x4f\x5f\xe6\x6e\x9f\x23\x66\xb5\x41\xd9\x3a\x2e\xbc\x47\x98\xff\x4a\xfe\x5d\xe5\x50\x4e\x28\x99\x1b\xfb\x55\xc4\xb5\xb3\x19\xef\x00\xf7\x0f\xdb\x63\x56\xdd\x3c\x13\xfc\x2a\x7b\x51\x0e\xfb\x7e\x80\x4f\x9f\xcf\xce\x16\xa3\xbd\x70\xd2\xaa\x4b\x37\x3b\xf9\xc4\x6f\xd5\xe4\x9f\xde\x52\xbd\x57\xde\x2a\xeb\x85\xbc\x43\xe3\x84\x6d\x57\x9c\x7a\x6e\x73\x52\x50\x17\x5c\x2d\xea\xb4\xb0\x69\xe6\xbb\x48\xde\x51\xbe\x72\xda\xeb\x8e\xad\xd2\x39\xc4\x9a\x18\xda\xb8\x2c\x20\x7b\x4a\xea\x3c\x8e\x4b\x12\x4a\x48\x13\xcf\x46\x15\xf1\x6b\x95\xe5\xe1\xb4\xba\xca\x8b\xe1\x1b\x11\x93\x2a\x8a\x05\x28\xad\xce\x5f\x95\x16\x07\x9b\x7f\x78\x63\x54\xe4\x77\x2a\x6f\xa2\xd2\xf8\xb7\xd3\x53\xe6\x77\x21\xa2\xcb\xba\x63\xf7\x2b\xf9\x00\xf7\xab\xf4\xc0\x06\x08\x28\x97\xa2\x7a\x4a\xe7\xaa\x4e\x1a\xce\xf1\x50\xca\x95\x4a\xab\x08\x27\xdd\x29\x27\x59\x4f\x49\xe3\x77\x11\x42\x57\x98\xfe\x0f\x00\x00\xff\xff\xd6\xf6\x2e\xec\xe9\x03\x00\x00")

func mainGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_mainGoTmpl,
		"main.go.tmpl",
	)
}

func mainGoTmpl() (*asset, error) {
	bytes, err := mainGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "main.go.tmpl", size: 1001, mode: os.FileMode(420), modTime: time.Unix(1498528957, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _modelAllAllGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8c\x31\x8e\xc2\x30\x10\x45\xfb\x39\xc5\x57\xb4\x45\xd2\xd8\xfd\x4a\x5b\x6c\x48\x8d\x10\x37\x30\xc9\x60\x22\x1c\x3b\x32\xa6\x40\xa3\xb9\x3b\x22\x81\x54\x94\x33\xff\xbd\x67\xad\xe7\xc8\xd9\x15\x1e\x70\x7a\xc0\x8f\x71\x2e\x34\xbb\xfe\xea\x3c\x63\x4a\x03\x07\xa2\x71\x9a\x53\x2e\xa8\x09\x10\xc1\x4f\x4e\xa9\x1c\x5c\xb9\xe0\xf7\x0f\xe6\xf8\x39\x54\xd7\x39\xbb\xe8\x19\xa6\x6b\xb1\xbe\x2a\x91\xcd\x50\xb5\x4b\xd2\x8a\x98\xbd\x9b\x58\xb5\x82\x08\xc7\x41\x95\x1a\xa2\xf3\x3d\xf6\xd8\x85\x74\xe3\xff\x10\xea\x06\xf2\xbd\xb8\xc9\x66\x61\xbb\xf6\x85\xbe\x2b\x4a\xcf\x00\x00\x00\xff\xff\x61\xa1\xa3\x6c\xd1\x00\x00\x00")

func modelAllAllGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_modelAllAllGoTmpl,
		"model/all/all.go.tmpl",
	)
}

func modelAllAllGoTmpl() (*asset, error) {
	bytes, err := modelAllAllGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "model/all/all.go.tmpl", size: 209, mode: os.FileMode(420), modTime: time.Unix(1497967059, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _modelDatabaseDbGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x94\x4f\x6b\xe3\x3a\x14\xc5\xd7\xd2\xa7\xb8\xcf\x3c\x8a\xfd\x48\xe5\xb7\x0e\x74\x31\x89\x87\xa1\xb4\x99\xce\x90\xc2\x2c\x8b\x6c\x5f\xbb\xa2\xb2\xe4\xc8\x72\x69\x10\xfa\xee\x83\xfc\xa7\x71\x87\x74\xa0\xfb\xae\x62\xdd\x73\xcf\xb1\x74\xfd\x53\xd2\x14\x6a\x54\x68\xb8\xc5\x12\xf2\x23\xd4\x42\xb5\x96\xb6\xbc\x78\xe2\x35\x82\x73\xbd\x2a\xd1\x48\xa1\x10\xd8\x77\xde\xa0\xf7\xd4\x39\xf8\xd7\xf2\x5c\x22\xac\xaf\x80\xdd\x0f\x4f\xde\x53\x2a\x9a\x56\x1b\x0b\x31\x25\x51\xc9\x2d\xcf\x79\x87\x69\x77\x90\x11\xa5\x24\xaa\x85\x7d\xec\x73\x56\xe8\x26\x2d\x75\x83\xb5\x4e\x6b\x6d\xda\x88\x92\x07\x58\x6a\xb5\xbe\xec\x0e\xf2\xb2\x34\xe2\x19\x4d\xda\x1c\x07\xfb\x79\xf7\x93\xb0\x5d\x2a\x75\xfd\x8e\x2e\xd4\xd0\x50\xe6\x11\x4d\x28\x7d\xe6\x26\xec\x2b\xdb\x6c\xb5\xaa\x44\x0d\xff\x95\x39\x9b\x17\x94\x92\x32\xbf\x56\x95\x26\x64\x2c\x87\x67\x4a\x3a\x34\x82\xcb\x71\xb5\x14\x12\x4a\xab\x5e\x15\xf0\x0d\x6d\xb6\x89\x13\x38\x49\xe0\x28\x11\x15\x8c\x61\xf0\xcf\x15\x28\x21\x43\x8d\x18\xb4\xbd\x51\x93\x40\x89\xa7\x64\x9e\xcf\x0a\xd0\x98\x30\xc6\xee\x20\xd9\x5d\x8b\x2a\x8e\xc6\x43\xaf\x60\xde\x1e\xdb\xeb\xde\x14\x18\x27\xc9\x90\x1e\xfa\x17\xd1\x2d\x57\xa2\x88\xd1\x98\x64\xc8\x15\xd5\xc9\xb7\xe3\x2f\xd7\xa5\xc4\xad\x56\xaa\x0b\x96\xff\x07\xc3\xfc\x66\xb6\x47\xbb\xec\x88\xcf\xfa\xce\xa6\x86\x7d\xfe\x3d\xf5\xb5\x23\x3e\xeb\x1b\x53\xcb\xbc\xe1\x6d\x38\xfb\x45\x40\x81\x65\xf9\x8e\xb7\x2e\xcb\xd7\x70\x1a\x4e\x26\xb8\xc4\xc2\xae\x61\xe8\xd8\x1d\xf7\x3f\x6f\xa7\x92\xf3\x9e\x12\xe7\xc0\x70\x55\xe3\x04\xe1\x0e\x2d\xef\xbc\x07\xe7\x40\x54\xd0\x70\x5b\x3c\xce\xa0\x8e\x0d\x23\xbf\x30\xbc\x98\x7d\x29\xcb\xa1\xf8\x4b\xd8\xc7\x20\xc4\xce\x41\xdf\xb6\x68\xb6\xbc\x41\xf9\xc6\xe1\xfc\x0a\x22\xe7\x96\xa5\x28\x09\x47\xbd\xc1\x63\x17\x3b\x17\x3e\xcb\x01\xd8\x0f\x23\x1a\x6e\x8e\x37\x78\xfc\xfa\x62\x0d\x87\x88\xf7\x56\x3f\x08\x55\x18\x6c\x50\xd9\xc8\x7b\x6b\x7a\x74\x0e\x65\x87\xde\x57\x5c\x76\x61\xa1\x4a\x3f\xc5\x9f\xfc\x21\x9f\x12\x32\xa9\x30\xfd\x4e\x23\x63\xf7\x86\x17\x78\xa7\xe2\x28\x5a\x81\xd4\x35\xcb\xb0\xe2\xbd\xb4\xb7\xba\xae\xd1\x24\xaf\x38\xc3\x15\x5c\xbc\xb2\xe9\xb2\xcd\x7a\x31\xd7\x30\xeb\xf5\x38\x08\x4f\xff\xe0\xd3\x9f\x00\xdf\x4f\x77\xe0\x1c\xe6\xcb\xfb\x71\x06\xf6\xa5\xfc\x71\xe4\x07\xf3\x27\xf8\x9f\xe0\x7f\x0c\xfc\x37\x4c\x7e\x14\xff\x37\xc4\xce\x97\x60\x2b\x75\x87\xc3\x05\x78\xef\xcf\x7d\xac\xb0\x6c\xc3\x86\xde\x78\xfe\xc6\xd3\x1e\x94\x90\xd4\xd3\xdf\x01\x00\x00\xff\xff\x69\xaa\x17\x12\x65\x07\x00\x00")

func modelDatabaseDbGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_modelDatabaseDbGoTmpl,
		"model/database/db.go.tmpl",
	)
}

func modelDatabaseDbGoTmpl() (*asset, error) {
	bytes, err := modelDatabaseDbGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "model/database/db.go.tmpl", size: 1893, mode: os.FileMode(420), modTime: time.Unix(1498529049, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _modelDatabaseTableGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\x6d\x4f\xe3\x46\x10\xfe\xee\x5f\x31\x5d\xa1\x93\x7d\xa4\x1b\x5d\x5b\xf5\x03\x52\x54\x29\x2f\xd0\xa8\x90\x3b\x41\x4e\xa7\xaa\xaa\xc8\xda\x1e\x87\x05\x7b\xd7\xac\xd7\x40\x64\xf9\xbf\x57\xbb\x9b\x18\x07\x12\xce\x7d\x93\xa0\x2a\x5f\x70\x66\xd6\xe3\x79\x9e\x67\x76\x66\xb7\xdf\x87\x25\x0a\x54\x4c\x63\x0c\xe1\x0a\x96\x5c\xe4\xda\xcb\x59\x74\xc3\x96\x08\x55\x55\x8a\x18\x55\xca\x05\x02\x1d\x0f\x67\x2c\xc3\xba\xf6\x3c\x9e\xe5\x52\x69\xf0\x3d\x00\x12\x33\xcd\x42\x56\x60\xbf\xb8\x4d\x89\x31\x68\x9e\xa1\x7d\x48\x32\x6d\xff\x87\x2b\x8d\x05\xf1\xcc\xe3\x92\xeb\xab\x32\xa4\x91\xcc\xfa\xb1\xcc\x70\x29\xfb\x4b\xa9\x72\xb2\xc7\xc5\xc5\x0d\xd7\x45\x3f\x0e\xf7\x2d\x90\xc6\x4f\xbc\xc0\xf3\xee\x98\x82\x4b\x18\x80\xf9\x38\x9d\xc9\xfb\xc6\x90\x64\x9a\x7e\x52\x5c\xe8\x54\x34\xb6\xe2\x36\xa5\x13\xa5\x66\xf2\x5c\xde\x17\x8d\xd5\xa6\x49\x67\x78\x3f\x2c\x93\x04\x55\x63\x2f\x35\x4f\x0b\x3a\x15\xfa\xfb\xef\x3c\x4f\xaf\x72\xcb\x4a\x9e\xa3\x1a\xb1\x0c\x53\xa0\x73\x16\xa6\xe8\x98\x81\x42\xab\x32\xd2\x50\x79\x00\x55\x05\x8a\x89\x25\xc2\x01\xef\xc1\xc1\x1d\x1c\x0d\x80\x8e\x64\x5a\x66\xa2\x80\xba\xb6\x0b\x5a\x51\x0e\xee\xe8\x31\xc7\x34\xae\x6b\xa8\xaa\x83\x3b\x7a\x22\xe7\xab\xdc\x44\x5c\xc4\xe1\x11\xb1\xa6\xb5\x9f\xc0\x75\x21\xc5\x13\xdb\x02\xaa\x0a\x45\x6c\xb4\x01\xb8\x8c\x43\x78\x1f\x87\x74\x3c\x9c\x8a\x44\x82\xf9\xb3\x51\xbe\x25\x0b\xe3\xd5\x0f\xd6\x3b\x57\x4c\x14\x2c\xd2\x5c\x8a\x96\xbf\xf6\xbc\xa4\x14\x11\xcc\xf0\x7e\x2f\x4a\x3f\x80\xf7\xfb\x29\x30\xd8\x65\x78\x6d\xf0\xbe\x14\xe4\x0b\xd7\x57\xb2\xd4\xe3\xa1\x1f\xb8\x17\xe8\x05\x6a\x97\xb1\xb5\x28\xd4\xa5\x12\xc6\xd1\x29\xa7\x56\xb8\x8e\xc9\xbd\xdb\xbb\xca\x2c\x32\xf2\x38\xf9\x36\xa2\x59\xcd\x9c\x83\x27\x40\xc7\x98\xb0\x32\xd5\xd0\x98\x8d\xa3\x40\xbd\xb1\x53\x68\xbd\xe0\xa4\xd9\x7e\xae\x77\x83\xf4\x4d\x7a\xfb\x11\x04\xd0\xa2\xc9\xc2\xd9\xec\x3f\x83\xe9\x04\xdb\x84\x9a\x3a\x18\xc0\xbb\xc7\x4a\x70\xc0\xc6\xc3\xa3\xe6\x25\x3a\x1e\xf6\x9c\x31\x3c\x63\x79\xdb\x6e\x7e\xf7\x6c\x9a\x7f\x26\xb3\x56\x51\xf9\xcf\xeb\x2c\xd8\xf0\x4f\x4d\x11\x0e\x40\x3f\x74\x0f\x7e\x82\xfa\xe2\x36\x9d\x3c\x60\x54\x6a\xa9\xfc\x00\x4c\xe3\xa0\x2d\x93\x8d\xcd\x93\x26\xfc\x37\x03\x10\x3c\x5d\x63\x7e\x24\xda\xf8\x9e\xb1\x6f\xa8\xea\x9e\xca\x54\x14\xa8\xf4\x0e\xfa\x4d\xa4\xa7\x79\xb6\x3b\xc1\xf6\xfe\xb7\x0e\x9e\x00\xde\x82\xdb\xc7\x40\x22\x85\xa6\x0f\x5f\x32\x4d\xda\xe5\xb3\x59\x35\x2b\xd3\x14\xc8\xaf\x93\x0b\xb2\x55\x75\x6b\xaf\x6b\x19\x40\xb8\xe9\x54\xa4\x5d\x98\x8e\xf3\x6d\x44\x9b\x6e\x33\x00\xbb\xde\x77\x5d\xee\x04\xf5\xe7\xf9\x68\xce\x33\xf4\x03\xfa\x59\xf0\x07\x3f\x08\x5a\x5f\xc2\xb4\xc0\xce\x71\x9b\x88\xa7\x32\x62\xa9\x8b\xd9\x8e\xd5\xec\x89\x1d\x91\x5f\x2d\xa8\xcd\x7c\x79\x09\xca\x8e\xcd\xbe\x43\xeb\x32\x8f\xff\xd7\xfa\x55\x83\xfa\x8b\x5a\x3f\x3e\xa1\x52\xa6\x2d\x34\x4d\x75\xdd\x38\x64\x78\x1d\xb8\x5e\x65\x16\x6c\xf5\xa9\x9c\x09\x1e\xf9\xa8\x54\xb0\xd5\x7a\x4f\x50\xef\x9f\x7c\x57\xa8\xd0\x8f\xa4\x88\xcd\xe9\x83\x8b\x65\x0f\x98\x5a\x16\x40\x29\xe5\x42\xa3\x4a\x58\x84\x55\x1d\xc0\x6f\xbf\x7f\x7d\x2e\x16\x26\xdd\x97\x16\x56\xf5\xde\x89\x73\xd9\x7b\x86\xf7\x02\x53\x8c\xb4\xff\xce\x44\xee\x01\xb9\x98\x9c\x4e\x46\x73\xa8\xaa\x25\x6a\x1b\xd3\x52\x6d\x02\x17\xad\x29\x0b\xc7\xe7\x1f\xcf\x60\x51\x55\xed\xef\x2e\xe0\xcb\xcf\x93\xf3\x09\x10\x38\x04\x03\xd5\x61\xa4\x94\x76\x24\x12\x20\x91\xca\xa4\xb8\x1e\xfe\xae\x1f\x5b\xc4\x6e\xfd\xf3\x03\xc8\xf6\x84\x28\x3a\x69\x31\x92\xa5\xd0\x1d\xb4\xe0\x42\xff\xf8\xc3\x0b\xd3\x3b\x12\x7a\x1f\x9b\x53\xa1\xfd\x0d\x93\x91\xfd\xdc\x87\xe0\x5f\xa2\x6c\x8d\x3e\x12\xba\x13\xf8\x63\xae\x8a\x2e\xe0\xff\xfe\xf9\x6c\x7f\x11\xee\xe6\xec\xa3\x40\xb3\xe9\xfe\xa1\x12\x3c\xb4\x10\x0f\x81\xc0\xe9\xf4\x6c\x3a\x87\x0f\xe4\xab\xd4\x3a\x9b\xb9\x73\xd8\xa3\xcb\xe0\xc9\x15\xa4\x71\x54\xeb\xc6\xb2\x66\x5e\xf0\xd4\x1a\xea\xdd\xfa\x74\x3a\x34\xbf\xa4\xd8\x70\x65\xe1\xfb\x82\x65\xd8\x68\x96\xd8\xf1\xb4\xa5\xd8\xdb\x16\x6c\x61\x76\x80\x85\x78\x08\x64\x31\xf8\x89\xac\x31\xbe\x4d\xb1\xcc\x45\x67\x24\x45\xcc\xed\x19\xdb\xc0\xea\xc1\xd6\x9e\xfb\xcf\xeb\xd7\xea\x68\x6f\x51\x47\xff\x06\x57\x50\x55\xf4\x93\xe2\x19\x53\xab\x5f\x70\xe5\x6e\xfa\x6f\x5e\xa6\x2d\x4c\x75\xed\x76\xda\x0d\xae\x5e\x95\x3e\x7f\x04\x00\x00\xff\xff\xc4\x16\x32\x46\xef\x12\x00\x00")

func modelDatabaseTableGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_modelDatabaseTableGoTmpl,
		"model/database/table.go.tmpl",
	)
}

func modelDatabaseTableGoTmpl() (*asset, error) {
	bytes, err := modelDatabaseTableGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "model/database/table.go.tmpl", size: 4847, mode: os.FileMode(420), modTime: time.Unix(1498529071, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _routerGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x14\xc6\xc1\x09\x82\x31\x0c\x06\xd0\x7b\xa6\x08\x3d\xa9\x60\x9d\xc2\x05\xdc\xa0\x86\x18\x3f\xa4\x49\xa9\xe9\x49\xdc\xfd\xa7\x97\xc7\x1b\x4d\x3e\xcd\x94\x7b\x83\x13\xa1\x8f\x98\xc9\xc5\x90\xef\xf5\xac\x12\xfd\x66\xf0\xab\x85\x43\xf6\x0a\xd1\x6b\xb9\xf0\x54\xc3\x37\x1f\xb1\x52\x4f\x73\xcb\x17\x83\xd7\xbb\x1b\x5c\xcf\xfc\xa3\x3f\x1d\x01\x00\x00\xff\xff\x5c\x5c\x62\x07\x59\x00\x00\x00")

func routerGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_routerGoTmpl,
		"router.go.tmpl",
	)
}

func routerGoTmpl() (*asset, error) {
	bytes, err := routerGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "router.go.tmpl", size: 89, mode: os.FileMode(420), modTime: time.Unix(1497324800, 0)}
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
	".gitignore": Gitignore,
	"Makefile": makefile,
	"NAME": name,
	"README.md": readmeMd,
	"app.go.tmpl": appGoTmpl,
	"config/config.yaml": configConfigYaml,
	"config.go.tmpl": configGoTmpl,
	"controller/controller.go.tmpl": controllerControllerGoTmpl,
	"controller/gen_controller.go.tmpl": controllerGen_controllerGoTmpl,
	"controller/gen_route.go.tmpl": controllerGen_routeGoTmpl,
	"controller/route.go.tmpl": controllerRouteGoTmpl,
	"fresh.conf": freshConf,
	"gen_types.go.tmpl": gen_typesGoTmpl,
	"main.go.tmpl": mainGoTmpl,
	"model/all/all.go.tmpl": modelAllAllGoTmpl,
	"model/database/db.go.tmpl": modelDatabaseDbGoTmpl,
	"model/database/table.go.tmpl": modelDatabaseTableGoTmpl,
	"router.go.tmpl": routerGoTmpl,
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
	".gitignore": &bintree{Gitignore, map[string]*bintree{}},
	"Makefile": &bintree{makefile, map[string]*bintree{}},
	"NAME": &bintree{name, map[string]*bintree{}},
	"README.md": &bintree{readmeMd, map[string]*bintree{}},
	"app.go.tmpl": &bintree{appGoTmpl, map[string]*bintree{}},
	"config": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{configConfigYaml, map[string]*bintree{}},
	}},
	"config.go.tmpl": &bintree{configGoTmpl, map[string]*bintree{}},
	"controller": &bintree{nil, map[string]*bintree{
		"controller.go.tmpl": &bintree{controllerControllerGoTmpl, map[string]*bintree{}},
		"gen_controller.go.tmpl": &bintree{controllerGen_controllerGoTmpl, map[string]*bintree{}},
		"gen_route.go.tmpl": &bintree{controllerGen_routeGoTmpl, map[string]*bintree{}},
		"route.go.tmpl": &bintree{controllerRouteGoTmpl, map[string]*bintree{}},
	}},
	"fresh.conf": &bintree{freshConf, map[string]*bintree{}},
	"gen_types.go.tmpl": &bintree{gen_typesGoTmpl, map[string]*bintree{}},
	"main.go.tmpl": &bintree{mainGoTmpl, map[string]*bintree{}},
	"model": &bintree{nil, map[string]*bintree{
		"all": &bintree{nil, map[string]*bintree{
			"all.go.tmpl": &bintree{modelAllAllGoTmpl, map[string]*bintree{}},
		}},
		"database": &bintree{nil, map[string]*bintree{
			"db.go.tmpl": &bintree{modelDatabaseDbGoTmpl, map[string]*bintree{}},
			"table.go.tmpl": &bintree{modelDatabaseTableGoTmpl, map[string]*bintree{}},
		}},
	}},
	"router.go.tmpl": &bintree{routerGoTmpl, map[string]*bintree{}},
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

