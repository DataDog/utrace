// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ebpf/bin/probe.o

package assets


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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataProbeO = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x4b\x68\x5c\xd5\x1b\xff\xee\xe4\x35\x79\x34\x4d\xd2\x4e\xff\xe9\x34\xff\x7f\xb3\xfb\x87\x51\xf3\xb2\x4a\x2a\x88\xa1\xf8\x08\x52\x21\x14\xc1\x74\xd3\x9b\xc9\xcd\x6d\x67\x4c\x32\x19\x67\x6e\x6a\x1e\x15\xe3\xa2\x10\xba\xea\xc2\x47\x28\x45\x13\xab\xd2\x45\x0b\xd9\x48\x02\x16\xa6\x60\x17\x41\x5c\x54\x70\x11\x28\x62\x56\x12\x10\x31\x42\xc1\x2c\x6a\xae\xdc\x73\x7e\x67\xee\xb9\xe7\xde\x9b\x4c\xb1\xd2\x2e\x72\xa0\xfd\xe5\xfb\xce\xf9\x5e\xe7\x7c\xdf\x77\xce\xdc\xf7\x5f\x39\xf9\x6a\x44\xd3\x48\x0c\x8d\xfe\x24\x97\x72\xc7\x5c\xdc\xfd\xbb\x17\xff\x1f\x26\x8d\x0a\x87\x38\xaf\x59\x11\x32\xe2\x0f\x6c\x07\x0b\x9f\x73\xba\x2a\x42\xf4\xc0\xb6\x6d\x75\xdd\x45\x66\x93\x28\x46\x3d\x8c\x5e\xd6\x38\x7d\xaf\x21\x1a\xa8\xf7\x9c\x46\xd4\x4a\x44\xef\x02\x63\x5a\x17\x97\xab\xe2\xf3\x33\xd3\x9b\x36\xc7\x0d\xe0\x3a\x70\x0d\x78\x17\xb8\x6a\x0b\xfb\xf5\x0e\x5d\xc3\xe9\xe4\x22\xf7\xdb\x88\x73\x3d\x85\x24\xe2\x8b\x78\xfd\x58\x2e\x73\x82\xe2\xf2\x47\x9c\xf5\x35\x5c\xef\x54\x32\xc5\xe6\xa7\x86\xb2\xdc\x9f\x04\xec\xc7\xb9\x3f\x53\xc6\x3c\xe3\x17\x16\xb1\x2f\x1a\xd1\x9a\x6d\xdb\xcb\x11\xa2\x28\xfc\x29\x07\x3a\xf4\x39\x42\xbc\xc0\xc2\x17\x90\x2b\x27\x5a\xb5\x6d\x3b\xcc\xbf\x42\x19\xfc\xac\x20\xea\x84\xbe\xc3\x01\xfa\x62\x54\xcd\xcf\x6b\xfa\x8e\xef\xbc\xee\xec\x78\x5e\x11\xe5\xbc\x60\x0f\xeb\x3e\x92\xd6\x57\x38\x7a\x2b\x5d\xba\x5e\xb2\x6f\xd4\xf0\xfd\xde\xfd\x9c\xb9\x82\xd2\xf3\x89\x27\x66\x73\xd4\x3b\x6f\x5c\xda\xf2\xc5\xb9\xb5\x83\x9e\x02\xf2\x2a\x56\x55\xc7\xcf\xf3\xd2\x86\x92\x37\x9b\x3e\x7d\x1b\x3b\xfa\xc5\x0f\x66\x0a\xf3\x47\x51\x3f\xf7\xce\x05\xef\x9f\xa8\x2f\xb1\x8f\x33\x35\x5b\x1e\xfb\xc2\xbf\xc2\x3b\x1c\xc5\x7e\x19\xf1\xfb\x76\x50\xfc\x85\x0f\x4a\xdf\xdf\xfb\x3b\xc6\x71\x9a\xef\xaf\xaf\xee\xfd\xfb\xb1\xb9\xa3\x9e\x7e\x1e\xb7\x52\xf7\xa2\x7e\x93\x8b\xa8\xe7\xf8\xba\x4f\xef\xba\x93\xff\xd7\x40\x97\xf1\x73\x14\x7d\x40\xb5\xb7\x5c\xee\xda\x8d\x94\x94\x6f\x2f\xc2\x8f\xe0\xbe\x31\x33\x7d\x1b\xb8\xc2\xeb\x7a\x11\xfd\x26\xbe\x04\x7a\x1d\x7d\x64\x6d\xf7\x3e\x52\x29\xf7\x91\x5b\x76\x49\xe7\x54\x72\x5f\x5a\x29\xe6\xcb\x6e\xfd\x64\xe9\x91\xf4\x13\x9e\xb0\xcb\x95\xfc\x3c\x8d\xb7\x6f\xf8\xce\xed\xc6\x8e\xf9\xc0\x27\xee\x0d\x22\xfe\x28\xf4\x03\x63\xd1\x1a\xae\x1f\xf2\x46\xfc\x61\xf5\xff\xb3\xbe\x15\x5e\x6f\xff\x56\x1f\x8b\xf1\xf3\xa8\xf6\xce\x1b\x57\x1e\xb2\x8f\xa1\x0f\xc4\xa2\xb5\x0c\x8b\xf7\xe4\x95\x8d\x47\xd3\xbf\x52\xc1\xfb\x58\x38\xc0\x31\xd9\xc5\x1d\x58\x46\x5e\x8d\x24\x70\xcf\x26\x78\x1c\x33\x09\xf8\x93\x40\x1d\x89\x7b\x33\x81\xfa\x4b\xa0\xfe\x12\xa8\xbf\x04\xea\x2f\xb1\x02\x5c\x02\x5e\x07\x2e\x00\xe7\x81\x97\x81\x73\xc0\x59\xe0\x24\x30\x0b\x4c\x01\x07\x81\x03\xc0\x7e\x60\x1f\xb0\x17\xd8\x03\xec\x04\xb6\x01\x5b\x81\xcd\xc0\x06\x60\x14\x48\xc0\xad\x6d\xc4\x0f\xdc\x00\xae\x03\xd7\x80\x77\x81\xab\xc0\xdb\xc0\x15\xe0\x12\xf0\x3a\x70\x81\xe1\x5b\x1a\x91\x6d\x13\xed\xc7\x39\x14\xbe\xc2\x39\x57\x12\x2d\x6c\xbb\xf5\xee\x9c\xcb\x19\x8d\x9f\xf3\x33\x01\x79\x5c\x18\x52\xf3\x60\x5f\x91\xf6\xde\xe7\x57\x99\x5d\x51\x5f\x46\x7c\x7e\x5b\xcd\xaf\xab\xdb\xde\xbe\x3d\xbf\xed\xcf\x37\xb5\x5f\xab\x79\x55\x89\x77\x52\x03\xfc\xd9\x1b\xee\xd0\xb0\x2f\xac\xe2\x1a\x1e\xb7\x37\x4f\xd6\x88\x20\x6f\xa2\xf8\x7b\x6f\xb8\xa3\x0a\x7b\xb3\xe5\xb4\xf8\xb6\xc7\xed\xcd\x93\x35\xca\x1f\xb7\x03\x4f\xf0\x70\xfa\x8d\x73\x77\x94\xef\xf5\x62\xdf\xd0\x90\x3b\x2c\x7f\xf6\x92\xc8\x33\x5e\xeb\x3f\x49\xdb\xb6\xcd\xde\x41\xe2\x9a\xd2\xa6\x4f\x51\xf4\x42\xad\xe6\xfc\xda\x6e\xc6\x3f\x31\x5a\xa5\x86\xdd\x42\x44\xc7\xa4\xb9\x5e\x25\xf1\x9c\xf9\x37\x25\x7a\xb2\xcc\x3f\x6f\x49\x74\x5f\x80\xfc\x87\x12\x3d\x50\xe2\x65\xf1\x29\xab\x83\x3f\x6c\x95\xff\x03\xd3\x5f\x46\x0d\x8a\x9e\xdf\x18\xbf\x82\x66\x95\xf5\xdf\x81\xdf\xa9\xf8\x75\x9f\xd1\x55\x94\x52\xe2\xf9\x05\xfc\x05\x45\xcf\x2d\xf0\xfb\x14\xbb\x5f\x83\x4f\x8a\x9e\x9f\x19\xbf\x9a\x2e\x2b\x7a\xbe\x05\xbf\x47\xf1\x67\x8d\xd1\xb5\x3e\xfd\x3c\xda\x32\x3f\x93\xf1\x2b\x42\xf8\x55\x21\xfc\xea\x10\x7e\xad\x8f\xf7\x31\x11\x35\xd2\xc1\x22\x2d\x4a\xee\x02\xe3\xc7\x7c\xfc\x4b\x8c\xdf\x54\xe4\xb7\x20\xbe\x97\x18\x7f\x9f\x8f\xaf\x85\x3c\xa9\x5e\xd7\x82\xf9\x3f\x11\x51\x93\x14\x97\xc8\xc3\xef\x19\xdf\x8d\x4b\xe4\xdf\x1b\xcc\xee\x7e\x9f\xdd\xe3\x21\xfa\xcf\xb0\xf5\x0d\xbe\xf5\x5d\x8c\x5f\xef\xe3\x1f\x65\xfc\x46\x1f\x3f\xcb\xf8\x07\x7c\xfc\x53\x9a\xe3\xa7\xbb\xcf\xa2\x0e\x6e\x32\xff\xdd\xf3\x15\xb5\xf9\x0d\xe3\xbb\xe7\x2b\xea\x52\x7c\x86\x72\x48\xe7\x27\x63\xa7\x44\x3b\xa7\x35\x28\xd1\x71\x22\xda\x90\xe8\x23\x44\x34\x27\xed\xff\x7f\x89\x68\x5d\xa2\x0f\x4b\xcf\x15\xfe\xdb\xc4\xab\xef\x7f\x44\xc5\xfa\x12\xf6\xaf\x2b\xf6\xb7\x24\xba\xa5\xf8\x8d\xdb\xd5\x37\xab\xe8\x5b\x50\xf4\xad\x2a\xfa\x7a\x34\xaf\xbe\x01\xcd\x1b\xdf\x5d\xcd\x1b\x9f\xd8\x3f\x41\x4f\x46\xbc\xf1\xae\x45\xbc\xf1\xf6\x94\x29\xb4\xe2\x6f\x56\xf1\xf7\xb2\xe2\xef\x8a\xe2\x2f\x29\xfe\xae\x48\xb4\x93\x15\x6d\x11\xaf\xbe\x16\x49\xbe\x4e\xb1\xef\x54\xcd\x80\x44\x3b\x59\x38\x29\xd1\x4e\x76\xcf\x4b\x34\xcb\xde\xb3\xe9\x51\xcb\xcc\xe9\x13\x79\x33\xa7\x0f\xa5\x33\xc9\xdc\x14\xb5\x5b\xe6\xa4\x45\x63\xc9\x6c\xbe\xc3\xca\x25\x0d\x53\x37\xcf\x9b\x19\x2b\xcf\x39\x79\x2b\x99\xb3\x74\x41\x19\xe3\x13\x19\xcb\xcc\x81\x1a\x1d\xcf\x5b\x3a\x13\x71\x17\x1b\x23\x1e\x0e\xfb\x7b\x58\xcf\xa6\x87\xf3\xa4\x9f\x37\x73\xf9\xf4\x78\x86\xcf\x70\xe3\x7a\x36\x69\xa5\x48\x1f\x4d\x1b\x66\x26\x6f\x52\x7b\xce\x1c\x6d\x37\x53\xfa\xd9\x5c\x72\xcc\xa4\x89\x6c\x6e\x7c\xc8\xd4\x27\x98\x12\x9a\xc8\x99\x96\x87\x31\xe2\x63\x78\x28\x47\x17\xd7\xd0\x21\x73\x84\x90\xcc\x1c\x09\x64\x7a\x38\x79\x33\x33\xac\x4b\xf1\xd1\xd9\x89\x8c\xa1\xa7\x87\x89\x51\xd9\xf1\x74\xc6\xd2\xf3\x46\xca\x1c\xc6\xff\xd9\xdc\xb8\x61\xe6\xf3\xba\x39\x69\x1a\x4c\x9d\xbb\xae\x83\xad\xe8\x08\x58\x37\x96\x4c\x67\xda\x0d\x6a\xcf\x5b\x39\x2b\x39\x44\xed\xf9\xa9\x31\x07\x4f\x9e\x38\xd1\xad\x1f\xe7\xf0\xbc\x03\xcf\xea\xcf\x39\xd0\xc5\xa1\x93\xc3\x31\xbd\x9b\xcf\x75\xf3\x95\xdd\x7c\x09\xa8\xae\x6e\x7a\x14\xe3\x4b\xcd\x77\x95\xb1\x31\x87\x07\xc4\x03\xe5\xde\x52\xdf\xa9\x07\xc1\xab\x54\xf8\xbd\x21\xf6\xd4\xa7\xdc\xef\xbb\xc8\xb7\x2a\xf7\xa3\xf2\x59\x9e\x7e\x65\xb7\xab\x7f\xa4\xf0\x7d\x4d\xf4\xb3\x66\xe2\x71\x0a\xf9\xe2\xbb\x49\x0b\xb6\x3f\x08\xbb\xea\xfb\x48\xb5\xff\x57\x88\xfd\x55\xd8\xef\x97\xec\x57\x04\xd8\x7f\x2a\xc4\xfe\x02\x0e\x45\x7d\x7f\xa9\xf6\xff\xaf\x05\xdb\x6f\xfe\x0f\xc7\x25\xc9\x7e\x55\x80\xfd\x43\x21\xf6\x9b\x71\xf5\xaa\xef\x3b\xd5\x7e\x53\x88\xfd\x55\xd8\x97\xe3\xaf\x0e\xb0\x3f\x1d\xb6\xff\x58\xa8\xbe\x1f\x55\xfb\xe7\xc3\xe2\x47\xfe\xb6\x4a\xf6\x6b\x03\xec\xbf\x80\xfc\x53\x6b\x60\x85\x7f\x2e\x2f\xde\xe7\x62\xa8\xf9\xfb\x74\x88\xfc\x8f\xb5\xa5\xc9\xf7\x85\xc8\x6f\xd6\x95\x26\x7f\x3a\x44\xbe\xae\xbe\x34\xf9\x78\x88\x7c\xdb\xfe\xd2\xe4\x2f\x86\xc8\xbf\xdc\x50\x9a\xfc\x48\x88\xfc\x60\x63\x69\xf2\x9f\x84\xc8\x5f\x68\x0a\x5e\xaf\xf6\xaf\xf7\x42\xe4\x67\x43\xe4\x55\xfa\x1a\xe4\xd5\x67\xfc\x1c\xe4\x97\x14\xbe\x9a\xbf\x9f\x85\xf4\x8f\x1e\xe4\xaf\x5c\x3f\x87\x02\xf2\xf7\xa6\x16\xfc\x3d\xae\x07\x4f\xd2\x4e\xe9\xbd\xd3\x28\xc9\x8b\xdf\x87\x7f\x07\x00\x00\xff\xff\xd4\xa1\x4a\xde\xc8\x20\x00\x00")

func bindataProbeOBytes() ([]byte, error) {
	return bindataRead(
		_bindataProbeO,
		"/probe.o",
	)
}



func bindataProbeO() (*asset, error) {
	bytes, err := bindataProbeOBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "/probe.o",
		size: 8392,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1642258739, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"/probe.o": bindataProbeO,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"": {Func: nil, Children: map[string]*bintree{
		"probe.o": {Func: bindataProbeO, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
