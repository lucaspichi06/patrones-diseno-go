package state

type State interface {
	AddItem(int) error
	RequestItem() error
	InsertMoney(money int) error
	DispenseItem() error
}
