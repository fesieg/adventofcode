import re

pattern = r"^mul\(\d+,\d+\)$"
test_strings = [
    "mul(123,456)",
    "mul(1,2)",
    "mul(123456789,987654321)",
    "mul(123,)",
    "mul(,123)",
    "amull(123,456)",
]

for test in test_strings:
    if re.match(pattern, test):
        print(f"'{test}' matches the pattern!")
    else:
        print(f"'{test}' does NOT match the pattern.")
