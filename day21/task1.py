import re

pattern1 = re.compile("([a-z]+): (\\d+)")
pattern2 = re.compile("([a-z]+): ([a-z]+)\\s([+\\-*/])\\s([a-z]+)")


def parse_line(calc_table, line):
    match = pattern1.match(line)
    if match:
        var_name = match.group(1)
        var_value = match.group(2)
        calc_table[var_name] = {'value': int(var_value)}
        return

    match = pattern2.match(line)
    if match:
        var_name = match.group(1)
        left_varname = match.group(2)
        op = match.group(3)
        right_varname = match.group(4)
        calc_table[var_name] = {'left': left_varname, 'op': op, 'right': right_varname, 'value': '?'}
        return

    raise ValueError(f"Don't know how to parse {line}")


def read_input(filename):
    calc_table = {}
    with open(filename, 'r') as fr:
        for line in fr.readlines():
            parse_line(calc_table, line.strip())
    return calc_table


def calc_root(calc_table: dict, node: str) -> int:
    if isinstance(calc_table[node]['value'], str):
        left_val = calc_root(calc_table, calc_table[node]['left'])
        right_val = calc_root(calc_table, calc_table[node]['right'])
        op = calc_table[node]['op']

        if op == "+":
            calc_table[node]['value'] = left_val + right_val
        if op == "-":
            calc_table[node]['value'] = left_val - right_val
        if op == "/":
            calc_table[node]['value'] = int(left_val / right_val)
        if op == "*":
            calc_table[node]['value'] = left_val * right_val

    return calc_table[node]['value']


calc_table = read_input('input2')
val = calc_root(calc_table, 'root')
print(val)
