with open('input.txt') as f:
    lines = f.readlines()

cubes_on = set()

for line in lines:
    cmd = line.split(" ")[0]
    xstart = int(line.split("x=")[1].split("..")[0])
    xend = int(line.split("x=")[1].split("..")[1].split(",")[0])
    ystart = int(line.split("y=")[1].split("..")[0])
    yend = int(line.split("y=")[1].split("..")[1].split(",")[0])
    zstart = int(line.split("z=")[1].split("..")[0])
    zend = int(line.split("z=")[1].split("..")[1].split(",")[0])

    if cmd == "on":
        for i in range(xstart, xend+1):
            for j in range(ystart, yend+1):
                for k in range(zstart, zend+1):
                    cubes_on.add((i, j, k))
    else:
        for i in range(xstart, xend+1):
            for j in range(ystart, yend+1):
                for k in range(zstart, zend+1):
                    if (i, j, k) in cubes_on:
                        cubes_on.remove((i, j, k))
print(len(cubes_on))
