// Code generated by go-bindata. (@generated) DO NOT EDIT.

 //Package main generated by go-bindata.// sources:
// internal/protoc/protoc-gen-authoption/templates/auth_method_mapping.go.tmpl
package main

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesAuth_method_mappingGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\x4f\x6f\xdb\x20\x14\x3f\x8f\x4f\xf1\x84\x7c\xd8\xa2\x16\xb4\x6b\xa4\x1c\xa6\x76\x9d\x76\x68\x63\x69\xbd\x57\xd4\xbc\x62\x14\xf3\x47\x40\xb2\xad\xc8\xdf\x7d\x02\x3b\x8d\xbd\x4d\x93\xc6\xc5\x60\xde\xef\xaf\xcd\x39\xdc\x38\x89\xa0\xd0\x62\x10\x09\x25\x3c\xff\x04\x1f\x5c\x72\xdd\xb5\x42\x7b\x2d\x8e\xa9\x37\x98\x7a\x27\x19\xdc\xee\xe1\x61\xff\x08\x9f\x6f\xbf\x3e\x32\x42\xbc\xe8\x0e\x42\x21\xe4\xcc\xee\xf4\x80\xec\x8b\x6b\x0f\x8a\x3d\x08\x83\xe3\x48\x08\xd1\xc6\xbb\x90\xe0\x3d\x01\x00\xa0\xca\x39\x35\x20\x53\x6e\x10\x56\x31\x17\x14\x57\xc1\x77\xb4\x5e\x92\x77\x54\xe9\xd4\x1f\x9f\x59\xe7\x0c\x7f\xd5\x49\x48\x1c\xde\x9e\xda\x26\x0c\x56\x0c\x5c\x78\xcd\x8b\x9b\x57\xfa\x1f\x80\xa2\xc2\x23\x86\x13\x06\x6e\xb4\x94\x03\x7e\x17\x01\x29\xf9\x40\x48\xce\x10\x84\x55\x08\x4d\x84\xed\x0e\xa6\x10\xdf\x30\x9c\x74\x87\x11\x4a\x06\xbe\xd9\x10\xd8\x40\xce\x4d\x3c\xe7\x82\x0d\x27\xa4\x73\x36\xa6\xe5\xeb\xa7\xfb\xda\x50\x1b\xf0\x45\xff\x80\x1d\xd0\x9c\x9b\x89\xaf\x9d\x4a\x1a\x47\xb6\x18\xa7\x84\x9c\x44\x58\x11\x7c\x3a\xa6\x7e\x22\x89\xb0\x83\x1a\x93\x4d\xe7\x7b\xe1\xbd\xb6\x0a\x72\xed\xea\x62\xda\x14\xd3\x4d\x9c\xa7\x8a\xb5\x79\xe5\x0c\x8d\x29\x7c\x7b\x9f\xca\x8c\xf3\x49\x3b\x0b\x8d\x61\xfb\xba\x8b\x40\xe7\xaa\xd8\xe9\x23\x2b\x52\x4f\xd3\x08\x85\x35\x8b\x7e\x01\x61\xe5\x82\xec\x6d\xc7\x5a\x0c\x46\xc7\x58\x78\x17\x98\xfa\xa5\xf9\xbf\xb3\xf3\x9c\xcf\x2d\x6c\xe7\x9c\x93\xad\xbc\xa2\x29\xeb\x22\xb2\xad\x8d\xfe\x45\x7d\x1c\xe9\xd5\x1f\xb8\x9b\x1e\xbb\x43\x2b\x82\x30\xbf\xe1\xea\xc5\x9d\xc6\x41\xce\x0e\xd6\xd8\xf1\x6a\x91\x1e\xed\xb9\xd3\x9c\x61\x3a\x8c\xf5\x9f\x41\x2b\x4b\xe6\x5f\x01\x00\x00\xff\xff\x39\x1d\x49\x8d\x38\x03\x00\x00")

func templatesAuth_method_mappingGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesAuth_method_mappingGoTmpl,
		"templates/auth_method_mapping.go.tmpl",
	)
}

func templatesAuth_method_mappingGoTmpl() (*asset, error) {
	bytes, err := templatesAuth_method_mappingGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/auth_method_mapping.go.tmpl", size: 824, mode: os.FileMode(420), modTime: time.Unix(1666096359, 0)}
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
	"templates/auth_method_mapping.go.tmpl": templatesAuth_method_mappingGoTmpl,
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
		"auth_method_mapping.go.tmpl": &bintree{templatesAuth_method_mappingGoTmpl, map[string]*bintree{}},
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
