# Topological sorting algorithm

Topological sorting for Directed Acyclic Graph (DAG) is a linear ordering of vertices such that for every directed edge u v, vertex u comes before v in the ordering. Topological Sorting for a graph is not possible if the graph is not a DAG.

For example, a topological sorting of the following graph is “5 4 2 3 1 0”. There can be more than one topological sorting for a graph. For example, another topological sorting of the following graph is “4 5 2 3 1 0”. The first vertex in topological sorting is always a vertex with in-degree as 0 (a vertex with no incoming edges).

<img src="https://media.geeksforgeeks.org/wp-content/uploads/20200818211917/Topological-Sorting-1.png"/>

<a href="https://www.geeksforgeeks.org/topological-sorting/">Source</a>
