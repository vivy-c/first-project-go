package services

import "github.com/vivy-c/first-project-go/pkg/dto"

type Services interface {
	SaveMahasiswaAlamat(req *dto.MahasiswaReqDTO) error
	SaveAlamatId(req *dto.AlamatIdReqDTO) error
	UpdateMahasiswaNama(req *dto.UpadeMahasiswaNamaReqDTO) error
	ShowAllMahasiswaAlamat() (string, error)
}
