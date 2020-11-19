package Controller

import (
	"Book/Business"
	"Book/Controller/payload"
	"Book/Controller/response"
	"Book/Database/Models"
	"encoding/json"
	"github.com/go-chi/render"
	"net/http"
)

type RoleController interface {
	CreateRole(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	GetList(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Active(w http.ResponseWriter, r *http.Request)
	DeActive(w http.ResponseWriter, r *http.Request)
}
type roleController struct {
	roleBusiness Business.RoleBusiness
}
// Create new book godoc
// @tags Role
// @Summary Create new Role
// @Description create new role
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product body payload.Role true "role information"
// @Success 200 {object} response.RoleResponse
// @Router /roles [post]
func (r2 roleController) CreateRole(w http.ResponseWriter, r *http.Request) {
	var pl payload.Role

	err := json.NewDecoder(r.Body).Decode(&pl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}
	res := &response.RoleResponse{}

	role := Models.Role{
		Name:   pl.Name,
		Status: true,
	}
	data, err := r2.roleBusiness.Create(role)

	if err != nil {
		res = &response.RoleResponse{
			Data: nil,
			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
		}
	} else {
		res = &response.RoleResponse{
			Meta: response.Meta{
				Success: true,
				Message: "Success",
			},
			Data: data,
		}
	}
	render.JSON(w, r, res)
}
// Create new book godoc
// @tags Role
// @Summary Create new Role
// @Description create new role
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product body payload.Role true "role information"
// @Success 200 {object} response.RoleResponse
// @Router /roles [put]
func (r2 roleController) Update(w http.ResponseWriter, r *http.Request) {
	var pl payload.Role
	err := json.NewDecoder(r.Body).Decode(&pl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}
	res := &response.RoleResponse{}
	role := Models.Role{
		Name:   pl.Name,
		Status: true,
	}
	data, err := r2.roleBusiness.Update(role)
	if err != nil {
		res = &response.RoleResponse{
			Data: nil,
			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
		}
	} else {
		res = &response.RoleResponse{
			Meta: response.Meta{
				Success: true,
				Message: "Success",
			},
			Data: data,
		}
	}
	render.JSON(w, r, res)
}
// Filter get list roles godoc
// @tags Role
// @Summary get list roles
// @Description  list roles
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.RoleResponse
// @Router /roles [get]
func (r2 roleController) GetList(w http.ResponseWriter, r *http.Request) {
	data, err := r2.roleBusiness.GetAll()
	res := &response.ListRoleResponse{}
	if err != nil {
		res = &response.ListRoleResponse{
			Data: nil,
			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
		}
	} else{
		res = &response.ListRoleResponse{
			Data: data,
			Meta: response.Meta{
				Success: true,
				Message: "success",
			},
		}
	}
	render.JSON(w, r, res)
}
// Filter get role godoc
// @tags Role
// @Summary get role
// @Description  role
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query string true "id"
// @Success 200 {object} response.ListBooksResponse
// @Router /roles/{id} [get]
func (r2 roleController) GetById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == ""{
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}
	data, err := r2.roleBusiness.Get(id)
	res := &response.RoleResponse{}
	if err != nil {
		res = &response.RoleResponse{
			Data: nil,
			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
		}
	} else{
		res = &response.RoleResponse{
			Data: data,
			Meta: response.Meta{
				Success: true,
				Message: "success",
			},
		}
	}
	render.JSON(w, r, res)
}
// Filter active role godoc
// @tags Role
// @Summary active role
// @Description active role
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query string true "id"
// @Success 200 {object} response.ListBooksResponse
// @Router /roles/active/{id} [put]
func (r2 roleController) Active(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == ""{
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}
	role, err := r2.roleBusiness.Get(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	role.Status = true
	data, err := r2.roleBusiness.Update(*role)
	res := &response.RoleResponse{}
	if err != nil {
		res = &response.RoleResponse{
			Data: nil,
			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
		}
	} else{
		res = &response.RoleResponse{
			Data: data,
			Meta: response.Meta{
				Success: true,
				Message: "success",
			},
		}
	}
	render.JSON(w, r, res)
}
// Filter active role godoc
// @tags Role
// @Summary active role
// @Description active role
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query string true "id"
// @Success 200 {object} response.ListBooksResponse
// @Router /roles/deactivate/{id} [put]
func (r2 roleController) DeActive(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == ""{
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}
	role, err := r2.roleBusiness.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	role.Status = false
	data, err := r2.roleBusiness.Update(*role)

	res := &response.RoleResponse{}
	if err != nil {
		res = &response.RoleResponse{
			Data: nil,
			Meta: response.Meta{
				Success: false,
				Message: err.Error(),
			},
		}
	} else{
		res = &response.RoleResponse{
			Data: data,
			Meta: response.Meta{
				Success: true,
				Message: "success",
			},
		}
	}
	render.JSON(w, r, res)
}

func NewRoleController() RoleController {
	roleBusiness := Business.NewRoleBusiness()
	return &roleController{
		roleBusiness: roleBusiness,
	}
}