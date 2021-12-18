import math

with open('input.txt') as f:
    lines = f.readlines()


class Pair:
    a = None
    b = None
    index_a = None
    index_b = None

    def __str__(self):
        # return "[%s:%s,%s:%s]" % (self.index_a, self.a, self.index_b, self.b)
        return "[%s,%s]" % (self.a, self.b)

    def magnitude(self):
        m = 0
        if type(self.a) == Pair:
            m += 3*self.a.magnitude()
        else:
            m += 3*self.a
        if type(self.b) == Pair:
            m += 2*self.b.magnitude()
        else:
            m += 2*self.b
        return m


def pairify(x):
    p = Pair()
    if type(x[0]) == list:
        p.a = pairify(x[0])
    else:
        p.a = x[0]
    if type(x[1]) == list:
        p.b = pairify(x[1])
    else:
        p.b = x[1]
    return p


index = 0


def assign_index(x: Pair):
    global index
    index = 0

    def traverse(v: Pair):
        global index
        if type(v.a) == Pair:
            traverse(v.a)
        else:
            v.index_a = index
            index += 1
        if type(v.b) == Pair:
            traverse(v.b)
        else:
            v.index_b = index
            index += 1
    traverse(x)


def increment_index(x: Pair, i: int, by: int):
    global index
    if i < 0 or i >= index:
        return

    def traverse(v):
        if type(v.a) == Pair:
            traverse(v.a)
        elif v.index_a == i:
            v.a += by
        if type(v.b) == Pair:
            traverse(v.b)
        elif v.index_b == i:
            v.b += by
    traverse(x)


def explode(x) -> bool:
    def traverse(v, depth) -> bool:
        if type(v.a) == Pair:
            if depth >= 3 and type(v.a.a) == int and type(v.a.b) == int:
                # explode
                increment_index(x, v.a.index_a-1, v.a.a)
                increment_index(x, v.a.index_b+1, v.a.b)
                v.a = 0
                assign_index(x)
                return True
            if traverse(v.a, depth+1):
                return True
        if type(v.b) == Pair:
            if depth >= 3 and type(v.b.a) == int and type(v.b.b) == int:
                increment_index(x, v.b.index_a-1, v.b.a)
                increment_index(x, v.b.index_b+1, v.b.b)
                v.b = 0
                assign_index(x)
                return True
            if traverse(v.b, depth+1):
                return True
        return False
    return traverse(x, 0)


def split(x) -> bool:
    def traverse(v) -> bool:
        if type(v.a) == Pair:
            if traverse(v.a):
                return True
        elif v.a >= 10:
            old = v.a
            v.a = Pair()
            v.a.a = math.floor(old / 2)
            v.a.b = math.ceil(old / 2)
            assign_index(x)
            return True
        if type(v.b) == Pair:
            if traverse(v.b):
                return True
        elif v.b >= 10:
            old = v.b
            v.b = Pair()
            v.b.a = math.floor(old / 2)
            v.b.b = math.ceil(old / 2)
            assign_index(x)
            return True
        return False
    return traverse(x)


def reduce(x: Pair):
    while True:
        if explode(x):
            continue
        if split(x):
            continue
        break


# part 1
lhs = pairify(eval(lines[0]))
for line in lines[1:]:
    rhs = pairify(eval(line))
    adder = Pair()
    adder.a = lhs
    adder.b = rhs
    assign_index(adder)
    reduce(adder)
    lhs = adder

print(lhs.magnitude())

# part 2
max_mag = 0
for lhs in lines:
    for rhs in lines:
        if lhs == rhs:
            continue
        adder = Pair()
        adder.a = pairify(eval(lhs))
        adder.b = pairify(eval(rhs))
        assign_index(adder)
        reduce(adder)
        max_mag = max(max_mag, adder.magnitude())
print(max_mag)
