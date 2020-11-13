package Business

import (
	"Book/Database"
	"Book/Database/DAL"
	"Book/Database/Models"
)

type categoryBusiness struct {
	CategoryRepository Models.CategoryRepository 
}

func (c categoryBusiness) CreateCategory(category Models.Category) (*Models.Category, error) {
	cate, err := c.CategoryRepository.Create(category)
	if err != nil {
		Database.ErrLog.Print(err)
		return nil, err
	}
	return cate, err
}

func (c categoryBusiness) GetAllActive() ([]*Models.Category, error) {
	categories, err := c.CategoryRepository.Gets()
	if err != nil {
		Database.ErrLog.Print(err)
		return nil, err
	}
	return categories, err
}

func (c categoryBusiness) Update(category Models.Category) (*Models.Category, error) {
	cate, err := c.CategoryRepository.Update(category)
	if err != nil {
		Database.ErrLog.Print(err)
		return nil, err
	}
	return cate, err
}

type CategoryBusiness interface {
	CreateCategory(category Models.Category) (*Models.Category, error)
	GetAllActive() ([]*Models.Category, error)
	Update(category Models.Category) (*Models.Category, error)
}
func NewCategoryBusiness() CategoryBusiness{
	categoryRepository := DAL.NewCategoryRepository(Database.Db)
	return categoryBusiness{CategoryRepository: categoryRepository}
}
