# RedBlackTree

In computer science, a red–black tree is a kind of self-balancing binary search tree. Each node stores an extra bit representing "color" ("red" or "black"), used to ensure that the tree remains balanced during insertions and deletions.[3]

When the tree is modified, the new tree is rearranged and "repainted" to restore the coloring properties that constrain how unbalanced the tree can become in the worst case. The properties are designed such that this rearranging and recoloring can be performed efficiently.

The re-balancing is not perfect, but guarantees searching in {\displaystyle {\mathcal {O}}(\log n)}{\mathcal {O}}(\log n) time, where {\displaystyle n}n is the number of nodes of the tree. The insertion and deletion operations, along with the tree rearrangement and recoloring, are also performed in {\displaystyle {\mathcal {O}}(\log n)}{\mathcal {O}}(\log n) time.[4]

Tracking the color of each node requires only one bit of information per node because there are only two colors. The tree does not contain any other data specific to its being a red–black tree, so its memory footprint is almost identical to that of a classic (uncolored) binary search tree. In many cases, the additional bit of information can be stored at no additional memory cost.

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/Red-black_tree_example.svg/632px-Red-black_tree_example.svg.png"/>

<a href="https://en.wikipedia.org/wiki/Red%E2%80%93black_tree">Source</a>
