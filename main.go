package main

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var TWILIO_ACCOUNT_SID string = os.Getenv("TWILIO_ACCOUNT_SID")
var TWILIO_AUTH_TOKEN string = os.Getenv("TWILIO_AUTH_TOKEN")
var VERIFY_SERVICE_SID string = os.Getenv("VERIFY_SERVICE_SID")
var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: TWILIO_ACCOUNT_SID,
	Password: TWILIO_AUTH_TOKEN,
})

func kirimOTP(to string) {
	params := &openapi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(VERIFY_SERVICE_SID, params)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Verifikasi terkirim: '%s'\n", *resp.Sid)
	}
}

func verifikasiOtp(to string) {
	var code string
	fmt.Println("Silakan periksa ponsel Anda dan masukkan kode:")
	fmt.Scanln(&code)

	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(to)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(VERIFY_SERVICE_SID, params)

	if err != nil {
		fmt.Println(err.Error())
	} else if *resp.Status == "approved" {
		fmt.Println("Selamat Anda terverifikasi!")
	} else {
		fmt.Println("Kode OTP salah!")
	}
}

func main() {
	to := "<masukan nomor telepon Anda di sini>"

	kirimOTP(to)
	verifikasiOtp(to)
}
