import numpy as np

VER = 2

if VER == 1:
    areas = [[8, 12, 12, 16], [4, 8, 8, 12], [8, 12, 8, 12], [0, 4, 8, 12], [4, 8, 0, 4], [4, 8, 4, 8]]
if VER == 2:
    areas = [
        [0, 50, 100, 150], [50, 100, 50, 100], [100, 150, 50, 100],
        [0, 50, 50, 100], [150, 200, 0, 50], [100, 150, 0, 50]
    ]


class Player:
    def __init__(self):
        self.x: int = 0
        self.y: int = 0
        self.facing: int = 0
        self.side: int = 4

    def is_wall(self, grid, x, y):
        return grid[y][x] == '#'

    def rotate(self, right: bool):
        if right:
            self.facing = (self.facing + 1) % 4
        else:
            self.facing = (self.facing - 1) % 4

    def move_one_step(self, grid):
        direction = "right"
        if self.facing == 1:
            direction = "down"
        elif self.facing == 2:
            direction = "left"
        elif self.facing == 3:
            direction = "up"

        char, _x, _y, _side, _direction = player.next_step(grid, self.x, self.y, self.side, direction)
        if char == '#':
            return

        facing = 0
        if _direction == "down":
            facing = 1
        elif _direction == "left":
            facing = 2
        elif _direction == "up":
            facing = 3

        self.x = _x
        self.y = _y
        self.side = _side
        self.facing = facing

    def next_step(self, grid, x, y, side, direction):
        edge = len(grid['1'])

        nside, ndir = side, direction
        wrap = False

        if direction == "right":
            x += 1
            if x == edge:
                nside, ndir = get_next_direction(side, direction)
                wrap = True
                if ndir == "down":
                    x = edge - y - 1
                if ndir == "left":
                    y = edge - y - 1
                if ndir == "up":
                    x = y

        elif direction == "down":
            y += 1
            if y == edge:
                nside, ndir = get_next_direction(side, direction)
                wrap = True
                if ndir == "right":
                    y = edge - x - 1
                if ndir == "left":
                    y = x
                if ndir == "up":
                    x = edge - x - 1

        elif direction == "left":
            x -= 1
            if x == -1:
                nside, ndir = get_next_direction(side, direction)
                wrap = True
                if ndir == "right":
                    y = edge - y - 1
                if ndir == "down":
                    x = y
                if ndir == "up":
                    x = edge - y - 1

        elif direction == "up":
            y -= 1
            if y == -1:
                nside, ndir = get_next_direction(side, direction)
                wrap = True
                if ndir == "right":
                    y = x
                if ndir == "down":
                    x = edge - x - 1
                if ndir == "left":
                    y = edge - x - 1

        if wrap:
            if ndir == "right":
                x = 0
            if ndir == "left":
                x = edge - 1
            if ndir == "down":
                y = 0
            if ndir == "up":
                y = edge - 1

        return grid[str(nside)][y][x], x, y, nside, ndir

    def __str__(self):
        score = (self.y+1) * 1000 + (self.x+1) * 4 + self.facing
        return f"x: {self.x+1}, y: {self.y+1}, facing: {self.facing}\nscore: {score}"


def get_next_direction(side, direction):
    if VER == 1:
        if direction == "right":
            sides = {1: (4, "left"), 5: (6, "right"), 6: (2, "right"), 2: (1, "down"), 4: (1, "left"), 3: (1, "right")}
        if direction == "down":
            sides = {4: (2, "down"), 2: (3, "down"), 3: (5, "up"), 5: (3, "up"), 6: (3, "right"), 1: (5, "right")}
        if direction == "left":
            sides = {5: (1, "up"), 1: (3, "left"), 2: (6, "left"), 6: (5, "left"), 4: (6, "down"), 3: (6, "up")}
        if direction == "up":
            sides = {4: (5, "down"), 5: (4, "down"), 3: (2, "up"), 2: (4, "up"), 6: (4, "right"), 1: (2, "left")}

    if VER == 2:
        if direction == "right":
            sides = {1: (3, "left"), 2: (1, "up"), 6: (3, "right"), 5: (3, "up"), 4: (1, "right"), 3: (1, "left")}
        if direction == "down":
            sides = {4: (2, "down"), 2: (3, "down"), 3: (5, "left"), 5: (1, "down"), 6: (5, "down"), 1: (2, "left")}
        if direction == "left":
            sides = {5: (4, "down"), 6: (4, "right"), 2: (6, "down"), 1: (4, "left"), 4: (6, "right"), 3: (6, "left")}
        if direction == "up":
            sides = {4: (5, "right"), 5: (6, "up"), 3: (2, "up"), 2: (4, "up"), 6: (2, "right"), 1: (5, "up")}

    return sides[side]


def rotate_side(mat: list, right: bool):
    if right:
        return np.flip(mat, axis=0).T
    return np.flip(mat.T, axis=0)


def grid_to_cube(grid):
    if VER == 1:
        side4 = grid[0:4, 8:12]
        side2 = grid[4:8, 8:12]
        side3 = grid[8:12, 8:12]
        side1 = grid[8:12, 12:16]
        side6 = grid[4:8, 4:8]
        side5 = grid[4:8, 0:4]

    if VER == 2:
        side1 = grid[0:50, 100:150]
        side4 = grid[0:50, 50:100]
        side2 = grid[50:100, 50:100]
        side3 = grid[100:150, 50:100]
        side5 = grid[150:200, 0:50]
        side6 = grid[100:150, 0:50]

    return {
        "1": side1,
        "2": side2,
        "3": side3,
        "4": side4,
        "5": side5,
        "6": side6,
    }


def parse_map_row(line: str):
    row = []
    for char in line:
        row.append(char)
    return row


def parse_path(line: str):
    instructions = []
    lline = len(line)
    buffer = ""
    is_number = True
    for i in range(0, lline):
        if is_number:
            if ord(line[i]) < 58:
                buffer += line[i]
            else:
                instructions.append(int(buffer))
                buffer = line[i]
                is_number = False
        else:
            if ord(line[i]) < 58:
                instructions.append(buffer)
                buffer = line[i]
                is_number = True
            else:
                buffer += line[i]

    if is_number:
        instructions.append(int(buffer))
    else:
        instructions.append(buffer)

    return instructions


def normalize_grid(grid):
    max_width = 0
    for i in range(len(grid)):
        if len(grid[i]) > max_width:
            max_width = len(grid[i])

    for i in range(len(grid)):
        remaining = max_width - len(grid[i])
        for j in range(remaining):
            grid[i].append(' ')

    return grid


def read_input(filename):
    grid = []
    with open(filename, 'r') as fr:
        is_map = True
        for line in fr.readlines():
            if line[-1] == "\n":
                line = line[:-1]

            if len(line) == 0:
                is_map = False
                continue

            if is_map:
                grid.append(parse_map_row(line))
            else:
                instructions = parse_path(line)

    grid = normalize_grid(grid)

    return grid, instructions


def test_one_direction(side, direction):
    pside = side
    pdir = direction
    for i in range(4):
        _side, _dir = get_next_direction(pside, pdir)
        assert _side + pside != 7
        pside, pdir = _side, _dir

    try:
        assert pside == side
    except AssertionError:
        print(f"I did not get from {side} (is {pside}) back after 4 steps going {direction} initially")


def test_cube():
    for direction in ["right", "down", "left", "up"]:
        for i in range(1, 7):
            test_one_direction(i, direction)


def visualize(grid, player):
    for i in range(len(grid['1'])):
        for k in range(1, 7):
            for j in range(len(grid[str(k)][0])):
                if player.side == k and player.x == j and player.y == i:
                    print("P ", end='')
                else:
                    print(grid[str(k)][i][j] + " ", end='')
            print("   ", end='')
        print()
    print()


test_cube()

grid, instructions = read_input(f'input{VER}')
grid = grid_to_cube(np.array(grid))

side = 4
player = Player()
x, y, side, direction = (player.x, player.y, side, "right")

for i, instruction in enumerate(instructions):
    if isinstance(instruction, int):
        for j in range(instruction):
            player.move_one_step(grid)
    if isinstance(instruction, str):
        player.rotate(instruction == "R")

visualize(grid, player)
y = player.y
x = player.x
win_coords = [areas[player.side - 1][0] + y + 1, areas[player.side - 1][2] + x + 1]
print(win_coords[0] * 1000 + win_coords[1] * 4 + player.facing)
print(player.side)
print(player)
