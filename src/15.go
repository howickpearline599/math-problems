def find_factors(n):
    factors = []
    for i in range(1, int(n**0.5) + 1):
        if n % i == 0:
            factors.append(i)
            if n // i != i:
                factors.append(n // i)
    return sorted(factors)

def is_prime(num):
    if num < 2:
        return False
    for i in range(2, int(num**0.5) + 1):
        if num % i == 0:
            return False
    return True

def sum_of_divisors(n):
    total = 0
    for i in range(1, n+1):
        if is_prime(i) and n % i == 0:
            total += i
    return total

# Example usage
n = 28
print("Factors:", find_factors(n))
print("Is Prime:", is_prime(n), "and Sum of Divisors:", sum_of_divisors(n))
