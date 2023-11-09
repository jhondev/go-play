# Mad Aliens Simulation
The mad alien's simulation helps you to simulate an invasion in a world built based on provided custom arguments.

Aliens fight in cities and destroy them.

## Basic Rules
- Map info is given in a text file (check [world map](pkg/data/providers/testdata/world_map.txt))
    - 1-4 directions (north, south, east, or west)
- Aliens are created and located randomly based on the given info
- Each iteration, the aliens can travel in any of the directions leading out of a city
- When aliens end up in the same place, they fight 
    - In the process kill each other and destroy the city
- When a city is destroyed, it is removed from the map, and so are any roads that lead into or out of it
- Once a city is destroyed, aliens can no longer travel to or through it

## Additional Rules (Assumptions)
- More than 2 aliens can end up in the same city
- Each iteration will move all aliens and will check for battles after it
- An alien gets trapped when its city doesn't have more paths

## Play

### Prerequisites
- [Go 1.20](https://golang.org/dl/)

### Print help info
```
make build

./mad -h

./mad run -h
```

### Run with default values (10 aliens, 15 cities, 10 max moves)
```
make run
```

### Run with 3 cities (2 aliens)
```
make run-3-cities
``` 

### Run with custom values
```
make build

./mad run -p <file_path> -n <number of aliens> -m <max moves>
```

### Tests
```
make test
```

### Debug
```
a default debug set of arguments is already configured. check .vscode/launch.json
```

## Technical

### Implementation
The main idea of the implementation is to provide a plug-and-play experience in the code. Applying the `inversion of control principle` we can loose coupling data sources and events stores, adding flexibility to changes, maintainability, and testability.

The primary data structure used is the `map` that gives us nice performant options to query and updates
battlefield states during the different events of the simulation.

A state machine is also built to add the ability to recreate an entire simulation or particular scenarios, giving the option, for instance, to go back "in time" and have a different simulation result from a particular point in history. The state machine is managed in a records memory implementation that can be easily changed.

The IoC principle is also applied for the random function to be used in the location and alien's movements, this way we can inject different types of random function implementations depending on the scenario. In the default setup, the `math/rand` from the std library is injected.

### Structure
Domain packages are separated from the `cli` that way we can have different types of outputs (`cli`, `api`, etc), it just needs to implement the core interfaces and high-order functions. The current `cli` uses the external cobra package to improve the user experience

### Improvements
Lots of stuff can be improved, some examples:
- Adding new providers to manage the state machine and data source
- Adding new options to interact with the simulation (the state machine will help here)
- Adding an output interface to avoid printing from the core packages
- Adding more unit tests
- Adding benchmark tests
    - Improve performance