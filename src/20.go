def square_root(x):
    if x < 0: 
        raise ValueError("Cannot find square root of a negative number")
    else:
        return x ** 0.5

try:
    print(square_root(-4)) # This should throw an exception
except ValueError as e:
    print(e) # Output: Cannot find square root of a negative number
