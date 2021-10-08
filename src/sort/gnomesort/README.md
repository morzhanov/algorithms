# GnomeSort

Gnome Sort also called Stupid sort is based on the concept of a Garden Gnome sorting his flower pots. A garden gnome sorts the flower pots by the following method-

- He looks at the flower pot next to him and the previous one; if they are in the right order he steps one pot forward, otherwise he swaps them and steps one pot backwards.
- If there is no previous pot (he is at the starting of the pot line), he steps forwards; if there is no pot next to him (he is at the end of the pot line), he is done.

```bash
If you are at the start of the array then go to the right element (from arr[0] to arr[1]).
If the current array element is larger or equal to the previous array element then go one step right
                   if (arr[i] >= arr[i-1])
                      i++;
If the current array element is smaller than the previous array element then swap these two elements and go one step backwards
                       if (arr[i] < arr[i-1])
                       {
                           swap(arr[i], arr[i-1]);
                           i--;
                       } 
Repeat steps 2) and 3) till ‘i’ reaches the end of the array (i.e- ‘n-1’)
If the end of the array is reached then stop and the array is sorted.

```

<a href="https://www.geeksforgeeks.org/gnome-sort-a-stupid-one/">Source</a>
