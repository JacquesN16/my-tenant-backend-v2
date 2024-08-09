package pdf

import (
	"bytes"
	"fmt"
	"encoding/json"
	"html/template"
	"net/smtp"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type PDFData struct {
	today string
	title string
	month int32
	year int32
	lastName string
	firstName string
	email string
	startDate int64
	rent float64
	charge float64
	totalRent float64
	totalRentInText string
	dateOfPayment string
	startPeriod int64
	endPeriod int64
}

func GeneratePDF(c *fiber.Ctx){
	var data PDFData
	id := c.Params("id")

	err := json.Unmarshal(c.Body(), &data);
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"msg": err,
			"data":nil,
		})
		return
	}

	tmpl, err := template.ParseFiles("receiptTemplate.html")
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"msg": err,
			"data":nil,
		})
		return
	}


	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, data); err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"msg": err,
			"data":nil,
		})
		return
	}


	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"msg": err,
			"data":nil,
		})
		return
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(&buffer))
	pdf, err := pdfg.GeneratePDF()
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"msg": err,
			"data":nil,
		})
		return
	}

	if err := sendEmail(pdf, "recipient@example", "Receipt");
	err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"msg": err,
			"data":nil,
		})
		return
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"msg": "PDF generated and sent successfully",
		"data":nil,
	})
}

func sendEmail(pdf []byte, to, subject string) error {
	auth := smtp.PlainAuth("", "username", "pwd", "smtp.gmail.com")

	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\nMIME-version: 1.0;\r\nContent-Type: multipart/mixed; boundary=\"BOUNDARY\"\r\n\r\n--BOUNDARY\r\nContent-Type: text/plain\r\n\r\nPlease find attached the receipt\r\n\r\n--BOUNDARY\r\nContent-Type: application/pdf\r\nContent-Disposition: attachment; filename=\"receipt.pdf\"\r\n\r\n", to, subject))
	msg = append(msg, pdf...)

	err:= smtp.SendMail("smtp.gmail.com:587", auth, "sender@mail.com", []string{to}, msg)
	if err != nil {
		return err
	}

	return nil
}
