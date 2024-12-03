import { getFileContent } from "../utils/tsutils";

const content = await getFileContent("input.txt");

function getReultsOfMulsInValidSections(inputContent: string){
  const onlyValidSections: string[] = [];

  const isOdd = (num: number) => Math.abs(num % 2) === 1;
  
  // our approach here is to split the input string multiple times and just remove all the parts after "don't"
  // TODO: our approach doesnt work unfortunately
  const splitAlongDo = inputContent.split("do()");

  splitAlongDo.forEach(item => {
    if (!item.includes("don't")){
      onlyValidSections.push(item);
    } else {
      const parts = item.split("don't");
      // this is the problem, I assume this only works for one do(), one don't()
      for (const [index, part] of parts.entries()){
        if (!isOdd(index)){
          onlyValidSections.push(part);
        }
      }
    }
  })

  const mulResultsUsingInstructions = getResultsOfMulOperations(onlyValidSections.join(""))
  return mulResultsUsingInstructions
}


function getResultsOfMulOperations(inputContent: string){
  let totalMulResults = 0;
  
  // this pattern catches all the mul(digit, digit) occurences in the input string
  const pattern = /mul\(\d+,\d+\)/g;
  
  const matches = inputContent.match(pattern);
  
  if (matches){
    for (const match of matches) {
      totalMulResults += validateAndCalculate(match);
    }
  }
  
  return totalMulResults
}

console.log("total results of valid muls (part 1): ", getResultsOfMulOperations(content));
console.log("total results of valid muls (part 2): ", getReultsOfMulsInValidSections(content));

function validateAndCalculate(mul: string): number {
  const regex = /\d+/g;

  // get all full numbers from mul string
  const matches = mul.match(regex);

  if (matches?.length === 2) {
    const numOne = parseInt(matches[0]);
    const numTwo = parseInt(matches[1]);
    return numOne * numTwo;
  } else {
    console.log("Error: invalid mul operation -> ", mul);
    return 0
  }
}

