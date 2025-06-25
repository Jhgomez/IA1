package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Item struct {
	Value  int
	Weight int
}

const (
	populationSize = 10
	numGenerations = 10
	mutationRate   = 0.05
	maxWeight      = 10
	maxItems       = 3
)

func main() {
	rand.Seed(time.Now().UnixNano())

	items := []Item{
		{Value: 10, Weight: 5},
		{Value: 40, Weight: 4},
		{Value: 30, Weight: 6},
		{Value: 50, Weight: 3},
		{Value: 35, Weight: 2},
		{Value: 25, Weight: 3},
		{Value: 15, Weight: 1},
	}

	best, fitness := geneticAlgorithm(items)
	fmt.Println("\n✅ Final best individual:", best)
	fmt.Println("🏆 Total fitness:", fitness)
}

func geneticAlgorithm(items []Item) ([]int, int) {
	itemCount := len(items)
	population := initializePopulation(itemCount)

	fmt.Println("🔁 Initial Population:")
	for i, ind := range population {
		fmt.Printf("  %2d: %v (fitness=%d)\n", i, ind, fitness(ind, items))
	}

	for gen := 0; gen < numGenerations; gen++ {
		fmt.Printf("\n=== Generation %d ===\n", gen+1)
		population = sortByFitness(population, items)

		bestFitness := fitness(population[0], items)
		fmt.Printf("🏅 Best fitness: %d (individual: %v)\n", bestFitness, population[0])

		newPop := [][]int{population[0], population[1]} // Elitism

		for len(newPop) < populationSize {
			p1 := tournamentSelection(population, items, 3)
			p2 := tournamentSelection(population, items, 3)

			fmt.Println("🔀 Selected parents:")
			fmt.Println("    P1:", p1)
			fmt.Println("    P2:", p2)

			c1, c2 := crossover(p1, p2)

			fmt.Println("👶 Offspring before mutation:")
			fmt.Println("    C1:", c1)
			fmt.Println("    C2:", c2)

			mutate(c1)
			mutate(c2)

			fmt.Println("🧬 Offspring after mutation:")
			fmt.Println("    C1:", c1)
			fmt.Println("    C2:", c2)

			newPop = append(newPop, c1, c2)
		}

		population = newPop[:populationSize]
	}

	best := getBest(population, items)
	return best, fitness(best, items)
}

func initializePopulation(n int) [][]int {
	pop := make([][]int, populationSize)
	for i := 0; i < populationSize; i++ {
		ind := make([]int, n)
		for j := 0; j < n; j++ {
			ind[j] = rand.Intn(2)
		}
		pop[i] = ind
	}
	return pop
}

func fitness(ind []int, items []Item) int {
	totalWeight := 0
	totalValue := 0
	selectedItems := 0

	for i := range ind {
		if ind[i] == 1 {
			totalWeight += items[i].Weight
			totalValue += items[i].Value
			selectedItems++
		}
	}

	if totalWeight > maxWeight || selectedItems > maxItems {
		return 0
	}
	return totalValue
}

func crossover(p1, p2 []int) ([]int, []int) {
	n := len(p1)
	point := rand.Intn(n-1) + 1

	child1 := append([]int(nil), p1[:point]...)
	child1 = append(child1, p2[point:]...)

	child2 := append([]int(nil), p2[:point]...)
	child2 = append(child2, p1[point:]...)

	return child1, child2
}

func mutate(ind []int) {
	for i := range ind {
		if rand.Float64() < mutationRate {
			fmt.Printf("💥 Mutation at index %d in %v\n", i, ind)
			ind[i] = 1 - ind[i]
		}
	}
}

func sortByFitness(pop [][]int, items []Item) [][]int {
	sorted := make([][]int, len(pop))
	copy(sorted, pop)

	for i := 0; i < len(sorted); i++ {
		for j := i + 1; j < len(sorted); j++ {
			if fitness(sorted[j], items) > fitness(sorted[i], items) {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}
	return sorted
}

func tournamentSelection(pop [][]int, items []Item, k int) []int {
	best := pop[rand.Intn(len(pop))]
	for i := 1; i < k; i++ {
		contender := pop[rand.Intn(len(pop))]
		if fitness(contender, items) > fitness(best, items) {
			best = contender
		}
	}
	return append([]int(nil), best...)
}

func getBest(pop [][]int, items []Item) []int {
	best := pop[0]
	for _, ind := range pop {
		if fitness(ind, items) > fitness(best, items) {
			best = ind
		}
	}
	return best
}

// in machine learning there is a type of algorithms called genetic algorithms, I now they can use a concept
// of convergence and I can give the a value to target to maximize, what is it called the distance between the
// value being evaluated and the target value, is it called hope or something else?

// Ok I want to solve a variation of the Knapsack Problem using genetic algorithms, In my problem I have a two
// constraints, the number of items that I'm allowed to put in the bag as well as the total weight of the items
// in the bag are limited by two given variables, as usual I want to maximize the total value of the items while
// complying to the limits given by the two variables I mentioned, I need a genetic algorithm written
// in golang that helps me find the best combination

// In this solution we exit on generation 10,
