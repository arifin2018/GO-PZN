package repositories

type CategoryRepository interface {
	Get()
	GetById()
	Insert()
	Update()
	Delete()
}
