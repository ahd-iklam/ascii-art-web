package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func excuteInexPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	templates.ExecuteTemplate(w, "index.html", nil)
}

func excuteAsciiArtResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	art, err := generateAsciiArt(text, banner)
	if err != 0 {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Art string
	}{
		Art: art,
	}

	templates.ExecuteTemplate(w, "result.html", data)
}

func generateAsciiArt(text, bannerType string) (string, int) {
	bannerFilePath := filepath.Join("banners", bannerType+".txt")
	er := 0
	file, err := os.ReadFile(bannerFilePath)
	if err != nil {
		er = 1
		return "", er
	}
	// count how much of new lines
	count := strings.Count(text, "\r\n")
	// split the argumant by new lines
	testLines := strings.Split(text, "\r\n")
	var asciiChars []string
	// split the banner by double new line
	if bannerType == "thinkertoy" {
		asciiChars = strings.Split(string(file), "\r\n\r\n")
	} else {
		asciiChars = strings.Split(string(file), "\n\n")
	}
	characters := make([][]string, len(asciiChars))
	// stock every line of instance of banner in a case
	for i, char := range asciiChars {
		if bannerType == "thinkertoy" {
			characters[i] = strings.Split(char, "\r\n")
		} else {
			characters[i] = strings.Split(char, "\n")
		}
	}
	for _, v := range text {
		if (v < 32 || v > 126) && (v != 10 && v != 13) {
			er = 1
			return "", er
		}
	}
	// print a result with all requirements
	result := ""
	counter := 1
	for _, line := range testLines {
		if line == "" {
			if counter <= count {
				result += "\r\n"
			}
			counter++
			continue
		}
		for l := 0; l < 8; l++ {
			for _, char := range line {
				if char == ' ' { // a space character
					result += characters[char-32][l+1]
				} else {
					index := char - 32
					result += characters[index][l]
				}
			}
			result += "\r\n"
		}
	}
	return result, er
}
