package handler

import (
	config "assignment_1/config"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func UniversityHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUniversity(w, r)
	default:
		http.Error(w, "Method is not supported. Only GET method is supported", http.StatusNotImplemented)
	}

}

// This method responds to the GET request sent by the UniversitysHandler and generates a JSON response containing the relevant universities.
func getUniversity(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 5 || parts[3] != config.UNI_INFO {
		http.Error(w, "Malformed URL", http.StatusBadRequest)
		return
	}
	uniName := strings.ReplaceAll(parts[4], " ", "%20")
	universities := getUniversityByName(uniName, w)

	//Enhancing the universities array with supplementary data that was not provided by the Hipolabs University API.
	for i := range universities {
		temp := getCountryByName(universities[i].CountryName, w)
		universities[i].Languages = temp[0].Languages
		universities[i].Map.OpenStreetMaps = temp[0].Map.OpenStreetMaps
	}

	json.Marshal(universities)
	w.Header().Add("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(&universities); err != nil {
		log.Println("An ERROR occurred while encoding json", err)
	}
}

// Function for searching universities in the hipolabs API based on the name and county parameters.
// Returns an array of type University
func getUniversityByNameAndCountry(university string, country string, w http.ResponseWriter) []config.University {
	response, err := http.Get("http://universities.hipolabs.com/search?name=" + university + "&country=" + country)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	defer response.Body.Close()
	var temp []config.University
	if err := json.NewDecoder(response.Body).Decode(&temp); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return temp
}

// Function for searching universities in the hipolabs API based on the name parameter.
// Returns an array of type University
func getUniversityByName(university string, w http.ResponseWriter) []config.University {
	response, err := http.Get("http://universities.hipolabs.com/search?name=" + university)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	defer response.Body.Close()
	var temp []config.University
	if err := json.NewDecoder(response.Body).Decode(&temp); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return temp
}
