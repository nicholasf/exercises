# Exercise: pytest Fundamentals

A series of progressively deeper challenges covering the parts of pytest that experienced developers often don't discover until much later.

Work through each section in order. Each builds on the last.

---

## 1. Assertions and failure output

Write a test that intentionally fails by asserting two dicts are equal when they differ. Run it and read the output carefully.

pytest rewrites assertion expressions to produce detailed diffs — no `assertEqual`, `assertDictEqual`, or message strings needed. Notice how much information you get for free compared to other test frameworks.

**Learn:** pytest's assertion introspection; why you never need `unittest`-style assertion methods.

---

## 2. Testing exceptions

Write a function that raises `ValueError` when given bad input. Test it with `pytest.raises()` as a context manager, and assert on the exception message via `excinfo.value`.

```python
with pytest.raises(ValueError, match="must be positive"):
    my_function(-1)
```

**Learn:** `pytest.raises()`, the `match` parameter, inspecting exception details.

---

## 3. Fixtures

Write a fixture that returns a pre-populated list, and use it across two different tests. Then write a second fixture that depends on the first one — pytest injects fixtures into other fixtures automatically.

```python
@pytest.fixture
def sample_data():
    return [1, 2, 3]

def test_length(sample_data):
    assert len(sample_data) == 3
```

**Learn:** fixtures as dependency injection; composing fixtures from other fixtures.

---

## 4. Fixture scope and teardown

Write a fixture that prints "setup" and "teardown" around a `yield`. Run the tests with `-v -s` to see the lifecycle. Then change the `scope` parameter to `"module"` and observe when setup/teardown fires relative to the tests.

```python
@pytest.fixture(scope="module")
def expensive_resource():
    print("setup")
    yield some_resource
    print("teardown")
```

Scopes in order of lifetime: `function` (default) → `class` → `module` → `package` → `session`.

**Learn:** `yield` fixtures for teardown; fixture scoping to control how often setup runs.

---

## 5. conftest.py

Move your fixtures into a `conftest.py` file at the directory level. Observe that all test files in that directory can use them without importing anything.

Then create a second `conftest.py` one level up and put a broader fixture there. Understand how pytest resolves fixtures by walking up the directory tree.

**Learn:** `conftest.py` as the standard place for shared fixtures; fixture scoping by directory.

---

## 6. Parametrize

Replace any test that repeats similar assertions with `@pytest.mark.parametrize`. Each parameter set becomes a separate named test case in the output.

```python
@pytest.mark.parametrize("input,expected", [
    (0, 0),
    (1, 1),
    (10, 55),
    (20, 6765),
])
def test_fibonacci(input, expected):
    assert calculate_value(input) == expected
```

**Learn:** data-driven tests; readable output per case; combining with fixtures.

---

## 7. Marks

Mark a test as expected to fail with `@pytest.mark.xfail`, and mark another to skip with `@pytest.mark.skip`. Then create a custom mark (e.g. `@pytest.mark.slow`) and configure it in `pyproject.toml` to avoid the unknown-mark warning.

Run with `-m "not slow"` to exclude marked tests.

```toml
[tool.pytest.ini_options]
markers = ["slow: marks tests as slow to run"]
```

**Learn:** built-in marks; custom marks; filtering test runs with `-m`.

---

## 8. Built-in fixtures worth knowing

Explore these fixtures that pytest provides out of the box — no imports needed:

- `tmp_path` — gives you a temporary `pathlib.Path` directory, unique per test
- `monkeypatch` — safely patch attributes, environment variables, or dict entries; automatically reverts after the test
- `capsys` — capture stdout/stderr and assert on printed output

Write one test using each.

**Learn:** the standard fixtures that replace manual mocking in many common cases.

---

## 9. Putting it together

Write a small module with a class that reads from an environment variable and writes output to a file. Test it using:
- a fixture with `yield` for setup/teardown
- `monkeypatch` to set the environment variable
- `tmp_path` for the output file
- `parametrize` for multiple input cases

This is a representative test you'd write in a real project.
