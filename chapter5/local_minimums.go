package chapter5

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Value  int
}

var totalSteps int

func LocalMinimums(guess *Node) []*Node {
	totalSteps = totalSteps + 1
	var r []*Node
	//Cases for returning
	//1.1. You are node above one leaf and left leaf is null.  In this case compare the node with its parent and the non-null leaf and return it if less or null
	//1.2 You are node above one leaf and right leaf is null.  In this case compare the node with its parent and the non-null leaf and return it if less or null
	//2. You are node above two leaves. -- In this case compare the node value with its parent, and right and left children and return it if less or null
	//3. You are leaf node with no children  -- In this case this node is local minimum if < parent. Return it or null
	if guess.Left != nil &&
		guess.Right != nil &&
		guess.Parent != nil &&
		guess.Value < guess.Left.Value && guess.Value < guess.Right.Value && guess.Value < guess.Parent.Value { //This is case 2
		return []*Node{guess}
	} else if guess.Right != nil &&
		guess.Left == nil &&
		guess.Parent != nil &&
		guess.Value < guess.Right.Value && guess.Value < guess.Parent.Value { //This is case 1.1
		return []*Node{guess}
	} else if guess.Left != nil &&
		guess.Right == nil &&
		guess.Parent != nil &&
		guess.Value < guess.Left.Value && guess.Value < guess.Parent.Value { //This is case 1.2
		return []*Node{guess}
	} else if guess.Left == nil &&
		guess.Right == nil &&
		guess.Parent != nil &&
		guess.Value < guess.Parent.Value { //This is case 3
		return []*Node{guess}
	} else if guess.Left == nil &&
		guess.Right == nil &&
		guess.Parent != nil &&
		guess.Value >= guess.Parent.Value { //If you are at a leaf and not minimum return nil
		return nil
	} else { //Divide and Conquer until you get to a node less than all its neighbors
		r = append(r, LocalMinimums(guess.Left)...)
		r = append(r, LocalMinimums(guess.Right)...)
		return r
	}
}
