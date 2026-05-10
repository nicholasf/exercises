from functools import cache

@cache
def _fib(n: int) -> int:
    if n < 2:
        return n

    return _fib(n -2) + _fib(n - 1)

def calculate_value(n: int) -> int:
    """
    Calculates the value of the nth integer in the Fibonacci sequence
    """
    return _fib(n)