package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/tidwall/gjson" // requierment for ez json parsing
)

func get_url() string {
	return "https://coronavirus-tracker-api.herokuapp.com/v2/locations" // return the url for c19 api
}

func get_old_deaths() string {
	out_old_deaths, err_old_deaths := exec.Command("bash", "-c", "cat ss.txt").Output() // read from file (ez way with bash) the last value

	if err_old_deaths != nil {
		fmt.Println(err_old_deaths, "not k", out_old_deaths)
		return "0"
	}

	return string(out_old_deaths)
}

func write_latest_deaths(latest string) {
	out, err := exec.Command("bash", "-c", "echo "+latest+" > ss.txt").Output() // write file (ez bash wasy) the latest value

	if err != nil {
		fmt.Println(err, "not k", out)
		return
	}

	fmt.Println(out, "ok")
}

func get_diff_btw_latest_and_older_ones(latest string, old string) string {
	out_dif, err_dif := exec.Command("bash", "-c", "echo $(("+latest+" - "+old+"))").Output() // make some ez math:  compute Latest_nr_of Deaths - Old_nr_of_Deaths

	if err_dif != nil {
		fmt.Println(err_dif, "not k", out_dif)
		return "0"
	}

	fmt.Println(out_dif, "ok")
	return string(out_dif)
}

func get_latest_info_form_c19_api() (string, string, string, string, string, string) {
	req, _ := http.NewRequest("GET", get_url(), nil) // some api stuff
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	country := gjson.Get(string(body), "locations.186.country")
	country_population := gjson.Get(string(body), "locations.186.country_population")
	last_updated := gjson.Get(string(body), "locations.186.last_updated")
	latest_confirmed := gjson.Get(string(body), "locations.186.latest.confirmed")
	latest_deaths := gjson.Get(string(body), "locations.186.latest.deaths")
	latest_recovered := gjson.Get(string(body), "locations.186.latest.recovered")

	fmt.Println("Service is on cov")
	fmt.Println(res)

	return country.String(), country_population.String(), last_updated.String(), latest_confirmed.String(), latest_deaths.String(), latest_recovered.String()
}

func get_public_ip() string {
	out, err := exec.Command("bash", "-c", "curl ifconfig.me").Output() // get some more api

	if err != nil {
		fmt.Println("ERR: %s\n", err)
		return "?.?.?.?"
	}

	return string(out)
}

func kill_8000() string {
	_, err := exec.Command("bash", "-c", "kill -9 $(lsof -t -i:8000)").Output() // hard kill

	if err != nil {
		return "Fail to kill 8000 or it's aleady killed"
	}

	return "8000 kiled"
}
