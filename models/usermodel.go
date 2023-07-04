package models

import (
	"database/sql"
	"fmt"

	"github.com/cehuda1/go-crud1/config"
	"github.com/cehuda1/go-crud1/entities"
)

type UserModel struct {
	conn *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &UserModel{
		conn: conn,
	}
}

func (p *UserModel) FindAll() ([]entities.User,error)  {

	rows, err := p.conn.Query("select * from user")
	if err != nil {
		return[]entities.User{},err
	}
	defer rows.Close()

	var dataUser []entities.User
	for rows.Next() {
		var user entities.User
		rows.Scan(&user.Id,
		&user.NamaLengkap, 
		&user.Usia, 
		&user.JenisKelamin, 
		&user.Alamat, 
		&user.No_Hp, 
		&user.Tanggal_Lahir)

		if user.JenisKelamin == "1" {
			user.JenisKelamin = "Laki - Laki"
		} else {
			user.JenisKelamin = "2"
			user.JenisKelamin = "Perempuan"
		}

		dataUser = append(dataUser, user)
	}
	return dataUser, nil

}

func (p *UserModel) Create(user entities.User) bool {

	result, err := p.conn.Exec("insert into user (nama_lengkap, usia, jenis_kelamin, alamat, no_hp, tanggal_lahir) value(?,?,?,?,?,?)",
	user.NamaLengkap, user.Usia ,user.JenisKelamin, user.Alamat, user.No_Hp, user.Tanggal_Lahir)

	if err!= nil{
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()
	return lastInsertId > 0
}
/////