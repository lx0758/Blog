package service

import (
	"blog/bean/po"
	"blog/database"
	"errors"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var gFileUrlPrefix = "/files/"
var gFileUploadDirName = "files"
var gFileUploadDir string

func init() {
	pwd, _ := os.Getwd()
	dirSeparator := filepath.Separator
	gFileUploadDir = pwd + string(dirSeparator) + gFileUploadDirName + string(dirSeparator)
}

type FileService struct {
	Service
	db *gorm.DB
}

func (s *FileService) OnInitService() {
	s.db = database.GetDB()
}

func (s *FileService) Count() int {
	var count int64
	s.db.Model(&po.File{}).
		Count(&count)
	return int(count)
}

func (s *FileService) GetFileUrl(file po.File) string {
	return gFileUrlPrefix + strings.ReplaceAll(file.Path, string(filepath.Separator), "/")
}

func (s *FileService) PaginationByAdmin(
	fileName *string,
	fileType *string,
	pageNum int,
	pageSize int,
	orderName *string,
	orderMethod *string,
) po.Pagination[po.File] {
	db := s.db.Model(&po.File{}).
		Preload("Author")
	if fileName != nil && *fileName != "" {
		db = db.Where("upper(?) LIKE upper(?)", clause.Column{Name: "display_name"}, "%"+*fileName+"%")
	}
	if fileType != nil && *fileType != "" {
		db = db.Where("upper(?) LIKE upper(?)", clause.Column{Name: "type"}, "%"+*fileType+"%")
	}
	db = db.Order(database.TableOrderBy(orderName, orderMethod, "id", true))
	return database.Pagination[po.File](db, pageNum, pageSize)
}

func (s *FileService) ListNewFile(count int) []po.File {
	var files []po.File
	s.db.Model(&po.File{}).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
			Desc:   true,
		}).
		Limit(count).
		Find(&files)
	return files
}

func (s *FileService) AddByAdmin(userId int, file multipart.FileHeader) *po.File {
	filename := file.Filename
	fileSize := file.Size
	fileType := file.Header.Get("Content-Type")
	localFilePath := s.generatorPath(filename)
	localFileStorePath := gFileUploadDir + localFilePath

	uploadFile, err := file.Open()
	defer func() {
		if uploadFile != nil {
			_ = uploadFile.Close()
		}
	}()
	if err != nil {
		return nil
	}
	_ = os.MkdirAll(filepath.Dir(localFileStorePath), 0644)
	localFile, err := os.Create(localFileStorePath)
	defer func() {
		if localFile != nil {
			_ = localFile.Close()
		}
	}()
	if err != nil {
		return nil
	}
	_, err = io.CopyN(localFile, uploadFile, file.Size)
	if err != nil {
		return nil
	}

	filePO := &po.File{
		DisplayName: filename,
		Length:      fileSize,
		Type:        fileType,
		Path:        localFilePath,
		AuthorId:    userId,
		CreateTime:  time.Now(),
	}
	s.db.Model(&po.File{}).
		Create(&filePO)
	if filePO.Id == 0 {
		return nil
	}
	refreshBlogCacheInfo()
	return filePO
}

func (s *FileService) UpdateByAdmin(id int, file multipart.FileHeader) error {
	filePO := po.File{}
	s.db.Model(&po.File{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		First(&filePO)
	if filePO.Id == 0 {
		return errors.New("文件信息不存在")
	}

	fileSize := file.Size
	fileType := file.Header.Get("Content-Type")
	localFileStorePath := gFileUploadDir + filePO.Path

	if filePO.Type != fileType {
		return errors.New("文件类型不匹配")
	}

	uploadFile, err := file.Open()
	defer func() {
		if uploadFile != nil {
			_ = uploadFile.Close()
		}
	}()
	if err != nil {
		return nil
	}
	localFile, err := os.Create(localFileStorePath)
	defer func() {
		if localFile != nil {
			_ = localFile.Close()
		}
	}()
	if err != nil {
		return nil
	}
	_, err = io.CopyN(localFile, uploadFile, file.Size)
	if err != nil {
		return nil
	}

	updateTime := time.Now()
	db := s.db.Model(&po.File{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		Select("length", "update_time").
		Updates(&po.File{
			Length:     fileSize,
			UpdateTime: &updateTime,
		})
	if db.RowsAffected > 0 {
		refreshBlogCacheInfo()
	}
	return db.Error
}

func (s *FileService) DeleteByAdmin(id int) bool {
	file := po.File{}
	db := s.db.Model(&po.File{}).
		Clauses(clause.Returning{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		Delete(&file)
	if db.RowsAffected > 0 {
		s.cleanFile(file.Path)
		refreshBlogCacheInfo()
	}
	return db.RowsAffected > 0
}

func (s *FileService) generatorPath(filename string) string {
	layout := "2006" + string(filepath.Separator) + "01" + string(filepath.Separator)
	dateDir := time.Now().Format(layout)
	uuidDir := s.generateShortUuid() + string(filepath.Separator)
	return dateDir + uuidDir + filename
}

var (
	uuidRandom = rand.New(rand.NewSource(time.Now().UnixNano()))
	uuidChars  = []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}
)

func (s *FileService) generateShortUuid() string {
	uuid := make([]byte, 16)
	_, _ = uuidRandom.Read(uuid)
	// Set version to 0100
	uuid[6] = uuid[6]&0x0f | 0x40
	// Set bits 6-7 to 10
	uuid[8] = uuid[8]&0x3f | 0x80

	var result strings.Builder
	var height, low uint8
	var partValue int
	for i := 0; i < len(uuid); i += 2 {
		height = uuid[i]
		low = uuid[i+1]
		partValue = int(height)<<8 | int(low)
		result.WriteString(uuidChars[partValue%62])
	}
	return result.String()
}

func (s *FileService) cleanFile(path string) {
	fileUploadDir := filepath.Dir(gFileUploadDir)
	localFileStorePath := gFileUploadDir + path
	currentPath := localFileStorePath
	parentPath := filepath.Dir(currentPath)
	for {
		_ = os.Remove(currentPath)

		entries, _ := os.ReadDir(parentPath)
		if len(entries) == 0 {
			currentPath = parentPath
			parentPath = filepath.Dir(currentPath)
		} else {
			break
		}

		if fileUploadDir == currentPath {
			break
		}
	}
}
