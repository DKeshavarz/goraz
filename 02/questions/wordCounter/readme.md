# Concurrent Word Frequency Counter

## Problem

Read a large text file and find the top 10 most frequent words using concurrency.

Input: Text file path
Output: Top 10 words with their frequencies


## Architecture

```
project/
├── main.go              # YOU WRITE - Main program
├── wordcounter/
│   └── counter.go       # YOU WRITE - Concurrent implementation
├── test.txt             # Sample text file
└── go.mod
```


## Task

Design a concurrent solution that:
1. Reads a file line by line (or in chunks)
2. Processes lines concurrently to count word frequencies
3. Merges results from all workers
4. Returns top 10 most frequent words

## Implementation Strategies
You can choose any Strategy.
### Option A: Line-based Processing
```
File → Read Lines → Channel → Workers → Merge → Top 10
```

### Option B: Chunk-based Processing
```
File → Chunk 1 → Worker 1
File → Chunk 2 → Worker 2  → Merge → Top 10
File → Chunk 3 → Worker 3
```

### Option C: Pipeline
```
Reader → Splitter → Workers (fan-out) → Merger (fan-in) → Top 10
```

## Example

**Output:**
```
Top 10 most frequent words:
1. go (3)
2. hello (2)
3. world (2)
4. is (1)
5. awesome (1)
6. of (1)
```

