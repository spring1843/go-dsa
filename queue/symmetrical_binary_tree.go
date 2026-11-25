package queue

import (
	"github.com/spring1843/go-dsa/tree"
)

type queue struct {
	data []*tree.BinaryTreeNode
}

// IsTreeSymmetrical solves the problem in O(n) time and O(n) space.
func IsTreeSymmetrical(root *tree.BinaryTreeNode) bool {
	if root == nil {
		return false
	}
	q1 := new(queue)
	q2 := new(queue)

	q1.enqueue(root)
	q2.enqueue(root)

	for q1.len() != 0 {
		tmp1 := q1.dequeue()
		tmp2 := q2.dequeue()

		if !areMirrored(tmp1, tmp2) {
			return false
		}

		if tmp1.Left != nil {
			q1.enqueue(tmp1.Left)
			q2.enqueue(tmp2.Right)
		}

		if tmp1.Right != nil {
			q1.enqueue(tmp1.Right)
			q2.enqueue(tmp2.Left)
		}
	}
	return true
}

// areMirrored returns true if the nodes have the same value and their children mirror each other.
func areMirrored(a, b *tree.BinaryTreeNode) bool {
	return a.Val == b.Val && (a.Left == nil) == (b.Right == nil) && (a.Right == nil) == (b.Left == nil)
}

func (q *queue) len() int                         { return len(q.data) }
func (q *queue) enqueue(val *tree.BinaryTreeNode) { q.data = append(q.data, val) }
func (q *queue) dequeue() *tree.BinaryTreeNode {
	tmp := q.data[0]
	q.data = q.data[1:len(q.data)]
	return tmp
}
