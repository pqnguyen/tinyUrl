package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/pqnguyen/tinyUrl/services/sample"
)

type sampleRepository struct {
	db *gorm.DB
}

func NewSampleRepository(db *gorm.DB) sample.Repository {
	return &sampleRepository{db}
}
