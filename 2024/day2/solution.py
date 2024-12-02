from typing import List

from utils import pyutils as adventutils


def formatContent(content: List[str]) -> List[List[int]]:
    return [list(map(int, line.split(" "))) for line in content]


content = formatContent(
    adventutils.getTextFileContentAsListOfLines("./day2/input")
)


def reportIsSafe(report: List[int]) -> bool:
    ascending = all(report[i] < report[i + 1] for i in range(len(report) - 1))
    descending = all(report[i] > report[i + 1] for i in range(len(report) - 1))

    if not (ascending or descending):
        return False

    for i in range(len(report) - 1):
        if abs(report[i] - report[i + 1]) > 3 or report[i] == report[i + 1] <= 0:
            return False

    return True


def part1(reports: List[str]) -> int:
    numberOfSafeReports = 0
    for report in reports:
        numberOfSafeReports += 1 if reportIsSafe(report) else 0

    return numberOfSafeReports


print("number of safe reports (part 1): ", part1(content))
