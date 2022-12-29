import re
from time import time
from typing import Optional, Dict, List
from itertools import product
import functools


cache = {}
TIME = 10
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


def cmp_cache(c1, c2):
    #if c1[0][2] < c2[0][2]:
    #    return -1
    #if c1[0][2] > c2[0][2]:
    #    return 1
    #return c1[1][0] - c2[1][0]
    return c2[1][0] - c1[1][0]


def get_opened_valves_from_visited(arr):
    rated = set(['FI', 'EA', 'AT', 'SA', 'LR', 'TO', 'YD', 'TH', 'QN', 'YG', 'UD', 'GW', 'NA', 'XX', 'KB'])
    arr = set(arr)
    r1 = arr.difference(rated.difference(arr))
    return r1, arr.difference(rated), r1.difference(arr.difference(rated))


def get_cache_prospective_paths(count):
    best_paths = {k: v for k, v in sorted(cache.items(), key=functools.cmp_to_key(cmp_cache))[:count]}
    for cache_key, cache_value in best_paths.items():
        res = get_opened_valves_from_visited(cache_value[1])
        print(f"{cache_value[0]}: standing at {cache_key[0]}, {cache_key[1]},"
              f"visited: {cache_value[1]}\nopened: {res[0]}\nredundant: {res[1]}\ncleaned: {res[2]}\n")


def dfs(valves: Dict[str, dict], valve1: str, valve2: str, opened: List[str], rem_time: int):
    if rem_time == 0:
        return 0, set(), ""

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
    winning_open = set()
    winning_prev_path = ""
    winning_path1 = ""
    winning_path2 = ""

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

        sub_score, sub_path, wp = dfs(valves, sub_valve1, sub_valve2, sub_opened, rem_time - 1)

        if a1 == "open":
            if valves[valve1]["flow"] * (rem_time - 1) + sub_score > score:
                score = valves[valve1]["flow"] * (rem_time - 1) + sub_score
                winning_open = sub_path
                winning_open.add(valve1)
                winning_prev_path = wp
                winning_path1 = a1
                winning_path2 = a2
        else:
            if sub_score > score:
                score = sub_score
                winning_open = sub_path
                winning_prev_path = wp
                winning_path1 = a1
                winning_path2 = a2

        if a2 == "open":
            if valves[valve2]["flow"] * (rem_time - 1) + sub_score > score:
                score = valves[valve2]["flow"] * (rem_time - 1) + sub_score
                winning_open = sub_path
                winning_open.add(valve2)
                winning_prev_path = wp
                winning_path1 = a1
                winning_path2 = a2
        else:
            if sub_score > score:
                score = sub_score
                winning_open = sub_path
                winning_prev_path = wp
                winning_path1 = a1
                winning_path2 = a2

        if a1 == "open" and a2 == "open" and valve1 != valve2:
            if valves[valve1]["flow"] * (rem_time - 1) + valves[valve2]["flow"] * (rem_time - 1) + sub_score > score:
                score = valves[valve1]["flow"] * (rem_time - 1) + valves[valve2]["flow"] * (rem_time - 1) + sub_score
                winning_open = sub_path
                winning_open.add(valve1)
                winning_open.add(valve2)
                winning_prev_path = wp
                winning_path1 = a1
                winning_path2 = a2

    if winning_path1 == "" and winning_path2 == "":
        winning_path = "Finish"
    else:
        winning_path = f"{winning_path1},{winning_path2} -> {winning_prev_path}"

    cache[key] = (score, winning_open, winning_path)
    return score, winning_open, winning_path


start = time()
valves = get_valves()
score, opened, path = dfs(valves, "AA", "AA", ["AA"], TIME)

# I got 6 candidates from steps=10, long-live the manual genetic algorithm!
#score, path = dfs(valves, "TH", "EA", ['YG', 'SA', 'EA', 'QN', 'KB', 'TH', 'GW', 'LR', 'FI', 'AA'], TIME)  # 338
#score, path = dfs(valves, "PL", "GJ", ['YG', 'SA', 'EA', 'QN', 'KB', 'TH', 'GW', 'AT', 'LR', 'FI', 'AA'], TIME)  # 338
#score, path = dfs(valves, "AA", "AA", ['YG', 'SA', 'EA', 'QN', 'KB', 'TH', 'GW', 'AT', 'LR', 'FI', 'AA'], TIME)  # 338
#score, path = dfs(valves, "RC", "PL", ['YG', 'SA', 'EA', 'QN', 'TH', 'GW', 'AT', 'LR', 'FI', 'AA'], TIME)  # 322
#score, path = dfs(valves, "EA", "SA", ['YG', 'SA', 'EA', 'QN', 'TH', 'GW', 'LR', 'FI', 'AA'], TIME)  # 322
#score, path = dfs(valves, "AA", "EA", ['YG', 'SA', 'EA', 'QN', 'TH', 'GW', 'LR', 'FI', 'AA'], TIME)

with open("results.txt", "w") as fw:
    fw.write(str(score))

print(score)
print(opened)
print(path)
print(time() - start)

#get_cache_prospective_paths(10)
# outputs 414: {'JJ', 'DD', 'BB', 'AA', 'CC', 'HH', 'EE'}