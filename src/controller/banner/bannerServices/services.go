package bannerservices

import (
	bannerrepository "growdo/src/controller/banner/bannerRepository"
	"growdo/src/helpers/componen"
	"growdo/src/model"
	"os"
)

type Service interface {
	Create(payload *model.Banner) (*model.Banner, string, error)
	All(st string, f *model.FilterCari) (*[]model.Banner, string, error)
}

type service struct {
	repository bannerrepository.Repository
}

func NewService(repository bannerrepository.Repository) *service {
	return &service{repository: repository}
}

func (s *service) Create(payload *model.Banner) (*model.Banner, string, error) {

	// cek emial
	var mes = "terjadi kesalahan ketika create"

	upload, err := componen.UploadData(payload.Images, "public/banner", ".png,.jpeg,.jpg")
	if err != nil {
		return nil, "ada kesalahan ketika upload file", err
	}

	payload.Images = upload

	data, err := s.repository.Create(payload)
	if err != nil {
		os.Remove(upload)
		return nil, mes, err
	}

	dataEmail, err := s.repository.FindDetail("id", &model.FilterCari{Id: data.Id})
	if err != nil {
		os.Remove(upload)
		return nil, mes, err
	}

	return dataEmail, "berhasil create", nil
}

func (s *service) All(st string, f *model.FilterCari) (*[]model.Banner, string, error) {
	data, err := s.repository.All(st, f)
	if err != nil {
		return nil, "ada kesalahan ketika upload file", err
	}
	return data, "successful get data", nil
}
