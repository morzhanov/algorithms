# PigeonSort

Pigeonhole sorting is a sorting algorithm that is suitable for sorting lists of elements where the number of elements and the number of possible key values are approximately the same.
It requires O(n + Range) time where n is number of elements in input array and ‘Range’ is number of possible values in array.

```bash
Find minimum and maximum values in array. Let the minimum and maximum values be ‘min’ and ‘max’ respectively. Also find range as ‘max-min+1’. 
Set up an array of initially empty “pigeonholes” the same size as of the range.
Visit each element of the array and then put each element in its pigeonhole. An element arr[i] is put in hole at index arr[i] – min.
Start the loop all over the pigeonhole array in order and put the elements from non- empty holes back into the original array.
```

<a href="https://www.geeksforgeeks.org/pigeonhole-sort/">Source</a>
