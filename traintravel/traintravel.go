package traintravel

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

func init() {
	Timezone, _ = time.LoadLocation("Europe/Warsaw")
}

var (
	Timezone *time.Location
	Country          = "PL"
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
		DepartureTime:   time.Date(2025, time.January, 18, 9, 17, 0, 0, time.UTC),
		ArrivalTime:     time.Date(2025, time.January, 18, 10, 0, 0, 0, time.UTC),
		Operator:        "KD",
		Price:           15.51,
		BicyclesAllowed: true,
		DogsAllowed:     true,
	}
	bookings = newStore[uuid.UUID, Booking]()
	payments = newStore[uuid.UUID, Payment]()
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
	Price           float32
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

func GetBooking(id uuid.UUID) (Booking, bool) {
	return bookings.get(id)
}

func NewBooking(tripID uuid.UUID, passengerName string, hasDog bool) Booking {
	b := Booking{
		ID:            uuid.Must(uuid.NewV7()),
		TripID:        tripID,
		PassengerName: passengerName,
		HasDog:        hasDog,
	}
	bookings.set(b.ID, b)
	return b
}

type Payment struct {
	ID        uuid.UUID
	BookingID uuid.UUID
	Amount    float32
	Currency  string
	Card      Card
	Status    string
}

func NewPayment(bookingID uuid.UUID, amount float32, currency string, card Card) Payment {
	p := Payment{
		ID:        uuid.Must(uuid.NewV7()),
		BookingID: bookingID,
		Amount:    amount,
		Currency:  currency,
		Card:      card,
		Status:    "pending",
	}
	payments.set(p.ID, p)
	return p
}

type Card struct {
	Name, Number, CVC string
	ExpMonth, ExpYear int64
}

type store[K comparable, V any] struct {
	data map[K]V
	mu   sync.Mutex
}

func newStore[K comparable, V any]() *store[K, V] {
	return &store[K, V]{
		data: make(map[K]V),
		mu:   sync.Mutex{},
	}
}

func (s *store[K, V]) get(k K) (V, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	v, ok := s.data[k]
	return v, ok
}

func (s *store[K, V]) set(k K, v V) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[k] = v
	return
}
