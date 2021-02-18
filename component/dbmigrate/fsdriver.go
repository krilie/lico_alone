package dbmigrate

import (
	"embed"
	"github.com/golang-migrate/migrate/v4/source"
	"io"
)

//go:embed migrations
var sqlFiles embed.FS

type FsDriver struct {
	fs *embed.FS
}

func (f *FsDriver) Open(url string) (source.Driver, error) {
	panic("implement me")
}

func (f *FsDriver) Close() error {
	panic("implement me")
}

func (f *FsDriver) First() (version uint, err error) {
	panic("implement me")
}

func (f *FsDriver) Prev(version uint) (prevVersion uint, err error) {
	panic("implement me")
}

func (f *FsDriver) Next(version uint) (nextVersion uint, err error) {
	panic("implement me")
}

func (f *FsDriver) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	panic("implement me")
}

func (f *FsDriver) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	panic("implement me")
}
