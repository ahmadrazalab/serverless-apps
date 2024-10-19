Sure, I'll guide you through creating a simple frontend application for your static content hosted on S3, and a small backend hosted on your EC2 server to handle email sending.

### Frontend Application (HTML, CSS, JavaScript)

Create an `index.html` file with a simple contact form:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Contact Us</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 50px;
        }
        form {
            max-width: 600px;
            margin: auto;
        }
        label {
            display: block;
            margin: 10px 0 5px;
        }
        input, textarea {
            width: 100%;
            padding: 10px;
            margin: 5px 0 20px;
            border: 1px solid #ccc;
            border-radius: 4px;
        }
        button {
            padding: 10px 20px;
            background-color: #28a745;
            border: none;
            border-radius: 4px;
            color: white;
            cursor: pointer;
        }
        button:hover {
            background-color: #218838;
        }
    </style>
</head>
<body>
    <h1>Contact Us</h1>
    <form id="contactForm">
        <label for="name">Name:</label>
        <input type="text" id="name" name="name" required>

        <label for="email">Email:</label>
        <input type="email" id="email" name="email" required>

        <label for="phone">Phone Number (Optional):</label>
        <input type="tel" id="phone" name="phone">

        <label for="query">Query:</label>
        <textarea id="query" name="query" required></textarea>

        <button type="submit">Submit</button>
    </form>

    <script>
        document.getElementById('contactForm').addEventListener('submit', function(event) {
            event.preventDefault();

            const formData = {
                name: document.getElementById('name').value,
                email: document.getElementById('email').value,
                phone: document.getElementById('phone').value,
                query: document.getElementById('query').value,
            };

            fetch('http://your-ec2-server-ip:3000/send-email', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formData),
            })
            .then(response => response.json())
            .then(data => {
                if (data.success) {
                    alert('Your query has been submitted!');
                } else {
                    alert('There was an error submitting your query.');
                }
            })
            .catch(error => {
                console.error('Error:', error);
                alert('There was an error submitting your query.');
            });
        });
    </script>
</body>
</html>
```

### Backend Application (Node.js)

1. Set up your EC2 instance to run a Node.js application.
2. Create a new directory for your backend application and initialize it:

```bash
mkdir backend
cd backend
npm init -y
```

3. Install the necessary dependencies:

```bash
npm install express nodemailer body-parser
```

4. Create a `server.js` file with the following content:

```javascript
const express = require('express');
const nodemailer = require('nodemailer');
const bodyParser = require('body-parser');

const app = express();
const port = 3000;

app.use(bodyParser.json());

app.post('/send-email', async (req, res) => {
    const { name, email, phone, query } = req.body;

    let transporter = nodemailer.createTransport({
        service: 'gmail',
        auth: {
            user: 'your-email@gmail.com',
            pass: 'your-email-password'
        }
    });

    let mailOptions = {
        from: 'your-email@gmail.com',
        to: 'your-email@gmail.com',
        subject: 'New Contact Form Submission',
        text: `Name: ${name}\nEmail: ${email}\nPhone: ${phone}\nQuery: ${query}`
    };

    try {
        await transporter.sendMail(mailOptions);
        res.status(200).json({ success: true, message: 'Email sent successfully' });
    } catch (error) {
        res.status(500).json({ success: false, message: 'Failed to send email', error: error.toString() });
    }
});

app.listen(port, () => {
    console.log(`Server running on http://localhost:${port}`);
});
```

### Deploy and Test

1. **Upload your `index.html` to S3**:
   - Go to your S3 bucket.
   - Upload the `index.html` file.
   - Set the appropriate permissions to make it publicly accessible.

2. **Deploy your Node.js application on EC2**:
   - SSH into your EC2 instance.
   - Ensure Node.js is installed.
   - Copy your backend code to the EC2 instance.
   - Run the backend application using `node server.js`.

3. **Update the fetch URL** in the frontend JavaScript to point to your EC2 server's public IP and port.

4. **Test the Contact Form**:
   - Open your S3 URL in a browser.
   - Fill out the contact form and submit it.
   - Verify that the email is sent successfully and that the frontend works as expected even if the EC2 server is down (emails won't be sent, but the form will still work).

By following these steps, you will have a static frontend hosted on S3 that interacts with a backend on EC2 for email sending. If the EC2 server is down, the frontend will still be accessible, but email functionality will be temporarily unavailable.