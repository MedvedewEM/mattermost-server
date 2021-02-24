// Code generated for package migrations by go-bindata DO NOT EDIT. (@generated)
// sources:
// mysql/000001_create_teams.down.sql
// mysql/000001_create_teams.up.sql
// mysql/000002_create_team_members.down.sql
// mysql/000002_create_team_members.up.sql
// mysql/000003_create_compliances.down.sql
// mysql/000003_create_compliances.up.sql
// postgres/000001_create_teams.down.sql
// postgres/000001_create_teams.up.sql
// postgres/000002_create_team_members.down.sql
// postgres/000002_create_team_members.up.sql
// postgres/000003_create_compliances.down.sql
// postgres/000003_create_compliances.up.sql
package migrations

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

// Mode return file modify time
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

var _mysql000001_create_teamsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x91\x51\x6b\xbb\x30\x14\xc5\xdf\xf3\x29\xee\x9b\xfa\xe7\xcf\xd8\x9e\xa5\x63\xa9\x5e\xd7\x80\x9a\x62\x22\xeb\x5b\xc9\x6c\xc6\x04\x8d\x62\xd2\x6d\x1f\x7f\xb4\x4a\xd9\xda\xbd\x14\x0a\x83\xd1\x3c\x84\x40\xce\x39\x39\x37\x3f\x81\x12\x1e\xfa\x41\xf7\x6a\xd0\x1b\xe1\x94\xd3\xad\x36\x0e\x66\xe0\x0b\x4c\x31\x92\xc0\x12\x9f\x00\x00\x8c\xfb\x6e\x4d\x17\x11\x2f\x73\xe9\xff\x0b\x20\x29\x78\x06\x2c\x4f\x78\x91\x51\xc9\x78\xbe\x16\xd1\x02\x33\x7a\x13\xf1\xb4\xcc\x72\x71\xf0\x3d\x2d\xb0\x40\x70\xea\xb9\xd1\x6b\xa3\x5a\x0d\x33\xf0\xa4\x56\xad\xf5\x0e\x12\x9a\xc7\x93\xc0\x56\xaf\xba\x55\x30\x83\x98\x4a\x3a\xa7\x02\xfd\xe0\x9b\xaa\xea\x9a\x6d\x6b\x0e\x39\xb4\x69\xba\x77\xde\x6b\xc3\xcc\x5b\xed\xf4\x98\x18\xc0\x3d\xdc\xfe\xdf\x1f\x3d\x9a\x4a\x2c\x40\xd2\x79\x8a\xb0\x7f\x14\xe2\x82\x2f\x61\xec\x08\x47\xf6\xd0\x9b\x5c\xd3\xa8\x77\x1e\x09\x82\x90\x90\x65\x81\x4b\x5a\x20\xa8\xc6\xe9\x81\xbd\xe0\x47\x6d\x9d\x1d\xe7\x3f\xfd\xc3\x90\xe0\x0a\xa3\x52\x1e\xc9\x43\x12\x23\x4d\x53\x1e\x51\x89\xf0\x63\x60\x48\xc8\x1f\xa1\x92\x2a\xeb\x76\x59\xac\xea\x4c\xd9\x6f\xd4\xf9\x60\x4e\x13\xae\x6c\x2e\xc3\x26\xd6\xb6\x1a\xea\xde\xd5\x9d\x39\x17\xca\x17\xeb\x95\xc6\x65\x68\x3c\x0e\xdd\xb6\x8f\x3a\x63\xdd\xa0\x6a\xa3\x37\xe7\x22\x39\xf6\xff\x3a\x97\x7d\xb9\xb1\x2d\x4b\x00\x57\x4c\x48\x31\xf6\x0e\xc9\x67\x00\x00\x00\xff\xff\xf1\x20\x25\x99\x74\x06\x00\x00")

func mysql000001_create_teamsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql000001_create_teamsDownSql,
		"mysql/000001_create_teams.down.sql",
	)
}

func mysql000001_create_teamsDownSql() (*asset, error) {
	bytes, err := mysql000001_create_teamsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/000001_create_teams.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mysql000001_create_teamsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x96\x51\x6f\x9b\x30\x14\x85\xdf\xf9\x15\xf7\x2d\x30\x4d\x53\x9b\xb5\x55\xa5\x28\xd3\x5c\xb8\x69\xad\x11\xe8\xc0\x68\xed\x53\xe4\x04\x77\xb5\x04\x36\x02\xa7\x6b\xff\xfd\x14\x43\x58\xd7\xd0\x2d\x6a\x1f\x26\x4d\x7d\x41\x0a\xf7\xf3\xd1\xf5\x3d\xc7\x26\x7e\x82\x84\x21\x30\x72\x16\x22\xd0\x19\x44\x31\x03\xbc\xa2\x29\x4b\x81\x09\x5e\x36\xe0\x3a\x00\x00\x34\x87\x3b\x5e\xaf\x6e\x79\xed\x8e\x4f\x3c\x4b\x45\x59\x18\xbe\xb7\x45\xbf\x16\xdc\x08\x62\x60\x29\xbf\x4b\x65\xdc\xf1\x81\x07\x01\xce\x48\x16\x3e\xa6\xb2\x2a\xdf\x83\x0a\x44\x21\xf6\xa0\x64\x53\x15\xfc\x21\xe2\xa5\xe8\xfb\x3a\x39\x1a\x22\xf7\x40\x02\xd1\xac\x6a\x59\x19\xa9\xd5\xaf\x4d\x1e\x1f\x0f\xa1\x58\x72\x59\xf4\xd0\xe1\xf8\x74\x08\x62\x0f\x95\xf8\x9b\x90\xaf\xcb\x8a\xab\x7d\x36\x40\x8a\x42\xff\x10\x79\xa0\x4b\x2e\x55\x03\x46\xdc\x9b\xb6\x40\xd5\x9d\x34\xe2\x91\x31\x1f\xc7\x43\xeb\xd3\xd5\xad\x28\xc5\x13\xff\x76\xb1\xcb\x84\xce\x49\x72\x0d\x5f\xf0\x1a\x5c\x9a\x7b\x9d\x67\x11\xfd\x9a\xa1\x7d\x69\x5b\x75\x37\xcf\xae\xb6\x79\x29\xf3\xfb\x85\xd9\xc4\x64\xa1\xfe\x5c\x96\xb6\xd9\x85\xcc\xc1\xdd\xf6\x3d\xc8\xad\x6d\x48\x16\xdc\x80\xbb\xcd\xcb\x20\xb7\xb2\x91\xb3\xdc\x36\x7d\x83\x5c\x6e\xe3\x64\xb9\x6d\xb2\x06\xb9\xc6\x4e\xc9\xf6\xb7\x1d\x98\xe7\x78\x80\xd1\x39\x8d\x70\x4a\x95\xd2\xc1\x59\x3f\x35\xff\x82\x24\x29\xb2\xe9\xda\xdc\x9c\x96\xcb\xa3\x89\xe3\xa4\xc8\xe0\x73\x55\x8b\x8a\xd7\x22\x4f\x0d\x37\xa2\x14\xca\xc0\x14\xdc\x14\x43\xf4\x19\xd0\x59\x7b\x8e\xda\xa7\xf5\xa5\x2d\xf8\x71\x16\x31\xf7\x9d\x07\xb3\x24\x9e\x03\x8d\x66\x71\x32\x27\x8c\xc6\xd1\x22\xf5\x2f\x70\x4e\x3e\xf8\x71\x98\xcd\xa3\xb4\x5f\xf7\xed\x02\x13\x04\xc3\x97\x85\x68\x87\x3e\x85\x91\x3d\xa9\xa3\x1e\x21\x51\xd0\x01\x76\x5b\x1c\xa6\x10\x10\x46\xce\x48\x8a\xae\xf7\x1b\xb5\xd2\xc5\xba\x54\xbd\x8e\xcd\x5a\x5c\x09\xd5\x7a\xd4\x2a\x7a\xf0\x09\x0e\xda\x99\x8d\xba\xa6\x0f\x47\xdd\x6f\x12\x32\x4c\xba\xcb\xa3\xbd\x2e\x48\x10\xc0\x13\x19\x58\x6a\x5d\x4c\x46\x8e\xe7\x4d\x1c\xe7\x32\xc1\x4b\x92\x20\xf0\xc2\x88\x9a\xde\x44\xda\xe0\xbd\x6c\x4c\xd3\xee\x7f\x77\x86\x13\x07\xaf\xd0\xcf\xd8\xee\x8a\x89\x13\x20\x09\xc3\xd8\xdf\xdc\x5f\xcf\xc9\xfe\x3f\xde\x84\xbc\x31\x1b\x2d\xba\xd2\xaa\x3d\x1a\x2f\xb5\x67\x57\xa9\xbb\x6c\xdf\x3c\x7a\xa5\x47\x8f\xbe\x24\x2f\x35\xe7\xb9\x8f\xd1\x9b\x37\xaf\xf4\xe6\xbc\xd6\xeb\xca\xd7\xaa\x31\x35\x97\x4a\xe4\x2f\x35\xe8\xa9\x0e\x18\xa9\x1e\x36\xff\x54\x0e\xff\x9d\x47\x3f\x03\x00\x00\xff\xff\x63\x30\x29\x71\xc5\x09\x00\x00")

func mysql000001_create_teamsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql000001_create_teamsUpSql,
		"mysql/000001_create_teams.up.sql",
	)
}

func mysql000001_create_teamsUpSql() (*asset, error) {
	bytes, err := mysql000001_create_teamsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/000001_create_teams.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mysql000002_create_team_membersDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x91\xdf\x6b\xbb\x30\x14\xc5\xdf\xf3\x57\xdc\x37\xf5\xcb\x97\xb1\x3d\x4b\xc7\x52\xbd\xae\x82\x9a\x92\x44\xd6\xb7\x92\xb6\x77\xac\x60\x5c\xd1\x14\xf6\xe7\x8f\xaa\x73\xdd\x8f\xa7\xb5\x30\x18\xcd\x43\x08\xe4\x9c\x93\x93\xfb\x51\xa8\xe1\x6e\xd7\xd0\xce\x34\xb4\x51\xce\x38\xb2\x54\x3b\x98\x80\xaf\x30\xc3\x48\x43\x9a\xf8\x0c\x00\xa0\xdf\x0f\x6b\xb8\x88\x44\x59\x68\xff\x5f\x00\x89\x14\x39\xa4\x45\x22\x64\xce\x75\x2a\x8a\xa5\x8a\x66\x98\xf3\xab\x48\x64\x65\x5e\xa8\xd1\xf7\x30\x43\x89\xe0\xcc\xaa\xa2\x65\x6d\x2c\xc1\x04\x3c\x4d\xc6\xe6\x64\x57\xd4\xb4\xde\x28\xe4\x45\x3c\xc8\xda\xf5\x13\x59\x03\x13\x88\xb9\xe6\x53\xae\xd0\x0f\x3e\xa8\xd6\xcf\xd5\xde\xd6\x63\x9a\x3a\xc8\xa9\x6c\xa9\xe9\xc3\x02\xb8\x85\xeb\xff\xdd\xd1\xe3\x99\x46\x09\x9a\x4f\x33\x84\xa3\x57\x21\x96\x62\x0e\x7d\x55\x78\xf7\x87\xde\x60\x1b\x3e\x7b\xe3\xb1\x20\x08\x19\x9b\x4b\x9c\x73\x89\x60\x2a\x47\x4d\xfa\x88\x2f\xdb\xd6\xb5\xfd\x04\xbe\x4e\x31\x64\xb8\xc0\xa8\xd4\x9f\xe4\x21\x8b\x91\x67\x99\x88\xb8\x46\xf8\x36\x30\x64\xec\x0f\x72\xe1\x1b\xbb\xad\x4f\x01\xd3\x05\x5c\xc8\x9c\x9f\xcc\xfd\x9e\x5a\x77\x0a\x99\x2e\xe0\x42\xe6\x9c\x64\x62\xaa\xc8\x11\xff\x21\x96\x37\xf7\xaf\x33\xe9\x4a\xf5\x5d\xd3\x04\x70\x91\x2a\xad\x8e\x5b\x87\xec\x35\x00\x00\xff\xff\x11\xda\xf6\x5b\x82\x06\x00\x00")

func mysql000002_create_team_membersDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql000002_create_team_membersDownSql,
		"mysql/000002_create_team_members.down.sql",
	)
}

func mysql000002_create_team_membersDownSql() (*asset, error) {
	bytes, err := mysql000002_create_team_membersDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/000002_create_team_members.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mysql000002_create_team_membersUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x95\x5f\x6b\xdb\x30\x14\xc5\xdf\xfd\x29\xee\x5b\xac\x51\x46\x57\x4a\x19\x84\x8c\x29\xf6\x4d\x23\x66\xcb\x41\x92\x59\xfb\x14\x94\xf8\x76\x35\x58\x4e\xb1\x95\xd1\x7d\xfb\xe1\x3f\xc9\x3a\xba\x75\x0f\xeb\xd8\x28\xf3\x83\xb1\x75\x8f\x8e\x2e\xe7\x77\xb1\x23\x85\xdc\x20\x18\x3e\x4f\x10\xc4\x02\x64\x66\x00\xaf\x84\x36\x1a\x0c\x59\x97\x92\xdb\x50\xd3\x42\x18\x00\x40\xbf\x22\x0a\xf8\x6c\x9b\xed\xad\x6d\xc2\xb3\x0b\xd6\xeb\x65\x9e\x24\x27\xbd\x20\x6f\xa9\x79\x52\xa0\x76\x15\xb5\xc7\xfa\xc5\x39\x1b\x96\x63\xaa\xc8\x13\xf7\xb0\x29\x3f\x95\xb5\x0f\xcf\x4e\xc7\xc2\x4a\x89\x94\xab\x6b\xf8\x80\xd7\x10\x0e\xc7\x9f\x8c\xa7\x8c\x8a\xae\x52\x16\xf7\x6b\x4f\xd6\xb9\xa1\xdb\xfe\x79\x5d\x16\x87\x1d\x4f\x28\xf7\x2d\x35\xbd\xf2\x97\x9e\x45\xdf\xe2\xda\x7a\x08\x0f\xdd\xb2\x80\x01\xca\x4b\x21\x71\x26\xea\x7a\x17\xcf\x21\xc6\x05\xcf\x13\x03\xd1\x92\x2b\x8d\x66\xb6\xf7\x37\x6f\xdd\xe6\x7c\x1a\x04\x1a\x0d\xbc\xbf\x6b\xe8\xce\x36\x54\x68\x6f\x3d\x39\xaa\x3d\xcc\x20\xd4\x98\x60\x64\x40\x2c\x86\x8c\x87\x7b\x77\x8d\x85\x28\xcb\xa5\x09\x5f\x31\x58\xa8\x2c\x05\x21\x17\x99\x4a\xb9\x11\x99\x5c\xeb\x68\x89\x29\x7f\x1d\x65\x49\x9e\x4a\x7d\xdc\xf7\x71\x89\x0a\xc1\xdb\x4d\x45\xeb\xda\x3a\x82\x19\x4c\x1e\xb0\x9c\x1c\x85\x5c\xc6\xa3\xac\xdd\xde\x92\xb3\x30\x83\x98\x1b\x3e\xe7\x1a\x43\xf6\x9d\x6a\xbb\xab\xf6\xae\x3e\xba\xe9\x4e\x4e\x5d\x64\x83\x19\x83\x77\x70\x3a\x44\x37\x19\xbb\x7e\x33\x19\xdf\x79\x62\x50\x8d\xf3\xf5\x70\xa2\x78\x1c\xc3\x37\x1f\xf0\x65\xfd\xa5\x23\x7f\xce\xa6\x93\x80\xb1\x69\x10\xac\x14\xae\xb8\x42\xb0\x95\xa7\x46\xdc\xc8\x9d\xc7\xfb\xb2\xf5\xed\x90\xc3\xe3\x2c\xa7\x01\x5e\x61\x94\x9b\xc7\x3b\xa6\x41\x8c\x3c\x49\xb2\xa8\x9b\xf4\x9f\xd9\xbe\x4c\x46\xbc\x70\x65\xfd\x1c\x90\x7a\xa3\xff\x94\xfe\x0c\xa5\xcb\x3d\xb5\xfe\x39\x28\xf5\x46\xff\x04\xa5\x97\x85\xe9\xf0\xc5\xff\x3d\x46\x3f\xf8\xcb\xfd\x35\x40\x5f\x03\x00\x00\xff\xff\x57\xa0\xc2\x52\xf8\x07\x00\x00")

func mysql000002_create_team_membersUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql000002_create_team_membersUpSql,
		"mysql/000002_create_team_members.up.sql",
	)
}

func mysql000002_create_team_membersUpSql() (*asset, error) {
	bytes, err := mysql000002_create_team_membersUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/000002_create_team_members.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mysql000003_create_compliancesDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x70\xce\xcf\x2d\xc8\xc9\x4c\xcc\x4b\x4e\x2d\xb6\xe6\x02\x04\x00\x00\xff\xff\x4d\x2a\x2d\xf2\x22\x00\x00\x00")

func mysql000003_create_compliancesDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql000003_create_compliancesDownSql,
		"mysql/000003_create_compliances.down.sql",
	)
}

func mysql000003_create_compliancesDownSql() (*asset, error) {
	bytes, err := mysql000003_create_compliancesDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/000003_create_compliances.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mysql000003_create_compliancesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8e\x41\x4b\xc3\x40\x10\x85\xef\xf9\x15\x73\x4c\xc0\x83\x94\x52\x04\xe9\x61\xbb\x99\xea\xd2\x18\x25\xbb\x01\x7b\xeb\x36\x19\xeb\x42\xb3\x29\xbb\x13\xb5\xff\x5e\x68\xa1\x04\x09\x78\x7d\xdf\xf7\x66\x9e\xac\x50\x18\x04\x23\x56\x05\x82\x5a\x43\xf9\x6a\x00\xdf\x95\x36\x1a\x64\xdf\x9d\x8e\xce\xfa\x86\x22\xa4\x09\x00\x80\x6a\xe1\xcb\x86\xe6\xd3\x86\x74\xb6\xc8\x2e\x6e\x59\x17\xc5\xdd\x05\xca\x40\x96\x49\x30\xec\xdd\xc1\x79\x4e\x67\xf7\xd9\x15\xd4\x91\xc2\x9f\x66\x8e\x6b\x51\x17\xe3\xb6\x66\xcb\x43\xbc\x49\x8b\xf9\x94\x24\xfb\xc1\x33\x38\xcf\x74\xa0\x30\xc1\x77\x39\xc5\x66\x07\x4c\x3f\x7c\x0d\xcc\xf9\x44\xff\xdc\xd4\x6c\x03\x4f\xac\x46\xdf\x4e\xa4\x1b\x3a\x7f\xf7\xa1\x8d\xa3\x1f\xd8\x59\x77\x1c\x07\x6f\x95\x7a\x11\xd5\x16\x36\xb8\x85\x54\xb5\x59\x92\x01\x96\x4f\xaa\xc4\xa5\xf2\xbe\xcf\x57\xb7\x11\xf2\x59\x54\x1a\xcd\x72\xe0\x8f\x87\x6e\x3f\x7f\x4c\x7e\x03\x00\x00\xff\xff\xdc\x08\x89\x8a\x8d\x01\x00\x00")

func mysql000003_create_compliancesUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql000003_create_compliancesUpSql,
		"mysql/000003_create_compliances.up.sql",
	)
}

func mysql000003_create_compliancesUpSql() (*asset, error) {
	bytes, err := mysql000003_create_compliancesUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/000003_create_compliances.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres000001_create_teamsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xd0\xb1\x0e\xc2\x20\x10\x06\xe0\x9d\xa7\xe0\x3d\x98\xaa\xc5\xa4\x49\x6d\x4d\x5b\x93\x6e\x84\xc0\x45\x2f\xa1\x40\x80\xaa\x8f\x6f\x2c\x83\x6e\xc5\xf9\xbe\xff\x72\xff\x55\xed\xc4\x07\x3a\x55\x87\x96\xd3\x04\x72\x89\xb4\x1e\xfa\x0b\x3d\xf6\xed\xf5\xdc\xd1\xe6\x44\xf9\xdc\x8c\xd3\x48\xa5\x31\xee\xe9\x3c\x58\xb4\x0f\x4c\xc0\x48\x69\xd0\xc8\x98\x3e\x73\x54\xce\xae\x5e\xcb\x7f\xb2\x1a\xa2\x0a\xe8\x13\x3a\x5b\x1e\xba\x05\xb7\x7a\xe5\x6c\x4c\x41\xa2\x05\xcd\x08\xd9\x60\xd3\xd5\x7c\xfe\x71\xa8\x5f\x62\xdb\x23\xac\x5c\x80\xed\xa1\xdc\x5b\xa0\xde\x95\xb9\xa5\x90\x69\x57\xaa\x00\x85\x52\x83\x81\x32\x19\xd5\x1d\x96\x7c\x67\xa6\xf9\x63\x5f\xba\x31\x46\xde\x01\x00\x00\xff\xff\x5c\x01\x7e\x15\xf9\x01\x00\x00")

func postgres000001_create_teamsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres000001_create_teamsDownSql,
		"postgres/000001_create_teams.down.sql",
	)
}

func postgres000001_create_teamsDownSql() (*asset, error) {
	bytes, err := postgres000001_create_teamsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/000001_create_teams.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres000001_create_teamsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\x4f\x6f\xb2\x40\x10\xc6\xef\x7c\x8a\x39\x4a\xf2\x1e\x7c\x6d\x35\x4d\x3c\x51\xdd\xa6\xa4\x8a\x2d\x62\xa3\x27\xb2\xb2\x13\x3b\x09\xec\x12\x58\xdb\xfa\xed\x1b\xd9\x45\xc5\x3f\x09\x5e\x9f\x67\xf9\xcd\x30\xcf\xcc\x28\x64\x5e\xc4\x20\xf2\x9e\x27\x0c\xfc\x17\x08\x66\x11\xb0\xa5\x3f\x8f\xe6\xa0\x91\x67\x25\x74\x1c\x00\x00\x12\xf0\xe9\x85\xa3\x57\x2f\xec\xf4\x06\x2e\xbc\x87\xfe\xd4\x0b\x57\xf0\xc6\x56\xff\x2a\x3f\x29\x90\x6b\xe4\x1a\xd6\xb4\x21\xa9\x8d\xb8\xcd\xc5\xa5\x28\x30\xc5\x4b\x91\xca\x3c\xe5\x3b\xc9\x33\x3c\xd4\x19\x3c\xba\xc6\xbc\xae\x0a\x2c\x93\x82\x72\x4d\x4a\x1e\x5b\xeb\xf7\xad\x8b\x19\xa7\xf4\xa0\xff\xef\x3d\x59\x5d\xef\x72\xbc\xf2\x3c\x51\x59\xce\xe5\x8d\xfa\x3c\x4d\xd5\x0f\x0a\xa1\x32\x4e\xb2\x3c\x42\xbb\xdd\xae\x7d\x41\xf2\x9b\x34\x9e\xcc\xe8\xa1\x67\x9d\x32\xf9\xc2\x0c\x9b\xd3\x33\xce\x22\xf0\x3f\x16\xac\xb3\x2f\xe9\x3a\xee\xd0\x71\x6c\x12\x7e\x30\x66\xcb\xb3\x24\x48\xfc\xc6\x55\x1a\x71\xd5\xe1\x2c\xa8\xb3\xa9\xbe\x86\x61\xbb\x6f\x4d\x97\x31\x89\x13\x40\xdd\xb9\xdb\x92\x61\x22\x8d\xb9\x3e\x61\xd4\x31\xb7\x65\x98\x5d\x69\x32\xea\xfd\x69\xcb\x30\x5b\xd4\x64\xd4\x9b\xd5\x96\x61\xb2\x69\xce\xa3\xce\x6b\x1f\x88\x37\x89\x58\x68\x2f\xc3\xd8\xde\x78\x0c\xa3\xd9\x64\x31\x0d\xce\xa0\xd5\x8a\xa8\x1c\xa5\x99\x27\xac\x95\x4a\x91\xcb\xe1\x3d\x8c\x94\x97\x7a\xff\x84\x12\x25\xcd\x48\xed\x89\xdc\x45\xb9\x75\x16\x77\x41\x36\x85\xda\xe6\x89\x92\xa5\x2e\x38\x49\x14\xc7\xff\xf9\x0b\x00\x00\xff\xff\xd9\x46\x12\x20\x2f\x04\x00\x00")

func postgres000001_create_teamsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres000001_create_teamsUpSql,
		"postgres/000001_create_teams.up.sql",
	)
}

func postgres000001_create_teamsUpSql() (*asset, error) {
	bytes, err := postgres000001_create_teamsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/000001_create_teams.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres000002_create_team_membersDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x49\x4d\xcc\xcd\x4d\xcd\x4d\x4a\x2d\x2a\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x4e\xce\x48\xcd\x4d\x2d\x2d\x4e\x2d\xb2\xe6\x22\x47\x67\x62\x4a\x6e\x66\x1e\x79\x5a\xd3\x4b\x53\x8b\x4b\x48\xd5\x9a\x92\x9a\x93\x5a\x92\x9a\x58\x62\xcd\xc5\x05\x56\xe0\xe9\xe7\xe2\x1a\x81\x24\x9f\x99\x52\x11\x8f\x64\x0a\x98\x1d\x9f\x99\x62\x4d\x9c\x6a\x50\x38\x10\xaf\x1a\xe2\x96\x78\x84\x63\x20\x9e\x40\x68\x08\x49\x4d\xcc\xf5\x85\x28\xb6\xe6\x02\x04\x00\x00\xff\xff\x79\x56\x1c\x35\x98\x01\x00\x00")

func postgres000002_create_team_membersDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres000002_create_team_membersDownSql,
		"postgres/000002_create_team_members.down.sql",
	)
}

func postgres000002_create_team_membersDownSql() (*asset, error) {
	bytes, err := postgres000002_create_team_membersDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/000002_create_team_members.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres000002_create_team_membersUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\x31\x6b\xf3\x30\x10\x86\x77\xff\x8a\x1b\x2d\xc8\xf4\xf1\x91\xc5\x93\x6a\xab\xd4\x54\x96\x8b\xa2\x94\x64\x32\x72\x75\xa4\x02\xcb\x06\x4b\x81\xfe\xfc\x62\xcb\x29\xc1\x2d\x85\x34\x9b\x38\x1d\xcf\x3d\xaf\x4e\xb9\x64\x54\x31\x50\xf4\x81\x33\x28\x1f\x41\xd4\x0a\xd8\xa1\xdc\xa9\x1d\x04\xd4\xce\xa1\x6b\x71\xf4\x90\x26\x00\x30\x57\xac\x81\x57\x2a\xf3\x27\x2a\xd3\x7f\x5b\x32\xf7\x8b\x3d\xe7\x9b\xb9\xe1\xec\x71\xfc\xb5\x61\x1c\x3a\xf4\x5f\xf7\xdb\xff\x24\x96\x0d\x76\x18\x50\x07\x68\xed\xc9\xf6\x21\x16\x5f\x64\x59\x51\x79\x84\x67\x76\x84\x34\x8e\xde\x2c\x13\x48\x42\xb2\x24\x59\xdc\x4b\x51\xb0\xc3\xca\xdd\x9a\x8f\xe6\xca\x7f\x3e\x37\xd6\x40\x2d\x40\xa1\x76\xd5\x25\x56\xc4\x92\xec\x16\xd6\xa4\xf0\x13\x6b\x51\xbb\x89\x15\x83\x37\x3a\x7c\xa3\x5d\x9e\x64\x0a\x4a\xb9\x62\x72\xd9\xd1\xf5\x56\x68\x51\x40\x5e\xf3\x7d\x25\x56\x63\xfc\xdb\x3b\x3a\x9c\x84\xa0\x1d\x86\x0e\x75\x9f\xfd\x15\xa2\x8d\xb3\xfd\xdd\x94\xd3\x19\x7d\xb8\x83\xb2\xfa\x20\x59\xf2\x19\x00\x00\xff\xff\x00\x64\x8f\x9f\xb8\x02\x00\x00")

func postgres000002_create_team_membersUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres000002_create_team_membersUpSql,
		"postgres/000002_create_team_members.up.sql",
	)
}

func postgres000002_create_team_membersUpSql() (*asset, error) {
	bytes, err := postgres000002_create_team_membersUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/000002_create_team_members.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres000003_create_compliancesDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\xce\xcf\x2d\xc8\xc9\x4c\xcc\x4b\x4e\x2d\xb6\xe6\x02\x04\x00\x00\xff\xff\xb8\x61\x53\x41\x22\x00\x00\x00")

func postgres000003_create_compliancesDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres000003_create_compliancesDownSql,
		"postgres/000003_create_compliances.down.sql",
	)
}

func postgres000003_create_compliancesDownSql() (*asset, error) {
	bytes, err := postgres000003_create_compliancesDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/000003_create_compliances.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres000003_create_compliancesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\x41\x4b\x87\x40\x10\x05\xf0\xbb\x9f\x62\xf8\x9f\x14\x3a\xf4\x97\xf2\xd2\x69\x93\x8d\x24\xb3\x58\xb7\xc8\xe3\xb6\x3b\xc8\x92\xae\xb2\x3b\x12\x7e\xfb\x40\xcd\xc8\xba\xfe\xde\x9b\xc7\xe4\x82\x33\xc9\x41\xb2\xdb\x92\x43\x71\x07\xd5\x93\x04\xfe\x56\xd4\xb2\x06\x3d\xf4\x63\x67\x95\xd3\x18\x20\x8e\x00\x00\xac\x81\x57\x26\xf2\x7b\x26\xe2\x34\x4b\x96\x6e\xf5\x52\x96\x17\x4b\xa8\x3d\x2a\x42\x45\xf0\x6e\x5b\xeb\x68\xc5\x29\xa0\xff\x7d\xb5\x7a\x20\x45\x53\xd8\x3d\xbb\xda\x5c\x0f\x93\x23\xb0\x8e\xb0\x45\xbf\xd2\xc9\x60\xd0\xa7\xbd\x7a\x7d\x4e\xb7\x2e\xcd\x23\xfe\x5d\x08\xa4\x3c\x1d\xbe\x40\x67\x0e\xf2\x81\xf3\xe7\xe0\x4d\xf8\x67\x16\x7b\x65\xbb\x9f\xe0\x7c\x99\x7e\x4f\x3f\x8b\xe2\x91\x89\x06\x1e\x78\x03\xb1\x35\x49\x94\xdc\x44\x5f\x01\x00\x00\xff\xff\xda\xde\xa3\xec\x40\x01\x00\x00")

func postgres000003_create_compliancesUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres000003_create_compliancesUpSql,
		"postgres/000003_create_compliances.up.sql",
	)
}

func postgres000003_create_compliancesUpSql() (*asset, error) {
	bytes, err := postgres000003_create_compliancesUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/000003_create_compliances.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"mysql/000001_create_teams.down.sql":           mysql000001_create_teamsDownSql,
	"mysql/000001_create_teams.up.sql":             mysql000001_create_teamsUpSql,
	"mysql/000002_create_team_members.down.sql":    mysql000002_create_team_membersDownSql,
	"mysql/000002_create_team_members.up.sql":      mysql000002_create_team_membersUpSql,
	"mysql/000003_create_compliances.down.sql":     mysql000003_create_compliancesDownSql,
	"mysql/000003_create_compliances.up.sql":       mysql000003_create_compliancesUpSql,
	"postgres/000001_create_teams.down.sql":        postgres000001_create_teamsDownSql,
	"postgres/000001_create_teams.up.sql":          postgres000001_create_teamsUpSql,
	"postgres/000002_create_team_members.down.sql": postgres000002_create_team_membersDownSql,
	"postgres/000002_create_team_members.up.sql":   postgres000002_create_team_membersUpSql,
	"postgres/000003_create_compliances.down.sql":  postgres000003_create_compliancesDownSql,
	"postgres/000003_create_compliances.up.sql":    postgres000003_create_compliancesUpSql,
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
	"mysql": &bintree{nil, map[string]*bintree{
		"000001_create_teams.down.sql":        &bintree{mysql000001_create_teamsDownSql, map[string]*bintree{}},
		"000001_create_teams.up.sql":          &bintree{mysql000001_create_teamsUpSql, map[string]*bintree{}},
		"000002_create_team_members.down.sql": &bintree{mysql000002_create_team_membersDownSql, map[string]*bintree{}},
		"000002_create_team_members.up.sql":   &bintree{mysql000002_create_team_membersUpSql, map[string]*bintree{}},
		"000003_create_compliances.down.sql":  &bintree{mysql000003_create_compliancesDownSql, map[string]*bintree{}},
		"000003_create_compliances.up.sql":    &bintree{mysql000003_create_compliancesUpSql, map[string]*bintree{}},
	}},
	"postgres": &bintree{nil, map[string]*bintree{
		"000001_create_teams.down.sql":        &bintree{postgres000001_create_teamsDownSql, map[string]*bintree{}},
		"000001_create_teams.up.sql":          &bintree{postgres000001_create_teamsUpSql, map[string]*bintree{}},
		"000002_create_team_members.down.sql": &bintree{postgres000002_create_team_membersDownSql, map[string]*bintree{}},
		"000002_create_team_members.up.sql":   &bintree{postgres000002_create_team_membersUpSql, map[string]*bintree{}},
		"000003_create_compliances.down.sql":  &bintree{postgres000003_create_compliancesDownSql, map[string]*bintree{}},
		"000003_create_compliances.up.sql":    &bintree{postgres000003_create_compliancesUpSql, map[string]*bintree{}},
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
