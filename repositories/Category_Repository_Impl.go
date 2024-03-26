package repositories

import (
	"GoResfulApiPribadi/helpers"
	"GoResfulApiPribadi/model"
	"context"
	"database/sql"
	"fmt"
)

type CategoryRepositoryImpl struct {
}

func CategoryRepositoryConstruct() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Get(ctx context.Context, sqltx *sql.Tx) {
	var category model.Category
	sqlRow, err := sqltx.QueryContext(ctx, "SELECT * FROM Category")
	helpers.PanicIfError(err)
	for sqlRow.Next() {
		if err := sqlRow.Scan(&category.Id, &category.Name); err != nil {
			panic(err.Error())
		}
	}
	fmt.Println(category)
}

func (repository *CategoryRepositoryImpl) GetById() {

}

func (repository *CategoryRepositoryImpl) Insert() {

}

func (repository *CategoryRepositoryImpl) Update() {

}

func (repository *CategoryRepositoryImpl) Delete() {
}
