// Code generated by go-bindata.
// sources:
// template/data/footer.html
// template/data/header.html
// template/data/index.html
// template/data/not_found.html
// DO NOT EDIT!

package template

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

var _templateDataFooterHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x80\x02\x1b\xfd\x94\xcc\x32\x3b\x2e\x08\x33\x29\x3f\xa5\xd2\x8e\xcb\x46\x3f\xa3\x24\x37\xc7\x8e\x0b\x10\x00\x00\xff\xff\x1a\x69\xbe\x97\x23\x00\x00\x00")

func templateDataFooterHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templateDataFooterHtml,
		"template/data/footer.html",
	)
}

func templateDataFooterHtml() (*asset, error) {
	bytes, err := templateDataFooterHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/data/footer.html", size: 35, mode: os.FileMode(420), modTime: time.Unix(1516623186, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateDataHeaderHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x55\x4d\x6f\xe3\x36\x10\xbd\xf7\x57\x4c\xd5\x8f\xd8\xe8\x8a\x4c\x36\xdb\x2e\x10\xcb\x5e\x04\xdd\xee\xa2\x80\xd1\x43\x8b\x9c\x8a\x22\x18\x51\x63\x89\x0e\x4d\x0a\xe4\xc8\x8e\xa1\xf5\x7f\x2f\x24\xc7\xb6\x2c\x2b\xed\xea\x90\x90\xe3\x37\xf3\x1e\xe7\x8d\xa8\xa4\xe0\x95\x99\x7d\x03\x00\x90\x14\x84\xd9\x7e\xd9\x6e\x59\xb3\xa1\xd9\x83\x37\x10\x0a\xe7\x99\x7c\x22\xf7\xa1\x13\xc4\x68\xfb\x04\x9e\xcc\x34\x0a\xbc\x35\x14\x0a\x22\x8e\xa0\xf0\xb4\x98\x46\x05\x73\x19\xee\xa4\x5c\xe1\xb3\xca\xac\x48\x9d\xe3\xc0\x1e\xcb\x66\xa3\xdc\x4a\x1e\x03\xf2\x56\xdc\x8a\xf7\x52\x85\x70\x8a\x89\x95\xb6\x42\x85\x10\x75\xc8\xbe\x8d\x63\xf8\x6c\x5c\x8a\x06\x82\x66\x02\xc6\x1c\x46\x39\x63\x2e\x96\x61\x0c\x31\x7c\x76\x2e\x37\x04\xf7\x16\xcd\x96\xb5\x0a\x10\xc7\x9d\xec\xa0\xbc\x2e\x19\x30\x6c\xad\x82\xe0\xd5\x49\xe0\x66\xb3\x11\x79\x9b\xcb\x98\xaf\xd0\x62\x4e\xbe\x55\xd8\xd4\x96\xcb\xf0\x41\x67\xd3\x87\xfb\xf8\xe6\xe6\xed\xfb\x77\xef\x6e\xaf\x7f\x89\x6f\xa2\x59\x22\xf7\xf5\x2e\x08\x4e\x81\xe6\xd9\x68\x9b\xb9\x8d\xc8\x90\x71\x8e\x5b\xf2\x30\xbd\x0c\x7d\xf9\x02\x7f\xff\x33\x39\x4b\x5b\x54\x56\xb1\x76\x16\x1a\x05\xa3\x71\x7d\x04\x8b\xb2\x0a\xc5\x08\x7d\x5e\xad\xc8\x72\x18\x4f\x76\x67\x79\x2d\xfc\x6a\x19\xae\xde\x80\xa5\x0d\x7c\x44\xa6\xd1\x78\x3c\x19\xc0\x28\x67\x17\x3a\xbf\x7a\x03\x57\xe7\x27\xbb\xea\xa0\x5f\x3d\xe3\xbe\x7d\x32\x30\xb2\x56\x72\x19\xa4\xb6\x19\x3d\x8b\x65\xf8\xea\xbe\x1c\x0f\xe8\x29\xd3\x9e\x14\xdf\x57\x5c\x8c\xc6\x50\x9f\xa1\x3a\x1d\x34\x4e\x61\x93\x20\x3c\x95\x06\x15\x8d\xce\xdc\x5b\xa0\xa2\xd4\xb9\xa7\xd6\xb5\xf5\x5b\x71\x73\x2d\x33\x8d\xc6\xe5\xd2\x61\xc5\xc5\x87\x08\x7e\xba\x28\xdc\x3c\x91\x32\x9a\x2c\x3f\xea\x6c\x5a\xd7\xf0\xbd\xf8\xb5\xed\x8a\xf8\xf4\x52\xee\xbe\x2c\x7f\xff\x08\xbb\xdd\xab\xe9\x3f\x1e\xe4\x3f\x56\x5e\x0f\x96\x98\xbb\x5c\xdb\x87\x3f\xe7\xff\x53\x25\x94\xce\x06\x7a\xe4\x6d\x49\x53\xe5\x32\x8a\x2e\x90\x3d\x17\x77\xaf\xb8\x94\xc8\xd3\x4b\x9c\xa4\x2e\xdb\x76\x8c\xb0\xb8\x06\x65\x30\x84\x69\x64\x71\x9d\xa2\x87\xfd\xbf\x58\xdb\x35\xf9\x40\xd1\xb9\x47\x49\xa6\x8f\x78\xe5\x2c\xa3\xb6\xe4\xe3\x85\xa9\x74\xd6\x43\xf6\xd1\x2f\x65\x1b\x25\xe4\x07\xb0\x2d\x1e\x7b\xe8\xd4\xa3\xcd\x0e\x77\xc7\x77\x51\x7b\xef\xfc\xd5\xdc\x3b\xb6\xb9\x78\x70\x80\x51\x66\x7a\xdd\x93\x5c\x99\x4e\xd5\xc3\xf1\x3a\x4b\xaf\xf3\x82\x3b\x8a\xea\x1a\xf4\x02\xc4\x43\x20\x0f\xbb\xdd\x25\x85\xd1\x3d\x95\x4c\xcf\x1c\xcd\x5a\xa7\x9b\x24\xf1\x49\xfb\xc0\x7f\xe0\x8a\x60\xb7\x83\x53\x78\x8e\xc7\x68\x22\x8d\x1e\x10\x6f\xf4\x2c\x49\x2b\x66\x67\x0f\x0c\x29\x5b\x48\xd9\xc6\x19\xda\x9c\x8e\xde\xa4\x6c\x23\x70\x56\x19\xad\x9e\xa6\x91\x71\xb9\xab\x78\x34\x8e\x66\xf3\x76\x95\xc8\x7d\x8d\xd9\x39\x4b\x5d\x03\x99\x40\xaf\x1c\x69\x88\x78\x90\xed\xfc\xe5\x6c\x39\xb5\x85\x8d\xe6\x02\x0e\xf3\xfd\x1f\x02\x6c\xd6\xe7\x4f\x64\x65\x7a\x86\x9d\x7b\x98\x48\x8b\xdd\xed\xd0\x04\x46\xd0\x7e\x6a\xa6\xd1\x46\x67\x5c\xdc\xc1\xcf\xd7\x3f\x4c\x80\x5d\x79\x07\xb7\xcd\xaa\x74\x41\x37\x17\xc5\x5d\xf3\x59\x42\xd6\x6b\x9a\xf4\x26\xb0\xae\xf7\x47\x10\xbf\x79\xef\x7c\xb8\x10\xd9\x21\x45\x43\x9e\xa1\xfd\xfb\x62\xcb\xc0\x34\xd7\x35\xf8\xe6\x37\x10\x83\xfd\x2e\x3d\x35\xf3\x22\xda\x51\x68\x36\x43\x05\x06\x9b\x75\x31\xde\x27\xe0\xbf\x01\x00\x00\xff\xff\xee\x88\xb2\x65\xb6\x07\x00\x00")

func templateDataHeaderHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templateDataHeaderHtml,
		"template/data/header.html",
	)
}

func templateDataHeaderHtml() (*asset, error) {
	bytes, err := templateDataHeaderHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/data/header.html", size: 1974, mode: os.FileMode(420), modTime: time.Unix(1516623203, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateDataIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x52\xc1\xae\xda\x30\x10\xbc\xf3\x15\x5b\x4b\x4f\x6a\x0f\xc1\x95\x50\x2f\xe0\xe4\xd6\x1b\x87\x8a\x2a\xbd\x3b\xf1\x42\xac\x67\xec\x68\xbd\x40\x51\xc4\xbf\x57\x09\x41\x38\xbc\x1c\x7a\x78\xa7\xc4\xbb\xb3\x33\x3b\xa3\x55\xc6\x9e\xa1\x76\x3a\xc6\x5c\xec\x03\x1d\xb3\x03\x85\x53\x0b\xc7\xbf\x59\x3c\x66\x2b\x51\x2c\x00\x00\x94\xd3\x15\x3a\xd8\x07\xca\x85\xf5\xed\x89\x4b\x72\xe2\x31\x15\x29\x0b\xde\x5d\x45\xf1\x4b\xc7\x78\x09\x64\x94\x1c\xe0\xe3\xe8\x80\x9f\x28\xd4\xc1\x33\x05\x27\xc0\x9a\x94\xae\x75\xba\xc6\x26\x38\x83\x94\x8b\x9f\x9e\x91\xa0\xdc\x6d\x1f\x1b\x54\x27\xe6\xe0\x1f\x3c\x15\x7b\xa8\xd8\x67\x2d\xd9\xa3\xa6\xab\x80\xe0\x6b\x67\xeb\xf7\x5c\x1c\x90\x7f\x37\x81\x7a\xca\xaf\xdf\x04\x44\xbe\x3a\xcc\xc5\xc5\x1a\x6e\xd6\xb0\xfa\xfe\xb6\x01\x87\x7b\x5e\xc3\xea\xc7\xdb\x06\xda\x10\x2d\xdb\xe0\xd7\x40\xe8\x34\xdb\x33\x6e\x44\x31\x8c\x83\xe5\x2f\x4a\xde\x45\xc7\x0d\x92\xa0\xa6\xe9\x24\x36\x7b\x47\x71\x94\x17\xf3\xa6\x09\xb5\xe9\xf3\x1a\x59\xa5\xb1\xe7\x62\x31\x7e\xba\x0e\x2e\x96\x1b\x58\x6e\xad\x7f\x8f\x70\xbb\x2d\x52\x55\xd6\x95\xc3\x8c\x30\xb6\xc1\x47\x7b\xc6\x47\x34\x43\x7d\x02\x4a\x17\xe3\x06\xb5\x49\xdf\xf4\x7c\x8c\x80\xa2\xdc\x6d\x95\xe4\xe6\x63\xe3\x9e\x45\x49\x6e\xbe\xfd\xc7\xe2\x25\x4e\x5b\x4a\xa6\x02\x7d\xef\x45\xbe\x0a\x66\xf4\xde\x75\x40\xda\x1f\x10\x96\xbd\xd3\x29\xf7\xcb\x92\xf7\xa2\xf9\x58\x1c\x1a\x1a\x1a\xc2\x7d\x2e\xba\x0e\x96\xe5\x6e\x0b\xb7\x9b\x28\x9e\xff\x4a\xea\x19\x32\x39\xc7\xf6\x7f\x12\xf7\xfb\x4a\x74\x92\xc2\x67\x88\xf5\x94\x43\xb2\xaf\xb1\xcc\x73\x3d\x13\xef\x3a\x40\x6f\xd2\x29\x25\x93\xbc\x95\x1c\x8e\x23\xbd\xb6\x11\xfe\x2f\x00\x00\xff\xff\xc9\xef\xcf\x79\x05\x04\x00\x00")

func templateDataIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templateDataIndexHtml,
		"template/data/index.html",
	)
}

func templateDataIndexHtml() (*asset, error) {
	bytes, err := templateDataIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/data/index.html", size: 1029, mode: os.FileMode(420), modTime: time.Unix(1516623201, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateDataNot_foundHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\xc9\x30\xb4\xf3\xcb\x2f\x51\x70\xcb\x2f\xcd\x4b\xb1\xd1\xcf\x30\xb4\xe3\xe2\x02\x04\x00\x00\xff\xff\x6b\xf2\x53\x78\x14\x00\x00\x00")

func templateDataNot_foundHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templateDataNot_foundHtml,
		"template/data/not_found.html",
	)
}

func templateDataNot_foundHtml() (*asset, error) {
	bytes, err := templateDataNot_foundHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/data/not_found.html", size: 20, mode: os.FileMode(420), modTime: time.Unix(1516623198, 0)}
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
	"template/data/footer.html": templateDataFooterHtml,
	"template/data/header.html": templateDataHeaderHtml,
	"template/data/index.html": templateDataIndexHtml,
	"template/data/not_found.html": templateDataNot_foundHtml,
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
	"template": &bintree{nil, map[string]*bintree{
		"data": &bintree{nil, map[string]*bintree{
			"footer.html": &bintree{templateDataFooterHtml, map[string]*bintree{}},
			"header.html": &bintree{templateDataHeaderHtml, map[string]*bintree{}},
			"index.html": &bintree{templateDataIndexHtml, map[string]*bintree{}},
			"not_found.html": &bintree{templateDataNot_foundHtml, map[string]*bintree{}},
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

