from utils import pyutils as adventutils
from typing import List, Dict
import re
from collections import defaultdict

stones = [int(s) for s in adventutils.file_content_as_string("./day11/input.txt").split(" ")]

def get_changed_stones(stone: str) -> List[int]:
    if stone == "0": return [1]

    if len(stone) % 2 == 0:
        mid = int(len(stone) / 2)
        firstHalf = stone[:mid]
        secondHalf = stone[mid:]
        return [int(remove_leading_zeroes(firstHalf)), int(remove_leading_zeroes(secondHalf))]
    
    return [int(stone) * 2024]

def blink_map_n_times(input_stones: List[int], blinks: int) -> Dict[int, int]:
    stone_map = defaultdict(int)
    
    for stone in input_stones:
        # Initialize map with each input stone as key, and value as the amount of that stone present
        stone_map[stone] += 1

    while blinks > 0:
        # For each blink, go through all keys in the map
        # Get the changed stones for each key, and increment value in the map for each changed stone
        # We use a new map to store the new values, so we don't mess up the current map
        new_map = defaultdict(int)
        for stone, count in stone_map.items():
            changed_stones = get_changed_stones(str(stone))
            for changed_stone in changed_stones:
                new_map[changed_stone] += count

        stone_map = new_map
        blinks -= 1

    return dict(stone_map)

def remove_leading_zeroes(with_zeroes: str) -> str:
    return re.sub(r'^0+(?!$)', '', with_zeroes)

print("PART 1: ", sum(blink_map_n_times(stones, 25).values()))
print("PART 2: ", sum(blink_map_n_times(stones, 75).values()))