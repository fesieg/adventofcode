from typing import List

from utils import pyutils as adventutils


def partOne(arr1: List[int], arr2: List[int]) -> int:
    totalDistance = 0

    for i in range(len(arr1)):
        totalDistance += abs(arr1[i] - arr2[i])

    return totalDistance


def partTwo(arr1: List[int], arr2: List[int]) -> int:
    similarityScore = 0

    for i in range(len(arr1)):
        countInArr2 = arr2.count(arr1[i])
        if countInArr2 > 0:
            similarityScore += arr1[i] * countInArr2

    return similarityScore


def getSortedListsFromInput(input: str) -> List[List[int]]:
    arr1, arr2 = [], []

    for line in adventutils.file_contest_as_list_of_lines(input):
        ids = line.split("   ")
        arr1.append(int(ids[0]))
        arr2.append(int(ids[1]))

    arr1.sort()
    arr2.sort()

    return arr1, arr2


arr1, arr2 = getSortedListsFromInput("./day1/input")
print("total distance (part1): ", partOne(arr1, arr2))
print("similarity score (part2): ", partTwo(arr1, arr2))
