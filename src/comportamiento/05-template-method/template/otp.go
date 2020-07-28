package template

type IOtp interface {
	GenRandomOTP(len int) string
	SaveOTPCache(string)
	GetMessage(string) string
	SendNotification(string) error
	PublishMetric()
}

type Otp struct {
	IOtp IOtp
}

func (o *Otp) GenAdnSendOTP(otpLength int) error {
	otp := o.IOtp.GenRandomOTP(otpLength)
	o.IOtp.SaveOTPCache(otp)
	message := o.IOtp.GetMessage(otp)
	if err := o.IOtp.SendNotification(message); err != nil {
		return err
	}
	o.IOtp.PublishMetric()
	return nil
}
