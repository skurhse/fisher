# Reverse Linked List

In computer science, a linked list is a linear collection of data elements whose order, unlike an array, is not inherent to their placement in memory. Instead, each element points to the next, and together this collection of nodes represents a sequence.

In a *singly* linked list, each node contains a data value and a 'next' reference. This structure allows for removal of (or insertions around) accessed elements in constant or O(1) time. A drawback of linked lists however, is that raw access requires iteration which is done in linear or O(N) time.

In general, more memory is required to store a linked list than is required for an array.

**Challenge**

Consider the following class definition, exemplified here in python:
```python
# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
```
Given the head of a singly linked list, return the reversed list.

Example One:
```
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
```
Example Two:
```
Input: head = [1]
Output: [1]
```
Example Three:
```
Input: head = []
Output: []
```

Both recursive and iterative solutions exist. Which one is faster?
