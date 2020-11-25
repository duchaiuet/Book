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
	"net/http"
)

type userController struct {
	userBusiness Business.UserBusiness
}

// Create new user godoc
// @tags User
// @Summary Create new user
// @Description create new user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user body payload.User true "user information"
// @Success 200 {object} response.CreateBookResponse
// @Router /users [post]
func (u userController) Create(w http.ResponseWriter, r *http.Request) {
	var pl payload.User

	err := json.NewDecoder(r.Body).Decode(&pl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	res := &response.CreateUserResponse{}
	if err != nil {
		res = &response.CreateUserResponse{
			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		}
	}
	role, err := primitive.ObjectIDFromHex(pl.Role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	add := Models.Address{
		City:     pl.Address.City,
		District: pl.Address.District,
		Ward:     pl.Address.Ward,
		Text:     pl.Address.Text,
	}
	user := Models.User{
		Name:        pl.Name,
		Role:        role,
		UserName:    pl.UserName,
		Password:    pl.Password,
		PhoneNumber: pl.PhoneNumber,
		Address:     add,
		Status:      true,
	}
	data, err := u.userBusiness.Create(user)

	if err != nil {
		res = &response.CreateUserResponse{
			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		}
	} else {
		res = &response.CreateUserResponse{
			Meta: response.Meta{
				Success: true,
				Message: "Success",
			},
			Data: data,
		}
	}
	render.JSON(w, r, res)
}
// update user by user godoc
// @tags User
// @Summary DeActive Book by code
// @Description update Book
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "id"
// @Param user body payload.UserUpdate true "user information"
// @Success 200 {object} response.UpdateUserResponse
// @Router /users/{id} [put]
func (u userController) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}
	var pl payload.User
	err := json.NewDecoder(r.Body).Decode(&pl)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}
	res := &response.UpdateUserResponse{}
	user, err := u.userBusiness.Get(id)
	if err != nil {
		res = &response.UpdateUserResponse{
			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		}
	}
	userUpdate := Models.User{}
	newRole, err := primitive.ObjectIDFromHex(pl.Role)
	newAdd := Models.Address{
		City:     pl.Address.City,
		District: pl.Address.District,
		Ward:     pl.Address.Ward,
		Text:     pl.Address.Text,
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if user == nil {
		res = &response.UpdateUserResponse{
			Meta: response.Meta{
				Success: false,
				Message: "not found",
			},
			Data: nil,
		}
	} else{

		userUpdate = Models.User{
			Id:          user.Id,
			Name:        pl.Name,
			Role:        newRole,
			UserName:    pl.UserName,
			PhoneNumber: pl.PhoneNumber,
			Address:     newAdd,
		}
	}

	result, err := u.userBusiness.Update(id, userUpdate)

	if err != nil {
		res = &response.UpdateUserResponse{
			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
			 Data: nil,
		}
	} else {
		res = &response.UpdateUserResponse{
			Meta: response.Meta{
				Success: true,
				Message: "Success",
			},
			Data: result,
		}
	}
	render.JSON(w, r, res)
}

// Filter get list user godoc
// @tags User
// @Summary get list user
// @Description  list user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query integer false "page"
// @Param page_size query integer false "page size each page"
// @Param text query string false "code or name"
// @Success 200 {object} response.ListUserResponse
// @Router /users [get]
func (u userController) GetList(w http.ResponseWriter, r *http.Request) {
	//var page, pageSize int
	//var err error
	//page, err = strconv.Atoi(r.URL.Query().Get("page"))
	//if err != nil  {
	//	page = 1
	//}
	//pageSize, err = strconv.Atoi(r.URL.Query().Get("page_size"))
	//if err != nil {
	//	pageSize = 10
	//}
	//text := r.URL.Query().Get("text")
	users, err := u.userBusiness.GetAll()
	res := &response.ListUserResponse{}
	if err != nil {
		res = &response.ListUserResponse{
			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		}
	} else {
		res = &response.ListUserResponse{
			Meta: response.Meta{
				Success: true,
				Message: "success",
			},
			Data: users,
		}
	}
	render.JSON(w, r, res)
}

// Filter get list books godoc
// @tags User
// @Summary get  book by id
// @Description  book by id
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string false "id"
// @Success 200 {object} response.GetUserByIdResponse
// @Router /users/{id} [get]
func (u userController) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}
	user, err := u.userBusiness.Get(id)
	res := &response.GetUserByIdResponse{}
	if err != nil {
		res = &response.GetUserByIdResponse{
			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		}
	} else {
		res = &response.GetUserByIdResponse{
			Meta: response.Meta{
				Success: true,
				Message: "Success",
			},
			Data: user,
		}
	}
	render.JSON(w, r, res)
}

// update user by code godoc
// @tags User
// @Summary Active user by code
// @Description update user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} response.StatusUserResponse
// @Router /users/activate/{id} [put]
func (u userController) Active(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}
	res := &response.StatusUserResponse{}

	err := u.userBusiness.UpdateStatus(id, true)
	if err != nil {
		res = &response.StatusUserResponse{

			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
		}
	} else {
		res = &response.StatusUserResponse{
			Meta: response.Meta{
				Success: true,
				Message: "Success",
			},
		}
	}
	render.JSON(w, r, res)
}

// update user by code godoc
// @tags User
// @Summary Deactivate user by code
// @Description update user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} response.StatusUserResponse
// @Router /users/deactivate/{id} [put]
func (u userController) DeActive(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}
	res := &response.StatusUserResponse{}

	err := u.userBusiness.UpdateStatus(id, false)
	if err != nil {
		res = &response.StatusUserResponse{

			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
		}
	} else {
		res = &response.StatusUserResponse{
			Meta: response.Meta{
				Success: true,
				Message: "Success",
			},
		}
	}
	render.JSON(w, r, res)
}

type UserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	GetList(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Active(w http.ResponseWriter, r *http.Request)
	DeActive(w http.ResponseWriter, r *http.Request)
}

func NewUserController() UserController {
	userBusiness := Business.NewUserBusiness()
	return &userController{
		userBusiness: userBusiness,
	}
}
