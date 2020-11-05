package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Country struct {
	Name string `json:"name"`
	Code string `json:"alpha2code"`
}

func main() {
	var inputs = []string{
		"United States",
		"Estonia",
		"Morocco",
		"Germany",
		"India",
		"Ukraine",
		"Australia",
		"Indonesia",
		"Switzerland",
		"Pakistan",
		"Singapore",
		"Ghana",
		"Italy",
		"Armenia",
		"United Kingdom",
		"Belgium",
		"Colombia",
		"Israel",
		"Netherlands",
		"Brazil",
		"Spain",
		"Kazakhstan",
		"Afghanistan",
		"Georgia",
		"Russia",
		"Bangladesh",
		"Jordan",
		"Philippines",
		"France",
		"South Korea",
		"Kuwait",
		"Canada",
		"Hong Kong",
		"Vietnam",
		"Myanmar",
		"Belarus",
		"Costa Rica",
		"Romania",
		"Nigeria",
		"Peru",
		"United Arab Emirates",
		"Malaysia",
		"Latvia",
		"Taiwan",
		"Mexico",
		"Lebanon",
		"Japan",
		"Finland",
		"Croatia",
		"Austria",
		"New Zealand",
		"Greece",
		"Poland",
		"Sweden",
		"Egypt",
		"Slovenia",
		"Kenya",
		"Ireland",
		"Turkey",
		"Venezuela",
		"South Africa",
		"Saudi Arabia",
		"Sri Lanka",
		"Czech Republic",
		"Thailand",
		"Cambodia",
		"China",
		"Mongolia",
		"Uzbekistan",
		"Oman",
		"Ethiopia",
		"Denmark",
		"French Polynesia",
		"Iran",
		"Syria",
		"Iraq",
		"Lithuania",
		"Qatar",
		"Argentina",
		"Tanzania",
		"Wallis and Futuna",
		"Algeria",
		"Trinidad and Tobago",
		"Bahrain",
		"Moldova",
		"Portugal",
		"Tunisia",
		"Norway",
		"Nepal",
		"Azerbaijan",
	}

	var fixList = map[string]string{
		"United States":  "US",
		"United Kingdom": "GB",
		"Russia":         "RU",
		"South Korea":    "KR",
		"Vietnam":        "VN",
		"Venezuela":      "VE",
		"Iran":           "IR",
		"Syria":          "SY",
		"Tanzania":       "TZ",
		"Moldova":        "MD",
	}

	resp, err := http.Get("https://restcountries.eu/rest/v2/all")
	if err != nil {
		panic("Connection failed to https://restcountries.eu/rest/v2/all")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var countries []Country
	json.Unmarshal(body, &countries)
	listCountries := make(map[string]string)
	for _, c := range countries {
		listCountries[c.Name] = c.Code
	}

	for _, inp := range inputs {
		if listCountries[inp] != "" {
			outputQuery(inp, listCountries[inp])
		} else if fixList[inp] != "" {
			outputQuery(inp, fixList[inp])
		} else {
			fmt.Println(inp)
		}
	}

}

func outputQuery(country string, code string) {
	fmt.Printf("db.users.update({country : \"%v\"},{$set : {country : \"%v\"}}, false, true)\r\n", country, code)
}
