def SimpleSearch(array: list[int], element: int) -> tuple[int, int]:
    comparisonCount = 0
    for i in range(len(array)):
        comparisonCount += 1
        if array[i] == element:
            return i, comparisonCount
    return -1, comparisonCount

# Must be sorted in ascending order
def BinarySearch(array: list[int], element: int) -> tuple[int, int]:
    left = 0
    right = len(array) - 1
    comparisonCount = 0

    while left <= right:
        mid = (left + right) // 2
        
        if array[mid] < element:
            comparisonCount += 1
            left = mid + 1
        elif array[mid] > element:
            right = mid - 1
            comparisonCount += 2
        else:
            comparisonCount += 2
            return mid, comparisonCount

    return -1, comparisonCount

