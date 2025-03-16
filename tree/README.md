# Trees

A collection of zero or more nodes in which root, a distinguished node, is connected to all subtrees via a directed edge. In the real world, trees are like real trees but upside down.

Traversing a tree means exploring every node in a tree and performing some work on it. Picking the right order in which nodes are explored is key to efficiently solving a problem. The three main ways of traversing a tree are shown in Figure-1.

```ASCII
[Figure 1] Tree traversals - Full Binary Search Tree Example

     4			Traversal Name		Order 		        Example
    / \
   /   \		In-Order                left,self,right		1,2,3,4,5,6,7
  2     6		Pre-Order 		self,left,right		4,2,1,3,6,5,7
 / \   / \		Post-Order		left,right,self		1,3,2,5,7,6,4
1   3 5   7
```

Tree traversal methods are named after the order in which the nodes are explored. For example, in Pre-Order traversal, the node is explored before its children; in Post-Order traversal, the node is explored after its children. Children are always traversed from left to right.

There are many types of trees. Some important tree types include:

* **Binary Tree**: Every node can have no more than two children
* **Complete Binary Tree**: Each level is filled left to right, and all levels except the bottom are full
* **Full Binary Tree**: Every non-leaf node has exactly two children

```ASCII
[Figure 2] Important Tree Concepts

     1		 4	    1        * 	  1 - Is it a binary tree?
    /|\		/ \	   /        / \	    2 - Is it a complete binary tree?
   / | \       /   \	  2        /   \      3 - Is it a full binary tree?
  2  5  6     2     6	 / \	  *     +         1 2 3
 / \   / \     \   /	3   4    / \   / \	A 0 0 0
3   4 7   8  	3 5	    	5   5 2   2	B 1 0 0
						C 1 1 0
    (A)	        (B)	 (C)	    (D)		D 1 1 1
```

## Binary Search Tree (BST)

A Binary Search Tree (BST) is a sorted tree where, for every node n, the values of all nodes in its left subtree are less than n, and the values of all nodes in its right subtree are greater than n.

Performing an In-Order traversal of a binary search tree and outputting each explored node results in a sorted (In-Order) list of nodes. This is known as the tree sorting algorithm, which has a time complexity of O(NLogN).

### BST Complexity

The time complexity of operations such as Search, Deletion, Insertion, and finding the minimum and maximum values in a binary search tree is O(h), where h represents the tree height.

## AVL - Height Balanced BST

An AVL tree is a height-balanced binary search tree has a height of O(Log n), and its left and right subtrees of all nodes have equal heights. AVL stands for the name of the people who came up with the idea.

To maintain balance after an insertion, a single rotation is needed if the insertion is on the outer side, either left-left or right-right, while a double rotation is required if the insertion is on the inner side, either left-right or right-left.

### AVL Complexity

Same as a Binary Search Tree except that the tree height is known. So Search, Deletion, Insertion, and finding Min and Max in an AVL tree are all O(Log n) operations.

## Trie

A Trie, also known as a prefix tree, is a tree data structure commonly used for representing strings. Each descendant of a node in the Trie shares a common prefix, and a path from the root to a leaf node spells out a word held in the Trie.

### Trie Complexity

Insertion and Search are done in O(K), where K is the length word length.

## Application

Trees, such as Binary Search Trees (BSTs), offer O(Log n) time complexity for searches, which is superior to [linked lists](../linkedlist/)' and [array](../array/)'s linear access time. Trees are widely employed in search systems. Operating systems can represent file information using tree structures.

## Rehearsal

* [Serialize Binary Tree](./serialize_tree_test.go), [Solution](./serialize_tree.go)
* [Evaluate A Binary Expression Tree](./evaluate_expression_test.go), [Solution](./evaluate_expression.go)
* [Sorted Array to Balanced BST](./sorted_array_to_balanced_bsd_test.go), [Solution](./sorted_array_to_balanced_bsd.go)
* [Traverse Binary Tree](./traverse_binary_tree_test.go), [Solution](./traverse_binary_tree.go)
* [Reverse Binary Tree](./reverse_binary_tree_test.go), [Solution](./reverse_binary_tree.go)
* [Autocompletion](./autocompletion_test.go), [Solution](./autocompletion.go)
