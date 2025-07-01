package chapter5

import (
	"fmt"
	"github.com/greymatter-io/golangz/stack"
	"github.com/hashicorp/go-multierror"
)

func MergeSortWithInversionChecking[T any](xs []T, inversions int, isInversion func(l, r T) bool, lt func(l, r T) bool) ([]T, int) {
	merge := func(lxs, rxs []T, inversions int, lt func(l, r T) bool) ([]T, int) {
		//TODO Make ToList in Golangz iterate backwards for efficiency
		//TODO Make a stack in Golangz with push pop and peek.
		var rs = stack.FromArray(lxs)
		var ls = stack.FromArray(rxs)
		var r, l T
		result := make([]T, len(lxs)+len(rxs))
		var lok bool
		var rok bool
		for i := range result {
			l, lok = stack.Peek(ls)
			r, rok = stack.Peek(rs)
			if lok && rok {
				if lt(r, l) {
					rs = stack.Pop(rs)
					result[i] = r
					if isInversion(l, r) {
						inversions = inversions + 1 //If right is less than left that's an inversion
					}
				} else {
					ls = stack.Pop(ls)
					result[i] = l
				}
			} else if lok {
				ls = stack.Pop(ls)
				result[i] = l
			} else if rok {
				rs = stack.Pop(rs)
				result[i] = r
			}
		}
		return result, inversions
	}

	a := len(xs)
	if a == 0 {
		return xs, inversions
	} else if a == 1 {
		return xs, inversions
	} else if a == 2 {
		if lt(xs[0], xs[1]) {
			return xs, inversions
		} else {
			if isInversion(xs[0], xs[1]) {
				inversions = inversions + 1 //right less than left is an inversion
			}
			return []T{xs[1], xs[0]}, inversions
		}
	} else {
		i := a / 2
		var left = xs[0:i]
		var right = xs[i:]
		left, inversions = MergeSortWithInversionChecking(left, inversions, isInversion, lt)
		right, inversions = MergeSortWithInversionChecking(right, inversions, isInversion, lt)
		merged, inversions := merge(left, right, inversions, lt)
		return merged, inversions
	}
}

func NumberOfEquivalences[T any](xs []T, eq func(l, r T) bool, lt func(l, r T) bool) bool {
	ys := MergeSort(xs, lt)
	le := len(ys)
	var previous T
	var duplicates int
	for _, x := range ys {
		if eq(previous, x) {
			duplicates = duplicates + 1
		}
		previous = x
		if duplicates >= le/2 {
			return true
		}
	}
	return duplicates >= le/2

}

func MergeSort[T any](xs []T, lt func(l, r T) bool) []T {
	merge := func(lxs, rxs []T, lt func(l, r T) bool) []T {
		var rs = stack.FromArray(lxs)
		var ls = stack.FromArray(rxs)
		var r, l T
		result := make([]T, len(lxs)+len(rxs))
		var lok bool
		var rok bool
		for i := range result {
			l, lok = stack.Peek[T](ls)
			r, rok = stack.Peek(rs)
			if lok && rok {
				if lt(r, l) {
					rs = stack.Pop(rs)
					result[i] = r
				} else {
					ls = stack.Pop(ls)
					result[i] = l
				}
			} else if lok {
				ls = stack.Pop(ls)
				result[i] = l
			} else if rok {
				rs = stack.Pop(rs)
				result[i] = r
			}
		}
		return result
	}

	a := len(xs)
	if a == 0 {
		return xs
	} else if a == 1 {
		return xs
	} else if a == 2 {
		if lt(xs[0], xs[1]) {
			return xs
		} else {
			return []T{xs[1], xs[0]}
		}
	} else {
		i := a / 2
		left := xs[0:i]
		right := xs[i:]
		c := make(chan []T, 2)
		asyncMerge(left, lt, c)
		asyncMerge(right, lt, c)
		results := make([]T, 2)
		results[0], results[1] = <-c, <-c
		merged := merge(results[0], results[1], lt)
		return merged
	}
}

func asyncMerge[T any](xs []T, lt func(l, r T) bool, c chan []T) {
	go func() {
		merged := MergeSort(xs, lt)
		c <- merged
	}()
}

//func AggregateDias(client *soap.Client, userDN string) (UserAuthorizationObject, error) {
//	c := make(chan interface{}, 7)
//	asyncMerge(, c)
//	async(client, userDN, getAllUserGroups, c)
//	async(client, userDN, getSecurityAttributes, c)
//	async(client, userDN, getUserGIMMEEAttributesAndCommunityType, c)
//	async(client, userDN, getUserAdminOrganization, c)
//	async(client, userDN, getDutyOrganization, c)
//	async(client, userDN, getUserGIMMEEOrganizationalPath, c)
//	diasResults := make([]interface{}, 7)
//	diasResults[0], diasResults[1], diasResults[2], diasResults[3], diasResults[4], diasResults[5], diasResults[6] = <-c, <-c, <-c, <-c, <-c, <-c, <-c
//
//	var errors error
//
//	var user = UserAuthorizationObject{UserDN: userDN}
//	for _, diasValue := range diasResults {
//		switch v := diasValue.(type) {
//		case WhitePageAttributesResponse:
//			user.WhitePage = v
//		case AllUserGroupsResponse:
//			user.AllUserGroups = v
//		case ClearanceInfo:
//			user.ClearanceInfo = v
//		case GIMMEEAttributesAndCommunityTypeResponse:
//			user.GIMMEEAttributesAndCommunityType = v
//		case UserAdminOrganizationResponse:
//			user.AdminOrganization = v
//		case DutyOrganizationResponse:
//			user.DutyOrganization = v
//		case GIMMEEOrganizationPathResponse:
//			user.GIMMEEOrganizationPath = v
//		case error:
//			errors = multierror.Append(errors, v)
//		default:
//			errors = multierror.Append(errors, fmt.Errorf("Unknown type retured by DIAS web service: %T", v))
//		}
//	}
//	if errors != nil {
//		return UserAuthorizationObject{}, errors
//	}
//	return user, nil
//}
