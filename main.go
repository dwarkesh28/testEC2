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
    <title>Coming Soon</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
            font-family: Arial, sans-serif;
        }
        body {
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            background: linear-gradient(45deg, #232526, #414345);
            background-size: 400% 400%;
            animation: gradientBG 20s infinite ease-in-out;
            text-align: center;
            color: white;
        }
        @keyframes gradientBG {
            0% { background-position: 0% 50%; }
            50% { background-position: 100% 50%; }
            100% { background-position: 0% 50%; }
        }
        h1 {
            font-size: 4em;
            animation: fadeIn 5s ease-in-out infinite alternate;
        }
        @keyframes fadeIn {
            0% { opacity: 0.6; transform: scale(1); }
            100% { opacity: 1; transform: scale(1.05); }
        }
    </style>
</head>
<body>
    <h1>Coming Soon</h1>
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
