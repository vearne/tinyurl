package dao

import (
	"github.com/vearne/tinyurl/model"
	"github.com/vearne/tinyurl/resource"
)

func CreateTinyURL(url string) uint64 {
	item := models.TinyURL{}
	item.URL = url
	resource.MySQLClient.Create(&item)
	return item.ID
}

func GetURL(value uint64) string {
	item := models.TinyURL{ID: value}
	resource.MySQLClient.First(&item)
	return item.URL
}
