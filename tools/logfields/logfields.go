package logfields

type Entry struct {
	Name  string
	Value string
}

type Collection []*Entry

func (c *Collection) Add(entries ...*Entry) {
	for _, entry := range entries {
		if entry == nil {
			continue
		}

		*c = append(*c, entry)
	}
}

type Interface interface {
	LogFields() Collection
}

func Extract(m any) Collection {
	if m == nil {
		return nil
	}

	if f, ok := m.(Interface); ok {
		return f.LogFields()
	}

	return nil
}
