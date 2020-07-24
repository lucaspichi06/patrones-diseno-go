package main

import (
	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/02-chain-of-responsibility/domain"
	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/02-chain-of-responsibility/handler/implementation"
)

func main() {
	cashier := &implementation.Cashier{}
	//Set next for medical department
	medical := &implementation.Medical{}
	medical.SetNext(cashier)
	//Set next for doctor department
	doctor := &implementation.Doctor{}
	doctor.SetNext(medical)
	//Set next for reception department
	reception := &implementation.Reception{}
	reception.SetNext(doctor)
	patient := &domain.Patient{
		Name: "Anthony Kiddies",
	}
	//Patient visiting
	reception.Execute(patient)
}
