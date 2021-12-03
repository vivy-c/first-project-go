package services

import (
	"github.com/vivy-c/first-project-go/internal/repository"
	"github.com/vivy-c/first-project-go/pkg/dto"
	"github.com/vivy-c/first-project-go/pkg/dto/assembler"
)

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Services {
	return &service{repo}
}

func (s *service) SaveMahasiswaAlamat(req *dto.MahasiswaReqDTO) error {

	dtAlamat := assembler.ToSaveMahasiswaAlamats(req.Alamats)
	dtMahasiswa := assembler.ToSaveMahasiswa(req)

	err := s.repo.SaveMahasiswaAlamat(dtMahasiswa, dtAlamat)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) ShowAllMahasiswaAlamat() ([]*dto.MahasiswaAlamatResDTO, error) {

	getMahasiswaMap := make(map[int64]*dto.MahasiswaAlamatResDTO)
	DataMahasiswaAlamat, err := s.repo.ShowAllMahasiswaAlamat()
	if err != nil {
		return nil, err
	}
	for _, val := range DataMahasiswaAlamat {
		if _, ok := getMahasiswaMap[val.ID]; !ok {
			getMahasiswaMap[val.ID] = &dto.MahasiswaAlamatResDTO{
				ID:   val.ID,
				Nama: val.Name,
				Nim:  val.Nim,
			}
			getMahasiswaMap[val.ID].Alamats = append(getMahasiswaMap[val.ID].Alamats, &dto.AlamatResDTO{
				Jalan:   val.Jalan,
				NoRumah: val.NoRumah,
			})
		} else {
			getMahasiswaMap[val.ID].Alamats = append(getMahasiswaMap[val.ID].Alamats, &dto.AlamatResDTO{
				Jalan:   val.Jalan,
				NoRumah: val.NoRumah,
			})
		}

	}

	var Data []*dto.MahasiswaAlamatResDTO
	for _, datas := range getMahasiswaMap {
		Data = append(Data, datas)
	}

	return Data, err
}


func (s *service) UpdateMahasiswaNama(req *dto.UpadeMahasiswaNamaReqDTO) error {

	dtMhsiswa := assembler.ToUpdateMahasiswaNama(req)

	err := s.repo.UpdateMahasiswaNama(dtMhsiswa)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) SaveAlamatId(req *dto.AlamatIdReqDTO) error {
	dtAlamat := assembler.ToSaveAlamatId(req)

	err := s.repo.SaveAlamatId(dtAlamat)
	if err != nil {
		return err
	}

	return nil
}