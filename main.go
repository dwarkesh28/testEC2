package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// HomePageData holds data for the home page
type HomePageData struct {
	Title string
	Message string
}

func main() {
	// Serve static files (optional, if you have assets like CSS or JS)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Define routes and associate them with handlers
	http.HandleFunc("/", homeHandler)

	// Start the web server
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// homeHandler handles the "/" route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Define data to be passed into the template
	data := HomePageData{
		Title:   "Welcome to My Go App",
		Message: "Hello, this is a simple Go app using templates and routes!",
	}

	// Parse and execute the HTML template
	tmpl, err := template.New("home").Parse(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>{{.Title}}</title>
		</head>
		<body>
			<h1>{{.Title}}</h1>
			<p>{{.Message}}</p>
		</body>
		</html>
	`)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// Execute the template with the data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
