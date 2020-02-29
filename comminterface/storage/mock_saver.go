package storage

import (
	"ip_proxy/model"
)

type MockSaver struct {
	items []model.IPItem
}

func (m *MockSaver) SaveIPItems(items []model.IPItem) error {
	m.items = append(m.items, items...)
	return nil
}

func (m *MockSaver) GetTotalNum() (int, error) {
	return len(m.items), nil
}
