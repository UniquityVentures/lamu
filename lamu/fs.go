package lamu

import "io/fs"

type UsefulFilesystem interface {
	fs.FS
	fs.ReadDirFS
	fs.ReadFileFS
}
