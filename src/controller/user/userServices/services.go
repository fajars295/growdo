package userServices

import (
	"errors"
	"fmt"
	"growdo/src/controller/user/userRepository"
	"growdo/src/helpers/componen"
	"growdo/src/model"
	"strconv"
)

type Service interface {
	Create(payload *model.Users) (*model.Users, string, error)
	Login(payload *model.Login) (*model.Users, string, error)
	Profile(id int) (*model.Users, string, error)
}

type service struct {
	repository userRepository.Repository
}

func NewService(repository userRepository.Repository) *service {
	return &service{repository: repository}
}

func (s *service) Create(payload *model.Users) (*model.Users, string, error) {

	// cek emial
	var mes = "terjadi kesalahan ketika create"
	checkemail, err := s.repository.FindDetail("email", payload.Email)
	if err != nil {
		fmt.Println(err)
		return nil, mes, err
	}

	if checkemail.Id != 0 {
		return nil, "email sudah di gunakan", errors.New("email sudah di gunakan")
	}

	data, err := s.repository.Create(payload)
	if err != nil {
		return nil, mes, err
	}

	dataEmail, err := s.repository.FindDetail("email", data.Email)
	if err != nil {
		return nil, mes, err
	}

	return dataEmail, "berhasil create", nil
}
func (s *service) Login(payload *model.Login) (*model.Users, string, error) {

	dataEmail, err := s.repository.FindDetail("email", payload.Email)
	if err != nil {
		return nil, "error ketika mengambil data user", err
	}
	fmt.Println(dataEmail)

	if dataEmail.Id == 0 {
		return nil, "email tidak di temukan", errors.New("email tidak di temukan")
	}

	comparePassword := componen.ComparePassword(dataEmail.Password, payload.Password)
	if comparePassword != nil {
		return nil, "password yang anda gunakan salah", errors.New("password yang anda gunakan salah")
	}

	return dataEmail, "berhasil login", nil
}

func (s *service) Profile(id int) (*model.Users, string, error) {

	dataEmail, err := s.repository.FindDetail("id", strconv.Itoa(id))
	if err != nil {
		return nil, "error ketika mengambil data user", err
	}

	if dataEmail.Id == 0 {
		return nil, "data tidak ditemukan", errors.New("data tidak ditemukan")
	}

	return dataEmail, "berhasil", nil
}
