p1 = 10
p2 = 8

p1_points = 0
p2_points = 0

die_current = 1
die_rolls = 0


def die_next():
    global die_rolls
    global die_current
    d = die_current
    die_current += 1
    die_rolls += 1
    return d


current_p = 1
while True:
    rolls = die_next() + die_next() + die_next()
    if current_p == 1:
        p1 += rolls
        p1 = 1 + (p1 - 1) % (10)
        p1_points += p1
        if p1_points >= 1000:
            break
        current_p = 2
    else:
        p2 += rolls
        p2 = 1 + (p2 - 1) % (10)
        p2_points += p2
        if p2_points >= 1000:
            break
        current_p = 1

print(min(p1_points, p2_points) * die_rolls)
