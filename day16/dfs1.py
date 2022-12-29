import re
from time import time
from typing import Dict, Set


cache = {}
TIME = 30
INPUT_NAME = 'input2'


def get_cache_key(valve_name, opened, rem_time) -> tuple:
    return valve_name, tuple(sorted(opened)), rem_time


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


def dfs(valves: Dict[str, dict], valve_name: str, opened: Set[str], rem_time: int):
    if rem_time == 0:
        return 0

    key = get_cache_key(valve_name, opened, rem_time)
    if key in cache:
        return cache[key]

    score = 0
    if valve_name not in opened and valves[valve_name]["flow"] > 0:
        #tmp_opened = set()
        #tmp_opened.update(opened)
        #tmp_opened.add(valve_name)
        opened.add(valve_name)
        score = max(score, valves[valve_name]["flow"] * (rem_time - 1) +
                    dfs(valves, valve_name, opened, rem_time - 1))
        opened.remove(valve_name)

    for path in valves[valve_name]['valves']:
        score = max(score, dfs(valves, path, opened, rem_time - 1))

    cache[key] = score
    return score


start = time()
valves = get_valves()
suma = 0
for valve in valves.values():
    suma += valve["flow"]
print(suma)
#score = dfs(valves, "AA", set(["AA"]), TIME)
#print(score)
#print(time() - start)
