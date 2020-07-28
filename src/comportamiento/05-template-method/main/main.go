package main

import (
	"fmt"

	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/05-template-method/template"
	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/05-template-method/template/implementation"
)

func main() {
	smsOTP := &implementation.Sms{}
	o := template.Otp{
		IOtp: smsOTP,
	}

	o.GenAdnSendOTP(4)
	fmt.Println()
	emailOTP := &implementation.Email{}
	o = template.Otp{
		IOtp: emailOTP,
	}
	o.GenAdnSendOTP(4)
}
