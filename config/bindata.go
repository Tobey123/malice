// Code generated by go-bindata.
// sources:
// config/config.toml
// DO NOT EDIT!

package config

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

var _configConfigToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x53\x6d\x6b\xdb\x30\x10\xfe\xee\x5f\x71\x38\x30\x36\x58\x1c\xc7\xa1\xcd\x6a\x28\xa5\x74\xdd\x1b\xcd\x06\x6d\x61\xb0\xae\x0c\xd9\xba\xd8\x22\xb2\x64\x24\xb9\x6f\xbf\x7e\x77\x72\x9a\x2e\x1f\x07\x75\x02\xb2\xee\xed\x79\xee\xb9\xf3\x04\xce\x6c\xff\xe8\x54\xd3\x06\x78\x5b\xbf\x83\x22\x9f\x2f\x60\xca\xc7\x21\x54\x5a\xd4\x9b\x60\x7b\xf8\x66\x7d\x3b\x08\x58\x09\x65\xf0\x3d\x9c\x6a\x0d\x97\x9c\xe0\xe1\x12\x3d\xba\x3b\x94\x59\x32\x81\x2b\x44\xb8\xf8\x7a\x76\xfe\xfd\xea\x1c\xd6\xd6\x81\x56\x35\x1a\x8f\xa0\x0c\xdd\x3a\x11\x94\x35\x59\x92\x4c\x5e\xe7\x21\xbc\xd5\x29\xa3\x11\x7d\xb3\x56\xcd\xe0\x22\x00\xfc\x7f\x9d\x57\xe2\x93\x04\x15\x34\xc2\x31\xa4\x2b\xc1\x9d\xc3\xe5\x60\x82\xea\x70\x9f\x5f\x9a\x24\x37\x62\x08\xad\x75\xb7\x09\x80\x11\x5d\xcc\x78\xd6\x39\x25\xdb\x04\xac\x6b\x84\x51\x4f\x63\x3f\xe4\xfd\xac\xc2\x97\xa1\x1a\x7d\x95\xb2\x2f\x26\x2a\xbd\xb6\x83\x91\xe8\xe0\x0d\x9c\x9d\xff\xf8\x6d\x2e\xd4\x06\x3d\x04\x11\xc8\x14\x2c\xcd\x47\x18\x09\x15\xa2\xcb\xc6\x74\x69\x2b\x4a\x9f\x1f\x2d\x8f\xa6\xf9\xc1\xb4\x58\x5e\xe7\xcb\x72\x51\x94\x79\xfe\x8b\x9c\x9f\x94\xf3\x01\x6a\x2d\xbc\x07\x49\x25\xfc\x09\xfc\x6c\x1f\xc1\xd8\x70\x42\xac\xef\xb1\x62\xca\x83\xd3\x4c\x20\xcf\xe2\xaf\xfc\x90\x73\x61\x21\x3b\x65\xfe\x6c\x5d\xf3\x62\x19\x9d\xf3\x72\x41\x0f\x37\x8c\x9d\x50\x9a\x93\x5b\x4b\x00\x14\xe2\xbb\xd0\x67\xf8\x20\xba\x5e\x63\x56\xdb\x8e\x6b\xf4\xd6\xb1\xaf\x38\x60\x10\x5a\x2a\x8e\xe3\x93\x35\x8a\x7e\xa6\x45\x36\x3e\xef\xad\x93\x5c\x98\x58\x8a\x4a\x78\xfc\x57\xcb\x2e\xaa\x3f\x45\xbd\xe1\x2c\xd5\x89\x66\x4f\xe2\x19\x52\x7f\x41\xd5\x53\x1f\xc8\x50\x8e\xd1\x1c\x19\xf7\x38\xa2\x6a\x5b\x0b\xcd\x54\x53\x12\xe5\xba\x55\x1e\xe8\xcf\x0b\xed\x06\x63\x94\x69\x80\xc6\xf2\xd1\xd6\x1b\x8a\x66\xeb\x4a\xd4\xb4\xe2\x54\x4e\xeb\x38\x32\x9f\x6d\xbb\x61\xba\x37\x70\x54\xe4\x39\x30\x41\x5e\x06\x3b\x70\x8f\xf3\x9c\xae\x68\x44\xa5\x51\xd2\x35\xb8\x01\x59\x25\x73\xa7\x9c\x35\x1d\x9a\xc0\xe1\x04\xc6\x64\x24\xde\xa1\xb6\x3d\x5b\xc7\x11\x6e\xed\xbd\xb3\x72\xa8\x9f\x57\x4a\x46\x3a\x9c\xd6\x89\xba\xa5\x8f\x74\xba\x2f\x47\x1a\x01\x65\x6f\x95\x89\x03\x08\x75\x5f\xce\x66\xbb\x46\xcb\x62\xb1\x3c\xe4\x98\x4a\x19\xe9\x5f\xd2\xca\x19\x9d\xf7\xc2\x61\xe9\x2c\xbb\xb5\x32\x1b\xbf\x2f\x72\xb9\xd5\x33\xdd\xef\xb0\xc8\x89\x95\xb6\x4d\x33\xb2\x5a\x2b\x8d\xfb\x8c\x32\x72\xa6\x91\xef\x83\x57\x4f\xf8\xac\x0a\x5d\xc7\x81\x2d\xb6\xb7\x8a\xa6\x34\xf4\x0c\xba\x64\x02\xcc\x38\x7e\x55\xc7\xb0\x16\xda\xb3\x6e\xa4\xc4\xc3\xe3\xed\x4e\xd1\x9d\x87\xd6\x2d\x84\x9e\x11\xd3\xed\xbb\x1f\x2f\x7f\x03\x00\x00\xff\xff\x14\x18\x6d\x03\xf2\x04\x00\x00")

func configConfigTomlBytes() ([]byte, error) {
	return bindataRead(
		_configConfigToml,
		"config/config.toml",
	)
}

func configConfigToml() (*asset, error) {
	bytes, err := configConfigTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/config.toml", size: 1266, mode: os.FileMode(420), modTime: time.Unix(1485231213, 0)}
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
	"config/config.toml": configConfigToml,
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
	"config": &bintree{nil, map[string]*bintree{
		"config.toml": &bintree{configConfigToml, map[string]*bintree{}},
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

