package state

import (
	"fmt"
)

type HasItemState struct {
	VendingMachine *VendingMachine
}

func (i *HasItemState) RequestItem() error {
	if i.VendingMachine.ItemCount == 0 {
		i.VendingMachine.SetState(i.VendingMachine.NoItem)
		return fmt.Errorf("No item present")
	}
	fmt.Printf("Item requested\n")
	i.VendingMachine.SetState(i.VendingMachine.ItemRequested)
	return nil
}

func (i *HasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.VendingMachine.IncrementItemCount(count)
	return nil
}

func (i *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

func (i *HasItemState) DispenseItem() error {
	return fmt.Errorf("Please select item first")
}
