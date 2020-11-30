# 1668. Maximum Repeating Substring

https://leetcode.com/problems/maximum-repeating-substring/

# Description

For a string sequence, a string word is k-repeating if word concatenated k times is a substring of sequence. The word's maximum k-repeating value is the highest value k where word is k-repeating in sequence. If word is not a substring of sequence, word's maximum k-repeating value is 0.

Given strings sequence and word, return the maximum k-repeating value of word in sequence.

Example 1:

```
Input:
sequence = "ababc", word = "ab"
Output:
2
```

Example 2:

```
Input:
sequence = "ababc", word = "ba"
Output:
1
```

Example 3:

```
Input:
sequence = "ababc", word = "ac"
Output:
0
```
