from typing import List


def file_contest_as_list_of_lines(filePath: str) -> List[str]:
    with open(filePath) as inputFile:
        inputLines = inputFile.readlines()
    return [x.strip() for x in inputLines]


def file_content_as_string(filePath: str) -> str:
    with open(filePath) as inputFile:
        return inputFile.read().replace("\n", " ").replace("\r", "")
