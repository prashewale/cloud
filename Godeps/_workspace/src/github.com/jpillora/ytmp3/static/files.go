// Code generated by go-bindata.
// sources:
// files/index.html
// DO NOT EDIT!

package static

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _filesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x92\x41\x8f\x94\x40\x10\x85\xef\xf3\x2b\x2a\x9c\xdc\xa8\xb4\x71\xbd\x38\x30\x18\x13\x63\x34\x59\x13\xd7\xac\x51\x8f\x0d\x5d\x33\x94\xdb\xd0\x9d\xa6\x18\x86\x18\xff\xbb\x45\xb3\xcc\x62\x76\x2e\x04\xea\xbd\xfa\xaa\xfb\x51\x79\xcd\x8d\x2d\x36\x79\x8d\xda\x14\x1b\x80\x9c\x89\x2d\x16\xbf\x5c\xcf\x7d\x89\xc0\x0e\x1a\x7f\x9d\xab\xb9\xba\xc9\xd5\xec\xcb\x4b\x67\x46\xe8\x78\xb4\xb8\x4b\xbc\x36\x86\xda\xc3\x16\x5e\xbf\xf2\xa7\x0c\xf6\xae\xe5\x97\x7b\xdd\x90\x1d\xb7\x50\xb9\x3e\x10\x86\x2c\x89\xec\xfa\x7a\x02\xdf\xad\xc1\x52\xca\x55\x19\xa2\xec\xa7\x27\xc0\xd7\x80\x1e\x5b\x03\x1a\x16\xf3\xf7\x6f\x37\xe0\xc2\xf9\xf3\x48\x06\x1d\x7c\xfe\x20\x90\xed\xd4\xa7\xfc\xba\x3d\x2f\x8b\x9a\xd9\x77\x5b\xa5\x46\x96\x19\x69\x8d\xc1\xdd\xf7\xda\xfb\xb4\x72\x8d\x92\x8a\xfa\x23\xbd\xc2\x13\xec\x5f\x19\x5e\x3c\x65\x7c\x14\x15\x4f\xba\xf1\x16\x2f\x4d\xd0\x40\x66\x97\x90\xb1\xd4\xde\x27\x72\x7e\x7d\xbe\x03\x08\xf7\xf1\xfd\xc1\xb8\xef\xad\x7d\x6a\x5d\xa0\x5d\x15\xc8\x33\xf0\xe8\x25\x4b\xc6\x13\xab\xdf\xfa\xa8\xe7\x6a\x8c\xed\xa8\x03\x94\xba\x43\xd8\x81\x75\x95\x66\x72\x6d\xea\x02\x1d\xa8\x85\xe7\x90\xc4\x0b\x25\xd9\x83\x8f\x8c\xb8\xa2\x59\x24\x73\x3b\xbc\x19\xde\xfe\x38\xfc\xac\x6e\xcf\x86\x3e\xd8\x95\x63\x09\x6a\x18\x86\x74\x9c\x7f\x79\x0c\x69\xd0\x5c\xd5\xef\x8e\xbb\x35\x61\x23\x84\xf9\xca\x69\x87\xfc\x9e\x39\x50\xd9\x33\x3e\x4b\xea\x80\xfb\xe4\x85\x68\x57\xd9\xa3\x85\xda\x16\xc3\xa7\xbb\x2f\x37\x32\x8d\xcc\x24\x2c\x29\x5c\xee\x96\x73\x5d\xfd\xe7\x5a\x03\x44\xcc\x62\x62\x73\x2a\xd3\x1e\x4e\x0b\x18\xf7\x71\x5a\xdf\x7f\x01\x00\x00\xff\xff\x16\x7c\xbc\xd9\xc5\x02\x00\x00")

func filesIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_filesIndexHtml,
		"files/index.html",
	)
}

func filesIndexHtml() (*asset, error) {
	bytes, err := filesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/index.html", size: 709, mode: os.FileMode(420), modTime: time.Unix(1434541112, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"files/index.html": filesIndexHtml,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"files": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{filesIndexHtml, map[string]*bintree{
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
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
                err = RestoreAssets(dir, path.Join(name, child))
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

