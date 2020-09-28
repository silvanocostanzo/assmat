package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type TmplParams struct {
	Success bool
	Total   float64
}

func convertDetails(value string) float64 {
	el, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return el
}

func calculateTotal(details map[string]float64) float64 {
	return (details["hours"] * details["net"]) + (details["days"] * details["allocation"])
}

func main() {
	tmpl := template.Must(template.ParseFiles("template/form.html"))

	formDetails := make(map[string]float64)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		err := r.ParseForm()
		if err != nil {
			fmt.Errorf("an error happens")
			return
		}

		for k, v := range r.Form {
			formDetails[k] = convertDetails(v[0])
			fmt.Println(formDetails[k])
		}

		total := calculateTotal(formDetails)

		params := TmplParams{true, total}

		tmpl.Execute(w, params)
	})

	http.ListenAndServe(":8080", nil)
}
