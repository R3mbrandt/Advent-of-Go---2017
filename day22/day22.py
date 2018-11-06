#!/usr/bin/env python3
input_file = open("input_day22.txt", "r")

map = []
map_dict = {}

# Build an two-dimensional-array from input_file and close the file
for line in input_file:
    map.append(list(line.strip()))
input_file.close()

# Convert map to dict (x,y):"." or "#"
for y in range(len(map)):
    for x in range(len(map[y])):
        map_dict[(x, y)] = map[y][x]

# Calculate start position (middle of map)
y, x = (len(map)//2, len(map[0])//2)

directions = [(0, -1), (1, 0), (0, 1), (-1, 0)]
direction = 0
inf_counter = 0

for _ in range(10**7):
    if map_dict.setdefault((x, y), ".") == ".":     # clean Node
        map_dict[(x,y)]="W"
        direction = (direction+3) % len(directions)
    elif map_dict[(x,y)]=="W":                      # Weakened Node
        map_dict[(x,y)]="#"
        inf_counter += 1
    elif map_dict[(x,y)]=="#":                      # Infected Node
        map_dict[(x,y)]="F"
        direction = (direction+1) % len(directions)
    elif map_dict[(x,y)]=="F":                      # Flagged Node
        map_dict[(x,y)]="."
        direction = (direction+2) % len(directions)

    x += directions[direction][0]
    y += directions[direction][1]

print("Part2:",inf_counter)
