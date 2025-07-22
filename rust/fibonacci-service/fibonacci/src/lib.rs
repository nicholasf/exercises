use std::env;

pub fn fibonacci(number: u32) -> u32 {
    if number < 2 {
        return number
    }

    return fibonacci(number - 1) + fibonacci(number -2);
}

#[cfg(test)]
mod tests {
    #[test]
    fn various_fibonacci_positions() {
        assert_eq!(crate::fibonacci(1), 1);
        assert_eq!(crate::fibonacci(10), 55);
        assert_eq!(crate::fibonacci(20), 6765);
    }
}
