def getTextFileContentAsListOfLines(filePath):
    with open(filePath) as inputFile:
        inputLines = inputFile.readlines()
    return [x.strip() for x in inputLines]


def getTextFileContentAsString(filePath):
    with open(filePath) as inputFile:
        return inputFile.read().replace('\n', ' ').replace('\r', '')