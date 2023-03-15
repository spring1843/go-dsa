# Graph

## Rehearsal

### Iteratively Implement BFS and DFS

Implement BFS and DFS on a graph without using recursion. [Solution](iterative_traversal.go), [Tests](iterative_traversal_test.go)

### Is Graph a DAG

A directed acyclic graph (DAG) is a directed graph that has no cycles. Given a directed graph determine if it's a DAG or not. [Solution](is_dag.go), [Tests](is_dag_test.go)

### Topological Sort

In a DAG the topological order returns elements such that if there's a path from v(i) to v(j), then v(i) appears before v(j) in the ordering. Given a graph like (B) of _Figure 1_ output the graph in topological order like `{1,2,3,4}`. [Solution](topological_sort.go), [Tests](topological_sort_test.go)

### Employee Head Count

Head count of a person in an organization is 1 + the number of all their reports (direct and indirect). Given a list of employee IDs and their direct reports in each line like `1,4,5`, `5,7`, `4`, `7`. Where 1 has two direct reports (4 and 5); 5 has one report (7); 4 and 5 have no reports; Return the head count of a given employee ID. For example head count of 7 is 1, and headcount of 1 is 4. [Solution](employee_headcount.go), [Tests](employee_headcount_test.go)

### Remove Invalid Parentheses

Given a string containing parentheses and other alphabet letters like `(z)())()`, remove the minimum amount of parentheses to make the string valid like `(z())()` and `(z)()()`. [Solution](remove_invalid_parentheses.go), [Tests](remove_invalid_parentheses_test.go)

### Cheapest Flights

Given a list of flights each with a source and destination and a price, a maximum number of stops, source, and destination cities, return the cheapest costs not exceeding the maximum number of stops to reach from source city to destination. [Solution](cheapest_flights.go), [Tests](cheapest_flights_test.go)

### Word Ladder

Given a start word like `pop` and an end word like `car`, a dictionary of same length words like  `{"top","cop","cap","car"}` return the minimum number of transformations like 4 to get from start to end where each transformation between two words can happen when they are different by only one letter. [Solution](word_ladder.go), [Tests](word_ladder_test.go)

### Network Delay Time

Given `n`, the number of nodes in a network, travel times a list of `{source, destination, delay}` format, and `k` and a node number, return the time it will take for a message sent from k to be received by all nodes. [Solution](network_delay_time.go), [Tests](network_delay_time_test.go)

### Number of Islands

Given a grid in which 0 represents water and 1 represents a piece of land, return the number of islands. An Island is one or more piece of land that could be connected to other pieces of land on left, right, top or bottom. [Solution](number_of_islands.go), [Tests](number_of_islands_test.go)
