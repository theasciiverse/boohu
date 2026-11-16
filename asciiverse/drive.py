"""Plays Boohu for the catalog preview (asciiverse/record-preview.sh).

Reads the tmux pane and presses keys at a watchable pace: past the title,
then auto-explore (o), picking up simellas and items on the way. When a
creature is in view, the hero walks up to the nearest one and fights it
(moving into a monster attacks it), or waits for it to come. With hit
points low and nobody in view it rests (r); in a bad fight it backs
into a corridor, drinks the first potion it carries (once), or runs for
the stairs. When a level is explored it
travels to the stairs (G) and goes down (>). Quits without saving (Q)
after about 38 seconds.
    python3 -I drive.py <tmux socket> <session> [debug]
"""
import re
import subprocess
import sys
import time

SOCK, SESS = sys.argv[1], sys.argv[2]
DEBUG = len(sys.argv) > 3
LENGTH = 38.0
T0 = time.time()
KEYS = {(-1, 0): "h", (1, 0): "l", (0, -1): "k", (0, 1): "j",
        (-1, -1): "y", (1, -1): "u", (-1, 1): "b", (1, 1): "n"}
MAP_W, MAP_H = 79, 21              # the dungeon, at the top left
BLOCK = set("#¤ ")                 # walls, unseen cells


def tmux(*args):
    return subprocess.run(["tmux", "-L", SOCK, *args], capture_output=True,
                          text=True).stdout


def screen():
    return tmux("capture-pane", "-p", "-t", SESS).split("\n")


def key(k, wait=0.45):
    if DEBUG:
        print(round(time.time() - T0, 1), k, file=sys.stderr)
    tmux("send-keys", "-t", SESS, k)
    time.sleep(wait)


def cell(rows, x, y):
    if not (0 <= y < MAP_H and 0 <= x < MAP_W) or x >= len(rows[y]):
        return " "
    return rows[y][x]


def find(rows, test):
    return [(x, y) for y in range(MAP_H) for x in range(MAP_W)
            if test(cell(rows, x, y))]


def exposure(rows, pos, foes):
    """How many foes stand next to pos, then how open it is."""
    x, y = pos
    near = sum(max(abs(fx - x), abs(fy - y)) == 1 for fx, fy in foes)
    room = sum(cell(rows, x + dx, y + dy) not in BLOCK for dx, dy in KEYS)
    return (near, room)


def step_towards(rows, hero, goal, onto=False):
    """First step of the shortest walk next to goal (or onto it)."""
    seen = {hero: None}
    queue = [hero]
    reach = 0 if onto else 1
    for pos in queue:
        if max(abs(pos[0] - goal[0]), abs(pos[1] - goal[1])) <= reach and pos != hero:
            while seen[pos] != hero:
                pos = seen[pos]
            return (pos[0] - hero[0], pos[1] - hero[1])
        for d in KEYS:
            nxt = (pos[0] + d[0], pos[1] + d[1])
            c = cell(rows, *nxt)
            if nxt in seen or c in BLOCK or c.isalpha():
                continue
            seen[nxt] = pos
            queue.append(nxt)
    return None


def hp(rows):
    for row in rows:
        m = re.search(r"HP: ?(\d+)", row)
        if m:
            return int(m.group(1))
    return 99


time.sleep(2.5)                    # title screen
key("Space", 1.0)
start = time.time()
last = None
drank = False
while time.time() - start < LENGTH:
    rows = screen()
    text = "\n".join(rows[MAP_H:])
    lines = [r.rstrip() for r in rows[MAP_H:] if r.strip()]
    if lines and lines[-1].endswith("to continue]"):
        # a warning waits for x
        key("x", 0.8)
        continue
    hero = find(rows, lambda c: c == "@")
    if not hero:                   # a menu or the end screen
        key("Escape", 0.6)
        if not find(screen(), lambda c: c == "@"):
            if DEBUG:
                print("\n".join(screen()), file=sys.stderr)
            break
        continue
    hx, hy = hero[0]
    foes = find(rows, lambda c: c.isalpha())
    if foes and hp(rows) < 20 and not drank:
        drank = True               # one potion, whatever it is, for a bad fight
        key("q", 0.8)
        key("a", 1.0)
        key("Escape", 0.4)
        continue
    if foes:
        stairs = find(rows, lambda c: c == ">")
        if stairs and hp(rows) < 25:
            # hurt: run for the stairs and leave the fight behind
            if (hx, hy) in stairs:
                key(">", 1.5)
                continue
            step = step_towards(rows, (hx, hy), stairs[0], onto=True)
            if step:
                key(KEYS[step], 0.4)
                continue
        near = min(foes, key=lambda f: max(abs(f[0] - hx), abs(f[1] - hy)))
        dx, dy = near[0] - hx, near[1] - hy
        adjacent = [f for f in foes if max(abs(f[0] - hx), abs(f[1] - hy)) == 1]
        if len(adjacent) > 1:
            # surrounded: back into a spot where fewer can reach
            step = min((d for d in KEYS if cell(rows, hx + d[0], hy + d[1]) not in BLOCK
                        and not cell(rows, hx + d[0], hy + d[1]).isalpha()),
                       key=lambda d: exposure(rows, (hx + d[0], hy + d[1]), foes),
                       default=None)
            if step and exposure(rows, (hx + step[0], hy + step[1]), foes)[0] < len(adjacent):
                key(KEYS[step], 0.5)
                continue
        if max(abs(dx), abs(dy)) == 1:
            key(KEYS[(dx, dy)], 0.5)   # fight
            continue
        step = step_towards(rows, (hx, hy), near)
        key(KEYS[step] if step else ".", 0.4)
        continue
    if hp(rows) < 32:
        key("r", 1.0)
        continue
    if "Nothing left to explore" in text or rows == last:
        key("G", 1.5)
        key(">", 1.5)
        last = None
        continue
    key("o", 1.0)
    last = rows
key("Escape", 0.5)
key("Q", 1.0)
key("y", 2.0)
for _ in range(6):                 # the end screens, until the game exits
    if tmux("list-sessions").strip() == "":
        break
    key("Escape", 1.2)
