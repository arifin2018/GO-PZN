package Web_category

type UpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=200,min=1"`
}
