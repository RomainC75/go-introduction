package repos

type Flight struct {
	Id       int
	Capacity int
	Weight   float64
}

type Flights struct {
	flights []Flight
}

func NewFlights(flights []Flight) *Flights {
	return &Flights{
		flights: flights,
	}
}

func (f *Flights) GetIds() []int {
	ids := make([]int, 0, len(f.flights))
	for _, f := range f.flights {
		ids = append(ids, f.Id)
	}
	return ids
}
