package implementation

import (
	"fmt"

	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/05-template-method/template"
)

type Email struct {
	template.Otp
}

func (e *Email) GenRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (e *Email) SaveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (e *Email) GetMessage(otp string) string {
	return fmt.Sprintf("EMAIL OTP for login is %s", otp)
}

func (e *Email) SendNotification(message string) error {
	fmt.Printf("EMAIL: sending sms: %s\n", message)
	return nil
}

func (e *Email) PublishMetric() {
	fmt.Printf("EMAIL: publishing metrics\n")
}
