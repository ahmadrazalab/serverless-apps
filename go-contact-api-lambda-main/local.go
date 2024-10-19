package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
)

var smtpHost = "" // Change region if necessary
var smtpPort = "" // AWS SES supports 587 for TLS
var smtpUser = "" // os.Getenv("SMTP_USER") // AWS SES SMTP user
var smtpPass = "" //os.Getenv("SMTP_PASS") // AWS SES SMTP password

// ContactFormHandler handles the form submission
func ContactFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	subject := r.FormValue("subject")
	query := r.FormValue("query")

	if name == "" || email == "" || subject == "" || query == "" {
		http.Error(w, "Missing form fields", http.StatusBadRequest)
		return
	}

	// Email details
	from := "uptime@paytring.com"        // Replace with a verified AWS SES email address
	to := []string{"ahmad@paytring.com"} // Replace with the recipient's email address

	// Compose the email body
	msg := []byte("From: " + from + "\r\n" +
		"To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		"Name: " + name + "\n" +
		"Email: " + email + "\n" +
		"Query: " + query + "\r\n")

	// Send the email via AWS SES SMTP
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Email sent successfully!")
}

func main() {
	// Setup route to handle form submissions
	http.HandleFunc("/submit-query", ContactFormHandler)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	// 	tmpl.Execute(w, nil)
	// })

	// Start the HTTP server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
