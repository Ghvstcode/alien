## **Alien Invasion Challenge**

**Introduction**

In this challenge, you are tasked with writing a simulation of an alien invasion on Earth. The simulation involves multiple cities, and each city can have multiple aliens. The aliens move randomly from city to city and destroy cities they visit along the way. The simulation ends when all the aliens are dead or all cities have been destroyed.

**Solution**

The solution to this challenge consists of two main parts:

* Implementing the City and Alien types, which represent the cities and aliens in the simulation, respectively.
* Writing a simulation engine that runs the simulation by moving aliens from city to city and destroying cities as needed.
City and Alien Types
The City type represents a city in the simulation. It has the following methods:

  1. `AddNeighbor(direction string, neighbor *City)`: Adds a neighboring city in the given direction.
  2. `AddAlien(alien *Alien)`: Adds an alien to the city.
  3. `RemoveAlien(alien *Alien)`: Removes an alien from the city.
  4. `Destroy()`: Destroys the city and removes all aliens from it.
  5. `GetNeighbors() map[string]alien.City`: Returns a map of neighboring cities keyed by direction.
  6. `GetAliens() []*alien.Alien`: Returns a slice of aliens currently in the city.
  7. `GetName() string`: Returns the name of the city. <br/>
  

The simulation engine runs the simulation by moving aliens from city to city and destroying cities as needed. It has the following main steps:

1. Create the cities and aliens in the simulation.
2. Randomly assign aliens to cities.
3. Loop through all aliens and move them to a neighboring city.
4. If two or more aliens end up in the same city, destroy that city and all aliens in it.
5. If a city has no aliens, destroy it.
Repeat steps 3-5 until there are no more aliens or no more cities.

**Usage**

To run the simulation, run the main.go file in the root directory of this repository:

`go run main.go -map <map_file> -aliens <num_aliens>` <br/>

The map_file argument is a text file containing a map of cities and their neighbors. Each line in the file should contain the name of a city, followed by a list of its neighboring cities separated by commas. 

The `num_aliens` argument is the number of aliens to create in the simulation.


**Testing**

The project contains a set of unit tests for the alien and city packages. To run the tests, use the following command:

`go test ./...`

