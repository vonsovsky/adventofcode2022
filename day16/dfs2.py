import re
from time import time
from typing import Optional, Dict, List
from itertools import product


cache = {}
TIME = 26
INPUT_NAME = 'input2'


def get_cache_key(valve1, valve2, opened, rem_time) -> Optional[int]:
    agents = sorted([valve1, valve2])
    key = (agents[0], agents[1], rem_time, tuple(opened))
    return key


def get_valves():
    valve_map = {}

    with open(INPUT_NAME, 'r') as fr:
        pattern = re.compile("Valve ([A-Z]+) has flow rate=(\d+); tunnels? leads? to valves? ([A-Z\s,]+)")
        for line in fr.readlines():
            line = line.strip()
            match = pattern.match(line)
            valve_name = match.group(1)
            flow = int(match.group(2))
            valves = match.group(3)
            valves = [valve.strip() for valve in valves.split(',')]

            valve_map[valve_name] = {'name': valve_name, 'flow': flow, 'valves': valves}

    return valve_map


def dfs(valves: Dict[str, dict], valve1: str, valve2: str, opened: List[str], rem_time: int):
    if rem_time == 0:
        return 0

    key = get_cache_key(valve1, valve2, opened, rem_time)
    if key in cache:
        return cache[key]

    actions1 = []
    actions2 = []

    if valves[valve1]["flow"] > 0 and valve1 not in opened:
        actions1.append("open")
    if valves[valve2]["flow"] > 0 and valve2 not in opened:
        actions2.append("open")
    actions1.extend(valves[valve1]['valves'])
    actions2.extend(valves[valve2]['valves'])

    score = 0
    for a1, a2 in product(actions1, actions2):
        sub_opened = list(opened)
        sub_valve1 = valve1
        sub_valve2 = valve2

        if a1 == "open":
            sub_opened.append(valve1)
        else:
            sub_valve1 = a1

        if a2 == "open":
            if valve2 not in sub_opened:
                sub_opened.append(valve2)
        else:
            sub_valve2 = a2

        if a1 == "open" or a2 == "open":
            sub_opened = sorted(sub_opened)

        sub_score = dfs(valves, sub_valve1, sub_valve2, sub_opened, rem_time - 1)

        if a1 == "open":
            if valves[valve1]["flow"] * (rem_time - 1) + sub_score > score:
                score = valves[valve1]["flow"] * (rem_time - 1) + sub_score
        else:
            if sub_score > score:
                score = sub_score

        if a2 == "open":
            if valves[valve2]["flow"] * (rem_time - 1) + sub_score > score:
                score = valves[valve2]["flow"] * (rem_time - 1) + sub_score
        else:
            if sub_score > score:
                score = sub_score

        if a1 == "open" and a2 == "open" and valve1 != valve2:
            if valves[valve1]["flow"] * (rem_time - 1) + valves[valve2]["flow"] * (rem_time - 1) + sub_score > score:
                score = valves[valve1]["flow"] * (rem_time - 1) + valves[valve2]["flow"] * (rem_time - 1) + sub_score

    cache[key] = score
    return score


start = time()
valves = get_valves()
score = dfs(valves, "AA", "AA", ["AA"], TIME)
with open("results.txt", "w") as fw:
    fw.write(str(score))

print(score)
print(time() - start)
