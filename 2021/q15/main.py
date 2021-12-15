# -*- coding: utf-8 -*-

# part 2 faster dijkstra
from heapdict import heapdict

with open('input.txt') as f:
    lines = f.readlines()

mm = [[int(z) for z in x.strip()] for x in lines]

# part 2 map extender
m = [[0 for _ in range(5 * len(mm[0]))] for _ in range(5*len(mm))]
y_adder = 0
for y in range(5*len(mm)):
    if y > 0 and y % len(mm) == 0:
        y_adder += 1
    x_adder = 0
    for x in range(5*len(mm[0])):
        if x > 0 and x % len(mm[0]) == 0:
            x_adder += 1
        rl = (((mm[y % len(mm)][x % len(mm[0])])) + x_adder + y_adder)
        if rl > 9:
            rl = (rl % 10) + 1
        m[y][x] = rl


def risk_level(x, y):
    return m[y][x]


def neighbors(x, y):
    n = set()
    if x > 0:
        n.add((x-1, y))
    if x+1 < len(m[y]):
        n.add((x+1, y))
    if y > 0:
        n.add((x, y-1))
    if y+1 < len(m):
        n.add((x, y+1))
    return n


def dijkstra():
    Q = heapdict()
    dist = {}
    prev = {}
    goal = None
    for y in range(len(m)):
        for x in range(len(m[y])):
            dist[(x, y)] = float("inf")
            prev[(x, y)] = None
            Q[(x, y)] = dist[(x, y)]
            goal = (x, y)

    Q[(0, 0)] = 0

    while len(Q) > 0:
        u, _ = Q.popitem()

        if u == goal:
            return dist[goal]

        n = neighbors(u[0], u[1])
        for v in n:
            alt = dist[u] + risk_level(v[0], v[1])
            if alt == float("inf"):
                alt = risk_level(v[0], v[1])
            if alt < dist[v]:
                dist[v] = alt
                prev[v] = u
                Q[v] = alt


print(dijkstra())
