package main

import (
	"fmt"
	"net/http"
)

func setup_server() {
	http.HandleFunc("/", root_service)
	http.HandleFunc("/covid19_highlights", covid19_highlights)
	http.HandleFunc("/get_public_ip", get_ur_ip)
	http.ListenAndServe("0.0.0.0:8000", nil)

	fmt.Println("Web Server new")
}

func root_service(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service is on \n")
	fmt.Fprintf(w, "Services: \n -'/covid19_highlights'\n -'/get_public_ip'")
	fmt.Println("Service is on")
}

func covid19_highlights(w http.ResponseWriter, r *http.Request) {

	var country, country_population, last_updated, latest_confirmed, latest_deaths, latest_recovered string = get_latest_info_form_c19_api()

	fmt.Fprintf(w, "Country: %s \n", country)
	fmt.Fprintf(w, "Population: %s \n", country_population)
	fmt.Fprintf(w, "Last_update: %s \n", last_updated)
	fmt.Fprintf(w, "Confirmed: %s \n", latest_confirmed)
	fmt.Fprintf(w, "Deaths: %s \n", latest_deaths)
	fmt.Fprintf(w, "Recovered: %s \n", latest_recovered)
	fmt.Fprintf(w, "NewDeaths: %s \n", get_diff_btw_latest_and_older_ones(latest_deaths, get_old_deaths()))

	write_latest_deaths(latest_deaths)
}

func get_ur_ip(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ur public ip is: %s\n", get_public_ip())
}
