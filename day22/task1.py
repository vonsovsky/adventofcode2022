class Player:
    def __init__(self, x, y):
        self.x: int = x
        self.y: int = y
        self.facing: int = 0

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

        char, x, y = self.next_non_empty(grid, self.x, self.y, direction)
        if char == '.':
            self.x = x
            self.y = y

    def safe_op(self, var, maximum, add):
        var += add
        if var == maximum:
            var = 0
        if var == -1:
            var = maximum - 1
        return var

    def next_non_empty(self, grid, x, y, direction):
        height = len(grid)
        width = len(grid[0])

        if direction == "right":
            x = self.safe_op(x, width, +1)
            while grid[y][x] == ' ':
                x = self.safe_op(x, width, +1)
        if direction == "left":
            x = self.safe_op(x, width, -1)
            while grid[y][x] == ' ':
                x = self.safe_op(x, width, -1)
        if direction == "down":
            y = self.safe_op(y, height, +1)
            while grid[y][x] == ' ':
                y = self.safe_op(y, height, +1)
        if direction == "up":
            y = self.safe_op(y, height, -1)
            while grid[y][x] == ' ':
                y = self.safe_op(y, height, -1)

        return grid[y][x], x, y

    def __str__(self):
        score = (self.y+1) * 1000 + (self.x+1) * 4 + self.facing
        return f"x: {self.x+1}, y: {self.y+1}, facing: {self.facing}\nscore: {score}"


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


def visualize(grid, player):
    for i in range(len(grid)):
        for j in range(len(grid[0])):
            if player.x == j and player.y == i:
                print("P ", end='')
            else:
                print(grid[i][j] + " ", end='')
        print()
    print()


grid, instructions = read_input('input2')
for i in range(len(grid[0])):
    if grid[0][i] != ' ':
        player = Player(i, 0)
        break

for i, instruction in enumerate(instructions):
    if isinstance(instruction, int):
        for i in range(instruction):
            player.move_one_step(grid)
    if isinstance(instruction, str):
        player.rotate(instruction == "R")

visualize(grid, player)
print(player)
