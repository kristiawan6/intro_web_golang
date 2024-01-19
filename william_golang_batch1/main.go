package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id          int
	Name        string
	age         int
	Email       string
	Phonenumber int
}

type Chef struct {
	Id          int
	Username    string
	Name        string
	Phonenumber int
	Role        string
}

type Menu struct {
	Id    int
	Name  string
	Price int
	Stock int
}
type Order struct {
	Id       int
	UserId   int
	MenuId   int
	Quantity int
}

type Payment struct {
	Id      int
	UserId  int
	OrderId int
	Amount  int
}

// membuat fungsi bernama UserHandler dengan parameter w (http.ResponseWriter) dan r (*http.Request)
func UserHandler(w http.ResponseWriter, r *http.Request) {
	// validasi apakah metode HTTP adalah GET
	if r.Method == "GET" {
		// membuat array untuk memasukan contoh user
		listUser := []User{
			{1, "Willi", 19, "Willi@mail.com", 8129081769},
			{2, "Brodi", 22, "brodi@mail.com", 81290833329},
			{3, "Mail", 32, "mail@mail.com", 8143533329},
		}
		// konvert slice listUser ke format JSON
		res, err := json.Marshal(listUser)
		// validasi kesalahan selama proses marshaling JSON
		if err != nil {
			// Jika terjadi kesalahan, memberikan tanggapan dengan kesalahan server internal dan pesan kesalahan
			http.Error(w, "Gagal Konversi Ke JSON", http.StatusInternalServerError)
			return
		}

		// membuat header Content-Type untuk menunjukkan bahwa respons berformat JSON
		w.Header().Set("Content-Type", "application/json")
		// membuat kode status HTTP menjadi 200 OK
		w.WriteHeader(http.StatusOK)
		// menulis respons JSON ke http.ResponseWriter
		w.Write(res)
	} else {
		// jika metode HTTP bukan GET,maka akan menampilkan string metode tidak diizinkan
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
	}
}


func ChefHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		listChef := Chef{
			Id:          1,
			Username:    "Chef01",
			Name:        "Juna",
			Phonenumber: 81924939045,
			Role:        "Head Chef",
		}
		res, err := json.Marshal(listChef)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke JSON", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func MenuHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		listMenu := []Menu{
			{1, "Brokoli Bakar", 19000, 10},
			{2, "Ikan Gosong", 78000, 12},
			{3, "Es Teh Manis Panas", 6000, 9},
		}
		res, err := json.Marshal(listMenu)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke JSON", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		listOrder := []Order{
			{1, 1, 3, 2},
			{2, 2, 1, 1},
		}
		res, err := json.Marshal(listOrder)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke JSON", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		listPayment := []Payment{
			{1, 1, 1, 12000},
			{2, 2, 2, 19000},
		}
		res, err := json.Marshal(listPayment)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke JSON", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/user", UserHandler)
	http.HandleFunc("/chef", ChefHandler)
	http.HandleFunc("/menu", MenuHandler)
	http.HandleFunc("/order", OrderHandler)
	http.HandleFunc("/payment", PaymentHandler)
	fmt.Println("Server Running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
