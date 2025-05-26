def find_divisors(n):
    """ Find all divisors of n """
    divisors = []
    for i in range(1, n + 1):
        if n % i == 0:
            divisors.append(i)
    return divisors

n = 36 # Example number
print(find_divisors(n))
