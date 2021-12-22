import sys

sys.setrecursionlimit(5000)

p1_start = 10
p2_start = 9

# p1_start = 4
# p2_start = 8

# translate for modulo arithmetic
p1_start -= 1
p2_start -= 1

possible_outcomes = [i for i in range(3, 9 + 1)]
universes_spawned = [0, 0, 0, 1, 3, 6, 7, 6, 3, 1]

p1_wins = 0
p2_wins = 0

def branching_roll(turn, p1, p2, p1_score, p2_score, universes):
    if not turn and p1_score >= 21:
        global p1_wins
        p1_wins += universes
        return
    elif turn and p2_score >= 21:
        global p2_wins
        p2_wins += universes
        return

    if turn:
        for o in possible_outcomes:
            new_p1 = (p1 + o) % 10
            branching_roll(False, new_p1, p2, p1_score + new_p1 + 1, p2_score, universes * universes_spawned[o])
    else:
        for o in possible_outcomes:
            new_p2 = (p2 + o) % 10
            branching_roll(True, p1, new_p2, p1_score , p2_score + new_p2 + 1, universes * universes_spawned[o])

branching_roll(True, p1_start, p2_start, 0, 0, 1)
