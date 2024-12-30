package chapter1

import "fmt"

type ShipState = int //0 is at sea. Non-zero is port number
type Ship struct {
	Id               int
	ProposedSchedule []ShipState
	ActualSchedule   []ShipState
}

func (w Ship) String() string {
	return fmt.Sprintf("Id:%v, ProposedSchedule:%v, ActualSchedule:%v", w.Id, w.ProposedSchedule, w.ActualSchedule)
}

//Algorithm
//1. range loop over array of ships, index i
//2. For current ship range loop over all other ships and make sure there is not a port conflict with element i in any of their proposed schedules
//2.1  If there are no port conflicts, add ProposedSchedule element to end of ActualSchedule. A port conflict is same port at same element. At sea is no conflict
//2.2  Otherwise skip to next element in array of ships
//3 Return array of ships with ActualSchedule
func schedule(ships []*Ship) []*Ship {
	if len(ships) == 0 {
		return []*Ship{}
	}

	//Check the next ship's schedule to see if there are any conflicts
	fEarliestConflict := func(currentShipSchedule, otherShipSchedule []int) int {
		for i, shipState := range currentShipSchedule {
			atSea := shipState == 0
			shipPortConflict := shipState == otherShipSchedule[i]
			if !atSea && shipPortConflict {
				return i
			}
		}
		return len(currentShipSchedule)
	}

	//An invariant is that all ships have same calendar length reflected in the size of their proposed schedule
	for i, ship := range ships { //Range loop over array of all ships
		currentShipSchedule := ship.ProposedSchedule
		var earliestConflict = len(currentShipSchedule)
		var finalEarliestConflict = earliestConflict
		for j := i + 1; j < len(ships); j++ { //For current ship iterate over all ships later in array and truncate ship's Proposed Schedule at earliest conflict
			otherShipSchedule := ships[j].ProposedSchedule
			earliestConflict = fEarliestConflict(currentShipSchedule, otherShipSchedule)
			if earliestConflict < finalEarliestConflict {
				finalEarliestConflict = earliestConflict
			}
		}
		ship.ActualSchedule = ship.ProposedSchedule[0:finalEarliestConflict]
	}
	return ships
}
