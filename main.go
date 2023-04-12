package main

import (
	"bufio"
	"fmt"
	"github.com/Ghvstcode/alien/pkg/alien"
	"github.com/Ghvstcode/alien/pkg/city"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Parse command-line arguments
	if len(os.Args) != 2 {
		log.Fatal("Usage: alien-invasion <num-aliens>")
	}

	numAliens, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Invalid argument: num-aliens must be an integer")
	}

	// Read input file and create cities
	file, err := os.Open("map.txt")
	if err != nil {
		log.Fatal("Cannot open input file")
	}
	defer file.Close()

	cities := make(map[string]*city.City)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 2 {
			log.Fatalf("Invalid input: %s", line)
		}
		cityName := fields[0]
		if _, ok := cities[cityName]; !ok {
			cities[cityName] = city.NewCity(cityName)
		}
		for _, field := range fields[1:] {
			parts := strings.Split(field, "=")
			if len(parts) != 2 {
				log.Fatalf("Invalid input: %s", line)
			}
			direction, neighborName := parts[0], parts[1]
			if _, ok := cities[neighborName]; !ok {
				cities[neighborName] = city.NewCity(neighborName)
			}
			cities[cityName].AddNeighbor(direction, cities[neighborName])
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Cannot read input file")
	}

	// Create aliens and unleash them
	rand.Seed(time.Now().UnixNano())
	aliens := make([]*alien.Alien, numAliens)
	citiesSlice := make([]*city.City, 0, len(cities))
	for _, city := range cities {
		citiesSlice = append(citiesSlice, city)
	}
	for i := 0; i < numAliens; i++ {
		city := citiesSlice[rand.Intn(len(citiesSlice))]
		aliens[i] = alien.NewAlien(i+1, city)
	}

	// Run simulation until all aliens are dead or trapped
	for {
		// Move aliens
		for _, alien := range aliens {
			if !alien.Active {
				continue
			}
			neighborCities := make([]*city.City, 0, len(alien.City.GetNeighbors()))
			for _, neighbor := range alien.City.GetNeighbors() {
				neighborCities = append(neighborCities, neighbor.(*city.City))
			}
			if len(neighborCities) == 0 {
				alien.Active = false
				continue
			}
			if len(neighborCities) == 1 {
				neighbor := neighborCities[0]
				neighbor.AddAlien(alien)
				alien.City = neighbor
			} else {
				rand.Seed(time.Now().UnixNano())
				neighbor := neighborCities[rand.Intn(len(neighborCities))]
				neighbor.AddAlien(alien)
				alien.City.RemoveAlien(alien)
				alien.City = neighbor
			}
			// Check for destroyed cities
			if len(alien.City.GetAliens()) > 1 {
				for _, other := range alien.City.GetAliens() {
					if other.Id != alien.Id {
						alien.Destroy(other)
						break
					}
				}
			} else if len(alien.City.GetAliens()) == 1 {
				// Check for trapped aliens
				trapped := true
				for _, neighbor := range alien.City.GetNeighbors() {
					if len(neighbor.GetAliens()) == 0 {
						trapped = false
						break
					}
				}
				if trapped {
					alien.Active = false
				}
			} else {
				alien.Active = false
			}
		}

		// Check for end of simulation
		numActiveAliens := 0
		for _, alien := range aliens {
			if alien.Active {
				numActiveAliens++
			}
		}
		if numActiveAliens == 0 {
			fmt.Println("All aliens dead or trapped")
			break
		}
	}
}
