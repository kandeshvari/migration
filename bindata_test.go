// Code generated by go-bindata. (@generated) DO NOT EDIT.

 //Package migration generated by go-bindata.// sources:
// test-migrations/1_init.down.sql
// test-migrations/1_init.up.sql
// test-migrations/2_update.down.sql
// test-migrations/2_update.up.sql
// test-migrations/3_add_column.down.sql
// test-migrations/3_add_column.up.sql
package migration

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

var _testMigrations1_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func testMigrations1_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_testMigrations1_initDownSql,
		"test-migrations/1_init.down.sql",
	)
}

func testMigrations1_initDownSql() (*asset, error) {
	bytes, err := testMigrations1_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test-migrations/1_init.down.sql", size: 0, mode: os.FileMode(438), modTime: time.Unix(1557999818, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testMigrations1_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func testMigrations1_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_testMigrations1_initUpSql,
		"test-migrations/1_init.up.sql",
	)
}

func testMigrations1_initUpSql() (*asset, error) {
	bytes, err := testMigrations1_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test-migrations/1_init.up.sql", size: 0, mode: os.FileMode(438), modTime: time.Unix(1557999818, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testMigrations2_updateDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func testMigrations2_updateDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_testMigrations2_updateDownSql,
		"test-migrations/2_update.down.sql",
	)
}

func testMigrations2_updateDownSql() (*asset, error) {
	bytes, err := testMigrations2_updateDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test-migrations/2_update.down.sql", size: 0, mode: os.FileMode(438), modTime: time.Unix(1557999818, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testMigrations2_updateUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func testMigrations2_updateUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_testMigrations2_updateUpSql,
		"test-migrations/2_update.up.sql",
	)
}

func testMigrations2_updateUpSql() (*asset, error) {
	bytes, err := testMigrations2_updateUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test-migrations/2_update.up.sql", size: 0, mode: os.FileMode(438), modTime: time.Unix(1557999818, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testMigrations3_add_columnDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func testMigrations3_add_columnDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_testMigrations3_add_columnDownSql,
		"test-migrations/3_add_column.down.sql",
	)
}

func testMigrations3_add_columnDownSql() (*asset, error) {
	bytes, err := testMigrations3_add_columnDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test-migrations/3_add_column.down.sql", size: 0, mode: os.FileMode(438), modTime: time.Unix(1557999818, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testMigrations3_add_columnUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func testMigrations3_add_columnUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_testMigrations3_add_columnUpSql,
		"test-migrations/3_add_column.up.sql",
	)
}

func testMigrations3_add_columnUpSql() (*asset, error) {
	bytes, err := testMigrations3_add_columnUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test-migrations/3_add_column.up.sql", size: 0, mode: os.FileMode(438), modTime: time.Unix(1557999818, 0)}
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
	"test-migrations/1_init.down.sql":       testMigrations1_initDownSql,
	"test-migrations/1_init.up.sql":         testMigrations1_initUpSql,
	"test-migrations/2_update.down.sql":     testMigrations2_updateDownSql,
	"test-migrations/2_update.up.sql":       testMigrations2_updateUpSql,
	"test-migrations/3_add_column.down.sql": testMigrations3_add_columnDownSql,
	"test-migrations/3_add_column.up.sql":   testMigrations3_add_columnUpSql,
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
	"test-migrations": &bintree{nil, map[string]*bintree{
		"1_init.down.sql":       &bintree{testMigrations1_initDownSql, map[string]*bintree{}},
		"1_init.up.sql":         &bintree{testMigrations1_initUpSql, map[string]*bintree{}},
		"2_update.down.sql":     &bintree{testMigrations2_updateDownSql, map[string]*bintree{}},
		"2_update.up.sql":       &bintree{testMigrations2_updateUpSql, map[string]*bintree{}},
		"3_add_column.down.sql": &bintree{testMigrations3_add_columnDownSql, map[string]*bintree{}},
		"3_add_column.up.sql":   &bintree{testMigrations3_add_columnUpSql, map[string]*bintree{}},
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
