package state

import (
	"fmt"
)

type NoItemState struct {
	VendingMachine *VendingMachine
}

func (i *NoItemState) RequestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) AddItem(count int) error {
	i.VendingMachine.IncrementItemCount(count)
	i.VendingMachine.SetState(i.VendingMachine.HasItem)
	return nil
}

func (i *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) DispenseItem() error {
	return fmt.Errorf("Item out of stock")
}
