package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type URL struct {
	ID             int       `json:"id"`
	OriginalURL    string    `json:"original_url"`
	ShortUrl       string    `json:"short_url"`
	CreatetionDate time.Time `json:"creation_date"`
}

var urlDB = make(map[int]URL)

func generateShortUrl(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)

	return hash[:8]

}

var autoIncrement int = 1

func createURL(originalURL string) string {
	shortUrl := generateShortUrl(originalURL)
	id := autoIncrement
	autoIncrement++
	urlDB[id] = URL{
		ID:             id,
		OriginalURL:    originalURL,
		ShortUrl:       shortUrl,
		CreatetionDate: time.Now(),
	}
	return shortUrl
}

func getURl(id int) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}
func RootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}
func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	shortURL_ := createURL(data.URL)
	response := struct {
		ShortUrl string `json:"short_url"`
	}{ShortUrl: shortURL_}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/short/"):]
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid URL ID", http.StatusBadRequest)
		return
	}
	url, err := getURl(intID)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {
	http.HandleFunc("/", RootPage)
	http.HandleFunc("/shorten", ShortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)
	fmt.Println(`Server is running on http://localhost:3000`)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error on starting server: ", err)
	}
}
