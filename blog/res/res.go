package res

import (
	"embed"
	"io/fs"
	"strings"
)

var (
	//go:embed comic.ttf
	FontFile []byte

	//go:embed static/blog/*
	blogStaticFS embed.FS
	BlogStaticFS = createSafeFS(blogStaticFS, "static/blog/", false)

	//go:embed static/admin/*
	adminStaticFS embed.FS
	AdminStaticFS = createSafeFS(adminStaticFS, "static/admin/", true)

	//go:embed templates/*
	templatesFS embed.FS
	TemplatesFS = createFS(templatesFS, "templates/")
)

func createFS(fs embed.FS, pathPrefix string) fs.FS {
	return createSafeFS(fs, pathPrefix, true)
}

// resFS 将 embed.FS 的访问添加统一的前缀
func createSafeFS(fs embed.FS, relativePath string, listDirectory bool) fs.FS {
	if strings.HasPrefix(relativePath, "/") {
		panic("path must is a relative folder")
	}
	if !strings.HasSuffix(relativePath, "/") {
		panic("path must is a folder")
	}
	return resFS{fs, relativePath, listDirectory}
}

type resFS struct {
	embed.FS
	relativePath  string
	listDirectory bool
}

func (resFS resFS) Open(name string) (fs.File, error) {
	isDir, realName := resFS.handleRealName(name)
	if !resFS.listDirectory && isDir {
		return nil, fs.ErrPermission
	} else {
		return resFS.FS.Open(realName)
	}
}

func (resFS resFS) ReadDir(name string) ([]fs.DirEntry, error) {
	_, realName := resFS.handleRealName(name)
	return resFS.FS.ReadDir(realName)
}

func (resFS resFS) ReadFile(name string) ([]byte, error) {
	_, realName := resFS.handleRealName(name)
	return resFS.FS.ReadFile(realName)
}

func (resFS resFS) handleRealName(name string) (bool, string) {
	if name == "/" || name == "." {
		name = resFS.relativePath
	} else {
		name = resFS.relativePath + name
	}
	isDir := strings.HasSuffix(name, "/")
	realName := strings.TrimSuffix(name, "/")
	return isDir, realName
}
