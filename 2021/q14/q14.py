# -*- coding: utf-8 -*-
from collections import defaultdict

with open('input.txt') as f:
    lines = f.readlines()

template = lines[0]

rules = [v for i, v in enumerate(lines) if i > 1]
rules_map = {r.split(' -> ')[0]: r.split(' -> ')[1].strip() for r in rules}

pairs = zip(template, template[1:])
cts = defaultdict(lambda: 0)
for p in pairs:
    cts[(p[0]+p[1]).strip()] += 1


def step(cts):
    result = cts.copy()
    for k, v in cts.items():
        first = k[0]
        if len(k) == 1:
            continue
        second = k[1]
        if k in rules_map:
            insert = rules_map[k]
            result[first + insert] += v
            result[insert + second] += v
            result[first+second] -= v
            if result[first+second] <= 0:
                del result[first+second]
    return result


def count(cts):
    totals = defaultdict(lambda: 0)
    for k, v in cts.items():
        first = k[0]
        totals[first] += v
    return totals[max(totals, key=totals.get)] - totals[min(totals, key=totals.get)]


# part 1
c = cts
for i in range(10):
    c = step(c)
print(count(c))

# part 2
c = cts
for i in range(40):
    c = step(c)
print(count(c))
