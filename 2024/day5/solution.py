from utils import pyutils as adventutils

def valid_order(instruction, rules):
  for one, two in rules:
      if one in instruction and two in instruction:
          indexOne = instruction.index(one)
          indexTwo = instruction.index(two)
          if indexOne > indexTwo: 
              return False
  return True

def part_one(data_input):
    valid_instructions = []
    rules = get_rules_from_input(data_input)    
    for line in get_pages_from_input(data_input):
        instruction = [int(page) for page in line.split(",")]
        if valid_order(instruction, rules):
            valid_instructions.append(line)
            
    return get_middle_sum(valid_instructions)


def get_middle_sum(valid_instructions):
    middle_sum = 0
    for instruction in valid_instructions:
        # TODO: refactor this ugly type check
        if (type(instruction) == str):
          parts = instruction.split(",")
        else:
          parts = instruction
        middle_index = len(parts) // 2
        middle_value = int(parts[middle_index])
        middle_sum += middle_value

    return middle_sum

def get_rules_from_input(input_str):
    split_input = input_str.split("  ")
    rules_raw = split_input[0].replace(" ", ",").split(",")
    return [(int(a), int(b)) for a, b in (line.split('|') for line in rules_raw)]

def get_pages_from_input(input_str):
    pages_raw = input_str.split("  ")[1]
    return pages_raw.split(" ")

    
input_data = adventutils.file_content_as_string("./day5/simpleInput.txt")
print(f"Part 1 = {part_one(input_data)}")
