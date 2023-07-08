# Trees

A collection of zero or more nodes in which root, a distinguished node is connected to all subtrees via a directed edge. In the real world trees are like a real tree but upside down.

Traversing a tree means visiting every node in a tree and performing some work on it. Picking the right order in which we visit each node is key in solving a problem efficiently. Main three ways of traversing a tree are shown in figure-1.

```ASCII
[Figure 1] Tree traversals - Full Binary Search Tree Example

     4			Traversal Name		Order 		        Example
    / \
   /   \		In-Order                left,self,right		1,2,3,4,5,6,7
  2     6		Pre-Order,DFS		self,left,right		4,2,1,3,6,5,7
 / \   / \		Post-Order,BFS		left,right,self		1,3,2,5,7,6,4
1   3 5   7
```

The traversal methods of trees are named on when a particular operation is performed on the current node relative to its children. For example, in Pre-Order traversal, the operation is performed on the node prior to traversing its children.

There are many types of trees. Some important tree types include:

* **Binary Tree**: Every node can have no more than two children
* **Complete Binary Tree**: Each level is filled left to right, and all levels except the bottom are full
* **Full Binary Tree**: Every non-leaf node has exactly two children

```ASCII
[Figure 2] Important Tree Concepts

     1		 4	    1        * 	  1 - Is it a binary tree?
    /|\		/ \	   /        / \	    2 - Is it a complete binary tree?
   / | \       /   \	  2        /   \      3 - Is it a full binary tree?
  2  5  6     2     6	 / \	  *     +
 / \   / \     \   /	3   4    / \   / \	A 0 0 0
3   4 7   8  	3 5	    	5   5 2   2	B 1 0 0
						C 1 1 0
    (A)	        (B)	 (C)	    (D)		D 1 1 1
```

## Binary Search Tree (BST)

A Binary Search Tree (BST) is a type of sorted tree where, for every node n, the values of all nodes in its left subtree are less than n and the values of all nodes in its right subtree are greater than n.

Performing an In-Order traversal of a binary search tree and outputting each visited node results in a sorted (In-Order) list of nodes. This is known as the tree sort algorithm, which has a time complexity of O(NLogN). While there are other sorting algorithms available, none are more efficient than O(n*Log n).

### BST Complexity

The time complexity of operations such as Search, Deletion, Insertion, and finding the minimum and maximum values in a binary search tree is O(h), where h represents the height of the tree.

## AVL - Height Balanced BST

A height balanced binary search tree has a height of O(Log n) and its left and right subtrees of all nodes have equal heights.

In order to maintain balance after an insertion, a single rotation is needed if the insertion was on the outer side, either left-left or right-right, while a double rotation is required if the insertion was on the inner side, either left-right or right-left.

### AVL Complexity

Same as a Binary Search Tree except that the height of the tree is known. So Search, Deletion, Insertion, and finding Min and Max in an AVL tree are all O(Log n) operations.

## Trie

A Trie, also known as a prefix tree, is a tree data structure that is commonly used for representing strings. Each descendant of a node in the Trie shares a common prefix, and a path from the root to a leaf node spells out a word held in the Trie.

### Trie Complexity

Insertion and Search are done in O(K), where K is the length of the word.

## Application

Trees, such as Binary Search Trees (BSTs), can offer a time complexity of O(Log n) for searches, as opposed to the linear access time of linked lists. Trees are widely employed in search systems, and operating systems can represent file information using tree structures.

## Rehearsal

### Serialize Binary Tree

Write two functions to serialize and unserialize a binary tree like (B) in _Figure 2_ to and from a string like `4,2,6,nil,3,5,nil`. [Solution](serialize_tree.go), [Test](serialize_tree_test.go)

### Evaluate A Binary Expression Tree

Given an expression binary tree like (D) in _Figure2_ evaluate it to a float64 like 100. [Solution](evaluate_expression.go), [Test](evaluate_expression_test.go)

### Sorted Array to Balanced BST

Given a sorted array of integers like {1,2,3,4,5} return the root to a balanced BST.  [Solution](sorted_array_to_balanced_bsd.go), [Test](sorted_array_to_balanced_bsd_test.go)

### Traverse Binary Tree

Given a binary tree like _figure 1_, write three functions to traverse the tree in-order, pre-order, and post-order. [Solution](traverse_binary_tree.go), [Test](traverse_binary_tree_test.go)

### Implement Autocomplete

Given a word like "car" and a dictionary like {"car","caravan","card","carpet","cap","ca"}, return autocomplete suggestions where the given word is the prefix of a dictionary word like {"avan","d","pet"}. [Solution](auto_complete.go), [Test](auto_complete_test.go)
