package leetcode

type Node struct {
	Val  int
	Next *Node
}

func buildNodeList(arr []int) *Node {
	node := &Node{}
	c := node
	i := 0
	for i = 0; i < len(arr)-1; i++ {
		c.Val = arr[i]
		c.Next = &Node{}
		c = c.Next
	}
	c.Val = arr[i]
	return node
}

func isnodeListEqual(l1 *Node, l2 *Node) bool {
	if l1 == nil && l2 == nil {
		return true
	}

	c1 := l1
	c2 := l2
	for c1 != nil || c2 != nil {
		if (c1 == nil && c2 != nil) || (c2 == nil && c1 != nil) {
			return false
		}

		if c1.Val != c2.Val {
			return false
		}

		c1 = c1.Next
		c2 = c2.Next
	}

	return true
}
func mergeTwoSortList(l1 *Node, l2 *Node) *Node {
	if l1 == nil && l2 == nil {
		return nil
	}

	l := new(Node)

	c1 := l1
	c2 := l2
	c := l

	for c1 != nil || c2 != nil {
		if c1 == nil {
			c.Val = c2.Val
			c2 = c2.Next
		} else if c2 == nil {
			c.Val = c1.Val
			c1 = c1.Next
		} else {
			if c1.Val < c2.Val {
				c.Val = c1.Val
				c1 = c1.Next
			} else {
				c.Val = c2.Val
				c2 = c2.Next
			}
		}
		if !(c1 == nil && c2 == nil) {
			c.Next = &Node{}
			c = c.Next
		}
	}
	return l
}
