# -*- coding: utf-8 -*-
with open('input.txt') as f:
    lines = f.readlines()

coords = [tuple(map(int, v.strip().split(',')))
          for i, v in enumerate(lines) if i < lines.index('\n')]
cset = set(coords)
instructions = [v for i, v in enumerate(lines) if i > lines.index('\n')]


def fold_left(old_set, fold_x):
    new_set = old_set.copy()
    for c in old_set:
        x, y = c
        if x > fold_x:
            new_set.remove(c)
            new_set.add((2*fold_x-x, y))
    return new_set


def fold_up(old_set, fold_y):
    new_set = old_set.copy()
    for c in old_set:
        x, y = c
        if y > fold_y:
            new_set.remove(c)
            new_set.add((x, 2*fold_y - y))
    return new_set


def print_set(s):
    width = max((x for x, y in s))
    height = max((y for x, y in s))
    print(width, height)
    for y in range(height+1):
        for x in range(width+1):
            if (x, y) in s:
                print("■", end='')
            else:
                print(" ", end='')
        print("\n", end='')


s = cset
for instruction in instructions:
    print(instruction)
    if instruction.startswith("fold along x"):
        s = fold_left(s, int(instruction.replace("fold along x=", "")))
    else:
        s = fold_up(s, int(instruction.replace("fold along y=", "")))
    print_set(s)
