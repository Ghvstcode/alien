package city_test

import (
	"github.com/Ghvstcode/alien/pkg/alien"
	"github.com/Ghvstcode/alien/pkg/city"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCity_AddNeighbor(t *testing.T) {
	c1 := city.NewCity("City 1")
	c2 := city.NewCity("City 2")

	c1.AddNeighbor("north", c2)
	assert.Equal(t, c2, c1.GetNeighbors()["north"])
}

func TestCity_AddAlien(t *testing.T) {
	c1 := city.NewCity("City 1")
	a := alien.NewAlien(1, c1)

	c1.AddAlien(a)
	assert.Contains(t, c1.GetAliens(), a)
}

func TestCity_RemoveAlien(t *testing.T) {
	c1 := city.NewCity("City 1")
	a1 := alien.NewAlien(1, c1)
	a2 := alien.NewAlien(2, c1)

	c1.AddAlien(a1)
	c1.AddAlien(a2)

	c1.RemoveAlien(a1)
	assert.NotContains(t, c1.GetAliens(), a1)
	assert.Contains(t, c1.GetAliens(), a2)

	c1.RemoveAlien(a2)
	assert.NotContains(t, c1.GetAliens(), a2)
}

func TestCity_Destroy(t *testing.T) {
	c1 := city.NewCity("City 1")
	c2 := city.NewCity("City 2")
	a1 := alien.NewAlien(1, c1)
	a2 := alien.NewAlien(2, c1)

	c1.AddNeighbor("north", c2)
	c1.AddAlien(a1)
	c1.AddAlien(a2)

	c1.Destroy()

	assert.Len(t, c1.GetAliens(), 0)
	assert.False(t, a1.Active)
	assert.False(t, a2.Active)
	assert.Len(t, c2.GetAliens(), 0)
}

func TestCity_GetNeighbors(t *testing.T) {
	c1 := city.NewCity("City 1")
	c2 := city.NewCity("City 2")
	a := alien.NewAlien(1, c1)

	c1.AddNeighbor("north", c2)
	c1.AddAlien(a)

	neighbors := c1.GetNeighbors()

	assert.Len(t, neighbors, 1)
	assert.Contains(t, neighbors, "north")
	assert.Equal(t, c2, neighbors["north"])
}

func TestCity_GetAliens(t *testing.T) {
	c1 := city.NewCity("City 1")
	a1 := alien.NewAlien(1, c1)
	a2 := alien.NewAlien(2, c1)

	c1.AddAlien(a1)
	c1.AddAlien(a2)

	aliens := c1.GetAliens()

	assert.Len(t, aliens, 2)
	assert.Contains(t, aliens, a1)
	assert.Contains(t, aliens, a2)
}

func TestCity_GetName(t *testing.T) {
	c := city.NewCity("City 1")
	assert.Equal(t, "City 1", c.GetName())
}
