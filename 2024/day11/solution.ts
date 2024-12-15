import { getFileContent } from "../utils/tsutils";

const stones = (await getFileContent("input.txt")).split(" ").map(Number);

function getChangedStones(stone: string): number[] | BigInt[]{
  if (stone === "0") {
    return [1];
  }

  const length = stone.length;

  if (length % 2 === 0) {
    const mid = length / 2;
    const firstHalf = parseInt(leadingZeroesRemoved(stone.substring(0, mid)));
    const secondHalf = parseInt(leadingZeroesRemoved(stone.substring(mid)));
    return [firstHalf, secondHalf];
  }

  return [parseInt(stone) * 2024];
}


function blinkMapNTimes(inputStones: number[], blinks: number): Map<number, number> {
  let stoneMap = new Map();
  
  for (const stone of inputStones) {
    // initialize map with each input stones as key, and value as amount of that stone present
    stoneMap.set(stone, inputStones.filter(s => s === stone).length);
  }

  while (blinks > 0) {
    // TODO: people have told me a more efficient way would be recursion, so look into that
    // for each blink, go through all keys in the map
    // get the changed stones for each key, and increment value in the map for each changed stone
    // we use a new map to store the new values, so we don't mess up the current map
    const newMap = new Map();
    for (const [stone, count] of stoneMap) {
      let changedStones = getChangedStones(stone.toString());
      for (const changedStone of changedStones) {
        let newCount = newMap.get(changedStone) || 0;
        newMap.set(changedStone, newCount + count);
      }
    }

    stoneMap = newMap;

    blinks--;
  }

  return stoneMap;
}

console.log("PART 1: ", sumArray([...blinkMapNTimes(stones, 25).values()]));
console.log("PART 2: ", sumArray([...blinkMapNTimes(stones, 75).values()]));

function leadingZeroesRemoved(input: string): string {
  return input.replace(/^0+(?!$)/, '');
}

function sumArray(array: number[]): number {
  return array.reduce((acc, curr) => acc + curr, 0);
}