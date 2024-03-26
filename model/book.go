package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	Model
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ISBN      string         `json:"isbn"`
	Penulis   string         `json:"penulis"`
	Tahun     uint           `json:"tahun"`
	Judul     string         `json:"judul"`
	Gambar    string         `json:"gambar"`
	Stok      uint           `json:"stok"`
}

func (bk *Book)Create(db *gorm.DB)error  {
	err := db.Model(book{}).Create(&bk).Error
	if err != nil {
		return err
	}
	return err
}

func (bk *Book) GetByID(db *gorm.DB) (Book, error) {
	res := Book{}
	err := db.Model(Book{}). Where("id = ?", bk.Model.ID).Take(&res).Error

	if err != nil{
		return Book{}, err
	}
	return res, nil
}

func (bk *Book)GetAll(db *gorm.DB)([]Book, error)  {
	res := []Book{}

	err := db.
	Model(Book{}).
	Find(&res).
	Error

	if err != nil{
		return []Book{}, err
	}
	return res, nil
}


func (bk *Book)UpdateOne(db *gorm.DB)error{
	err := db.
	Model(Car{}).
	Select("id", "created_at", "updated_at", "deleted_at","isbn","penulis", "tahun", "judul", "gambar", "stok").Where("id = ?", bk.Model.ID).
	Updates(map[string]interface{}{
		"id": bk.ID, 
		"created_at" : bk.CreatedAt,
		 "updated_at" : bk.UpdatedAt,
		"deleted_at" : bk.DeletedAt,
		"isbn" : bk.ISBN,
		"penulis": bk.Penulis,
		"tahun" : bk.Tahun,
		"judul" : bk.Judul,
		"gambar" : bk.Gambar,
		"stok" : bk.Stok,
	}).
	Error

	if err != nil{
		return err
	}
	return nil
}

func (bk *Book) DeleteByID(db *gorm.DB) error{
	err := db.
		Model(Book{}).
		Where("id = ?", bk.Model.ID)
		Delete(&bk).
		Error

		if err != nil{
			return err
		}
		return nil
}