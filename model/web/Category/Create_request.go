package Web_category

type CreateRequest struct {
	Name string `validate:"required,max:200,min=1"`
}
