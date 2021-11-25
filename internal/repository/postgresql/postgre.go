package repository

import (
	"fmt"
	"log"

	"github.com/vivy-c/first-project-go/internal/models"
	"github.com/vivy-c/first-project-go/internal/repository"

	"github.com/jmoiron/sqlx"
	mhsErrors "github.com/vivy-c/first-project-go/pkg/errors"
)

const (
	SaveMahasiswa       = `INSERT INTO kampus.mahasiswas (nama, nim, created_at) VALUES ($1, $2, now()) RETURNING id`
	SaveMahasiswaAlamat = `INSERT INTO kampus.mahasiswa_alamats (jalan, no_rumah, created_at, id_mahasiswas) VALUES ($1,$2, now(), $3)`
	UpdateMahasiswaNama = `UPDATE kampus.mahasiswas SET nama = $1, updated_at = now() where id = $2`
)

var statement PreparedStatement

type PreparedStatement struct {
	updateMahasiswaNama *sqlx.Stmt
}

type PostgreSQLRepo struct {
	Conn *sqlx.DB
}

func NewRepo(Conn *sqlx.DB) repository.Repository {

	repo := &PostgreSQLRepo{Conn}
	InitPreparedStatement(repo)
	return repo
}

func (p *PostgreSQLRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Conn.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *PostgreSQLRepo) {
	statement = PreparedStatement{
		updateMahasiswaNama: m.Preparex(UpdateMahasiswaNama),
	}
}

func (p *PostgreSQLRepo) SaveMahasiswaAlamat(dataMahasiswa *models.MahasiswaModels, dataAlamat []*models.MahasiswaAlamatModels) error {

	tx, err := p.Conn.Beginx()
	if err != nil {
		log.Println("Failed Begin Tx SaveMahasiswa Alamat : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}
	var idMahasiswa int64
	err = tx.QueryRow(SaveMahasiswa, dataMahasiswa.Name, dataMahasiswa.Nim).Scan(&idMahasiswa)

	if err != nil {
		tx.Rollback()
		log.Println("Failed Query SaveMahasiswa: ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	for _, val := range dataAlamat {
		_, err = tx.Exec(SaveMahasiswaAlamat, val.Jalan, val.NoRumah, idMahasiswa)
		if err != nil {
			tx.Rollback()
			log.Println("Failed Query SaveMahasiswaAlamat : ", err.Error())
			return fmt.Errorf(mhsErrors.ErrorDB)
		}
	}

	return tx.Commit()
}

func (p *PostgreSQLRepo) UpdateMahasiswaNama(dataMahasiswa *models.MahasiswaModels) error {

	result, err := statement.updateMahasiswaNama.Exec(dataMahasiswa.Name, dataMahasiswa.ID)

	if err != nil {
		log.Println("Failed Query UpdateMahasiswaNama : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	rows, err := result.RowsAffected()

	if err != nil {
		log.Println("Failed RowAffectd UpdateMahasiswaNama : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	if rows < 1 {
		log.Println("UpdateMahasiswaNama: No Data Changed")
		return fmt.Errorf(mhsErrors.ErrorNoDataChange)
	}

	return nil
}
