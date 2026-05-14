# Exercise: Writing and Testing Decorators

A series of challenges that build from the mechanics of decorators to writing and testing them properly. Work through in order.

---

## 1. The wrapper pattern

Write a decorator `timer` that measures and prints how long the wrapped function takes to run. Apply it to a function that sleeps for a short duration.

```python
@timer
def slow_function():
    time.sleep(0.1)
```

**Learn:** the `wrapper(*args, **kwargs)` pattern; why you must call and return the original function; `time.perf_counter()` for timing.

---

## 2. functools.wraps

Without `@functools.wraps`, `slow_function.__name__` and `slow_function.__doc__` will reflect the wrapper, not the original. Add `@functools.wraps(func)` to your wrapper and verify that `__name__` and `__doc__` are preserved.

**Learn:** why `functools.wraps` is always required; what metadata is preserved; what breaks in logging, testing, and introspection without it.

---

## 3. Decorator with arguments

Write a decorator `retry(n)` that calls the wrapped function up to `n` times if it raises an exception, then re-raises on the final failure.

```python
@retry(3)
def flaky():
    ...
```

This requires a factory — a function that returns a decorator. Understand the three levels of nesting: `retry(n)` → decorator → wrapper.

**Learn:** the factory pattern for parameterised decorators; how `@retry(3)` differs from `@retry`.

---

## 4. Class-based decorator

Re-implement `retry` as a class with `__init__` and `__call__`. Both approaches are valid; the class form is often clearer when state is involved.

```python
class retry:
    def __init__(self, n):
        self.n = n

    def __call__(self, func):
        ...
```

**Learn:** `__call__` as the decorator protocol; when to prefer a class over a closure.

---

## 5. Testing decorators

This is where most developers get it wrong. Write tests that verify:

1. The decorated function still **returns the correct value** (your wrapper must `return func(*args, **kwargs)`).
2. `functools.wraps` is in place — assert `fn.__name__` and `fn.__wrapped__` are correct.
3. The decorator behaviour itself — for `retry`, assert the wrapped function is called the right number of times on failure, and exactly once on success. Use `unittest.mock.MagicMock` or `pytest`'s `monkeypatch` to control when the function raises.

```python
from unittest.mock import MagicMock

def test_retry_calls_once_on_success():
    fn = MagicMock(return_value=42)
    decorated = retry(3)(fn)
    result = decorated()
    assert result == 42
    fn.assert_called_once()

def test_retry_retries_on_failure():
    fn = MagicMock(side_effect=[ValueError, ValueError, 42])
    decorated = retry(3)(fn)
    result = decorated()
    assert result == 42
    assert fn.call_count == 3
```

**Learn:** testing decorator behaviour in isolation from the function it wraps; `MagicMock(side_effect=...)` to simulate intermittent failures; asserting call counts.

---

## 6. Stacking decorators

Apply two decorators to the same function and reason through the order of application (bottom-up) versus the order of execution (outer-first). Write a test that confirms which decorator runs first.

```python
@timer
@retry(3)
def operation():
    ...
```

**Learn:** decorator stacking order; how to write a test that detects execution order.

---

## Putting it together

Write a `rate_limit(calls_per_second)` decorator that raises `RuntimeError` if the decorated function is called more times per second than allowed. Test it by controlling time with `monkeypatch` on `time.perf_counter`.

This combines: the factory pattern, state in a closure, and non-trivial testing using time mocking.
