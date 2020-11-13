package Controller

import (
	"Book/Business"
	"Book/Controller/payload"
	"Book/Controller/response"
	"Book/Database/Models"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

type CategoryController interface {
	CreateCategory(w http.ResponseWriter, r *http.Request)
	GetListCategory(w http.ResponseWriter, r *http.Request)
	GetCategoryById(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
}
type categoryController struct {
	categoryBusiness Business.CategoryBusiness
}
// Create new book godoc
// @tags Category
// @Summary Create new book
// @Description create new book
// @Accept json
// @Produce json
// @Param product body payload.Category true "product information"
// @Success 200 {object} response.CategoryResponse
// @Router /categories [post]
func (c categoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var pl payload.Category

	err := json.NewDecoder(r.Body).Decode(&pl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}
	cate := Models.Category{
		Status: pl.Status,
		Name:   pl.Name,
	}
	result, err := c.categoryBusiness.CreateCategory(cate)
	res := &response.CategoryResponse{}
	if err != nil {
		res = &response.CategoryResponse{
			Category: nil,
			Meta:     response.Meta{
				Success: false,
				Message: err.Error(),
			},
		}
	} else {
		res = &response.CategoryResponse{
			Category: result,
			Meta:     response.Meta{
				Success: true,
				Message: "success",
			},
		}
	}


	render.JSON(w, r, res)
}
// Filter get list category godoc
// @tags Category
// @Summary get list Category
// @Description  list Category
// @Accept json
// @Produce json
// @Success 200 {object} response.ListCategoryResponse
// @Router /categories/ [get]
func (c categoryController) GetListCategory(w http.ResponseWriter, r *http.Request) {
	results, err := c.categoryBusiness.GetAllActive()
	res := &response.ListCategoryResponse{}
	if err != nil {
		res = &response.ListCategoryResponse{
			Category: nil,
			Meta:     response.Meta{
				Success: false,
				Message: err.Error(),
			},
		}
	} else {
		res = &response.ListCategoryResponse{
			Category: results,
			Meta:     response.Meta{
				Success: false,
				Message: "success",
			},
		}
	}

	render.JSON(w, r, res)
}

func (c categoryController) GetCategoryById(w http.ResponseWriter, r *http.Request) {

}
// update category by code godoc
// @tags Category
// @Summary Active Book by code
// @Description update Book
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param category body payload.Category true "category information"
// @Param code path string true "code"
// @Success 200 {object} response.CategoryResponse
// @Router /categories/{code} [put]
func (c categoryController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var pl payload.Category
	id := chi.URLParam(r, "code")
	log.Print(id)
	err := json.NewDecoder(r.Body).Decode(&pl)
	if err != nil {
		log.Print("1")
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}
	res := &response.CategoryResponse{}
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}
	cate := Models.Category{
		Id: Id,
		Status: pl.Status,
		Name:   pl.Name,
	}
	result, err := c.categoryBusiness.Update(cate)

	if err != nil {
		res = &response.CategoryResponse{
			Category: nil,
			Meta:     response.Meta{
				Success: false,
				Message: err.Error(),
			},
		}
	} else {
		res = &response.CategoryResponse{
			Category: result,
			Meta:     response.Meta{
				Success: true,
				Message: "success",
			},
		}
	}
	render.JSON(w, r, res)
}

func NewCategoryController() CategoryController {
	categoryBusiness := Business.NewCategoryBusiness()
	return &categoryController{
		categoryBusiness: categoryBusiness,
	}
}
