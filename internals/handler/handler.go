package handler

import "net/http"

func api(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.FormValue("name")
		phone := r.FormValue("phone")

		newPerson := Person{
			ID:          idCounter,
			Name:        name,
			PhoneNumber: phone,
		}
		idCounter++
		people = append(people, newPerson)

		http.Redirect(w, r, "/list", http.StatusSeeOther)
		return
	}
}
