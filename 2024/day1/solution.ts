import { getFileContent } from "../utils/tsutils";

const content = await getFileContent("input");

const array1: number[] = [];
const array2: number[] = [];

for (const pair of content.split("\r\n").map(line => line.split("   "))) {
  array1.push(Number(pair[0]));
  array2.push(Number(pair[1])); 
}

array1.sort((a, b) => a - b);
array2.sort((a, b) => a - b);

console.log("total distance (part 1): ", partOne(array1, array2));
console.log("similarity score (part 2): ", partTwo(array1, array2));

function partOne(array1: number[], array2: number[]): number {
  let totalDistance = 0;
  for (let i = 0; i < array1.length; i++) {
    totalDistance += Math.abs(array1[i] - array2[i]);
  }

  return totalDistance;
}

function partTwo(array1: number[], array2: number[]): number {
  let similarityScore = 0;
  for (let i = 0; i < array1.length; i++) {
    const countInArrayTwo = array2.filter(num => num === array1[i]).length;
    if (countInArrayTwo > 0) {
      similarityScore += countInArrayTwo * array1[i];
    }
  }

  return similarityScore;
}