package model_test

import (
	"MiniProject3-sekolahbeta/config"
	"MiniProject3-sekolahbeta/model"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("env not found, using global env")
	}
	config.OpenDB()
}


func TestCreateBook(t *testing.T)  {
	Init()

	bookData := model.Book{
		ISBN: "yoga",
		Penulis:"yoga",
		Tahun: 2012,
		Judul: "anu",
		Gambar: "anu",
		Stok: 10,
	}
	err := bookData.Create(config.Mysql.DB)
	assert.Nil(t, err)
	
}

func TestUpdate(t *testing.T)  {
	Init()

	bookData := model.Book{
		ISBN: "yoga",
		Penulis:"yoga",
		Tahun: 2012,
		Judul: "anu",
		Gambar: "anu",
		Stok: 10,
	}

	err := bookData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	bookData.ISBN = "sada"

	err = bookData.UpdateOneByID(config.Mysql.DB)
	assert.Nil(t, err)
	assert.NotEqual(t, bookData.ISBN,"yoga")	
	
}


func TestGetAll(t *testing.T)  {
	Init()

	bookData := model.Book{
		ISBN: "yoga",
		Penulis:"yoga",
		Tahun: 2012,
		Judul: "anu",
		Gambar: "anu",
		Stok: 10,
	}
	err := bookData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	res, err := bookData.GetAll(config.Mysql.DB)
	assert.Nil(t, err)
	assert.Greater(t, len(res), 0)
	// fmt.Println(res)
}

func TestDeleteByID(t *testing.T)  {
	Init()

	bookData := model.Book{
		Model: model.Model{
			ID: 25,
		},
	}
	err := bookData.DeleteByID(config.Mysql.DB)
	assert.Nil(t, err)
}