grid = []


def new_rock(typ: int, ypos: int, jet_pointer: int):
    tmpypos, jep = fallStone(typ, {'x': 2, 'y': ypos + 4}, jets, jet_pointer)
    ypos = max(ypos, tmpypos)
    #visualize(grid)
    while len(grid) < ypos + 10:
        add_row()

    return ypos, jep


def add_row():
    grid.append(['.'] * 7)


def position_modifiers(typ: int):
    if typ == 0:
        return [(0, 0), (0, 1), (0, 2), (0, 3)]
    elif typ == 1:
        return [(0, 1), (1, 0), (1, 1), (2, 1), (1, 2)]
    elif typ == 2:
        return [(2, 2), (1, 2), (0, 0), (0, 1), (0, 2)]
    elif typ == 3:
        return [(0, 0), (1, 0), (2, 0), (3, 0)]
    elif typ == 4:
        return [(0, 0), (0, 1), (1, 0), (1, 1)]

    raise ValueError(f"Unknown type: {typ}")


def collision_free(typ, pos) -> bool:
    for pmod in position_modifiers(typ):
        try:
            if grid[pos['y'] + pmod[0]][pos['x'] + pmod[1]] == '#':
                return False
        except IndexError as ex:
            print(ex)
            print(pos['y'], pmod[0], pos['x'], pmod[1])
            return False

    return True


def make_it_stone(gr, typ, pos) -> int:
    pmods = position_modifiers(typ)
    maximum = 0
    for pmod in pmods:
        if pos['y'] + pmod[0] > maximum:
            maximum = pos['y'] + pmod[0]
        gr[pos['y'] + pmod[0]][pos['x'] + pmod[1]] = '#'
    return maximum


def oneRoundMove(char: str, typ: int, pos: dict) -> bool:
    if char == "<" and pos['x'] > 0:
        if collision_free(typ, {'x': pos['x'] - 1, 'y': pos['y']}):
            pos['x'] -= 1
    lastX = position_modifiers(typ)[-1][1]
    if char == ">" and pos['x'] + lastX < 6:
        if collision_free(typ, {'x': pos['x'] + 1, 'y': pos['y']}):
            pos['x'] += 1

    if pos['y'] == 0:
        return False, pos

    free = collision_free(typ, {'x': pos['x'], 'y': pos['y'] - 1})
    if not free:
        return False, pos

    pos['y'] -= 1
    return True, pos


def fallStone(typ: int, pos: dict, jets, jet_pointer) -> int:
    moving = True
    while moving:
        moving, pos = oneRoundMove(jets[jet_pointer], typ, pos)

        jet_pointer += 1
        if jet_pointer == len(jets):
            jet_pointer = 0

    return make_it_stone(grid, typ, pos), jet_pointer


def visualize(gr):
    for i in range(len(grid)-1, -1, -1):
        print(f"{i + 1} ", end='')
        for j in range(len(gr[0])):
            print(f"{gr[i][j]} ", end='')
        print()
    print()


def findCycle(grid):
    # found cycle of size 2660, starting at 3300
    # it equals starting 2128 stopped rocks, with cycle of 1700 rocks
    block_size = 20

    for j in range(0, 3000):
        for i in range(j, len(grid) - block_size - j):
            found = True
            for k in range(0, block_size):
                if grid[i+j+k] != grid[j+k]:
                    found = False
                    break
            if found:
                print(j+i)


def calc_height(rocks):
    height = -1
    jet_pointer = 0
    if rocks >= 2128:
        rocks -= 2128
        height += 3300
        jet_pointer = 2476

        cycles = int(rocks / 1700)
        rocks %= 1700
        height += 2660 * cycles

    for i in range(10):
        add_row()
    grid[0] = ['.', '#', '#', '#', '#', '#', '.']
    grid[1] = ['#', '#', '#', '.', '#', '#', '#']
    grid[2] = ['.', '#', '.', '.', '.', '.', '#']
    grid[3] = ['.', '.', '.', '.', '.', '.', '#']

    ypos = 2
    print(f"Remaining rocks: {rocks}")
    for i in range(3, rocks + 3):
        ypos, jet_pointer = new_rock(i % 5, ypos, jet_pointer)

    return height + ypos - 2


def read_input(filename):
    with open(filename, 'r') as fr:
        for line in fr.readlines():
            jets = line.strip()
    return jets


jets = read_input('input2')

height = calc_height(rocks=1000000000000)  # for input2
print(height)

jet_pointer = 0
ypos = -1

grid = []
for i in range(10):
    add_row()

for i in range(2022):
    ypos, jet_pointer = new_rock(i % 5, ypos, jet_pointer)

#visualize(grid)

print(ypos)
# 2128
# 3828, 5528, 7228

# 5132
# 9312, 13457, 17648
#findCycle(grid)

"""
for i in range(120):
    ypos, jet_pointer = new_rock(i % 5, ypos, jet_pointer)

#visualize(grid)
print(ypos+1)

for i in range(120, 240):
    ypos, jet_pointer = new_rock(i % 5, ypos, jet_pointer)

print(ypos+1)

for i in range(240, 360):
    ypos, jet_pointer = new_rock(i % 5, ypos, jet_pointer)

print(ypos+1)
"""
