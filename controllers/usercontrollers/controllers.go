package usercontrollers

import (
	"html/template"
	"net/http"

	"github.com/cehuda1/go-crud1/entities"
	"github.com/cehuda1/go-crud1/models"
)

var UserModel = models.NewUserModel()

func Index(response http.ResponseWriter, request *http.Request) {
	
	user, _ := UserModel.FindAll()

	data := map[string]interface{}{
		"user":user,
	}

	temp, err := template.ParseFiles("views/user/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)

}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/user/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var user entities.User
		user.NamaLengkap = request.Form.Get("nama_lengkap")
		user.Usia = request.Form.Get("usia")
		user.JenisKelamin = request.Form.Get("jenis_kelamin")
		user.Alamat = request.Form.Get("alamat")
		user.No_Hp = request.Form.Get("no_hp")
		user.Tanggal_Lahir = request.Form.Get("tanggal_lahir")
		
		UserModel.Create(user)
		data := map[string]interface{}{
			"pesan": "Data User Berhasil Disimpan",
		}

		temp, _ := template.ParseFiles("views/user/add.html")
	temp.Execute(response, data)
	}

	
}

func Edit(response http.ResponseWriter, request *http.Request) {
	
}

func Delete(response http.ResponseWriter, request *http.Request) {
	
}