package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define Booking Struct
type Booking struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}

// Store bookings in memory (Replace with a DB in production)
var bookings []Booking

// Handle new booking requests
func createBooking(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newBooking Booking
	err := json.NewDecoder(r.Body).Decode(&newBooking)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	bookings = append(bookings, newBooking)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Booking successful!"})
}

// Serve the frontend HTML
func serveFrontend(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {

	// Serve static files (images, CSS, JS)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveFrontend)
	http.HandleFunc("/book", createBooking)

	fmt.Println("Server started on http://localhost:8090")
	http.ListenAndServe(":8090", nil)
}
