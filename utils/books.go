package utils

import (
	"MiniProject3-sekolahbeta/config"
	"MiniProject3-sekolahbeta/model"
	"time"
)


func InserCarData(data model.Book) error {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	return data.Create(config.Mysql.DB)
	
}