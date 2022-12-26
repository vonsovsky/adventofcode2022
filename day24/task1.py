blizzards = []
blchars = [">", "v", "<", "^"]
bliz_map_cache = {}
visited_nodes = set()


def parse_map_row(line: str, index: int):
    arr = list(line)
    for i in range(len(arr)):
        if arr[i] in blchars:
            blizzards.append({'y': index - 1, 'x': i - 1, 'facing': blchars.index(arr[i])})
    return arr


def blizzardAt(x, y, max_x, max_y, facing, time):
    if facing == 0:
        return y, (x + time) % max_x
    if facing == 2:
        return y, (x - time) % max_x
    if facing == 1:
        return (y + time) % max_y, x
    if facing == 3:
        return (y - time) % max_y, x


def convertBlizzardsToMap(max_x, max_y, time):
    if time in bliz_map_cache:
        return bliz_map_cache[time]

    bliz_map = {}
    for bliz in blizzards:
        y, x = blizzardAt(bliz['x'], bliz['y'], max_x, max_y, bliz['facing'], time)
        pos = (y, x)
        bliz_map[pos] = bliz['facing']
    bliz_map_cache[time] = bliz_map

    return bliz_map


def to_direction(facing: int):
    direction = ">"
    if facing == 1:
        direction = "v"
    if facing == 2:
        direction = "<"
    if facing == 3:
        direction = "^"
    return direction


def is_valid_pos(bliz_map, x, y, max_x, max_y):
    if x == 0 and y == -1:
        return True
    if x < 0 or y < 0:
        return False
    if x >= max_x or y >= max_y:
        return False

    pos = (y, x)
    return pos not in bliz_map


def calc_visited_node(x, y, max_x, max_y, time):
    return (x + time) % max_x, (y + time) % max_y, time


def add_node(queue, x, y, max_x, max_y, t):
    node_name = calc_visited_node(x, y, max_x, max_y, t)
    if node_name not in visited_nodes:
        queue.append((x, y, t))
        visited_nodes.add(node_name)


def bfs(grid, x, y):
    max_x = len(grid[0]) - 2
    max_y = len(grid) - 2
    queue = [(x, y, 0)]

    while len(queue) > 0:
        _x, _y, _t = queue.pop(0)
        if _x == max_x - 1 and _y == max_y - 1:
            return _t + 1

        bliz_map = convertBlizzardsToMap(max_x, max_y, _t + 1)

        if is_valid_pos(bliz_map, _x + 1, _y, max_x, max_y):
            add_node(queue, _x + 1, _y, max_x, max_y, _t + 1)
        if is_valid_pos(bliz_map, _x, _y + 1, max_x, max_y):
            add_node(queue, _x, _y + 1, max_x, max_y, _t + 1)
        if is_valid_pos(bliz_map, _x - 1, _y, max_x, max_y):
            add_node(queue, _x - 1, _y, max_x, max_y, _t + 1)
        if is_valid_pos(bliz_map, _x, _y - 1, max_x, max_y):
            add_node(queue, _x, _y - 1, max_x, max_y, _t + 1)
        if is_valid_pos(bliz_map, _x, _y, max_x, max_y):
            add_node(queue, _x, _y, max_x, max_y, _t + 1)


def visualize(grid, time):
    print(time)
    max_x = len(grid[0]) - 2
    max_y = len(grid) - 2
    bliz_map = convertBlizzardsToMap(max_x, max_y, time)
    for i in range(max_y):
        for j in range(max_x):
            pos = (i, j)
            if pos in bliz_map:
                print(to_direction(bliz_map[pos]), end=' ')
            else:
                print('. ', end='')
        print()
    print()


def read_input(filename):
    grid = []
    index = 0
    with open(filename, 'r') as fr:
        for line in fr.readlines():
            if line[-1] == "\n":
                line = line[:-1]

            grid.append(parse_map_row(line, index))
            index += 1

    return grid


grid = read_input('input2')

#visualize(grid, time=1)
steps = bfs(grid, 0, -1)
print(steps)
