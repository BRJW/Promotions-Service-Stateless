package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	gorilla "github.com/gorilla/schema"
	//"github.com/json-iterator/go"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var decoder = gorilla.NewDecoder()

type User struct {
	Years       int
	Balance     float32
	Rating      float32
	Age         int
	AccountType string
}

func main() {
	fmt.Println("Service started")
	mux := http.NewServeMux()
	mux.HandleFunc("/Promotions", getPromotions)
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:8080", nil))
	//}()
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func getPromotions(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Input", http.StatusBadRequest)
		return
	}

	var inputUser User
	err = decoder.Decode(&inputUser, r.Form)

	if err != nil {
		http.Error(w, "Bad Input", http.StatusBadRequest)
		return
	}

	if !validateUser(inputUser) {
		http.Error(w, "Bad Input", http.StatusBadRequest)
		return
	}

	var results []string

	if ruleMillennial(inputUser) {
		results = append(results, "Millennial Madness")
	}

	if ruleOldies(inputUser) {
		results = append(results, "Golden Oldies")
	}

	if ruleLoyalty(inputUser) {
		results = append(results, "Loyalty Bonus")
	}

	if ruleValued(inputUser) {
		results = append(results, "Valued Customer")
	}

	if len(results) == 0 {
		results = append(results, "No Promotions!")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func validateUser(user User) bool {
	if user.Years > 0 && user.Age > 0 && user.Rating > 0 && isValidAccountType(user.AccountType) {
		return true
	}
	return false
}

func isValidAccountType(category string) bool {
	switch category {
	case
		"Blue",
		"Gold",
		"Platinum":
		return true
	}
	return false
}

func ruleMillennial(user User) bool {
	if 21 <= user.Age && user.Age <= 35 {
		if user.Rating >= 600 || user.Balance > 10000 {
			return true
		}
	}
	return false
}

func ruleOldies(user User) bool {
	if user.Age >= 65 {
		if user.Rating >= 500 || user.Balance > 5000 {
			if user.Years >= 10 || user.AccountType == "Gold" || user.AccountType == "Platinum" {
				return true
			}
		}
	}
	return false
}

func ruleLoyalty(user User) bool {
	if user.Years > 5 {
		return true
	}
	return false
}

func ruleValued(user User) bool {
	if ruleGoodStanding(user) && !(ruleMillennial(user) || ruleOldies(user) || ruleLoyalty(user)) {
		return true
	}
	return false
}

func ruleGoodStanding(user User) bool {
	if user.AccountType == "Platinum" || user.Rating > 500 || user.Balance >= 0 {
		return true
	}
	return false
}
