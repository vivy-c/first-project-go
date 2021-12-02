package assembler

import (
	"github.com/vivy-c/first-project-go/internal/models"
	"github.com/vivy-c/first-project-go/pkg/dto"
)

func ToSaveMahasiswa(d *dto.MahasiswaReqDTO) *models.MahasiswaModels {
	return &models.MahasiswaModels{
		Name: d.Nama,
		Nim:  d.																							Nim,
	}
}

func ToSaveMahasiswaAlamat(d *dto.AlamatReqDTO) *models.MahasiswaAlamatModels {
	return &models.MahasiswaAlamatModels{
		Jalan:   d.Jalan,
		NoRumah: d.NoRumah,
	}
}

func ToSaveMahasiswaAlamats(datas []dto.AlamatReqDTO) []*models.MahasiswaAlamatModels {
	var mds []*models.MahasiswaAlamatModels
	for _, m := range datas {
		mds = append(mds, ToSaveMahasiswaAlamat(&m))
	}
	return mds
}

func ToUpdateMahasiswaNama(d *dto.UpadeMahasiswaNamaReqDTO) *models.MahasiswaModels {
	return &models.MahasiswaModels{
		Name: d.Nama,
		ID:   d.ID,
	}
}

// func ToSaveAlamatId(d *dto.AlamatIdReqDTO) *models.AlamatIdModels {
// 	return &models.AlamatIdModels{
func ToSaveAlamatId(d *dto.AlamatIdReqDTO) *models.MahasiswaAlamatModels {
	return &models.MahasiswaAlamatModels{
		Jalan:        d.Jalan,
		NoRumah:      d.NoRumah,
		IDMahasiswas: d.IDMahasiswas,
	}
}