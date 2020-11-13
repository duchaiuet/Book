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
	"strconv"
)

type BookController interface {
	CreateBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
	GetListBook(w http.ResponseWriter, r *http.Request)
	GetBookById(w http.ResponseWriter, r *http.Request)
	ActiveBook(w http.ResponseWriter, r *http.Request)
	DeActiveBook(w http.ResponseWriter, r *http.Request)
}
type bookController struct {
	bookBusiness Business.BookBusiness
}

func (b bookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

// Create new book godoc
// @tags Book
// @Summary Create new book
// @Description create new book
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product body payload.CreateBookPayload true "product information"
// @Success 200 {object} response.CreateBookResponse
// @Router /books [post]
func (b bookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var pl payload.CreateBookPayload

	err := json.NewDecoder(r.Body).Decode(&pl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}
	categoryIds := make([]primitive.ObjectID, len(pl.CategoryIDs))
	if len(pl.CategoryIDs) > 0 {
		for i := 0; i < len(pl.CategoryIDs); i++ {
			log.Print("category: ",pl.CategoryIDs[i])
			categoryIds[i], err = primitive.ObjectIDFromHex(pl.CategoryIDs[i])
			if err != nil {
				log.Print(err.Error())
				break
			}
		}
	}
	res := &response.CreateBookResponse{}
	if err != nil {
		res = &response.CreateBookResponse{
			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		}
	}

	book := Models.Book{
		Code:        pl.Code,
		Name:        pl.Name,
		CategoryIDs: categoryIds,
		Images:      pl.Images,
		Price:       pl.Price,
		Author:      pl.Author,
		Status:      pl.Status,
		Quantity:    pl.Quantity,
	}
	data, err := b.bookBusiness.Create(book)

	if err != nil {
		res = &response.CreateBookResponse{
			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		}
	} else {
		res = &response.CreateBookResponse{
			Meta: response.Meta{
				Success: true,
				Message: "Success",
			},
			Data: data,
		}
	}
	render.JSON(w, r, res)
}

// Filter get list books godoc
// @tags Book
// @Summary get list books
// @Description  list books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query integer false "page"
// @Param page_size query integer false "page size each page"
// @Param text query string false "code or name"
// @Success 200 {object} response.ListBooksResponse
// @Router /books/ [get]
func (b bookController) GetListBook(w http.ResponseWriter, r *http.Request) {
	var page, pageSize int
	var err error
	page, err = strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil  {
		page = 1
	}
	pageSize, err = strconv.Atoi(r.URL.Query().Get("page_size"))
	if err != nil {
		pageSize = 10
	}
	text := r.URL.Query().Get("text")
	results, total, err := b.bookBusiness.Filter(text, "", true, nil, page, pageSize)
	res := &response.ListBooksResponse{}
	if err != nil {
		res = &response.ListBooksResponse{
			Meta:       response.Meta{
				Success: false,
				Message: err.Error(),
			},
			Data:       nil,
			Pagination: response.Pagination{
				Page:     page,
				PageSize: pageSize,
				Total:    total,
			},

		}
	} else {
		res = &response.ListBooksResponse{
			Meta:       response.Meta{
				Success: true,
				Message: "success",
			},
			Data:       results,
			Pagination: response.Pagination{
				Page:     page,
				PageSize: pageSize,
				Total:    total,
			},

		}
	}
	render.JSON(w, r, res)
}

// Filter get list books godoc
// @tags Book
// @Summary get  book by id
// @Description  book by id
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string false "id"
// @Success 200 {object} response.GetBookByIdResponse
// @Router /books/{id} [get]
func (b bookController) GetBookById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}
	log.Print(id)
	book, err := b.bookBusiness.GetById(id)
	res := &response.GetBookByIdResponse{}
	if err != nil {
		res = &response.GetBookByIdResponse{
			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		}
	} else {
		res = &response.GetBookByIdResponse{
			Meta: response.Meta{
				Success: true,
				Message: "Success",
			},
			Data: book,
		}
	}
	render.JSON(w, r, res)
}

// update book by code godoc
// @tags Book
// @Summary Active Book by code
// @Description update Book
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param code path string true "code"
// @Success 200 {object} response.StatusProductResponse
// @Router /books/activate/{code} [put]
func (b bookController) ActiveBook(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err := b.bookBusiness.Active(code)
	res := &response.StatusProductResponse{}
	if err != nil {
		res = &response.StatusProductResponse{
			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
		}
	} else {
		res = &response.StatusProductResponse{
			Meta: response.Meta{
				Success: true,
				Message: "Success",
			},
		}
	}
	render.JSON(w, r, res)
}

// update book by code godoc
// @tags Book
// @Summary DeActive Book by code
// @Description update Book
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param code path string true "code"
// @Success 200 {object} response.UpdateBookResponse
// @Router /books/deactivate/{code} [put]
func (b bookController) DeActiveBook(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err := b.bookBusiness.DeActive(code)
	res := &response.StatusProductResponse{}
	if err != nil {
		res = &response.StatusProductResponse{
			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
		}
	} else {
		res = &response.StatusProductResponse{
			Meta: response.Meta{
				Success: true,
				Message: "Success",
			},
		}
	}
	render.JSON(w, r, res)
}

func NewBookController() BookController {
	bookBusiness := Business.NewBookBusiness()
	return &bookController{
		bookBusiness: bookBusiness,
	}
}
