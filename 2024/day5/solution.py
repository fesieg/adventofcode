from utils import pyutils as adventutils

def valid_order(instruction, rules):
  for one, two in rules:
      if one in instruction and two in instruction:
          indexOne = instruction.index(one)
          indexTwo = instruction.index(two)
          if indexOne > indexTwo: 
              return False
  return True


def find_valid(instruction, rules):
    if len(instruction) == 1:
        return instruction

    for i in range(len(instruction)):
        current_instruction = instruction[:i] + instruction[i+1:]
        
        if valid_order(current_instruction, rules):
            ordered = find_valid(current_instruction, rules)
            if ordered is not None:
                return [instruction[i]] + ordered

    return None


def topo_sort(instruction, rules):
    i = 0
    while i != len(instruction):
        i = len(instruction)
        for rule in rules:
            one, two = rule[0], rule[1]
            if one not in instruction or two not in instruction:
                continue
            index_one = instruction.index(one)
            index_two = instruction.index(two)
            if index_one > index_two:
                instruction.pop(index_one)
                instruction.insert(index_two, one)
                break  

    return instruction

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


def part_two(data_input):
    # broken
    rules = get_rules_from_input(data_input)
    invalid_instructions = []
    
    for line in get_pages_from_input(data_input):
        instruction = [int(page) for page in line.split(",")]
        if not find_valid(instruction, rules):
            invalid_instructions.append(instruction)
 
    sorted = []
    for instruction in invalid_instructions:
        sorted_instruction = topo_sort(instruction, rules)
        if sorted_instruction:
            sorted.append(sorted_instruction)

    return get_middle_sum(sorted)


def get_rules_from_input(input_str):
    split_input = input_str.split("  ")
    rules_raw = split_input[0].replace(" ", ",").split(",")
    return [(int(a), int(b)) for a, b in (line.split('|') for line in rules_raw)]

def get_pages_from_input(input_str):
    pages_raw = input_str.split("  ")[1]
    return pages_raw.split(" ")

    
input_data = adventutils.getTextFileContentAsString("./day5/input.txt")
print(f"Part 1 = {part_one(input_data)}")
print(f"Part 2 = {part_two(input_data)}")
