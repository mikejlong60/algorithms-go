package chapter4

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type UNode struct {
	Id       string   //The is the label of the DIT object
	Set      *UNode   //This is the parent set of this node, empty if top of DIT and non-empty otherwise
	Children []*UNode //This is empty for a leaf node, non-empty otherwise
}

// Returns a union-find data structure on set S(make sure S is really a set) where all elements are in separate sets.
func MakeUnionFind(S []string) []*UNode {
	var r = make([]*UNode, len(S))
	for i, j := range S {
		r[i] = &UNode{Id: j, Set: nil, Children: []*UNode{}}
	}
	return r
}

// For an element u that is a member of some set S, Find(u) returns the name of the set containing u
func Find(u *UNode) string {
	if u.Set == nil { //You have reached the set that is housing u
		return u.Id
	} else { //Keep looking up until you reach the containing set
		return Find(u.Set)
	}
}

// For two sets A and B Union(A, B) merges the set B into the set A
func Union(A, B *UNode) *UNode {
	if Find(A) != Find(B) { //Make B a member of set A because it is not a member of set A
		B.Set = A
		A.Children = append(A.Children, B)
	}
	return A
}

// Creates a set of RDNs from the whole list of User DNs(i.e. cn=test tester10,ou=people,ou=fred,ou=bigfoot,o=u.s. government,c=us)
func MakeSetOfRDNs(users []string) []string {
	var rdns = []string{}
	for _, j := range users {
		a := strings.Split(j, ",")
		for _, k := range a {
			rdns = append(rdns, k)
		}
	}
	ff := map[string]struct{}{}

	for _, b := range rdns {
		ff[b] = struct{}{}
	}
	r := []string{}
	for aa, _ := range ff {
		r = append(r, aa)
	}
	return r
}

// Makes a DIT from an array of User DNs(i.e. cn=test tester10,ou=people,ou=fred,ou=bigfoot,o=u.s. government,c=us)
// Growth of algorithm is linear, O(n) where n is the number of users.
// Returns a map of UNodes that represent the set of user DNs as a DIT with no duplication.
func ToDirectoryInformationTree(users []string) map[string]*UNode {
	start := time.Now()
	a := MakeSetOfRDNs(users) //Splits up big RDN for each user into a set of strings, meaning no duplicates
	s := MakeUnionFind(a)     //Makes a UNode for every member of set a above.

	var i = make(map[string]*UNode, len(s))
	//Turns set s into a map so you can lookup tokens in O(1)
	for _, k := range s {
		i[k.Id] = k
	}

	for _, j := range users { //Builds the whole DIT by unioning each object in the User DN.
		aa := strings.Split(j, ",")
		for index := range aa {
			if index+1 < len(aa) {
				Union(i[aa[index+1]], i[aa[index]])
			}
		}
	}
	log.Infof("ToDirectoryInformationTree for %v userDNs  took:%v", len(users), time.Since(start))
	return i
}

//Depth-First search
// A recursive algorithm for depth-first search. This algorithm is specialized for searching
// for leaves in a DIT starting at any node in the DIT.
//  It looks up the starting point (u) for every recursive call in O(1) time
// Depth-first search visits every vertex once and checks every edge in the graph once.
// Therefore, DFS complexity is O(V+E) where V is number of vertices(UNodes) and E is number of edges from UNode to UNode.
//TODO -- make the algorithm search for a list of UNodes which may themselves be regular expressions or something like it,
//TODO --- so that a user can put in pieces of a Node(leaf or parent) and produce the list of user DNs that union all those resulting sets.
//Params: same as Returns
//Returns:
//  u - *UNode the current node that gets explored by the algorithm.
//      This variable changes as the algorithm proceeds down the graph toward a leaf in DFS fashion.
//  seen - seen map[string]*UNode - the accumulated map of UNodes that the algorithm has seen thus far
//  userDN - the accumulating list of node Ids for a given path from the u above, to a single leaf.
//     This node can be anywhere in the DIT, a leaf or a parent.
//     This leaf is added to the allUserDNs once it's path has been totally reconstituted.
//  allUserDNs - the accumulating array of userDNs that match the entire DIT beneath the starting UNode u.

func DFSearch(u *UNode, seen map[string]*UNode, userDN string, allUserDNs []string) (*UNode, map[string]*UNode, string, []string) {
	seen[u.Id] = u
	if len(u.Children) == 0 { //you are at a leaf which means the full construction of userDN is now complete.
		allUserDNs = append(allUserDNs, userDN)
	}
	for _, connectedNode := range u.Children {
		_, explored := seen[connectedNode.Id]
		if !explored {
			newUserDN := fmt.Sprintf("%v,%v", connectedNode.Id, userDN)
			_, seen, newUserDN, allUserDNs = DFSearch(connectedNode, seen, newUserDN, allUserDNs)
		}
	}
	return u, seen, userDN, allUserDNs
}

func createPathAboveBaseDN(u *UNode, soFar string) string {
	if u == nil { //Caller called this function with the top node of the DIT
		return ""
	} else if u.Set == nil { //You have reached the top node which has no parent
		return fmt.Sprintf("%v,%v", soFar, u.Id)
	} else { //Keep looking up until you reach the top node
		return createPathAboveBaseDN(u.Set, fmt.Sprintf("%v,%v", soFar, u.Id))
	}
}

// Given a DIT produces the complete list of strings that produced the DIT.
// This is isomorphic with the ToDirectoryInformationTree function above
func FromDirectoryInformationTree(dit map[string]*UNode, baseDN string) []string {
	start := time.Now()
	//Start at the top of the tree
	root := dit[baseDN]

	//Do a Depth-first-search starting there and stop at leaf and add the whole path as the complete DN
	_, _, _, allUserDNs := DFSearch(root, make(map[string]*UNode), baseDN, []string{})
	//append the path above the baseDN to the end of each userDN
	pathAbove := createPathAboveBaseDN(root.Set, "")
	for i := 0; i < len(allUserDNs); i++ {
		allUserDNs[i] = fmt.Sprintf("%v%v", allUserDNs[i], pathAbove)
	}
	log.Infof("FromDirectoryInformationTree took:%v", time.Since(start))
	return allUserDNs
}
