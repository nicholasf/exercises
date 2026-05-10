# Exercise: Concurrent HTTP Requests

Fetch a list of URLs simultaneously using `asyncio` and `aiohttp`, then compare the performance against doing the same requests sequentially.

## Challenge

1. Install `aiohttp` as a dependency.
2. Write a sequential version that fetches each URL one at a time using `requests`, recording the total elapsed time.
3. Write an async version that fetches all URLs concurrently using `aiohttp` and `asyncio.gather()`, recording the total elapsed time.
4. Print each response's URL, status code, and elapsed time as it completes.
5. Print the total elapsed time for both approaches and observe the difference.

## Suggested URLs

Use a mix of public APIs that respond at different speeds, for example:
- `https://httpbin.org/delay/1`
- `https://httpbin.org/delay/2`
- `https://httpbin.org/delay/3`

## What you should learn

- How to write and call a `async def` coroutine
- How `asyncio.gather()` runs coroutines concurrently
- Why async I/O is faster than sequential I/O for network-bound work
- The difference between `requests` (blocking) and `aiohttp` (non-blocking)
