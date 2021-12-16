from typing import Union
import uuid
import functools

with open('input.txt') as f:
    lines = f.readlines()

raw_packet = lines[0].strip()


class Packet:
    raw = ""
    bin_string = ""
    version = 0
    type_id = 0
    length_type_id = 0
    length = 0
    data = 0
    bits_consumed = 0

    def __init__(self, raw: Union[str, None], bin_str: Union[str, None] = None):
        self.subpackets = []
        self.uuid = uuid.uuid4()
        if bin_str:
            self.bin_string = bin_str
        else:
            self.bin_string = "".join(
                [bin(int(char, 16))[2:].zfill(4) for char in raw])

    def parse(self):
        self.version = self.bin_to_dec(self.bin_string[0:3])
        self.type_id = self.bin_to_dec(self.bin_string[3:6])
        if self.type_id == 4:  # literal
            i = 6
            done = False
            literal_data = ""
            while not done:
                done = self.bin_string[i] == "0"
                literal_data += self.bin_string[i+1:i+5]
                i += 5
            self.data = self.bin_to_dec(literal_data)
            self.bits_consumed = i
        else:
            self.length_type_id = self.bin_to_dec(self.bin_string[6:7])
            if self.length_type_id == 0:
                # total length in bits of the sub-packets contained by this packet.
                self.length = self.bin_to_dec(self.bin_string[7:22])
                i = 22
                while i-22 < self.length:
                    next = Packet(None, self.bin_string[i:])
                    next.parse()
                    self.subpackets.append(next)
                    i += next.bits_consumed
                self.bits_consumed = i
            else:
                # number of sub-packets immediately contained by this packet.
                self.length = self.bin_to_dec(self.bin_string[7:18])
                i = 18
                subpacket_count = 0
                while subpacket_count < self.length:
                    next = Packet(None, self.bin_string[i:])
                    next.parse()
                    self.subpackets.append(next)
                    subpacket_count += 1
                    i += next.bits_consumed
                self.bits_consumed = i

    def bin_to_dec(self, bs: str):
        return int(bs, 2)

    def __str__(self):
        if self.type_id == 4:
            return "{%s literal v%d: %d}" % (str(self.uuid)[0:4], self.version, self.data)
        return "<%s operator v%d type %d lt %d l %d: [\n\t%s]>" % (str(self.uuid)[0:4], self.version, self.type_id, self.length_type_id, self.length, "\n\t".join([str(s) for s in self.subpackets]))

    def version_sum(self):
        return self.version + sum([p.version_sum() for p in self.subpackets])

    def exec(self) -> int:
        if self.type_id == 4:
            return self.data
        if self.type_id == 0:
            return sum((p.exec() for p in self.subpackets))
        if self.type_id == 1:
            if len(self.subpackets) == 1:
                return self.subpackets[0].exec()
            else:
                return functools.reduce(lambda a, b: a*b, (p.exec() for p in self.subpackets))
        if self.type_id == 2:
            return min((p.exec() for p in self.subpackets))
        if self.type_id == 3:
            return max((p.exec() for p in self.subpackets))
        if self.type_id == 5:
            return 1 if self.subpackets[0].exec() > self.subpackets[1].exec() else 0
        if self.type_id == 6:
            return 1 if self.subpackets[0].exec() < self.subpackets[1].exec() else 0
        if self.type_id == 7:
            return 1 if self.subpackets[0].exec() == self.subpackets[1].exec() else 0


p = Packet(raw_packet)
p.parse()
print(p.version_sum())
print(p.exec())
