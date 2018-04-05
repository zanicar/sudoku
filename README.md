# sudoku
A Golang terminal based sudoku solver, utilizing Dancing Links (Algorithm X) as described by Donald Knuth.

This was my first proper Golang endeavor to familiarize myself with the language and only some minor changes have been made over the years since it was originally written. The DLX implementation is available here: github.com/zanicar/dlx

The implementation is terminal based and it will find all solutions in the given search space (provided sufficient computing resources are available).

## Usage
    > sudoku 9 .5.2..7.3426.8.....3...52.6.8..40127.........60251..3.2.53...0.....5.4120.1..8.7.

* The first parameter denotes the board size and must be one of the following: 2, 4, 9, 16, 25;  
* The second parameter is the initialization string and it must be of length size^2;
* Each rune / string position denotes a single board cell and is indexed from left to right, starting at the top left corner;
* The values are zero indexed in order to support different variants such as Hexadoku;
* '.' represents an empty cell;

## Output
```
Initial Solution:

   6   | 3     | 8   4
 5 3 7 |   9   |      
   4   |     6 | 3   7
 - - - + - - - + - - -
   9   |   5 1 | 2 3 8
       |       |      
 7 1 3 | 6 2   |   4  
 - - - + - - - + - - -
 3   6 | 4     |   1  
       |   6   | 5 2 3
 1   2 |     9 |   8  

Solution 1:

 2 6 1 | 3 7 5 | 8 9 4
 5 3 7 | 8 9 4 | 1 6 2
 9 4 8 | 2 1 6 | 3 5 7
 - - - + - - - + - - -
 6 9 4 | 7 5 1 | 2 3 8
 8 2 5 | 9 4 3 | 6 7 1
 7 1 3 | 6 2 8 | 9 4 5
 - - - + - - - + - - -
 3 5 6 | 4 8 2 | 7 1 9
 4 8 9 | 1 6 7 | 5 2 3
 1 7 2 | 5 3 9 | 4 8 6

Found 1 solutions in 0.001000 seconds.
```

## Examples
### Single Solutions (9x9)
.5.2..7.3426.8.....3...52.6.8..40127.........60251..3.2.53...0.....5.4120.1..8.7.  
..64....8.0..3..7.8....60..2....41...5..8..2...80....4..23....5.4..0..6.7....18..  
7..........25......6..8.1...4...6.......346.....0...2...0....57..74...0..8....3..  
.4..0...5..6..5.1.7..1.6.482.7...58..0.....3..84...1.747.0.8..1.6.7..8..5...6..7.  
.....71.8.4.63..0.5.8..1.....3.6.52....3.......5...4......0....4.......220.7....1

### Multiple Solutions (9x9)
.7...8632.4...7.0..0.......7....4......7.3......2....5.......6..2.4...7.8613...4.  
0...8...3.1.....2...2...1.....3.0...3...4...5...8.5.....7...6...6.....7.5...0...8

### Hexadoku (16x16) - Single Solutions
....53.....0F.ED..D.CF6.27..5B..87.....456....011..0...9..C.....4.6.0..2...5.....C...9...81...54.....5..B...E1.F.....7...9.6.......F...D.15..0...2......7....6.8.9.41.0F..2E.3..E573....6.....B.B........E6...4AD4A2..7..F....39....2...D.....6.9F...4C.8A.2....

## Performance
Performance test from my laptop (Core i7-4720HQ with 8GB Memory) for some of the examples above:
* 9x9 multiple solutions, example 1: Average 0.001s for 8 solutions;
* 9x9 multiple solutions, example 2: Average 0.045s for 812 solutions;
* 16x16 single solution: Average 0.002s;

Performance tests do not include time taken for transforming, converting or printing of solutions.
