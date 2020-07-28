package implementation

import (
	"fmt"

	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/05-template-method/template"
)

type Sms struct {
	template.Otp
}

func (s *Sms) GenRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *Sms) SaveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *Sms) GetMessage(otp string) string {
	return fmt.Sprintf("SMS OTP for login is %s", otp)
}

func (s *Sms) SendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

func (s *Sms) PublishMetric() {
	fmt.Printf("SMS: publishing metrics\n")
}
