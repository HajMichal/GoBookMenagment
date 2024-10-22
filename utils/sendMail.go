package utils

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail(email string) {
	m := gomail.NewMessage()
	
	sender := os.Getenv("SENDER_MAIL")
	m.SetHeader("From", sender)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Welcome in the library")
	m.SetBody("text/html", emailBody)
	
	server := os.Getenv("SMTP_SERVER")
	serverPwd := os.Getenv("SMTP_PWD")

	d := gomail.NewDialer(server, 587, sender, serverPwd)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}

var emailBody = `
	<table width="100%" cellpadding="0" cellspacing="0" style="background-color: #ffffff; max-width: 600px; margin: auto; border: 1px solid #ddd;">
        <tr>
            <td style="padding: 20px; text-align: center; background-color: #4caf50; color: #ffffff;">
                <h1>Welcome to our library!</h1>
            </td>
        </tr>
        <tr>
            <td style="padding: 20px;">
                <h2>Hello!</h2>
                <p>Welcome to library, your new portal to a world of knowledge! We are excited to have you join our community of book lovers, researchers, and curious minds.</p>
                <p>As a member of library, you now have access to:</p>
                <ul>
                    <li>Thousands of eBooks, audiobooks, and journals</li>
                    <li>Exclusive recommendations based on your reading history</li>
                    <li>Personalized reading lists and book collections</li>
                    <li>24/7 access from any device</li>
                </ul>
                <p><strong>Getting started is easy:</strong></p>
                <ol>
                    <li><a href="[Login URL]">Log in to your account</a> using your registered email.</li>
                    <li>Explore categories or use our search tool to find your favorite books.</li>
                    <li>Start reading, listening, or adding books to your list!</li>
                </ol>
                <p>If you ever need assistance, our support team is just a click away at [Contact Email]. You can also visit our <a href="[Help Center URL]">Help Center</a> for FAQs and tutorials.</p>
                <p>Thank you for joining Library. We’re looking forward to being part of your reading journey!</p>
                <p>Happy reading!</p>
                <p><strong>The Library Team</strong></p>
            </td>
        </tr>
        <tr>
            <td style="padding: 20px; text-align: center; background-color: #f9f9f9; color: #555;">
                <p>Stay connected with us:</p>
                <p>
                    <a href="[Facebook URL]" style="margin-right: 10px;">Facebook</a> | 
                    <a href="[Twitter URL]" style="margin-right: 10px;">Twitter</a> | 
                    <a href="[Instagram URL]">Instagram</a>
                </p>
                <p>&copy; [2024] [Library]. All rights reserved.</p>
            </td>
        </tr>
    </table>
	`