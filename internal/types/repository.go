package types

import "gorm.io/gorm"

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Create(entity *T) error {
	return r.DB.Create(entity).Error
}

func (r *Repository[T]) Update(entity *T) error {
	return r.DB.Save(entity).Error
}

func (r *Repository[T]) Delete(entity *T) error {
	return r.DB.Delete(entity).Error
}

func (r *Repository[T]) CountByUUID(uuid any) (int64, error) {
	var total int64
	err := r.DB.Model(new(T)).Where("uuid = ?", uuid).Count(&total).Error
	return total, err
}

func (r *Repository[T]) FindByUUID(uuid any) error {
	var entity *T
	return r.DB.Where("uuid = ?", uuid).Take(entity).Error
}

func (r *Repository[T]) FindAll(entities *[]T) error {
	return r.DB.Find(entities).Error
}

func (r *Repository[T]) Find(entity *T) error {
	return r.DB.First(entity, entity).Error
}

func (r *Repository[T]) Exists(entity *T) bool {
	return r.DB.First(entity, entity).RowsAffected > 0
}
