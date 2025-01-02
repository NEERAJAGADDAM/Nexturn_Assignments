package main

import (
	"fmt"
	"strings"
)

// Define a struct to store city data
type City struct {
	Name        string
	AverageTemp float64
	Rainfall    float64
}

// Function to find city with highest temperature
func highestTemperature(cities []City) City {
	var highest City
	for _, city := range cities {
		if city.AverageTemp > highest.AverageTemp {
			highest = city
		}
	}
	return highest
}

// Function to find city with lowest temperature
func lowestTemperature(cities []City) City {
	var lowest City
	for _, city := range cities {
		if lowest.AverageTemp == 0 || city.AverageTemp < lowest.AverageTemp {
			lowest = city
		}
	}
	return lowest
}

// Function to calculate the average rainfall
func averageRainfall(cities []City) float64 {
	var totalRainfall float64
	for _, city := range cities {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(cities))
}

// Function to filter cities by rainfall threshold
func filterCitiesByRainfall(cities []City, threshold float64) []City {
	var filteredCities []City
	for _, city := range cities {
		if city.Rainfall > threshold {
			filteredCities = append(filteredCities, city)
		}
	}
	return filteredCities
}

// Function to search for a city by name
func searchCityByName(cities []City, name string) (City, bool) {
	for _, city := range cities {
		if strings.EqualFold(city.Name, name) {
			return city, true
		}
	}
	return City{}, false
}

func main() {
	var numCities int
	fmt.Print("Enter the number of cities: ")
	fmt.Scan(&numCities)

	// Create an empty slice of cities
	var cities []City

	// Input data for each city
	for i := 0; i < numCities; i++ {
		var name string
		var temp, rainfall float64

		// Take input for each city
		fmt.Printf("\nEnter details for city %d:\n", i+1)
		fmt.Print("City Name: ")
		fmt.Scan(&name)
		fmt.Print("Average Temperature (°C): ")
		fmt.Scan(&temp)
		fmt.Print("Rainfall (mm): ")
		fmt.Scan(&rainfall)

		// Append the city data to the cities slice
		cities = append(cities, City{Name: name, AverageTemp: temp, Rainfall: rainfall})
	}

	var choice int
	for {
		// Menu for user to choose an operation
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Find City with Highest Temperature")
		fmt.Println("2. Find City with Lowest Temperature")
		fmt.Println("3. Calculate Average Rainfall")
		fmt.Println("4. Filter Cities by Rainfall")
		fmt.Println("5. Search City by Name")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Highest Temperature
			highest := highestTemperature(cities)
			fmt.Printf("City with the highest average temperature: %s (%.2f°C)\n", highest.Name, highest.AverageTemp)

		case 2:
			// Lowest Temperature
			lowest := lowestTemperature(cities)
			fmt.Printf("City with the lowest average temperature: %s (%.2f°C)\n", lowest.Name, lowest.AverageTemp)

		case 3:
			// Average Rainfall Calculation
			avgRainfall := averageRainfall(cities)
			fmt.Printf("Average rainfall across all cities: %.2f mm\n", avgRainfall)

		case 4:
			// Filter Cities by Rainfall
			var threshold float64
			fmt.Print("Enter rainfall threshold (mm): ")
			_, err := fmt.Scanf("%f", &threshold)
			if err != nil {
				fmt.Println("Invalid input for threshold. Please enter a valid number.")
				break
			}

			filteredCities := filterCitiesByRainfall(cities, threshold)
			if len(filteredCities) > 0 {
				fmt.Println("Cities with rainfall above the threshold:")
				for _, city := range filteredCities {
					fmt.Printf("- %s: %.2f mm\n", city.Name, city.Rainfall)
				}
			} else {
				fmt.Println("No cities found with rainfall above the threshold.")
			}

		case 5:
			// Search City by Name
			var searchName string
			fmt.Print("Enter city name to search: ")
			fmt.Scan(&searchName)

			city, found := searchCityByName(cities, searchName)
			if found {
				fmt.Printf("City found: %s\nAverage Temperature: %.2f°C\nRainfall: %.2f mm\n", city.Name, city.AverageTemp, city.Rainfall)
			} else {
				fmt.Println("City not found.")
			}

		case 6:
			// Exit the program
			fmt.Println("Exiting the program...")
			return

		default:
			fmt.Println("Invalid choice! Please select a valid option.")
		}
	}
}
