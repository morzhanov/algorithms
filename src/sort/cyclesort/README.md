# CycleSort

Cycle sort is an in-place sorting Algorithm, unstable sorting algorithm, a comparison sort that is theoretically optimal in terms of the total number of writes to the original array.

- It is optimal in terms of number of memory writes. It minimizes the number of memory writes to sort (Each value is either written zero times, if it’s already in its correct position, or written one time to its correct position.)
- It is based on the idea that array to be sorted can be divided into cycles. Cycles can be visualized as a graph. We have n nodes and an edge directed from node i to node j if the element at i-th index must be present at j-th index in the sorted array.  Cycle in arr[] = {2, 4, 5, 1, 3}

```bash
 arr[] = {10, 5, 2, 3}
 index =  0   1   2   3
cycle_start = 0 
item = 10 = arr[0]

Find position where we put the item  
pos = cycle_start
i=pos+1
while(i<n)
if (arr[i] < item)  
    pos++;

We put 10 at arr[3] and change item to 
old value of arr[3].
arr[] = {10, 5, 2, 10} 
item = 3 

Again rotate rest cycle that start with index '0' 
Find position where we put the item = 3 
we swap item with element at arr[1] now 
arr[] = {10, 3, 2, 10} 
item = 5

Again rotate rest cycle that start with index '0' and item = 5 
we swap item with element at arr[2].
arr[] = {10, 3, 5, 10 } 
item = 2

Again rotate rest cycle that start with index '0' and item = 2
arr[] = {2, 3,  5, 10}  

Above is one iteration for cycle_stat = 0.
Repeat above steps for cycle_start = 1, 2, ..n-2
```

<a href="https://www.geeksforgeeks.org/cycle-sort/">Source</a>
