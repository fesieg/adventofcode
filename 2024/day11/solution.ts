import { getFileContent } from "../utils/tsutils";

const stones = (await getFileContent("input.txt")).split(" ");

function getChangedStones(stone: string): string[]{
  if (stone === "0") {
    return ["1"];
  }

  const length = stone.length;

  if (length % 2 === 0) {
    const mid = length / 2;
    const firstHalf = leadingZeroesRemoved(stone.substring(0, mid));
    const secondHalf = leadingZeroesRemoved(stone.substring(mid));
    return [firstHalf, secondHalf];
  }

  return [String(BigInt(stone) * 2024n)];
}

function blinkNTimes(inputStones: string[], blinks: number): string[] {
  // TODO: use map for part 2
  // loop through stones and insert result of getChangedStones for each stone
  let loopStones = structuredClone(inputStones)
  while (blinks > 0){
    let stoneCache: string[] = [];
    for (const stone of loopStones){
      stoneCache.push(...getChangedStones(stone));
    }

    loopStones = stoneCache;
    blinks--;
  }

  return loopStones
}

console.log("PART 1: ", blinkNTimes(stones, 25).length);
console.log("PART 2: ", blinkNTimes(stones, 75).length);

function leadingZeroesRemoved(input: string): string {
  return input.replace(/^0+(?!$)/, '');
}