# Data Structures and Algorithms in Go 🚀

[![tests](https://github.com/spring1843/go-dsa/actions/workflows/tests.yaml/badge.svg)](https://github.com/spring1843/go-dsa/actions/workflows/tests.yaml)
[![GitHub License](https://img.shields.io/badge/License-Apache%202.0-ff69b4.svg)](https://github.com/aws/karpenter/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/spring1843/go-dsa)](https://goreportcard.com/report/github.com/spring1843/go-dsa)
[![Coverage Report](https://coveralls.io/repos/github/spring1843/go-dsa/badge.svg?branch=main)](https://coveralls.io/github/spring1843/go-dsa?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/spring1843/go-dsa.svg)](https://pkg.go.dev/github.com/spring1843/go-dsa)

Welcome to **Data Structures and Algorithms in Go**! 🎉 This project is designed to serve as a dynamic, hands-on resource for learning and practicing data structures and algorithms in the Go programming language.

* Completely free, community-driven, and continuously evolving
* Executes and comes with 100% test coverage, ensuring a high level of quality
* Ability to use in your favorite IDE, editor, or web browser
* Pure Go, no third-party libraries, Guaranteed forward compatibility

## 📚 Table of Contents

* [Complexity Analysis](complexity.md)
* Data Structures
    * [Arrays](./array/README.md)
        * [Reverse Array In-place](./array/reverse_inplace_test.go)
        * [Add Two Numbers](./array/add_two_numbers_test.go)
        * [Find Duplicate in Array](./array/find_duplicate_in_array_test.go)
        * [Zero Sum Triplets](./array/zero_sum_triplets_test.go)
        * [Product of All Other Elements](./array/product_of_all_other_elements_test.go)
        * [Equal Sum Sub-arrays](./array/equal_sum_subarrays_test.go)
        * [Rotate K Times](./array/rotate_k_steps_test.go)
        * [Bubble Sort](./array/bubble_sort_test.go)
        * [Insertion Sort](./array/insertion_sort_test.go)
    * [Strings](./strings/README.md)
        * [The longest Dictionary Word Containing Key](./strings/longest_dictionary_word_test.go)
        * [Look and Tell](./strings/look_and_tell_test.go)
        * [In Memory Database](./strings/in_memory_database_test.go)
        * [Number in English](./strings/number_in_english_test.go)
        * [Reverse Vowels In a String](./strings/reverse_vowels_test.go)
        * [Longest Substring of Two Unique Characters](./strings/longest_substring_test.go)
    * [Linked Lists](./linkedlist/README.md)
        * [Linked List Serialization](./linkedlist/serialization_test.go)
        * [Reverse a Linked List In-place](./linkedlist/reverse_in_place_test.go)
        * [Join Two Sorted Linked Lists](./linkedlist/join_sorted_lists_test.go)
        * [Keep Repetitions](./linkedlist/keep_repetitions_test.go)
        * [Copy Linked List with Random Pointer](./linkedlist/copy_linklist_with_random_pointer_test.go)
        * [Implement LRU Cache](./linkedlist/lru_cache_test.go)
    * [Stacks](./stack/README.md)
        * [Max Stack](./stack/max_stack_test.go)
        * [Balancing Symbols](./stack/balancing_symbols_test.go)
        * [Infix to Postfix Conversion](./stack/infix_to_postfix_test.go)
        * [Evaluate Postfix](./stack/evaluate_postfix_test.go)
        * [Basic Calculator](./stack/basic_calculator_test.go)
        * [Longest Valid Parentheses](./stack/longest_valid_parentheses_test.go)
    * [Queues](./queue/README.md)
        * [A Queue Using Stacks](./queue/queue_using_stacks_test.go)
        * [Implement a Queue Using a Circular Array](./queue/queue_using_circular_array_test.go)
        * [Is Binary Tree Symmetrical](./queue/is_tree_symmetrical_test.go)
        * [Generate Binary Numbers](./queue/generate_binary_numbers_test.go)
        * [Find The Maximum Sub-array of Length K](./queue/maximum_of_sub_arrays_test.go)
    * [Hash Tables](./hashtable/README.md)
        * [Find Missing Number](./hashtable/missing_number_test.go)
        * [List Elements Summing Up to K](./hashtable/sum_up_to_k_test.go)
        * [Fastest Way to Cut a Brick Wall](./hashtable/find_anagrams_test.go)
        * [Find Anagrams](./hashtable/find_anagrams_test.go)
        * [Find Max Points on the Same Line](./hashtable/max_points_on_line_test.go)
    * [Trees](./tree/README.md)
        * [Serialize Binary Tree](./tree/serialize_tree_test.go)
        * [Evaluate A Binary Expression Tree](./tree/evaluate_expression_test.go)
        * [Sorted Array to Balanced BST](./tree/sorted_array_to_balanced_bsd_test.go)
        * [Traverse Binary Tree](./tree/traverse_binary_tree_test.go)
        * [Reverse Binary Tree](./tree/reverse_binary_tree_test.go)
        * [Implement Autocomplete](./tree/auto_complete_test.go)
    * [Heaps](./heap/README.md)
        * [Kth Largest Element](./heap/kth_largest_element_test.go)
        * [Merge Sorted Lists](./heap/merge_sorted_list_test.go)
        * [Median in a Stream](./heap/median_in_a_stream_test.go)
        * [Kth Closest Points to the Center](./heap/k_closest_points_to_origin_test.go)
        * [Sliding Maximum](./heap/sliding_maximum_test.go)
        * [Heap Sort](./heap/heap_sort_test.go)
* Algorithms
    * [Recursion](./recursion/README.md)
        * [Reverse an integer recursively](./recursion/reverse_number_test.go)
        * [Palindrome](./recursion/is_palindrome_test.go)
        * [Climbing Stairs](./recursion/climbing_stairs_test.go)
        * [Exponentiation](./recursion/exponentiation_test.go)
        * [Regular Expressions Matching](./recursion/)
    * [Divide and Conquer](./dnc//README.md)
        * [Binary Search](./dnc/binary_search_test.go)
        * [Square Root with Binary Search](./dnc/square_root_test.go)
        * [Rate Limit](./dnc/rate_limit_test.go)
        * [Towers of Hanoi](./dnc/towers_of_hanoi_test.go)
        * [Merge Sort](./dnc/merge_sort_test.go)
        * [Quick Sort](./dnc/quick_sort_test.go)
    * [Bit Manipulation](./bit//README.md)
        * [Division without multiplication or division operators](./bit/division_without_operators_test.go)
        * [Middle without division](./bit/middle_without_division_test.go)
        * [Addition without using plus (+) or any other arithmetic operators](./bit/addition_without_operators_test.go)
        * [Maximum without if conditions](./bit/max_function_without_conditions_test.go)
        * [Oddly Repeated Number](./bit/oddly_repeated_number_test.go)
    * [Backtracking](./backtracking//README.md)
        * [Permutations](./backtracking/permutations_test.go)
        * [Generate Parentheses](./backtracking/generate_parentheses_test.go)
        * [Phone Letter Combinations](./backtracking/phone_letter_combinations_test.go)
        * [Sudoku](./backtracking/sudoku_test.go)
        * [N Queens](./backtracking/n_queens_test.go)
    * [Graphs](./graph/README.md)
        * [Iteratively Implement BFS and DFS](./graph/iterative_traversal_test.go)
        * [Is Graph a DAG](./graph/is_dag_test.go)
        * [Topological Sort](./graph/topological_sort_test.go)
        * [Employee Head Count](./graph/employee_headcount_test.go)
        * [Remove Invalid Parentheses](./graph/remove_invalid_parentheses.go)
        * [Cheapest Flights](./graph/cheapest_flights_test.go)
        * [Word Ladder](./graph/word_ladder_test.go)
        * [Network Delay Time](./graph/network_delay_time_test.go)
        * [Number of Islands](./graph/number_of_islands_test.go)
    * [Greedy Algorithms](./greedy/README.md)
        * [Maximum Stock Profit](./greedy/max_stock_profit_test.go)
        * [Activity Selector](./greedy/activity_selector_test.go)
        * [Knapsack](./greedy/knapsack_test.go)
        * [Jump Game](./greedy/jump_game_test.go)
        * [Task Scheduling](./greedy/task_scheduling_test.go)
    * [Dynamic Programming](./dp/README.md)
        * [Rod Cutting](./dp/rod_cutting_test.go)
        * [Sum Up to Number](./dp/sum_up_to_integer_test.go)
        * [House Robber](./dp/house_robber_test.go)
        * [Minimum Deletion to Make a Palindrome](./dp/minimum_deletion_to_make_palindrome_test.go)
        * [Word Distance](./dp/word_distance_test.go)

## 🛠️ How to use

* Read the README.md file of each subject.
* Solve the rehearsal problems.
* Read the problem description.
* Copy the contents of each `*_test.go` file into [Go Playground](https://go.dev/play/).
* Change the first line from `package SOME_PACKAGE_NAME` to `package main`.
* Look at the sample input in the `*_test.go` file.
* Create a solution function by looking at the function signature in the test file.
* Implement your own solution.
* Review the provided solution and compare it to your own.

## 📋 Outline

All topics are discussed in README.md files in the corresponding directory. Each topic includes the following sections:

* 💡 **Implementation**: Detailed explanation of how the data structure or algorithm can be implemented, including code examples in Go.
* 📊 **Complexity**: Analysis of the time and space complexity of the data structure or algorithm.
* 🎯 **Application**: Discussion of problems that are commonly solved using the data structure or algorithm.
* 📝 **Rehearsal**: Practice problems with links to tests that provide 100% coverage and example inputs and outputs.
