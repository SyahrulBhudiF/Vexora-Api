package _interface

type IRepository[T any] interface {
	Create(entity *T) error
	Update(entity *T) error
	Delete(entity *T) error
	CountByUUID(uuid any) (int64, error)
	FindByUUID(uuid any) error
	FindAll(entities *[]T) error
	Find(entity *T) error
	Exists(entity *T) bool
}
