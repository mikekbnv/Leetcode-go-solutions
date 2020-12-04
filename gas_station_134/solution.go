package gas_station_134

func canCompleteCircuit(gas []int, cost []int) int {
	tank, index, sum := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
		tank += gas[i] - cost[i]

		if tank < 0 {
			index = i + 1
			tank = 0
		}

	}
	if sum < 0 {
		return -1
	}
	return index
}
