Certainly! Here is the modified backend code that uses your SMTP server (`smtp.mcsam.in`) to send emails.

### Backend Application (Node.js)

1. Ensure you have Node.js installed on your EC2 instance.
2. Set up your EC2 instance to run a Node.js application.
3. Create a new directory for your backend application and initialize it:

```bash
mkdir backend
cd backend
npm init -y
npm install dotenv

```

4. Install the necessary dependencies:

```bash
npm install express nodemailer body-parser
```

5. Create a `server.js` file with the following content:

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
        host: 'test.test.com',
        port: 1025,
        auth: {
            user: 'test',
            pass: 'test'
        }
    });

    let mailOptions = {
        from: 'test@test.in', // Update to the appropriate sender email
        to: 'your-email@gmail.com', // Update to your receiving email
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

1. **Copy your backend code to the EC2 instance**:
   - SSH into your EC2 instance.
   - Navigate to the directory where you want to place your backend code.
   - Copy the `server.js` file to this directory.

2. **Run the backend application**:
   - Install the necessary Node.js dependencies:
     ```bash
     npm install express nodemailer body-parser
     ```
   - Start the backend application:
     ```bash
     node server.js
     ```

3. **Ensure the EC2 instance allows inbound traffic on port 3000**:
   - Go to your EC2 instance settings.
   - Update the security group to allow inbound traffic on port 3000.

4. **Update the frontend fetch URL** in the `index.html` file to point to your EC2 server's public IP and port.

5. **Upload your `index.html` to S3**:
   - Go to your S3 bucket.
   - Upload the `index.html` file.
   - Set the appropriate permissions to make it publicly accessible.

6. **Test the Contact Form**:
   - Open your S3 URL in a browser.
   - Fill out the contact form and submit it.
   - Verify that the email is sent successfully using your SMTP server.

By following these steps, you will have a static frontend hosted on S3 that interacts with a backend on EC2 for email sending using your specified SMTP server. If the EC2 server is down, the frontend will still be accessible, but email functionality will be temporarily unavailable.