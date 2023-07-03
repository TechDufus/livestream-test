// Package main does the thing
package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const AppVersion = "0.0.3-test"

type catFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func writeHelloKitty(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "| | | |  ___  | || |  ___   | | / /(_)  _     _")
	fmt.Fprintln(w, "| |_| | / _ \\ | || | / _ \\  | |/ /  _ _| |_ _| |_  _  _")
	fmt.Fprintln(w, "|  _  |/ /_\\ \\| || |/ / \\ \\ |   /  | |_   _|_   _|| |/ /")
	fmt.Fprintln(w, "| | | |\\ ,___/| || |\\ \\_/ / | |\\ \\ | | | |_  | |_ | / /")
	fmt.Fprintln(w, "|_| |_| \\___/ |_||_| \\___/  |_| \\_\\|_| \\___| \\___||  /")
	fmt.Fprintln(w, "                       _           _              / /")
	fmt.Fprintln(w, "                      / \\_______ /|_\\             \\/")
	fmt.Fprintln(w, "                     /          /_/ \\__")
	fmt.Fprintln(w, "                    /             \\_/ /")
	fmt.Fprintln(w, "                  _|_              |/|_")
	fmt.Fprintln(w, "                  _|_  O    _    O  _|_")
	fmt.Fprintln(w, "                  _|_      (_)      _|_")
	fmt.Fprintln(w, "                   \\                 /")
	fmt.Fprintln(w, "                   _\\_______________/_")
	fmt.Fprintln(w, "                  /  \\/  (___)  \\/  \\")
	fmt.Fprintln(w, "                  \\__(  o     o  )__/")
	log.Printf("Endpoint Hit: %s", r.URL.Path)
}

// function to get catfact from api and unmarshall it from json to struct
func getCatFact() (catFact, error) {
	// http.Get("https://gatfact.ninja/fact")
	// This allows for untrusted x509 cert on api end
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Get("https://catfact.ninja/fact")
	if err != nil {
		return catFact{}, err
	}
	defer response.Body.Close()

	// get catfact from response body
	var catfact catFact
	err = json.NewDecoder(response.Body).Decode(&catfact)
	if err != nil {
		return catFact{}, err
	}
	// unmarshall catfact from json

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &catfact)
	if err != nil {
		return catFact{}, err
	}
	return catfact, nil
}

func getVersion() string {
	return AppVersion
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "Version: %s", getVersion())
	fmt.Printf("Version: %s", getVersion())
	writeHelloKitty(w, r)
}

func handleCatFact(w http.ResponseWriter, r *http.Request) {
	catfact, err := getCatFact()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Catfact: %s", catfact.Fact)
	fmt.Printf("Catfact: %s", catfact.Fact)
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "Length: %d", catfact.Length)
	fmt.Printf(" Length: %d", catfact.Length)
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "Version: %s", getVersion())
	fmt.Printf("Version: %s", getVersion())
	writeHelloKitty(w, r)
}

func handleRequests() {
	http.HandleFunc("/", handleCatFact)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
