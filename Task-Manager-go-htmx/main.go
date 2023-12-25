package main

// This is a simple web server that serves a static HTML page and handles
import (
	"html/template"
	"log"
	"net/http"
)

var tasks []string // This is a slice of tasks

// The main function starts the web server and listens for requests
func main() {
	http.HandleFunc("/", indexHandler)                                                         // This is the index page
	http.HandleFunc("/add-task", addTaskHandler)                                               // This is the handler for adding tasks
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // This is the handler for static files

	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server
}

// This handler handles the index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html") // Parse the HTML file

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // If there is an error, return an internal server error
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // If there is an error, return an internal server error
	}
}

// This handler handles adding tasks to the storage
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Here you can handle adding tasks to your storage
		// For demonstration, let's assume the task is sent via POST
		task := r.FormValue("task")
		// Handle the task as needed (e.g., store it in a slice, database, etc.)
		tasks = append(tasks, task) // Add the task to the slice
		log.Printf("Added task: %s", task)
	}
	http.Redirect(w, r, "/", http.StatusFound) // Redirect to the index page
}
