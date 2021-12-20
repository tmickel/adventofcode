with open('input.txt') as f:
    lines = f.readlines()

algorithm = lines[0]


def calc(lit: set, c, flip, min_x=None, max_x=None, min_y=None, max_y=None):
    (x_in, y_in) = c
    ns = ""
    for y in range(y_in-2, y_in+1):
        for x in range(x_in-2, x_in+1):
            if (x, y) in lit:
                ns += "1"
            else:
                if flip and (x < min_x or x > max_x or y < min_y or y > max_y):
                    ns += "1"
                else:
                    ns += "0"
    return algorithm[int(ns, 2)] == "#"


all_lit = set()

pic = lines[2:]
y = 0
x = 0

width = 0
height = 0
for line in pic:
    l = line.strip()
    x = 0
    for res in l:
        if res == '#':
            all_lit.add((x, y))
        x += 1
    y += 1
    width = len(l)
    height += 1

output = set()
for x in range(-2, width+2):
    for y in range(-2, height+2):
        if calc(all_lit, (x, y), False):
            output.add((x, y))

for i in range(2, 51):
    min_x = 100000
    max_x = -100000
    min_y = 100000
    max_y = -100000
    for pt in output:
        x, y = pt
        min_x = min(min_x, x)
        min_y = min(min_y, y)
        max_x = max(max_x, x)
        max_y = max(max_y, y)

    all_lit = output
    output = set()
    for x in range(-i*2, width+i*2):
        for y in range(-i*2, height+i*2):
            if calc(all_lit, (x, y), algorithm[0] == "#" and i % 2 == 0,  min_x, max_x, min_y, max_y):
                output.add((x, y))

print(len(output))
