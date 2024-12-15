from typing import List

from utils import pyutils as adventutils


def formatContent(content: List[str]) -> List[List[int]]:
    return [list(map(int, line.split(" "))) for line in content]


content = formatContent(adventutils.file_contest_as_list_of_lines("./day2/input"))


def reportIsSafe(report: List[int]) -> bool:
    ascending = all(report[i] < report[i + 1] for i in range(len(report) - 1))
    descending = all(report[i] > report[i + 1] for i in range(len(report) - 1))

    if not (ascending or descending):
        return False

    for i in range(len(report) - 1):
        if abs(report[i] - report[i + 1]) > 3:
            return False

    return True


def part1(reports: List[str]) -> int:
    return [reportIsSafe(report) for report in reports].count(True)


def part2(reports: List[str]) -> int:
    numberOfSafeReports = 0
    for report in reports:
        if not reportIsSafe(report):
            monte_carlo_reports = []
            # RIP memory (TODO: improve this)
            # create one list for each item in the report with that item missing
            for i in range(len(report)):
                new_list = report[:i] + report[i + 1 :]
                monte_carlo_reports.append(new_list)

            for shortened_report in monte_carlo_reports:
                if reportIsSafe(shortened_report):
                    numberOfSafeReports += 1
                    break
        else:
            numberOfSafeReports += 1

    return numberOfSafeReports


print("number of safe reports (part 1): ", part1(content))
print("number of safe reports with correction (part2)", part2(content))
