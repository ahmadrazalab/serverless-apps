# .env.lambda
SMTP_HOST	email-smtp.us-east-1.amazonaws.com
SMTP_PORT	587
SMTP_USER	Your SES SMTP username
SMTP_PASS	Your SES SMTP password
SMTP_FROM	your-verified-email@example.com
SMTP_TO	recipient@example.com
GOOS=linux GOARCH=amd64 go build -o main main.go

zip function.zip main


curl -X POST https://localhost:8080/submit-query \
-H "Content-Type: application/json" \
-d '{"name":"John Doe","email":"john.doe@example.com","subject":"Test Subject From Local","query":"This is a test query."}'

# go-contact-api-lambda
