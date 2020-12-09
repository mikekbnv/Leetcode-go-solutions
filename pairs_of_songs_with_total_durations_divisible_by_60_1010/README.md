# 1010. Pairs of Songs With Total Durations Divisible by 60

https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/

# Description

You are given a list of songs where the ith song has a duration of time[i] seconds.

Return the number of pairs of songs for which their total duration in seconds is divisible by 60. Formally, we want the number of indices i, j such that i < j with (time[i] + time[j]) % 60 == 0.

Example 1:

```
Input:
time = [30,20,150,100,40]
Output:
3
```

Example 2:

```
Input:
time = [60,60,60]
Output:
3
```

