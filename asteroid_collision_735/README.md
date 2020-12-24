# 735. Asteroid Collision

https://leetcode.com/problems/asteroid-collision/

# Description

We are given an array asteroids of integers representing asteroids in a row.

For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.

Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.

Example 1:

```
Input:
asteroids = [5,10,-5]
Output:
[5,10]
```

Example 2:

```
Input:
asteroids = [8,-8]
Output:
[]
```

Example 3:

```
Input:
asteroids = [10,2,-5]
Output:
[10]
```
