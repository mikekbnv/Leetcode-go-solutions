# 299. Bulls and Cows

https://leetcode.com/problems/bulls-and-cows/

# Description

You are playing the Bulls and Cows game with your friend.

You write down a secret number and ask your friend to guess what the number is. When your friend makes a guess, you provide a hint with the following info:

    °The number of "bulls", which are digits in the guess that are in the correct position.

    °The number of "cows", which are digits in the guess that are in your secret number but are located in the wrong position. Specifically, the non-bull digits in the guess that could be rearranged such that they become bulls.

Given the secret number secret and your friend's guess guess, return the hint for your friend's guess.

The hint should be formatted as "xAyB", where x is the number of bulls and y is the number of cows. Note that both secret and guess may contain duplicate digits.

Example 1:

```
Input:
secret = "1807", guess = "7810"
Output:
"1A3B"
```

Example 2:

```
Input:
secret = "1123", guess = "0111"
Output:
"1A1B"
```

Example 3:

```
Input:
secret = "1", guess = "0"
Output:
"0A0B"
```
