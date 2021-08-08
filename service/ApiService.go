package service

import (
	"focusapi/model"
)

type InstanceService struct {
	DAO model.InstanceDAO
}

// get api service
func (s *InstanceService) getSvc() *InstanceService {
	var m model.BaseModel
	return &InstanceService{
		DAO: &model.Instance{BaseModel: m},
	}
}

// find all apis
func (s *InstanceService) FindAllInstances() (interface{}, error) {
	var apis []model.Instance
	err := s.getSvc().DAO.FindAll(&apis)
	return apis, err
}

// find all Jobs with key
func (s *InstanceService) FindInstanceByEmail(email string) (*model.Instance, error) {
	api := model.Instance{}
	return &api, s.getSvc().DAO.FindByKeys(&api, map[string]interface{}{"email": email})
}

// find api by id
func (s *InstanceService) FindInstanceById(id uint64) (*model.Instance, error) {
	api := model.Instance{}
	return &api, s.getSvc().DAO.FindByKeys(&api, map[string]interface{}{"id": id})
}

// find all Jobs with key
func (s *InstanceService) FindAllInstancesWithKeys(keys map[string]interface{}) ([]model.Instance, error) {
	apis := []model.Instance{}
	return apis, s.getSvc().DAO.FindByKeys(&apis, keys)
}

// find all Jobs with key
func (s *InstanceService) FindAllInstanceByJobId(keys map[string]interface{}) (model.Instance, error) {
	api := model.Instance{}
	return api, s.getSvc().DAO.FindByKeys(&api, keys)
}

// create api
func (s *InstanceService) CreateInstance(api *model.Instance) error {
	return s.getSvc().DAO.Create(api)
}

// create api
func (s *InstanceService) GetInstanceByJobId(keys map[string]interface{}) (*model.Instance, error) {
	apis := model.Instance{}
	return &apis, s.getSvc().DAO.FindByKeys(&apis, keys)
}

// update api
func (s *InstanceService) UpdateInstance(id uint64, api *model.Instance) (int64, error) {
	rowsAffected, err := s.getSvc().DAO.Update(api, api.ID)
	return rowsAffected, err
}

// delete api
func (s *InstanceService) DeleteInstance(id uint64) (int64, error) {
	return s.getSvc().DAO.Delete(&model.Instance{}, id)
}

// find all apis
func (s *InstanceService) FindAllInstanceByPages(currentPage, pageSize int, totalRows *int64) ([]model.Instance, error) {
	apis := []model.Instance{}
	err := s.getSvc().DAO.Count(&apis, totalRows)
	if err != nil {
		return apis, err
	}
	return apis, s.getSvc().DAO.FindByPages(&apis, currentPage, pageSize)
}

//find apis by keys
func (s *InstanceService) FindAllInstanceByPagesWithKeys(keys,
	keyOpts map[string]interface{},
	currentPage, pageSize int,
	totalRows *int64,
	opts ...model.DAOOption) (interface{}, error) {
	//find api
	var apis []model.Instance
	err := s.getSvc().DAO.CountWithKeys(&apis, totalRows, keys, keyOpts, opts...)
	if err != nil {
		return apis, err
	}

	return apis, s.getSvc().DAO.FindByPagesWithKeys(&apis, keys, currentPage, pageSize, opts...)
}

//search apis by keys
func (s *InstanceService) SearchInstanceByPagesWithKeys(keys,
	keyOpts map[string]interface{},
	currentPage, pageSize int,
	totalRows *int64,
	opts ...model.DAOOption) (interface{}, error) {
	//search api
	var apis []model.Instance
	err := s.getSvc().DAO.CountWithKeys(&apis, totalRows, keys, keyOpts, opts...)
	if err != nil {
		return apis, err
	}

	return apis, s.getSvc().DAO.SearchByPagesWithKeys(&apis, keys, keyOpts, currentPage, pageSize, opts...)
}
