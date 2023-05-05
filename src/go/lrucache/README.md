# LRU Cache

In computing, a *cache* is a component that stores data so that future requests for that data can be served faster.

A *cache hit* occurs when the requested data can be found in a cache, while a *cache miss* occurs when it cannot.  The *hit ratio* of a cache describes how often a searched-for item is actually available.

*Cache replacement algorithms* are optimizing instructions that a program can utilize in order to manage a cache. When a cache is full, the algorithm must choose which items to discard to make room for the new ones. 

An *LRU cache* is a cache that uses one of the "Least Recently Used"  family of cache replacement algorithms.  LRU algorithms keep track of what was used when, and discard least recently used items first, using historical metadata in an attempt to maximize the hit ratio for better performance.

Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the `LRUCache` class:

- `LRUCache(int capacity)` Initialize the LRU cache with size `capacity`.
- `int get(int key)` Return the value of the `key` if the key exists, otherwise return `-1`.
- `void put(int key, int value)` Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity, evict the least recently used key.
- The functions get and put must each run in `O(1)` average time complexity.

Example 1:
```
Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4
```
Constraints:
- `1 <= capacity <= 3000`
- `0 <= key <= 104`
- `0 <= value <= 105`
- `At most 2 * 105 calls will be made to get and put.`
