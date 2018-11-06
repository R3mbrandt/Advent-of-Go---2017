#!/usr/bin/env python3
from collections import defaultdict


def isNum(z):
    try:
        int(z)
    except ValueError:
        return False
    return True


f = open("input_day23.txt", "r")

registers = defaultdict(int)
inp = []
mulcounter = 0
instrPointer = 0

for line in f:
    inp.append(line.strip().split(" "))

while True:
    if instrPointer >= len(inp):
        break
    i = inp[instrPointer]
    lv = int(i[1]) if isNum(i[1]) else registers[i[1]]
    rv = int(i[2]) if isNum(i[2]) else int(registers[i[2]])
    if i[0] == "set":
        registers[i[1]] = rv
    elif i[0] == "sub":
        registers[i[1]] = lv - rv
    elif i[0] == "mul":
        registers[i[1]] = lv * rv
        mulcounter += 1
    elif i[0] == "jnz":
        if lv != 0:
            instrPointer += rv
            continue
    instrPointer += 1

print(mulcounter)
