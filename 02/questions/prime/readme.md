# Concurrent Prime Finder

## Problem

Find all prime numbers up to N using concurrency.

Input: N (integer > 1)
Output: All prime numbers from 1 to N

## Task

Design a concurrent solution using goroutines and channels.

Choose your strategy:
- Split range into chunks
- Worker pool with queue
- Pipeline (fan-out/fan-in)
- Your own design

## Requirements

- User inputs N
- Find all primes up to N
- Return as []int
- Must use concurrency
- Print results

## Example Output

```
Enter a number: 100
Found 25 primes up to 100
Primes: [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97]
```
