package repository

import (
	"github.com/jinzhu/gorm"
	"talo.io/talo-api/services/sample"
)

type sampleRepository struct {
	db *gorm.DB
}

func NewSampleRepository(db *gorm.DB) sample.Repository {
	return &sampleRepository{db}
}
