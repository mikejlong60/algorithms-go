package chapter1

import (
	"fmt"
	"github.com/greymatter-io/golangz/linked_list"
)

//Perform the outer loop until every hospital has reached it's resident capacity.
//Inside that loop at first level make another loop that fills hospital resident capacity until complete.
type Hospital struct { //man
	Id                  string
	ResidentCapacity    int
	Residents           map[string]*Resident
	ResidentPreferences *linked_list.LinkedList[*Resident] //A stack of residents I want in order of preferences. When a resident is missing from it a hospital has already proposed to him.
}

type Resident struct { //woman
	Id string
	//HospitalPreferences map[string]int//A raking of hospital preferences for a resident. All the hospitals are in this map. Key is hospital ID and ranking is like a golf score, low is better.
	HospitalPreferences []*Hospital //An array of hospital preferences for a resident. All the hospitals are in this array.
	Hospital            *Hospital   //The hospital the resident is currently assigned to
}

func MatchResidentsToHospitals(hospitalsWithResidentOpenings *linked_list.LinkedList[*Hospital]) []*Resident {
	fmt.Printf("Size of list:%v\n", linked_list.Len(hospitalsWithResidentOpenings))

	//Invariants
	// 1. The sum of all hospital resident openings cannot exceed the total number of residents
	// 2. Due to preceding there can be more residents than there are resident openings at hospitals
	validateInvariants := func(hospitalsWithResidentOpenings *linked_list.LinkedList[*Hospital]) {
		xs := linked_list.ToArray(hospitalsWithResidentOpenings)
		var z int
		for _, y := range xs {
			z = z + y.ResidentCapacity
		}
		if z > linked_list.Len(xs[0].ResidentPreferences) {
			panic("Cannot have more hospital resident openings than you have residents")
		}
	}
	if linked_list.Len(hospitalsWithResidentOpenings) == 0 {
		return []*Resident{}
	}

	validateInvariants(hospitalsWithResidentOpenings)

	allResidents := linked_list.ToArray(linked_list.Head(hospitalsWithResidentOpenings).ResidentPreferences)

	mEq := func(m1 *Hospital, m2 *Hospital) bool {
		if m1.Id == m2.Id {
			return true
		} else {
			return false
		}
	}

	residentPrefersThisHospital := func(res *Resident, thisHospital *Hospital) bool { //Does resident prefer this hospital to the one to which he is currently assigned?
		var thisHospitalRanking = -1
		var currentHospitalRanking = -1
		for i, m := range res.HospitalPreferences {
			if mEq(m, thisHospital) {
				thisHospitalRanking = i
			} else if mEq(res.Hospital, m) {
				currentHospitalRanking = i
			}
			if thisHospitalRanking != -1 && currentHospitalRanking != -1 {
				break
			}
		}
		if thisHospitalRanking < currentHospitalRanking {
			fmt.Printf("resident:%v prefers this hospital:%v over current one:%v\n", res.Id, thisHospital.Id, res.Hospital.Id)
			return true
		} else {
			fmt.Printf("resident:%v prefers current hospital:%v over this one:%v\n", res.Id, res.Hospital.Id, thisHospital.Id)
			return false
		}
	}

	for hospitalsWithResidentOpenings != nil {
		thisHospital := linked_list.Head(hospitalsWithResidentOpenings)
		var hospitalResidentPreferences = thisHospital.ResidentPreferences
		for hospitalResidentPreferences != nil && len(thisHospital.Residents) < thisHospital.ResidentCapacity { //Loop over the hospital's resident preferences until the hospital has reached it' capacity
			resident := linked_list.Head(hospitalResidentPreferences)
			if resident.Hospital == nil {
				resident.Hospital = thisHospital
				hospitalResidentPreferences, _ = linked_list.Tail(hospitalResidentPreferences)
				thisHospital.Residents[resident.Id] = resident
			} else {
				//Does this resident prefer this hospital to the one to whom he is currently assigned? If so he
				//breaks his agreement to that hospital and you make hospital have an additional resident opening.
				//Otherwise just try the next hospital in the current resident's non-proposed-to(preferences) stack.
				if residentPrefersThisHospital(resident, thisHospital) {
					oldHospital := resident.Hospital
					delete(oldHospital.Residents, resident.Id)
					///Set up current resident with this hospital
					resident.Hospital = thisHospital
					thisHospital.Residents[resident.Id] = resident
					hospitalsWithResidentOpenings = linked_list.AddLast(oldHospital, hospitalsWithResidentOpenings)
				}
			}
			hospitalResidentPreferences, _ = linked_list.Tail(thisHospital.ResidentPreferences)
			thisHospital.ResidentPreferences = hospitalResidentPreferences
		} //end resident for
		hospitalsWithResidentOpenings, _ = linked_list.Tail(hospitalsWithResidentOpenings)
	} // end hospital with resident openings for
	return allResidents
}
