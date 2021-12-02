package repository

import "github.com/vivy-c/first-project-go/internal/models"

type Repository interface {
	SaveMahasiswaAlamat(dataMahasiswa *models.MahasiswaModels, dataAlamat []*models.MahasiswaAlamatModels) error
	UpdateMahasiswaNama(dataMahasiswa *models.MahasiswaModels) error
	// SaveAlamatId(dataAlamat *models.AlamatIdModels) error
	SaveAlamatId(dataAlamat *models.MahasiswaAlamatModels) error
	ShowAllMahasiswaAlamat() (string, error)
}
