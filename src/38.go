import numpy as np

def binary_search(arr, target):
    low = 0
    high = len(arr) - 1
    
    while low <= high:
        mid = (low + high) // 2
        if arr[mid] == target:
            return True
        elif arr[mid] < target:
            low = mid + 1
        else:
            high = mid - 1
            
    return False

# Example usage
data = [1, 3, 5, 7]
target = 3
print(binary_search(data, target))  # Output: True
