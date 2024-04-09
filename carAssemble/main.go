package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100.0)
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(carsPerHour / 60)
}

func CalculateCost(carsCount int) uint {
	groupsOfTen := carsCount / 10

	individualCars := carsCount % 10

	costWithGroups := uint(groupsOfTen) * 95000

	costWithIndividual := uint(individualCars) * 10000

	totalCost := costWithGroups + costWithIndividual
	return totalCost
}
