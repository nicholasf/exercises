# Exercise: Producer-Consumer Queue

Implement a producer-consumer pattern using `asyncio.Queue` and multiple concurrent consumers.

## Challenge

1. Write a producer coroutine that puts a series of work items (e.g. integers 1–10) into an `asyncio.Queue`, then signals completion with a sentinel value.
2. Write a consumer coroutine that reads items from the queue, simulates processing time with `await asyncio.sleep()`, and prints the result.
3. Run one producer and three consumers concurrently using `asyncio.gather()`.
4. Observe that consumers pick up work items as they become available rather than waiting for each other.

## What you should learn

- How `asyncio.Queue` coordinates work between coroutines
- How to run multiple coroutines concurrently with `asyncio.gather()`
- How to signal completion to consumers using a sentinel value
- Why this pattern appears constantly in real async Python code
