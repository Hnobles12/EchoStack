package repositories

import (
	"gorm.io/gorm"
)

// Repository provides CRUD operations for any model type.
type Repository[T any] struct {
	DB *gorm.DB
}

// NewRepository creates a new instance of Repository.
func NewRepository[T any](db *gorm.DB) *Repository[T] {
	return &Repository[T]{DB: db}
}

// Create inserts a new record in the database.
func (r *Repository[T]) Create(entity *T) error {
	return r.DB.Create(entity).Error
}

// FindAll retrieves all records of the model type with optional preloading of associations.
func (r *Repository[T]) FindAll(preloads ...string) ([]T, error) {
	var entities []T
	query := r.DB
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	if err := query.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

// FindByID retrieves a record by its ID with optional preloading of associations.
func (r *Repository[T]) FindByID(id any, preloads ...string) (*T, error) {
	var entity T
	query := r.DB
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	if err := query.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Update updates a record in the database.
func (r *Repository[T]) Update(entity *T) error {
	return r.DB.Save(entity).Error
}

// Delete removes a record from the database.
func (r *Repository[T]) Delete(entity *T) error {
	return r.DB.Delete(entity).Error
}

// DeleteByID removes a record from the database by its ID.
func (r *Repository[T]) DeleteByID(id any) error {
	return r.DB.Delete(new(T), id).Error
}
