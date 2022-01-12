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

var _bindataProbeO = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x5d\x68\x14\xd7\x17\x3f\xb3\x31\xd9\xcd\x87\xba\x46\x57\xe3\xea\x1f\xf3\xf0\x97\xbf\xe4\xdf\x6e\x3e\xb4\x6d\x6a\x1f\x1a\x84\xb6\xa1\xa4\x45\xa4\xd0\xf4\xc5\x71\x33\x3b\xba\xdb\x24\xbb\xdb\x99\x89\x26\x59\x4b\x2d\x54\x10\x9f\xf2\xd0\x8f\x20\xd2\xc6\xda\x16\x5f\x84\xbc\x94\x04\x2a\x6c\xa0\x3e\x84\xe2\xc3\x16\xfa\x10\xb0\x85\xbc\x35\x50\x68\x2d\x08\xdd\x07\xcd\x94\x7b\xef\x99\x9d\xbb\x67\x66\x92\xd5\x16\x92\x87\x5c\x48\x7e\x7b\x7e\x73\xcf\xc7\x3d\xf7\x9c\x3b\x77\xf7\x83\x57\x06\x5e\x0d\x29\x0a\x38\x43\x81\xbf\xc0\x95\xdc\xb1\x10\x77\x3f\xf7\xe1\xff\x38\x28\x50\xdc\x2b\xb8\xcb\x00\x50\x0f\x00\x85\xa6\xb2\xed\xc8\x3b\x00\xe0\x02\x00\xb4\x03\x80\xd6\xf4\x90\xf3\xc5\x2f\xc5\xfc\x70\x08\xe0\xa1\x6d\xdb\x6d\xc4\xd9\x65\x1e\x03\x40\x0c\x06\xb8\x9c\x0c\x0b\x7e\x4e\x01\x38\x9c\x6a\x02\x2d\x5e\xe2\x76\x98\x9e\x96\xcb\x0d\xb3\xcf\x19\xfd\x78\x7b\x21\xbe\x28\xec\xdf\x40\xfb\x0a\xc0\xa2\x6d\xdb\x73\x21\x80\x16\xc6\x9b\xae\xfd\x06\xd4\x97\x87\x16\x7f\xe0\x89\xef\xc1\x9a\xf1\xf5\x56\xe2\x62\xf2\xfd\x68\x44\xc4\x7b\x43\xd8\x29\xc4\x97\xab\xf2\x50\x68\x5a\xf1\xd8\x5f\xb6\x6d\xbb\x78\x13\xe5\x3a\x80\x32\x8b\x57\x98\xf1\xc4\x37\xb7\xcd\xf5\x1f\xf2\x79\x7e\x4e\x11\x79\xbe\x80\x18\x53\xfe\xcb\xf9\xc2\x55\x91\xaf\x89\x1b\x2b\x18\x97\xc8\x93\x36\xf9\x33\xf2\x22\x4e\x2d\xbe\x24\xe2\x4b\xa2\xff\x10\xf1\x5f\xc7\x82\x14\xfe\x0f\xf0\xfd\x2c\x55\xd6\xc7\x42\x3e\x87\xfb\xec\xec\x77\xf1\x2b\x5c\xd7\x36\xb1\x0f\x41\x76\x8b\x75\x68\xbf\x1e\xa0\x0d\xed\xed\xf7\xb1\x17\x83\xa6\x4a\xbe\x45\xbc\x77\x3d\xf9\xbc\xbb\xe6\x7e\x85\xc8\x7e\xa1\x5f\x9c\xf7\x89\x34\x9f\xd5\x71\xb1\x01\xe3\x8d\x54\xdb\xd3\xae\x3e\xf2\xf8\x7d\xb4\x86\xdf\x22\xd6\x6f\x2c\xdc\x82\xfb\xb1\x42\xea\xc2\x5b\x77\x2b\x6b\xae\x43\x24\x6c\x02\x9f\x1f\xc2\xfe\xbb\x7f\xce\x7f\x3d\xeb\xf5\xa7\x13\x5f\xf1\x3d\x81\x4e\xfd\x68\x71\xd1\xaf\x73\x8d\xfe\x79\x28\x7e\x28\xd0\x5b\x77\xc2\x60\xed\x7d\x3e\xc8\xe5\xa4\xe3\xe7\x1f\xf7\xe5\x9b\x22\x0f\xa4\x2f\x0b\x93\x22\xef\xb4\x3f\xd7\xed\xc7\xb0\x7f\x5c\x4f\xde\x8f\x2f\x61\x1c\xa2\xcf\x0a\x93\x01\x7d\x79\xed\x29\xfa\xb2\x41\xee\xcb\x9f\xec\xb5\xf7\x47\x14\xf6\xe6\xed\x73\x51\xb0\x73\x0d\x62\xff\xb4\x77\x9f\xb4\xcf\xc5\x83\xfb\x67\x70\xfd\x11\xb4\x8f\x18\x8b\x6c\xcc\x39\xb2\x7e\xbf\x3d\xaa\xbc\xd7\xe4\x11\xb8\x7f\x35\x9e\x3f\x31\x88\x89\xfd\x68\xac\x7e\xae\x5d\x2b\x7b\xd6\x5d\x5e\xeb\x1c\xc3\xfe\x8f\x45\x9a\x39\x16\x26\xb1\x8f\xae\x79\xdf\x6b\x4f\x75\x7e\xa5\xfd\xf3\x58\xdc\x2d\x30\xd9\x2d\x02\x98\xc3\xba\x1a\xee\x10\xf9\xd2\x3a\xc4\x3a\x0a\x1d\x18\x4f\x07\xf6\x53\xc7\x32\x22\xf6\x5b\x47\x09\x71\x11\x71\x01\x71\x1e\x71\x16\xf1\x16\xe2\x0c\xe2\x34\xe2\x14\xe2\x15\xc4\x4b\x88\xe3\x88\x79\xc4\x34\xe2\x19\xc4\x41\xc4\x93\x88\xfd\x88\x7d\x88\xbd\x88\x5d\x88\x47\x10\xdb\x11\xdb\x10\xa3\x88\x11\x44\x40\x2c\xaf\xe2\xfa\x11\x57\x10\x97\x11\x97\x10\x4b\x88\x8b\x88\x0b\x88\xf3\x88\xb3\x88\xb7\x10\x67\x38\xbe\xad\x00\xd8\x36\xc0\x4e\xdc\x87\xe2\x37\xb8\xcf\x0d\x00\x33\xab\x6e\xbf\xb3\x7d\x39\xad\x88\x7d\x7e\xd6\xa7\x8e\x8b\x43\xb4\x0e\xb6\x8b\x7d\xad\xf4\xe1\xf5\x55\xf0\xbd\x37\x4e\xaf\xd2\xfa\x9a\x5e\xad\x3e\xa7\xaf\xaf\x7a\xeb\x8d\x9e\xcf\xb4\xae\x58\x17\xb1\x29\x51\x8c\x67\x6b\xb8\xa3\x11\xf3\xc2\x3b\x2e\xba\xd1\xd1\x6c\xae\x11\xc2\xba\x89\xe0\xe7\xad\xe1\x8e\x30\xe6\xa6\xcc\x8e\xf8\x23\x1b\x1d\xcd\xe6\x1a\xdb\x36\x3a\x80\x4d\x3c\xd8\x79\xc3\xde\x1d\x3c\x47\x5b\x4d\x55\x35\x1a\xb1\x76\x78\x6e\xb6\xce\xe2\xaa\xa1\xe5\x72\xc3\x19\xfd\x78\xfb\xe1\x54\x13\xbc\x76\x72\x00\x56\x6d\xdb\x76\x52\xa4\x4c\x9e\x82\xc8\xc5\x66\x85\x7d\xeb\x6e\xc3\x3f\x67\x94\xa4\x1a\x3b\x08\x00\xc7\xe4\x67\xc4\x07\x7b\xfe\x96\x24\x8f\xd7\x79\x9f\x5b\x92\xdc\xaf\x78\x9f\x7f\x2c\xc9\x83\x35\xd6\xf7\xe7\xbc\x1f\xfe\xb4\x29\x7f\x87\xdb\xaf\x83\x05\x62\xe7\x47\xce\xd7\xc3\x02\x99\xff\x98\xf3\x61\x00\x12\xf7\x12\xf2\x69\xc2\xff\x80\xfc\x2c\xb1\xf3\x2d\xf2\xfd\xc4\xef\xaf\x9c\x6f\x84\x29\x32\xff\x7b\xe4\x7b\x49\x3e\x96\xb9\xdc\xec\xb1\x23\x56\x55\xe7\x25\x39\x5f\x1f\xc0\x87\x03\xf8\xc6\x00\xbe\xd9\xc3\x7d\x0a\x00\xbb\xf0\x3b\x12\x48\xe7\xf4\x45\xce\xef\xf5\xf0\x57\x39\xdf\x5a\xe1\x0f\xe2\xfa\x5e\xe6\xfc\x76\x0f\xaf\x04\xb4\xed\xeb\x8a\x3f\xff\x0b\x00\xb4\x4a\xeb\x72\xea\xed\x1e\xe7\xdd\x75\x39\x75\xf6\x06\xf7\xbb\xd3\xe3\xf7\xc5\x00\xfb\xa7\xf9\xfc\xa8\x67\x7e\x37\xe7\x77\x78\xf8\x43\x9c\xdf\xe5\xe1\xf3\x9c\xdf\xed\xe1\x4f\x29\x2c\x4e\x37\xcf\x4e\xbd\xdf\xe6\xf1\xbb\xfb\xeb\xf4\xe0\x77\x9c\x77\xf7\xd7\xe9\xbf\x3e\x44\x66\xf6\x3f\x00\x30\x2f\xc9\x6c\xb7\x96\x25\x99\x45\x31\x28\xe5\xfb\x00\xcb\x9b\x24\xef\xaf\xd4\xbb\x90\xe3\x00\xd0\x15\x72\x65\xd6\xa3\x53\x92\xdc\x26\xed\x99\xe3\xaf\x97\xf8\xbb\x02\xd5\xfe\xfa\x25\x79\x0f\x79\xce\xe2\x9f\x25\xf6\x96\x88\xbd\x7e\x12\x7f\x9a\xc4\x5f\x22\xf1\xb7\x87\xaa\xe5\x71\xb2\x9e\x25\xb2\x9e\xde\x3a\x22\x93\x78\xf3\x24\xde\x29\x12\xef\x3c\x89\x17\x48\xbc\xf3\x92\xcc\xba\xe3\x48\xa8\xda\xde\x41\x49\xbf\x99\xf8\x67\x67\xf4\xa0\x24\xb3\x2e\x1a\x97\x64\x56\x95\xd3\x92\xcc\xab\xfd\x6c\x66\xc4\xd2\x0d\x75\xcc\xd4\x0d\x75\x28\x93\x4d\x1a\x13\x90\xb0\xf4\x71\x0b\x46\x93\x79\xb3\xd3\x32\x92\x9a\xae\xea\xe7\xf5\xac\x65\x0a\xc6\xb4\x92\x86\xa5\x3a\x92\x96\x1b\xcb\x5a\xba\x81\xd2\x48\xce\xb4\x54\xae\xe2\x4e\xd6\x86\xab\x18\xfe\x39\xa5\xe6\x33\x29\x13\xd4\xf3\xba\x61\x66\x72\x59\xf1\x44\x38\x57\xf3\x49\x2b\x0d\xea\x48\x46\xd3\xb3\xa6\x0e\x09\x43\x1f\x49\xe8\x69\xf5\xac\x91\x1c\xd5\x61\x2c\x6f\xe4\x86\x74\x75\x8c\x1b\x81\x31\x43\xb7\xaa\x88\x61\x0f\x51\x25\x31\x5b\xc2\x42\xa7\xcc\x38\x4a\x32\x39\xec\x4b\x56\x31\xa6\x9e\x4d\xa9\xd2\xfa\xe0\xec\x58\x56\x53\x33\x29\xe0\x52\x3e\x97\xc9\x5a\xaa\xa9\xa5\xf5\x14\xfe\xcf\x1b\x39\x4d\x37\x4d\x55\x1f\xd7\x35\x6e\xce\x9d\xd7\xc9\x67\x74\xfa\xcc\x1b\x4d\x66\xb2\x09\x0d\x12\xa6\x65\x58\xc9\x21\x48\x98\x13\xa3\x0c\x07\x4e\x9c\xe8\x51\x5f\x60\xd0\xa5\x3e\xcf\xe0\xa8\xfa\x9c\x20\x8f\x32\xe8\x16\xd0\xa3\x76\x73\x3c\xa6\xf6\x88\x39\x3d\x90\x30\x72\xa9\xa4\x95\x64\xf6\xba\x13\xdd\x38\xa9\x0b\xfe\x85\xf1\xb5\xe2\x79\x35\xf2\x31\x8f\x17\x86\x7a\xf2\x9e\xa2\xbf\x15\xec\x41\xae\x81\xf0\x7d\x01\xfe\xe8\xf7\x80\x3f\xd6\xd1\x2f\x11\xff\xe4\x67\x78\xf8\x0d\xef\x88\x74\x4c\xe3\xab\xeb\x12\xca\xfb\x41\xac\xd3\xd1\xaf\x9c\x6f\x8a\xbf\xff\x28\x26\x85\xde\x87\xa8\xff\xc7\x01\xfe\xdb\xf6\x09\xec\x92\xfc\xd7\xfb\xf8\xff\x7f\x80\xff\x65\xf4\x4f\xef\x5b\xd4\xff\xff\x14\x7f\xff\xfd\xe8\x7f\x56\xf2\x1f\xf6\xf1\xbf\x37\xc0\xff\x20\xbe\x82\xe9\x7d\x8e\xfa\x6f\x0d\xf0\x5f\x46\xff\x27\x25\xff\x8d\x3e\xfe\x27\x03\xfc\xcf\xe0\x44\x7a\x5f\xa4\xfe\xcf\x07\xad\x1f\xeb\xb7\x5d\xf2\xdf\xec\xe3\xff\x38\xd6\x1f\xed\x81\xb2\xf8\x99\xbc\xf2\x5e\x77\x06\xad\xdf\x67\x02\xf4\xf7\xb4\xd4\xa6\xdf\x1f\xa0\xdf\xb5\xbd\x36\xfd\x77\x02\xf4\x07\x76\xd4\xa6\x1f\x0f\xd0\x4f\xef\xac\x4d\xff\x72\x80\xfe\x47\xd1\xda\xf4\x87\x03\xf4\x67\x76\xd5\xa6\xff\x3b\xd6\x4f\x0f\xe1\xef\xe0\x15\xb5\x85\xf0\x0a\xc1\xcf\x02\xfc\x2f\xb6\xfa\xfb\xa3\xe7\xdf\xfb\x01\xfa\xf7\x02\xf4\xa9\x7c\x13\xf5\xe9\xd7\x81\x12\xea\xd3\xef\x21\xb4\xfe\xbf\x08\x38\x7f\xd2\x58\xff\x72\xff\xed\xf3\xa9\xff\xdb\x8a\xff\x4f\x0e\x97\xf0\x6a\xdb\x26\xdd\x9b\xa2\x92\xbe\xf3\x7d\xf2\xef\x00\x00\x00\xff\xff\x56\xf0\x5f\x70\x40\x21\x00\x00")

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
		size: 8512,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1642017902, 0),
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
