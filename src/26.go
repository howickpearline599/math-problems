def gcd(a: int, b: int) -> int:
    while b != 0:
        a, b = b, a % b
    return a

def lcm(a: int, b: int) -> int:
    return abs(a * b) // gcd(a, b)

# Example usage:
a = 15
b = 28
print(f"GCD of {a} and {b} is {gcd(a, b)}")
print(f"LCM of {a} and {b} is {lcm(a, b)}")
