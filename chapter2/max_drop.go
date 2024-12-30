package chapter2

func findBreakingPointWithoutBreakingJar(originalLadder, slicedLadder []int, breakingPoint int) int {
	breakingPointIsHigher := slicedLadder[0] <= breakingPoint
	var wrungB4BreakingPoint = -1
	if slicedLadder[0] == breakingPoint { //special  case where highest wrung is at beginning of slice
		//Find breaking point in original ladder and get one before
		for i := 0; i < len(originalLadder); i++ {
			numberOfSingleSteps = numberOfSingleSteps + 1
			if originalLadder[i] == breakingPoint {
				return originalLadder[i-1]
			}
		}
	}
	if breakingPointIsHigher {
		for i := 0; i < len(slicedLadder); i++ {
			numberOfSingleSteps = numberOfSingleSteps + 1
			if slicedLadder[i] >= breakingPoint {
				wrungB4BreakingPoint = slicedLadder[i-1] //the wrung of the slicedLadder right before the breaking point
				break
			}

		}
	} else { //breaking point is lower
		for i := len(slicedLadder) - 1; i >= 0; i-- {
			if slicedLadder[i] >= breakingPoint {
				wrungB4BreakingPoint = slicedLadder[i-1] //the wrung of the slicedLadder right before the breaking point
				break
			}
		}
	}
	return wrungB4BreakingPoint
}

var numberOfSteps = 0
var numberOfSingleSteps = 0

//asymptotic lower bound is (n Log n)
//theta is (n + (n log n))
//asymptotic upper bound is (n)
//breaking point is assumed to be an ordered set of integers from low to high
func HighestBreakingPoint(originalLadder, ladder []int, breakingPoint, budget, usedBudget int) int {
	numberOfSteps = numberOfSteps + 1
	if usedBudget+1 == budget {
		return findBreakingPointWithoutBreakingJar(originalLadder, ladder, breakingPoint)
	} else { //Divide in half again
		halfWayPoint := len(ladder) / 2
		if ladder[halfWayPoint] == breakingPoint { //We are done under budget
			return ladder[halfWayPoint-1]
		}
		lowerHalf := ladder[0:halfWayPoint]
		upperHalf := ladder[halfWayPoint:]
		lastLowerHalfIdx := len(lowerHalf) - 1
		if breakingPoint <= lowerHalf[lastLowerHalfIdx] {
			return HighestBreakingPoint(originalLadder, lowerHalf, breakingPoint, budget, usedBudget+1)
		} else {
			return HighestBreakingPoint(originalLadder, upperHalf, breakingPoint, budget, usedBudget+1)
		}
	}
}
