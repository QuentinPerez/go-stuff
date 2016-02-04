## PIPELINE

- `gen()` generates a range of numbers
- `square()` returns the square


- `X` closes channel
- `->` sends over channel


## Diagram

```
/-------------\                  /-------------\                   /-----------\
|             |                  |             |                   |           |
|    Gen()    | ->  [X] 4 3 2 1  |   Square()  | -> [X] 16 9 4 1   |   main()  |
|             |                  |             |                   |           |
\-------------/                  \-------------/                   \-----------/
```
