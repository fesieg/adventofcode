from typing import List

def getTextFileContentAsListOfLines(filePath: str) -> List[str]:
    with open(filePath) as inputFile:
        inputLines = inputFile.readlines()
    return [x.strip() for x in inputLines]


def getTextFileContentAsString(filePath: str) -> str:
    with open(filePath) as inputFile:
        return inputFile.read().replace('\n', ' ').replace('\r', '')