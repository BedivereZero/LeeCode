package algorithms

import "log"

func canCompleteCircuit(gas []int, cost []int) int {
	index := 0
	for i := 0; i < len(gas); i++ {
		gas[i] = gas[i] - cost[i]
		if i > 0 {
			gas[i] += gas[i-1]
			if gas[i] < gas[index] {
				index = i
			}
		}
	}
	if gas[len(gas)-1] < 0 {
		return -1
	}
	return (index + 1) % len(gas)
}
func canCompleteCircuit(gas []int, cost []int) int {
	// gen remain
	for i := 0; i < len(gas); i++ {
		gas[i] = gas[i] - cost[i]
	}
	log.Println("A:", gas)

	index := 0
	for i := 1; i < len(gas); i++ {
		gas[i] += gas[i-1]
		if gas[i] < gas[index] {
			index = i
		}
	}
	log.Println("B:", gas)
	if gas[len(gas)-1] < 0 {
		return -1
	}
	return (index + 1) % len(gas)
}
func canCompleteCircuit(gas []int, cost []int) int {
	// gen remain
	for i := 0; i < len(gas); i++ {
		gas[i] = gas[i] - cost[i]
	}
	log.Println("A:", gas)

	index := 0
	for i := 1; i < len(gas); i++ {
		gas[i] += gas[i-1]
		if gas[i] < gas[index] {
			index = i
		}
	}
	log.Println("B:", gas)
	if gas[len(gas)-1] < 0 {
		return -1
	}
	return (index + 1) % len(gas)
}
