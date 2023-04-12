package alien_test

import (
	"github.com/Ghvstcode/alien/pkg/alien"
	"testing"
)

type MockCity struct {
	Name      string
	Neighbors map[string]alien.City
	Aliens    []*alien.Alien
	Destroyed bool
}

func (c *MockCity) AddAlien(alien *alien.Alien) {
	c.Aliens = append(c.Aliens, alien)
}

func (c *MockCity) RemoveAlien(alien *alien.Alien) {
	for i, a := range c.Aliens {
		if a == alien {
			c.Aliens = append(c.Aliens[:i], c.Aliens[i+1:]...)
			return
		}
	}
}

func (c *MockCity) Destroy() {
	c.Destroyed = true
}

func (c *MockCity) GetNeighbors() map[string]alien.City {
	return c.Neighbors
}

func (c *MockCity) GetName() string {
	return c.Name
}

func (c *MockCity) GetAliens() []*alien.Alien {
	return c.Aliens
}

func TestAlien_Move(t *testing.T) {
	// create some mock cities and aliens
	city1 := &MockCity{Name: "city1"}
	city2 := &MockCity{Name: "city2"}
	city1.Neighbors = map[string]alien.City{"north": city2}
	city2.Neighbors = map[string]alien.City{"south": city1}
	alien1 := alien.NewAlien(1, city1)

	// move the alien to a neighbor city
	alien1.Move()
	if len(city1.GetAliens()) != 0 {
		t.Errorf("expected city1 to have 0 aliens after alien1 moved, got %d", len(city1.GetAliens()))
	}
	if len(city2.GetAliens()) != 1 {
		t.Errorf("expected city2 to have 1 alien after alien1 moved, got %d", len(city2.GetAliens()))
	}
	if alien1.City != city2 {
		t.Errorf("expected alien1's city to be city2 after moving, got %v", alien1.City)
	}

	// move the alien to a destroyed city
	city2.Destroy()
	alien1.Move()
	if len(city1.GetAliens()) != 0 {
		t.Errorf("expected city1 to have 0 aliens after alien1 moved to destroyed city, got %d", len(city1.GetAliens()))
	}
	if alien1.Active {
		t.Errorf("expected alien1 to be inactive after moving to destroyed city, got %v", alien1.Active)
	}

	// move the alien with no neighbors
	city3 := &MockCity{Name: "city3"}
	alien2 := alien.NewAlien(2, city3)
	alien2.Move()
	if len(city3.GetAliens()) != 1 {
		t.Errorf("expected city3 to have 1 alien after alien2 moved, got %d", len(city3.GetAliens()))
	}
	if alien2.City != city3 {
		t.Errorf("expected alien2's city to be city3 after moving, got %v", alien2.City)
	}
	if alien2.Active {
		t.Errorf("expected alien2 to be inactive after moving to city with no neighbors, got %v", alien2.Active)
	}
}

func TestAlien_Destroy(t *testing.T) {
	// create some mock cities and aliens
	city1 := &MockCity{Name: "city1"}
	city2 := &MockCity{Name: "city2"}
	alien1 := alien.NewAlien(1, city1)
	alien2 := alien.NewAlien(2, city2)

	// Add aliens to cities
	city1.AddAlien(alien1)
	city2.AddAlien(alien2)

	// Destroy one alien
	alien1.Destroy(alien2)

	// Verify that the destroyed alien and the other alien are no longer active
	if alien1.Active {
		t.Errorf("Expected alien1 to be inactive after being destroyed, but it is active")
	}
	if alien2.Active {
		t.Errorf("Expected alien2 to be inactive after destroying alien1, but it is active")
	}

	// Verify that the destroyed alien has been removed from its city's list of aliens
	if len(city1.GetAliens()) != 0 {
		t.Errorf("Expected city1 to have no aliens after alien1 was destroyed, but it has %d aliens", len(city1.GetAliens()))
	}

}
