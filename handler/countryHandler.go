package handler

import (
	config "assignment_1/config"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func CountryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCountry(w, r)
	default:
		http.Error(w, "Method not supported", http.StatusNotImplemented)
	}
}

// This method searches for a country by NAME on the restcountries API and returns an array of type []Country.

func getCountryByName(countryName string, w http.ResponseWriter) []config.Country {
	response, err := http.Get("https://restcountries.com/v3.1/name/" + countryName)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	defer response.Body.Close()
	var temp []config.Country
	err = json.NewDecoder(response.Body).Decode(&temp)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	return temp
}

// This method searches for a country by its ALPHA code in the restcountries API and returns an array of type []Country.
func getCountryByCode(countryCode string, w http.ResponseWriter) []config.Country {
	response, err := http.Get("https://restcountries.com/v3.1/alpha/" + countryCode)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	defer response.Body.Close()
	var temp []config.Country
	err = json.NewDecoder(response.Body).Decode(&temp)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	return temp
}

// This function processes the GET request received by CountryHandler and generates a JSON response containing the relevant universities.
func getCountry(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 6 || parts[3] != "neighbourunis" {
		http.Error(w, "Malformed URL", http.StatusBadRequest)
		return
	}
	countryName := strings.ReplaceAll(parts[4], " ", "%20")
	countries := getCountryByName(countryName, w)
	for i := range countries[0].Neighbours {
		temp := getCountryByCode(countries[0].Neighbours[i], w)
		countries = append(countries, temp[0])
	}
	universityName := strings.ReplaceAll(parts[5], " ", "%20")
	universityArray := getUniversityByNameAndCountry(universityName, countries[0].CountryName.Name, w)
	limit := 0
	if strings.Contains(r.URL.RawQuery, "limit") {
		sLimit := r.URL.RawQuery
		regex, _ := regexp.Compile(`[^0-9]`)
		sLimit = regex.ReplaceAllString(sLimit, "")
		limit, _ = strconv.Atoi(sLimit)
	}
	for i := range countries {
		allCountries := getUniversityByNameAndCountry(universityName, countries[i].CountryName.Name, w)
		if len(allCountries) == 0 {
			log.Println("no universities attached to this country with that name")
		} else {
			if limit > 0 {
				for j := 0; j < limit; j++ {
					universityArray = append(universityArray, allCountries[j])
				}
			} else {
				for j := 0; j < len(allCountries); j++ {
					universityArray = append(universityArray, allCountries[j])
				}
			}
		}
	}

	for i := 0; i < len(universityArray); i++ {
		var additionalcountriesrmation = getCountryByName(universityArray[i].CountryName, w)
		universityArray[i].Languages = additionalcountriesrmation[0].Languages
		universityArray[i].Map.OpenStreetMaps = additionalcountriesrmation[0].Map.OpenStreetMaps
	}

	json.Marshal(universityArray)
	w.Header().Add("content-type", "application/json")
	err := json.NewEncoder(w).Encode(&universityArray)
	if err != nil {
		log.Println("An ERROR occurred while encoding json", err)
	}

}
