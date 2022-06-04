package masterDataServices

import (
	masterDatarepository "growdo/src/controller/masterData/masterDataRepository"
	"growdo/src/model"
)

type Service interface {
	Create(payload *model.MasterData) (*model.MasterData, string, error)
	All(st string, f *model.FilterCari) (*[]model.MasterData, string, error)
	FindAll(payload *model.FilterCari) (*[]model.MasterData, string, error)
}

type service struct {
	repository masterDatarepository.Repository
}

func NewService(repository masterDatarepository.Repository) *service {
	return &service{repository: repository}
}

func (s *service) Create(payload *model.MasterData) (*model.MasterData, string, error) {

	// cek emial
	var mes = "terjadi kesalahan ketika create"
	data, err := s.repository.Create(payload)
	if err != nil {
		return nil, mes, err
	}

	dataQuery, err := s.repository.FindDetail("id", &model.FilterCari{Id: data.Id})
	if err != nil {
		return nil, mes, err
	}

	return dataQuery, "berhasil membuat data", nil
}

func (s *service) All(st string, f *model.FilterCari) (*[]model.MasterData, string, error) {
	data, err := s.repository.All(st, f)
	if err != nil {
		return nil, "ada kesalahan ketika upload file", err
	}
	return data, "berhasil mengambil data", nil
}

func (s *service) FindAll(payload *model.FilterCari) (*[]model.MasterData, string, error) {

	// cek emial
	var mes = "terjadi kesalahan ketika Mengambil File"
	dataQuery, err := s.repository.FindAll(payload)
	if err != nil {
		return nil, mes, err
	}

	return dataQuery, "berhasil mengambil data", nil
}
