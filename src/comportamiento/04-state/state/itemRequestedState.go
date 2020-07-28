package state

import (
	"fmt"
)

type ItemRequestedState struct {
	VendingMachine *VendingMachine
}

func (i *ItemRequestedState) RequestItem() error {
	return fmt.Errorf("Item already requested")
}

func (i *ItemRequestedState) AddItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *ItemRequestedState) InsertMoney(money int) error {
	if money < i.VendingMachine.ItemPrice {
		fmt.Errorf("Inserted money is less. Please insert %d", i.VendingMachine.ItemPrice)
	}
	fmt.Println("Money entered is ok")
	i.VendingMachine.SetState(i.VendingMachine.HasMoney)
	return nil
}

func (i *ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("Please insert money first")
}
