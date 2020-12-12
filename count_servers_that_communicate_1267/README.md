# 1267. Count Servers that Communicate

https://leetcode.com/problems/count-servers-that-communicate/

# Description

You are given a map of a server center, represented as a m * n integer matrix grid, where 1 means that on that cell there is a server and 0 means that it is no server. Two servers are said to communicate if they are on the same row or on the same column.

Return the number of servers that communicate with any other server.

Example 1:

![alt text](https://assets.leetcode.com/uploads/2019/11/14/untitled-diagram-6.jpg)
```
Input:
grid = [[1,0],[0,1]]
Output:
0
```

Example 2:

![alt text](https://assets.leetcode.com/uploads/2019/11/13/untitled-diagram-4.jpg)
```
Input:
grid = [[1,0],[1,1]]
Output:
3
```

Example 3:

![alt text](https://assets.leetcode.com/uploads/2019/11/14/untitled-diagram-1-3.jpg)
```
Input:
grid = [[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]]
Output:
4
```
