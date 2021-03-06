package usecase

import (
	"github.com/pqnguyen/tinyUrl/services/sample"
)

type sampleUseCase struct {
	sampleRepo sample.Repository
}

func NewSampleUsecase(sample sample.Repository) sample.UseCase {
	return &sampleUseCase{
		sampleRepo: sample,
	}
}
