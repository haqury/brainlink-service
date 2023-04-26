package entity

type Map struct {
	Id   int
	maps map[int]*reaction
}

func (m Map) AddEvent(event Event) error {
	r := m.maps[event.EegDto[0].Input.LowAlpha]
	if r == nil {
		r = getReactionByEvent(event)
	}
	r.Events = append(r.Events, event)
	m.maps[event.EegDto[0].Input.LowAlpha] = r
	return nil
}

type reaction struct {
	Action   string
	Position [][2]int
	Value    []string
	Events   []Event
}

func getReactionByEvent(event Event) *reaction {
	r := reaction{
		Action:   "mouse",
		Position: [][2]int{{event.EegDto[0].System.ToX, event.EegDto[0].System.ToY}},
	}
	return &r
}
