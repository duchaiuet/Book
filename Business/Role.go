package Business

import (
	"Book/Database"
	"Book/Database/DAL"
	"Book/Database/Models"
)

type roleBusiness struct {
	RoleRepository Models.RoleRepository
}

func (r roleBusiness) Get(id string) (*Models.Role, error) {
	result, err := r.RoleRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r roleBusiness) Create(role Models.Role) (*Models.Role, error) {
	result, err := r.RoleRepository.Create(role)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r roleBusiness) GetAll() ([]*Models.Role, error) {
	result, err := r.RoleRepository.Gets()
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r roleBusiness) Update(role Models.Role) (*Models.Role, error) {
	result, err := r.RoleRepository.Update(role)
	if err != nil {
		return nil, err
	}
	return result, err
}

type RoleBusiness interface {
	Create(role Models.Role) (*Models.Role, error)
	GetAll() ([]*Models.Role, error)
	Get(id string)( *Models.Role, error)
	Update(role Models.Role) (*Models.Role, error)
}

func NewRoleBusiness() RoleBusiness {
	roleRepository := DAL.NewRoleRepository(Database.Db)
	return roleBusiness{RoleRepository: roleRepository}
}
