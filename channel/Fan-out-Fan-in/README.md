## Fan-In Fan-Out


### Fan-out
> Multiple functions can read from the same channel until that channel is closed; this is called fan-out. This provides a way to distribute work amongst a group of workers to parallelize CPU use and I/O.


### Fan-in
> A function can read from multiple inputs and proceed until all are closed by multiplexing the input channels onto a single channel that's closed when all the inputs are closed. This is called fan-in.



- `gen()` generates a range of numbers
- `square()` returns the square
- `merge()` merges square functions


- `X` closes channel
- `->` sends over channel


## Diagram

```
/-------------\              /-------------\              /------------\                  /------------\
|             |              |             |      c1      |            |                  |            |
|    Gen()    | -->  [X] 1   |   Square()  | [X]  1   ->  |   merge()  | -> [X] 1 4 16 9  |   main()   |
|             | \            |             |         /    |            |                  |            |
\-------------/  \           \-------------/        /     \------------/                  \------------/
                 / \         /-------------\    c2 / \
                /   \        |             |      /   \
               /     > [X] 3 |   Square()  | [X] 9     \
               \             |             |            \
                \            \-------------/             /
                 \           /-------------\         c3 /
                  \          |             |           /
                   > [X] 4 2 |   Square()  | [X] 16 4 /
                             |             |
                             \-------------/
```
