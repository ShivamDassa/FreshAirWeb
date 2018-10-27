package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	aMux := http.NewServeMux()

	aMux.Handle("/", http.FileServer(http.Dir("./static")))

	aMux.HandleFunc("/info/search", func(w http.ResponseWriter, r *http.Request) {
		// var res *http.Response
		if r.Method == "GET" {
			place := r.FormValue("place")
			fmt.Println("The value is", place)
			// // strReq := fmt.Sprintf("https://developers.zomato.com/api/v2.1/locations?query=%s", url.QueryEscape(city))
			// req, err := http.NewRequest(http.MethodGet, strReq, nil)
			// if err != nil {
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// }
			// // req.Header.Set("Accept", " application/json")
			// // req.Header.Set("user-key", "966b80a56bf2c010b2aef1a1f9c1b772")

			// clients := http.Client{}
			// if res, err = clients.Do(req); err != nil {
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// }

			// defer res.Body.Close()

			// if err != nil {
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// }
			// var location Location

			// data, _ := ioutil.ReadAll(res.Body)

			// if err := json.Unmarshal(data, &location); err != nil {
			// 	fmt.Println(err)
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// }

			// fmt.Println(location.LocationData)

			// var locData = location.LocationData

			// value = location.LocationData[0].EntityId

			// encoder := json.NewEncoder(w)
			// if err := encoder.Encode(locData); err != nil {
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// }

		}
	})
	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(aMux)

	// Start executing the application on port 8090
	err := http.ListenAndServe(":8080", n)
	if err != nil {
		fmt.Println("The error: ", err)
	}
}
