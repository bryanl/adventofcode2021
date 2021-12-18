from io import StringIO
from math import prod

solution1 = 0


def read_packet(buffer):
    """Read a packet."""
    global solution1
    version = buffer.read(3)
    solution1 += int(version, 2)
    type_id = buffer.read(3)

    if type_id == "100":
        chunks = []
        group = ""
        while group != "0":
            group = buffer.read(1)
            chunks.append(buffer.read(4))
        return int("".join(chunks), 2)

    packets = []
    if buffer.read(1) == "0":
        length = int(buffer.read(15), 2)
        target_length = buffer.tell() + length
        while buffer.tell() != target_length:
            packets.append(read_packet(buffer))
    else:
        num_packets = int(buffer.read(11), 2)
        for _ in range(num_packets):
            packets.append(read_packet(buffer))

    if type_id == "000": return sum(packets)
    if type_id == "001": return prod(packets)
    if type_id == "010": return min(packets)
    if type_id == "011": return max(packets)
    if type_id == "101": return int(packets[0] > packets[1])
    if type_id == "110": return int(packets[0] < packets[1])
    if type_id == "111": return int(packets[0] == packets[1])


data = "".join(
    [
        bin(int(x, 16))[2:].zfill(4)
        for x in open("input.txt").read().strip()
    ]
)
solution2 = read_packet(StringIO(data))


print(f"The first solution is {solution1}")  # 947
print(f"The second solution is {solution2}")  # 660797830937
