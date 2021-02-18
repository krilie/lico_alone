package dbmigrate

import (
	"embed"
	"errors"
	"github.com/golang-migrate/migrate/v4/source"
	"io"
	"sort"
	"strconv"
	"strings"
)

//go:embed migrations
var sqlFiles embed.FS

type FsDriver struct {
	fs       *embed.FS
	prefix   string // migrations
	fileInfo FileInfoLocals
}

type FileInfoLocal struct {
	Name    string
	Version uint
}
type FileInfoLocals []FileInfoLocal

func (f FileInfoLocals) Len() int {
	return len(f)
}

func (f FileInfoLocals) Less(i, j int) bool {
	return f[i].Version < f[j].Version
}

func (f FileInfoLocals) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f *FsDriver) Open(url string) (source.Driver, error) {
	dir, err := f.fs.ReadDir(f.prefix)
	if err != nil {
		return nil, err
	}
	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}
		version, err := parserUintVersion(entry.Name())
		if err != nil {
			return nil, err
		}
		f.fileInfo = append(f.fileInfo, FileInfoLocal{
			Name:    entry.Name(),
			Version: version,
		})
	}
	sort.Sort(f.fileInfo)

	return f, nil
}

func (f *FsDriver) Close() error {
	f.fileInfo = nil
	f.fs = nil
	return nil
}

func (f *FsDriver) First() (version uint, err error) {
	if f.fileInfo.Len() > 0 {
		return f.fileInfo[0].Version, nil
	}
	return 0, errors.New("no first version find")
}

func (f *FsDriver) Prev(version uint) (prevVersion uint, err error) {
	for index, local := range f.fileInfo {
		if local.Version == version {
			if index == 0 {
				return 0, errors.New("no prev")
			}
			return f.fileInfo[index-1].Version, nil
		}
	}
	return 0, errors.New("no file find")
}

func (f *FsDriver) Next(version uint) (nextVersion uint, err error) {
	for index, local := range f.fileInfo {
		if local.Version == version {
			if index == f.fileInfo.Len()-1 {
				return 0, errors.New("no next")
			}
			return f.fileInfo[index+1].Version, nil
		}
	}
	return 0, errors.New("no file find")
}

func (f *FsDriver) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	for _, local := range f.fileInfo {
		if local.Version == version {
			open, err := f.fs.Open(f.prefix + "/" + local.Name)
			if err != nil {
				return nil, "", err
			}
			return open, local.Name, nil
		}
	}
	return nil, "", errors.New("no file find")
}

func (f *FsDriver) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	panic("implement me")
}

func parserUintVersion(name string) (v uint, err error) {
	splitN := strings.SplitN(name, "_", 1)
	if len(splitN) < 1 {
		return 0, errors.New("migrations version error")
	}
	parseUint, err := strconv.ParseUint(splitN[0], 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(parseUint), nil
}
