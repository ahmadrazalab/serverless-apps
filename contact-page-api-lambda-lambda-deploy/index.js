const express = require('express');
const serverless = require('serverless-http');
const nodemailer = require('nodemailer');
const bodyParser = require('body-parser');
const dotenv = require('dotenv');

dotenv.config();

const app = express();
const port = 3000;



const cors = require('cors');
const corsOptions = {
    origin: process.env.ALLOWED_ORIGIN || 'https://kubecloud.in.net' // Replace with your website's URL
};
//  app.use(cors());
app.use(cors(corsOptions)); 

app.use(bodyParser.json());



app.post('/', async (req, res) => {
    const { name, email, phone, query } = req.body;

    let transporter = nodemailer.createTransport({
        host: process.env.SMTP_HOST,
        port: process.env.SMTP_PORT,
        auth: {
            user: process.env.SMTP_USER,
            pass: process.env.SMTP_PASS
        }
    });

    let mailOptions = {
        from: process.env.MAIL_FROM,
        to: process.env.MAIL_TO,
        subject: process.env.MAIL_SUBJECT,
        text: `Name: ${name}\nEmail: ${email}\nSubject: ${phone}\nQuery: ${query}`
    };

    try {
        await transporter.sendMail(mailOptions);
        res.status(200).json({ success: true, message: 'Email sent successfully' });
    } catch (error) {
        res.status(500).json({ success: false, message: 'Failed to send email', error: error.toString() });
    }
});

// app.listen(port, () => {
//     console.log(`Server running on http://localhost:${port}`);
// });
module.exports.handler = serverless(app);
