def is_prime(num):
    if num < 2:
        return False
    for i in range(2, int(num ** 0.5) + 1):
        if num % i == 0:
            return False
    return True

def prime_factors(num):
    factors = []
    i = 2
    while i <= num:
        if num % i == 0 and is_prime(i):
            factors.append(i)
            num //= i
        else:
            i += 1
    return factors

# Example usage
print(prime_factors(36)) # Output: [2, 2, 3, 3]
