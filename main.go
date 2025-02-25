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
	tmpl, err := template.New("home").Parse(`<!DOCTYPE html>
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
            font-family: 'Poppins', sans-serif;
        }
        body {
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            background: linear-gradient(-45deg, #0f2027, #203a43, #2c5364);
            background-size: 400% 400%;
            animation: gradientBG 15s infinite alternate ease-in-out;
            text-align: center;
            color: white;
            overflow: hidden;
            position: relative;
        }
        @keyframes gradientBG {
            0% { background-position: 0% 50%; }
            50% { background-position: 50% 50%; }
            100% { background-position: 100% 50%; }
        }
        h1 {
            font-size: 4.5em;
            font-weight: 700;
            text-transform: uppercase;
            letter-spacing: 8px;
            position: relative;
            z-index: 2;
            animation: fadeIn 3s ease-in-out;
        }
        @keyframes fadeIn {
            0% { opacity: 0; transform: scale(0.8); }
            100% { opacity: 1; transform: scale(1); }
        }
        h1::after {
            content: '';
            position: absolute;
            left: 0;
            bottom: -10px;
            width: 100%;
            height: 4px;
            background: rgba(255, 255, 255, 0.5);
            animation: underline 1.5s infinite alternate;
        }
        @keyframes underline {
            0% { width: 0; }
            100% { width: 100%; }
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
