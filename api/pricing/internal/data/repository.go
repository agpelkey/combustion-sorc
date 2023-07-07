package data

type DatabaseRepo interface {
	GetItemByID(id int) (*Item, error)
}
