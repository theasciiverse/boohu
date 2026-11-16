# Boohu: Break Out Of Hareka's Underground

**Every year, the elders send someone down for the simellas. This year
it's you.**

The medicinal simella plants grow only in Hareka's Underground, so every
year the village elders choose someone to go down and gather them. This
year the honour fell on you. Deep below, the elders say, a magical
monolith will carry you home. Between you and it: goblin bands, barking
hounds that follow your scent, tiny harpies that blink away when hurt,
cyclopes throwing rocks, mirror specters drinking your magic, and
earth dragons, which are said to have dug most of these tunnels in the
first place.

Boohu is anaseto's coffee-break roguelike, and every fight in it is a
question of where you stand. There are no experience points and no
character sheet to fuss over: you grow by what you find. A spear strikes
two foes in a row, an axe hits everyone around you, the dancing rapier
swaps places with your foe to strike the one behind, and the har-kar
gauntlets carry you through a line of enemies on the wind. Doors open only while
someone stands in them, blue foliage hides you from sight, and the
magical stones in the floor turn a hit into a teleport, a cloud of fog,
a confusing boom or a burst of temporary walls. Twice on the way down
you simply feel different and come away with an aptitude: good ears,
scales, fast legs, or the earth itself throwing your foes back.

Listen, too. Somewhere in the Underground you may hear faint music,
*larilon, larila*: that is Frundis, a musician and harmonist who also
happens to be a two-handed staff. And if someone teleports you away
for no reason at all, it is probably Marevor Helith, an ancient undead
nakrus, a repentant necromancer and expert in magaras, down here
looking for his old disciple Jaixel.

Reach the monolith on the eighth level and you are home. Or dive past
it into three optional depths, for more simellas and worse company.

## Playing

Run `boohu` in a terminal of at least 80x24. From 93x25 up it draws its
full layout with a side panel; smaller, it switches to a compact one by
itself.

- **h j k l y u b n**, the **arrow keys** or the **number pad**: move;
  walk into a creature to attack it
- **H J K L Y U B N**: travel in a direction
- **o**: explore automatically; **G**: go to the nearest stairs;
  **>**: descend
- **.** or **5**: wait a turn; **r**: rest
- **e** or **g**: equip what you stand on
- **q**: quaff a potion; **t**: throw; **v**: evoke a rod;
  **i**: inventory
- **x**: examine; **%**: your character and quest; **m**: messages
- **Mouse**: click to move, and the buttons along the bottom
- **?**: help; **=**: settings and key bindings; **S**: save and quit;
  **Q**: quit without saving

## Options

- `-s`: the 16-colour solarized palette (for a solarized terminal)
- `-o`: 8 colours only
- `-c`: keep the camera centred on you
- `-n`: no animations
- `-r _`: watch a replay of your last game

## Where things are kept

The save, settings, the replay of your last game and its statistics
(`#` writes them) live in `$XDG_STATE_HOME/boohu`
(`~/.local/state/boohu`), or `%LOCALAPPDATA%\boohu` on Windows.

## Credits

Boohu is by anaseto (Yon), 2017 to 2025:
<https://codeberg.org/anaseto/boohu>, with its website at
<https://anaseto.codeberg.page/games/boohu>. ISC licence, in `LICENSE`.
Its line goes on in Harmonist (2019), grown from Boohu's code and set in
the same Underground, and in Shamogu (2025).

## This package

The Asciiverse build of the terminal version (tcell), for Linux, macOS
and Windows, from upstream's `v0.14.1`. The game is unchanged; saves
moved to the state folder and the Windows data folder was fixed.
`PATCHES.md` lists every change, `Changes` is upstream's own history,
and `boohu.6` the manual page. Build it yourself with
`CGO_ENABLED=0 go build`.
