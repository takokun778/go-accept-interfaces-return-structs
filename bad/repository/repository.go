package repository

type Repository interface {
	Save(any) error
	Find(any) (any, error)
}
