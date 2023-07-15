package pkg

import (
	"fmt"

	"github.com/agpelkey/combustion-sorc/internal/data"
)

var mockItem = &data.Item{
	ID:    1,
	Name:  "testItem",
	Price: 5,
}

type ItemModel struct{}

func (m *ItemModel) Get(id int) (*data.Item, error) {
	switch id {
	case 1:
		return mockItem, nil
	default:
		return nil, fmt.Errorf("Record could not be found")
	}
}
