package model

import (
	"gorm.io/gorm"
)

type Book struct {
	Model
	ISBN    string `json:"isbn"`
	Penulis string `json:"penulis"`
	Tahun   uint   `json:"tahun"`
	Judul   string `json:"judul"`
	Gambar  string `json:"gambar"`
	Stok    uint   `json:"stok"`
}

func (bk *Book) Create(db *gorm.DB) error{
	err := db.Model(Book{}).Create(&bk).Error
	if err != nil{
		return err
	}
	return err
}

func (bk *Book) UpdateOneByID(db *gorm.DB) error{
	err := db.
	Model(Book{}).
	Select("isbn", "penulis", "tahun", "judul", "gambar", "stok").Where("id = ?", bk.Model.ID).
	Updates(map[string]interface{}{
		"isbn": bk.ISBN, "penulis": bk.Penulis, "tahun": bk.Penulis, "judul": bk.Judul, "gambar": bk.Judul, "stok": bk.Stok,
	}).
	Error

	if err != nil{
		return err
	}
	return nil
}

func (bk *Book) GetAll(db *gorm.DB) ([]Book, error){
	res := []Book{}

	err := db.Model(Book{}). Find(&res). Error

	if err != nil {
		return []Book{}, err
	}
	return res, nil
}


func (bk *Book) DeleteByID(db *gorm.DB) error  {
	err := db. 
	Model(Book{}).Where("id = ?", bk.Model.ID).Delete(&bk). 
	Error 	
	if err != nil{
		return err
	}
	return nil}

