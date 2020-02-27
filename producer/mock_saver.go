package producer

type MockSaver struct {
	items []IPItem
}

func (m *MockSaver) SaveIPItems(items []IPItem) error {
	m.items = append(m.items, items...)
	return nil
}

func (m *MockSaver) GetTotalNum() (int, error) {
	return len(m.items), nil
}
