def print_random_string():
    """
    This function generates a random string.
    """
    import random
    return ''.join(random.choices('abcdefghijklmnopqrstuvwxyz', k=10))

print(random_string)
