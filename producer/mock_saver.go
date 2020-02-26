package producer

type MockSaver struct {
	items []IPItem
}

func (m *MockSaver) SaveIpItems(items []IPItem) error {
	m.items = append(m.items, items...)
	return nil
}

func (m *MockSaver) GetTotalNum() int {
	return len(m.items)
}