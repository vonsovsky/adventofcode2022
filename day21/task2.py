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
        if calc_table[node]['value'] == "UNK":
            return "UNK"

        left_val = calc_root(calc_table, calc_table[node]['left'])
        right_val = calc_root(calc_table, calc_table[node]['right'])
        op = calc_table[node]['op']

        if isinstance(left_val, str) or isinstance(right_val, str):
            calc_table[node]['value'] = f"({left_val} {op} {right_val})"
        elif op == "+":
            calc_table[node]['value'] = left_val + right_val
        elif op == "-":
            calc_table[node]['value'] = left_val - right_val
        elif op == "/":
            calc_table[node]['value'] = int(left_val / right_val)
        elif op == "*":
            calc_table[node]['value'] = left_val * right_val

    return calc_table[node]['value']


def solve(exp, left, op_val, right_side):
    if op_val == "+":
        return right_side - exp
    if op_val == "*":
        return int(right_side / exp)
    if op_val == "-":
        if not left:
            return right_side + exp
        else:
            return (right_side - exp) * -1
    if op_val == "/":
        if not left:
            return right_side * exp
        else:
            return exp / right_side


def resolve_equation(left_side, right_side):
    brackets = 0
    op_val = ''
    op_pos = -1

    for i, char in enumerate(left_side):
        if char == '(':
            brackets += 1
        if char == ')':
            brackets -= 1

        if brackets == 1 and (char == '+' or char == '-' or char == '*' or char == '/'):
            op_val = char
            op_pos = i
            break

    ls = left_side[1:op_pos].strip()
    rs = left_side[op_pos+1:-1].strip()

    if '(' not in ls and ls != "UNK":
        right_side = solve(int(ls), True, op_val, right_side)
    if '(' not in rs and rs != "UNK":
        right_side = solve(int(rs), False, op_val, right_side)

    if '(' in ls:
        return resolve_equation(ls, right_side)
    if '(' in rs:
        return resolve_equation(rs, right_side)

    return f"UNK = {right_side}"


calc_table = read_input('input2')
calc_table["humn"] = {"value": "UNK"}
val2 = calc_root(calc_table, 'fzbp')
print(val2)
val1 = calc_root(calc_table, 'fglq')
print(val1)

res = resolve_equation(val1, int(val2))
print(res)

"""
# Tests
res = resolve_equation("(20 - ((UNK * 3) + 2))", 6)
print(res)
res = resolve_equation("(24 / (UNK * 3))", 2)
print(res)
res = resolve_equation("((UNK * 3) / 2)", 6)
print(res)
"""
