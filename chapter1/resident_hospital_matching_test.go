package chapter1

import (
	"github.com/greymatter-io/golangz/linked_list"
	"github.com/greymatter-io/golangz/sets"
	"testing"
)

func TestResidentHospitalMatching(t *testing.T) {
	res1 := &Resident{
		Id:                  "1",
		HospitalPreferences: nil,
		Hospital:            nil,
	}
	res2 := &Resident{
		Id:                  "2",
		HospitalPreferences: nil,
		Hospital:            nil,
	}
	res3 := &Resident{
		Id:                  "3",
		HospitalPreferences: nil,
		Hospital:            nil,
	}
	res4 := &Resident{
		Id:                  "4",
		HospitalPreferences: nil,
		Hospital:            nil,
	}
	res5 := &Resident{
		Id:                  "5",
		HospitalPreferences: nil,
		Hospital:            nil,
	}

	hosp1 := &Hospital{
		Id:                  "1",
		ResidentCapacity:    2,
		Residents:           make(map[string]*Resident),
		ResidentPreferences: nil,
	}
	hosp2 := &Hospital{
		Id:                  "2",
		ResidentCapacity:    2,
		Residents:           make(map[string]*Resident),
		ResidentPreferences: nil,
	}

	res1.HospitalPreferences = []*Hospital{hosp2, hosp1}
	res2.HospitalPreferences = []*Hospital{hosp2, hosp1}
	res3.HospitalPreferences = []*Hospital{hosp1, hosp2}
	res4.HospitalPreferences = []*Hospital{hosp1, hosp2}
	res5.HospitalPreferences = []*Hospital{hosp1, hosp2}

	var hospitalsWithResidentOpenings *linked_list.LinkedList[*Hospital]
	var hosp1ResPrefs *linked_list.LinkedList[*Resident]
	hosp1ResPrefs = linked_list.Push(res5, hosp1ResPrefs)
	hosp1ResPrefs = linked_list.Push(res4, hosp1ResPrefs)
	hosp1ResPrefs = linked_list.Push(res3, hosp1ResPrefs)
	hosp1ResPrefs = linked_list.Push(res2, hosp1ResPrefs)
	hosp1ResPrefs = linked_list.Push(res1, hosp1ResPrefs)

	hosp1.ResidentPreferences = hosp1ResPrefs
	hosp2.ResidentPreferences = hosp1ResPrefs //TODO Fix this to be different
	hospitalsWithResidentOpenings = linked_list.Push(hosp2, hospitalsWithResidentOpenings)
	hospitalsWithResidentOpenings = linked_list.Push(hosp1, hospitalsWithResidentOpenings)
	MatchResidentsToHospitals(hospitalsWithResidentOpenings)
	var hosp1Residents []string
	var hosp2Residents []string

	eq := func(l, r string) bool {
		if l == r {
			return true
		} else {
			return false

		}
	}
	for _, y := range hosp1.Residents {
		hosp1Residents = append(hosp1Residents, y.Id)
	}
	for _, y := range hosp2.Residents {
		hosp2Residents = append(hosp2Residents, y.Id)
	}
	if !sets.SetEquality(hosp1Residents, []string{"4", "3"}, eq) {
		t.Errorf("Actual:%v Expected:%v", hosp1Residents, []string{"4", "3"})
	}
	if !sets.SetEquality(hosp2Residents, []string{"2", "1"}, eq) {
		t.Errorf("Actual:%v Expected:%v", hosp2Residents, []string{"2", "1"})
	}
}
