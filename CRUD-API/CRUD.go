package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if r.Method == "POST" {
		var contact Person
		err := json.NewDecoder(r.Body).Decode(&contact)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		db, err := Connect()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		id := r.URL.Query().Get("id")
		_, err = db.Exec("UPDATE tb_crud SET name = ?, phone = ? WHERE id = ?", contact.Name, contact.Phone, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Success"))
		return
	}
	http.Error(w, "Wrong Method", http.StatusBadRequest)	
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if r.Method == "GET" {
		db, err := Connect()
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		id := r.URL.Query().Get("id")
		_, err = db.Exec("DELETE FROM tb_crud WHERE id = ?", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Success"))
		return
	}

	http.Error(w, "Bad Request", http.StatusBadRequest)
}


func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if r.Method == "POST" {
		var contact Person
		err := json.NewDecoder(r.Body).Decode(&contact)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		db, err := Connect()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		_, err = db.Exec("INSERT INTO tb_crud (name, phone) VALUES (?, ?)", contact.Name, contact.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Success"))
		return
	}
	http.Error(w, "Bad request", http.StatusBadRequest)
}


func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if r.Method == "GET" {
		db, err := Connect()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		data, err := db.Query("SELECT * FROM tb_crud")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var result []Person

		for data.Next() {
			var each = Person{}
			var err = data.Scan(&each.Id, &each.Name, &each.Phone)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			result = append(result, each)
		}

		hasil, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(hasil)
		return
	}
	http.Error(w, "Bad Request", http.StatusBadRequest)
}

func User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		
		db, err := Connect()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var result = Person{}
		err = db.QueryRow("SELECT * FROM tb_crud WHERE id =(?)", id).Scan(&result.Id, &result.Name, &result.Phone)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		hasil, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(hasil)
		return
	}
	http.Error(w, "Bad Request", http.StatusBadRequest)
}