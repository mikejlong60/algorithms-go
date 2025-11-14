package chapter1

import (
	"fmt"

	"github.com/greymatter-io/golangz/linked_list"
	"github.com/greymatter-io/golangz/propcheck"
)

//This algorithm accommodates both weak and strong instabilities.

// Definition of Instability - For a given married woman, does a man exist that she prefers over her current husband who also prefers her over his current wife.
// Given that definition of instability there can be no lying about preferences for a woman because a man can only propose once to a woman
// and that woman has to decide if that man is more preferred than her current husband on her list. It's a contradiction to expect that
// the preference list of the woman is ever disregarded. That would mean some matches would be unstable because a woman would be married to a man
// that violated the definition of a stable match given above.
type Man struct {
	Id                int
	HaveNotProposedTo *linked_list.LinkedList[*Woman] //A stack of women I want in order of preferences. When a woman is missing from it he has already proposed to her.
	Preferences       []*Woman                        //A list of preferences for man, lower index is a higher preference. It's an immutable representation of the Man's original preference list in HaveNotProposedTo above.
	EngagedTo         *Woman
}

type Woman struct {
	Id          int
	Preferences map[int]propcheck.Pair[int, *Man] // The key is the man's Id and the value is that man's ranking with 0 being the highest and a pointer to the complete Man.  No duplicate rankings are allowed.
	EngagedTo   *Man
}

//func (w Man) String() string {
//	var prefs []string
//	for _, iw := range w.Preferences {
//		prefs = append(prefs, iw.Id)
//	}
//	return fmt.Sprintf("Man{Id:%v, EngagedTo:%v, Preferences:%v}", w.Id, w.EngagedTo.Id, prefs)
//}
//
//func (w Woman) String() string {
//	var prefs = make(map[string]string) //map[string]int
//	for _, iw := range w.Preferences {
//		prefs[iw.B.Id] = fmt.Sprintf("Prefs:%v popp", iw.A)
//	}
//	return fmt.Sprintf("Woman{Id:%v, EngagedTo:%v, Preferences:%v}", w.Id, w.EngagedTo.Id, prefs)
//}

func womanPrefersMe(wp *Woman, courtier *Man) bool { //Does woman prefer this man to the one to which she is currently assigned?
	//This function assumes that the wp woman is already engaged.
	courtierRanking, courtierIsInPreferredList := wp.Preferences[courtier.Id]
	currentFianceeRanking, currentFianceeIsInPreferredList := wp.Preferences[wp.EngagedTo.Id]

	if !courtierIsInPreferredList && !currentFianceeIsInPreferredList { //Woman is indifferent to both men. So she will stick with her current fiancee.
		//fmt.Printf("Woman:%v is indifferent to both men. So she will stick with her current fiancee%v\n", wp.Id, wp.EngagedTo.Id)
		return false
	} else if courtierIsInPreferredList && !currentFianceeIsInPreferredList { //Woman is indifferent to her current husband but not the courtier. So she chooses the courtier.
		//fmt.Printf("Woman:%v is indifferent to her current husband:%v but not the courtier:%v. So she chooses the courtier.\n", wp.Id, wp.EngagedTo.Id, courtier.Id)
		return true
	} else if !courtierIsInPreferredList && currentFianceeIsInPreferredList { //Woman is indifferent to the courtier but she prefers her current fiancee because he is in her list of preferences.
		//So she sticks with her current fiancee
		//fmt.Printf("Woman:%v is indifferent to the courtier:%v but her current fiancee:%v in in her preference list. So she sticks with her current fiancee\n", wp.Id, courtier.Id, wp.EngagedTo.Id)
		return false
	} else if courtierRanking.A < currentFianceeRanking.A {
		//fmt.Printf("woman:%v prefers courtier:%v over current financee:%v\n", wp.Id, courtier.Id, wp.EngagedTo.Id)
		return true
	} else {
		//fmt.Printf("woman:%v prefers current fiancee:%v over courtier:%v\n", wp.Id, wp.EngagedTo.Id, courtier.Id)
		return false
	}
}

func Match(freeMen *linked_list.LinkedList[*Man]) []*Woman {
	fmt.Printf("Size of list:%v\n", linked_list.Len(freeMen))
	if linked_list.Len(freeMen) == 0 {
		return []*Woman{}
	}
	allWomen := linked_list.ToArray(linked_list.Head(freeMen).HaveNotProposedTo) //Every man must have every woman in his list of preferences

	for freeMen != nil {
		m := linked_list.Head(freeMen)
		for m.HaveNotProposedTo != nil {
			wp := linked_list.Head(m.HaveNotProposedTo)
			if wp.EngagedTo == nil {
				wp.EngagedTo = m
				m.EngagedTo = wp
				break
			} else {
				//Does this woman prefer me to whom she is currently engaged? If so she
				//breaks her engagement to that guy and you add that guy to free men.
				//Otherwise just try the next woman in the current man's non-proposed-to(preferences) stack.
				if womanPrefersMe(wp, m) {
					oldMan := wp.EngagedTo
					oldMan.EngagedTo = nil
					///Set up current man with this woman
					wp.EngagedTo = m
					m.EngagedTo = wp
					freeMen = linked_list.AddLast(oldMan, freeMen) //TODO Make a doubly-linked list in your library and use this instead for freeMen.
					break
				}
			}
			m.HaveNotProposedTo, _ = linked_list.Tail(m.HaveNotProposedTo)
		} //end woman for
		freeMen, _ = linked_list.Tail(freeMen)
	} // end man for
	return allWomen
}
