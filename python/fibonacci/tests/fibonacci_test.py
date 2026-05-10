import pytest
from fibonacci import calculate_value

def test_calculate_value():
    assert calculate_value(0) == 0
    assert calculate_value(1) == 1
    assert calculate_value(10) == 55
    assert calculate_value(20) == 6765