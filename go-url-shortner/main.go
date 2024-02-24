package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type URLShortener struct {
	// we're storing original urls and short urls in a map
	shortUrls map[string]string
}

func (us *URLShortener) handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/shorten", http.StatusPermanentRedirect)
	}

	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte(fmt.Sprintf(`<!DOCTYPE html>
<html>
<body>
<h2>URL SHORTENER</h2>
<form method="post" action="/shorten">
  <label> Enter your Link</label><br>
  <input type="text" id="url" name="url" placeholder="Enter your URL" required><br>
	<input type="submit" value="shorten">
</form> 
</body>
</html>
`)))
}

func (us *URLShortener) handleShorten(w http.ResponseWriter, r *http.Request) {
	/*
		1. show a form
		2. get url from form
		3. short that URL
		4. Present it
	*/

	// only cater to post method
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	if r.FormValue("url") == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	originalUrl := r.FormValue("url")

	shortUrl := generateShortKey()
	us.shortUrls[shortUrl] = originalUrl

	w.Write([]byte(fmt.Sprintf(`<!DOCTYPE html>
<html>
<body>
<h2>URL SHORTENER</h2>
<form method="post" action="/shorten">
  <label> Enter your Link</label><br>
  <input type="text" id="url" name="url" value="%s"><br>
</form> 
		<h3><a href="http://localhost:8080/short/%s">http://localhost:8080/short/%s</a></h3>
</body>
</html>
`, originalUrl, shortUrl, shortUrl)))

}

func (us *URLShortener) handleRedirect(w http.ResponseWriter, r *http.Request) {
	shortKey := r.URL.Path[len("/short/"):]

	http.Redirect(w, r, us.shortUrls[shortKey], http.StatusTemporaryRedirect)
}

func generateShortKey() string {
	var charset string = "abcdefghijklmnopqrstuvwxyz"
	const keylength = 6

	shortKey := make([]byte, keylength)

	rand.NewSource(rand.Int63())
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}

func main() {
	shortener := &URLShortener{
		shortUrls: make(map[string]string),
	}

	http.HandleFunc("/", shortener.handleForm)
	http.HandleFunc("/shorten", shortener.handleShorten)
	http.HandleFunc("/short/", shortener.handleRedirect)

	fmt.Println("URL Shortener is running on :8080")
	http.ListenAndServe(":8080", nil)
}
