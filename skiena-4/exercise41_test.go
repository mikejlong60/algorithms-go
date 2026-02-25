package skiena_4

func binarySearch(sortedArray []string, target string) bool {

	if len(sortedArray) == 0 {
		return false
	}

	if len(sortedArray) == 1 && sortedArray[0] == target {
		return true
	}

	if len(sortedArray) == 2 {
		if sortedArray[0] == target || sortedArray[1] == target {
			return true
		}
	}

	midpoint := len(sortedArray) / 2
	if sortedArray[midpoint] == target {
		return true
	}

	if sortedArray[midpoint] < target {
		return binarySearch(sortedArray[:midpoint], target)
	}
	
	if sortedArray[midpoint] > target {
		return binarySearch(sortedArray[midpoint:], target)
	}
	return false
}
