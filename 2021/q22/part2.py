from typing import List

with open('p2_input.txt') as f:
    lines = f.readlines()


class Cube:
    def __init__(self, cmd, x1, x2, y1, y2, z1, z2):
        self.cmd = cmd
        self.x1 = x1
        self.x2 = x2
        self.y1 = y1
        self.y2 = y2
        self.z1 = z1
        self.z2 = z2

    def volume(self):
        return abs(self.x2-self.x1+1)*abs(self.y2-self.y1+1)*abs(self.z2-self.z1+1)

    def intersects(self, other: 'Cube'):
        return self.intersection_volume(other) > 0

    def intersection_volume(self, other: 'Cube'):
        # thanks https://stackoverflow.com/a/5556796
        return max(1+min(other.x2, self.x2)-max(other.x1, self.x1), 0) \
            * max(1+min(other.y2, self.y2)-max(other.y1, self.y1), 0) \
            * max(1+min(other.z2, self.z2)-max(other.z1, self.z1), 0)

    def left_intersect(self, other: 'Cube'):
        return min(other.x2, self.x2)

    def right_intersect(self, other: 'Cube'):
        return max(other.x1, self.x1)

    def top_intersect(self, other: 'Cube'):
        return min(other.y2, self.y2)

    def bottom_intersect(self, other: 'Cube'):
        return max(other.y1, self.y1)

    def front_intersect(self, other: 'Cube'):
        return min(other.z2, self.z2)

    def back_intersect(self, other: 'Cube'):
        return max(other.z1, self.z1)

    def subtract(self, other: 'Cube'):
        return [
            Cube(self.cmd, self.x1, self.left_intersect(other),
                 self.y1, self.y2, self.z1, self.z2),  # left
            Cube(self.cmd, self.right_intersect(other), self.x2,
                 self.y1, self.y2, self.z1, self.z2),  # right
            Cube(self.cmd, self.x1, self.x2, self.y1, self.y2,
                 self.z1, self.front_intersect(other)),  # front
            Cube(self.cmd, self.x1, self.x2, self.y1, self.y2,
                 self.back_intersect(other), self.z2),  # back
            Cube(self.cmd, other.x1, other.x2, self.y1,
                 self.top_intersect(other), other.z1, other.z2),  # top
            Cube(self.cmd, other.x1, other.x2, self.bottom_intersect(
                other), self.y2, other.z1, other.z2),  # bottom
        ]

    def sign(self):
        if self.cmd == "off":
            return -1
        return 1

    def __str__(self):
        return "<%s: %s,%s %s,%s, %s,%s -> %s>" % (self.cmd, self.x1, self.x2, self.y1,
                                                   self.y2, self.z1, self.z2, self.volume())

    def __repr__(self):
        return self.__str__()


result = 0

cubes: List[Cube] = []

for line in lines:
    cmd = line.split(" ")[0]
    x1 = int(line.split("x=")[1].split("..")[0])
    x2 = int(line.split("x=")[1].split("..")[1].split(",")[0])
    y1 = int(line.split("y=")[1].split("..")[0])
    y2 = int(line.split("y=")[1].split("..")[1].split(",")[0])
    z1 = int(line.split("z=")[1].split("..")[0])
    z2 = int(line.split("z=")[1].split("..")[1].split(",")[0])
    cubes += [Cube(cmd, x1, x2, y1, y2, z1, z2)]


print([c for c in cubes[1].subtract(cubes[0])])
# for cube in cubes:
#     print(fun(cube, cubes, 0))
#     print("---")

print(result)
