def moveN(mixer, position_to_index, frm, by):
    num = mixer[frm]
    posVal = position_to_index[frm]
    to = ((frm + by - 1) % (len(mixer) - 1)) + 1

    mixer = mixer[:frm] + mixer[frm+1:]
    mixer = mixer[:to] + [num] + mixer[to:]
    position_to_index = position_to_index[:frm] + position_to_index[frm+1:]
    position_to_index = position_to_index[:to] + [posVal] + position_to_index[to:]

    return mixer, position_to_index, to


def mix(mixer, position_to_index):
    index_pointer = 0
    for i in range(len(position_to_index)):
        pos = position_to_index.index(index_pointer)
        val = mixer[pos]
        mixer, position_to_index, to = moveN(mixer, position_to_index, pos, val)
        index_pointer += 1

    return mixer, position_to_index


def read_input(filename):
    numbers = []
    with open(filename, 'r') as fr:
        for index, line in enumerate(fr.readlines()):
            numbers.append((index + 1, int(line.strip())))
    return numbers


initial_numbers = read_input('input2')
mixer = [num * 811589153 for index, num in initial_numbers]
position_to_index = [index - 1 for index, num in initial_numbers]

for i in range(10):
    mixer, position_to_index = mix(mixer, position_to_index)

suma = 0
grove = mixer.index(0)
pos = (grove + 1000) % len(mixer)
print(mixer[pos])
suma += mixer[pos]

pos = (grove + 2000) % len(mixer)
print(mixer[pos])
suma += mixer[pos]

pos = (grove + 3000) % len(mixer)
print(mixer[pos])
suma += mixer[pos]

print(suma)
