package mytravel

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

func init() {
	timezone, _ = time.LoadLocation("Europe/Warsaw")
}

var (
	timezone *time.Location
	Origin   Station = Station{
		ID:      uuid.Must(uuid.NewV7()),
		Name:    "Wrocław Główny",
		Address: "ul. Sucha",
	}
	Destination Station = Station{
		ID:      uuid.Must(uuid.NewV7()),
		Name:    "Sobótka",
		Address: "",
	}
	NextTrip Trip = Trip{
		ID:              uuid.Must(uuid.NewV7()),
		Origin:          Origin,
		Destination:     Destination,
		DepartureTime:   time.Date(2025, time.January, 18, 9, 17, 0, 0, timezone),
		ArrivalTime:     time.Date(2025, time.January, 18, 10, 0, 0, 0, timezone),
		Operator:        "KD",
		Price:           15.51,
		BicyclesAllowed: true,
		DogsAllowed:     true,
	}
	Bookings = NewStore[uuid.UUID, Booking]()
	Payments = NewStore[uuid.UUID, Payment]()
)

type Station struct {
	ID      uuid.UUID
	Name    string
	Address string
}

type Trip struct {
	ID              uuid.UUID
	Origin          Station
	Destination     Station
	DepartureTime   time.Time
	ArrivalTime     time.Time
	Operator        string
	Price           float64
	BicyclesAllowed bool
	DogsAllowed     bool
}

type Booking struct {
	ID            uuid.UUID
	TripID        uuid.UUID
	PassengerName string
	HasBicycle    bool
	HasDog        bool
}

type Payment struct {
	ID        uuid.UUID
	BookingID uuid.UUID
	Amount    float64
	Currency  string
	Card      Card
	Status    string
}

type Card struct {
	Name, Number, CVC string
	ExpMonth, ExpYear int64
}

type Store[K comparable, V any] struct {
	data map[K]V
	mu   sync.Mutex
}

func NewStore[K comparable, V any]() *Store[K, V] {
	return &Store[K, V]{
		data: make(map[K]V),
		mu:   sync.Mutex{},
	}
}

func (s *Store[K, V]) Get(k K) (V, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	v, ok := s.data[k]
	return v, ok
}

func (s *Store[K, V]) Set(k K, v V) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[k] = v
	return
}
