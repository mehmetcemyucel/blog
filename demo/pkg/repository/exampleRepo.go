package repository

import (
	"gorm.io/gorm"
	"helloWorld/pkg/model/example/entity"
)

type IExampleRepository interface {
	ActiveExamples(name string) ([]entity.Example, error)
}

type ExampleRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IExampleRepository {
	var repo IExampleRepository
	repo = &ExampleRepository{db}
	return repo
}

func (p *ExampleRepository) ActiveExamples(name string) ([]entity.Example, error) {
	var examples []entity.Example
	err := p.db.Where(`name = ? and deleted_at is null`, name).Find(&examples).Error
	return examples, err
}
