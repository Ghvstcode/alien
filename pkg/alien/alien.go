package alien

import (
	"fmt"
	"math/rand"
)

type City interface {
	AddAlien(alien *Alien)
	RemoveAlien(alien *Alien)
	Destroy()
	GetNeighbors() map[string]City
	GetName() string
	GetAliens() []*Alien
}
type Alien struct {
	Id     int
	City   City
	moves  int
	Active bool
}

func NewAlien(id int, city City) *Alien {
	return &Alien{Id: id, City: city, moves: 0, Active: true}
}

func (a *Alien) Move() {
	directions := make([]string, 0, len(a.City.GetNeighbors()))
	for direction := range a.City.GetNeighbors() {
		directions = append(directions, direction)
	}
	if len(directions) > 0 {
		direction := directions[rand.Intn(len(directions))]
		neighbor := a.City.GetNeighbors()[direction]
		a.City.RemoveAlien(a)
		neighbor.AddAlien(a)
		a.City = neighbor
		a.moves++
	} else {
		a.Active = false
	}
}

func (a *Alien) Destroy(other *Alien) {
	fmt.Printf("%s has been destroyed by alien %d and alien %d!\n", a.City.GetName(), a.Id, other.Id)
	a.City.Destroy()
	a.Active = false
	other.Active = false
}
