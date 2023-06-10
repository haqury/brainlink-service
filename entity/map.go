package entity

type Map map[int]*Reaction

func (m Map) AddEvent(event Event) error {
	r := m[event.EegDto[0].Input.LowAlpha]
	if r == nil {
		r = getReactionByEvent(event)
	}
	r.Events = append(r.Events, event)
	m[event.EegDto[0].Input.Meditation] = r
	return nil
}

type Reaction struct {
	Action   string
	Position [][2]int
	Value    []string
	Events   []Event
}

func getReactionByEvent(event Event) *Reaction {
	r := Reaction{
		Action:   "mouse",
		Position: [][2]int{{event.EegDto[0].System.ToX, event.EegDto[0].System.ToY}},
	}
	return &r
}
