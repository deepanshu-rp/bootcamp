package persistence

import (
	"exam/domain/entity"
	"exam/domain/repository"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Repositories struct {
	Student repository.StudentRepository
	Marks   repository.MarksRepository
	db      *gorm.DB
}

func NewRepositories(Dbdriver, DbUser, DbPassword, DbHost, DbName string, DbPort int) (*Repositories, error) {
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		DbUser,
		DbPassword,
		DbHost,
		DbPort,
		DbName)
	db, err := gorm.Open(Dbdriver, DBURL)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &Repositories{
		Student: NewStudentGormRepository(db),
		Marks:   NewMarksGormRepository(db),
		db:      db,
	}, nil
}

//closes the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

//This migrate all tables
func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&entity.Student{}).Error
}
