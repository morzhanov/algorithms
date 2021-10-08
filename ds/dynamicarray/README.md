# Dynamic Array

A Dynamic array (vector in C++, ArrayList in Java) automatically grows when we try to make an insertion and there is no more space left for the new item. Usually the area doubles in size.

A simple dynamic array can be constructed by allocating an array of fixed-size, typically larger than the number of elements immediately required. The elements of the dynamic array are stored contiguously at the start of the underlying array, and the remaining positions towards the end of the underlying array are reserved, or unused. Elements can be added at the end of a dynamic array in constant time by using the reserved space until this space is completely consumed.

<img src="https://media.geeksforgeeks.org/wp-content/uploads/dynamicarray.png"/>

<a href="https://www.geeksforgeeks.org/how-do-dynamic-arrays-work/">Source</a>
