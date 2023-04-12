package city

import (
	"github.com/Ghvstcode/alien/pkg/alien"
)

type City struct {
	name      string
	neighbors map[string]*City
	aliens    []*alien.Alien
}

func NewCity(name string) *City {
	return &City{name: name, neighbors: make(map[string]*City)}
}

func (c *City) AddNeighbor(direction string, neighbor *City) {
	c.neighbors[direction] = neighbor
}

func (c *City) AddAlien(alien *alien.Alien) {
	c.aliens = append(c.aliens, alien)
}

func (c *City) RemoveAlien(alien *alien.Alien) {
	for i, a := range c.aliens {
		if a == alien {
			c.aliens = append(c.aliens[:i], c.aliens[i+1:]...)
			return
		}
	}
}

func (c *City) Destroy() {
	for _, aliens := range c.aliens {
		aliens.Active = false
	}
	for _, neighbor := range c.neighbors {
		neighbor.RemoveAlien(c.aliens[0])
	}
}

func (c *City) GetNeighbors() map[string]alien.City {
	alienNeighbours := make(map[string]alien.City)
	for name, city := range c.neighbors {
		alienNeighbours[name] = alien.City(city)
	}
	return alienNeighbours
}

func (c *City) GetAliens() []*alien.Alien {
	return c.aliens
}

func (c *City) GetName() string {
	return c.name
}
