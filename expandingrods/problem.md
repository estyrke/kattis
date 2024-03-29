# Expanding Rods

![/problems/expandingrods/file/statement/en/img-0001.jpg](http://open.kattis.com/problems/expandingrods/file/statement/en/img-0001.jpg)

When a thin rod of length $L$ is heated $n$ degrees, it expands to a new
length $L’ = (1+n \cdot C) \cdot
L$, where $C$ is
the coefficient of heat expansion.

When a thin rod is mounted on two solid walls and then
heated, it expands and takes the shape of a circular segment,
the original rod being the chord of the segment.

Your task is to compute the distance by which the center of
the rod is displaced.

## Input

The input contains at most $20$ lines. Each line of input
contains three non-negative numbers:

- an integer $L$, the
     initial lenth of the rod in millimeters ($1 \le L \le 10^9$),

- an integer $n$, the
     temperature change in degrees ($0 \le n \le 10^5$),

- a real number $C$,
     the coefficient of heat expansion of the material
     ($0 \le C \le 100$, at
     most $5$ digits after
     the decimal point).


The input is such that the displacement of the center of any
rod is at most one half of the original rod length. The last
line of input contains three $-1$’s and it should not be
processed.

## Output

For each line of input, output one line with the
displacement of the center of the rod in millimeters with an
absolute error of at most $10^{-3}$ or a relative error of at
most $10^{-9}$.

Sample Input 1Sample Output 1

```
1000 100 0.0001
15000 10 0.00006
10 0 0.001
-1 -1 -1

```

```
61.328991534
225.020248568
0.000000000

```