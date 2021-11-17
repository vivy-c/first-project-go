//disesuain nama foldernya
package database

import (
	"fmt"

	"github.com/jmoiron/sqlx" //buat konek ke databse
	_ "github.com/lib/pq"  //import semuanya
)

type Database struct {
	Conn *sqlx.DB  //variable dan tipe data, untuk ambil dan terima data ke db
}

//nama fungsi nya bebas, untuk inisialisasi untuk manggil data dari si db
func Initialize(host, username, password, database, port string) (Database, error) {
	db := Database{}

	//perintah konek db postgre
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	//line dibawah mirip kek try catch
	conn, err := sqlx.Open("postgres", dsn) //sqlx db buat nerima db, sqlx open nerima driver dan source data name
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}