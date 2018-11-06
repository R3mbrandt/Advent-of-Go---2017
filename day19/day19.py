#!/usr/bin/env python3

input = []

f = open("input_day19.txt", "r")

for line in f.readlines():
    input.append(list(line))


# Search for Start in Row 0, Line 0
a_p_y = 0
a_p_x = 0

# Richtungen
dir = ((1, 0),  # runter (y+1)
       (-1, 0),  # hoch (y-1)
       (0, 1),  # rechts (x+1)
       (0, -1))  # links (x-1)


current_char =''

for a_p_x in range(0, len(input[a_p_y])):
    if input[a_p_y][a_p_x] == '|':
        current_char=input[a_p_y][a_p_x]
        start = (a_p_y, a_p_x)
        break


print(f"Start at: {start}: {current_char}")

while True:
    if current_char == '|':
        dy=a_p_y+1
        if dy > len (input):
            print(f"Error! Kann nicht weiter runter gehen")
        

print(dir[0])
