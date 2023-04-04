package queue

import (
	"github.com/spring1843/go-dsa/tree"
)

type queue struct {
	data []*tree.BinaryTreeNode
}

// IsTreeSymmetrical returns true if the given binary tree is symmetric and false otherwise.
func IsTreeSymmetrical(root *tree.BinaryTreeNode) (bool, error) {
	if root == nil {
		return false, nil
	}
	q1 := new(queue)
	q2 := new(queue)

	q1.enqueue(root)
	q2.enqueue(root)

	for q1.len() != 0 {
		tmp1 := q1.dequeue()
		tmp2 := q2.dequeue()

		unsymmetrical := tmp1.Val != tmp2.Val || tmp1.Left != nil && tmp2.Right == nil || tmp1.Left == nil && tmp2.Right != nil || tmp1.Right != nil && tmp2.Left == nil || tmp1.Right == nil && tmp2.Left != nil
		if unsymmetrical {
			return false, nil
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
	return true, nil
}

func (q *queue) len() int                         { return len(q.data) }
func (q *queue) enqueue(val *tree.BinaryTreeNode) { q.data = append(q.data, val) }
func (q *queue) dequeue() *tree.BinaryTreeNode {
	tmp := q.data[0]
	q.data = q.data[1:len(q.data)]
	return tmp
}
