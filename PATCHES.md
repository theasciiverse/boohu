# Changes from upstream

This is the Asciiverse packaging of
[Boohu: Break Out Of Hareka's Underground](https://codeberg.org/anaseto/boohu)
by anaseto (Yon), a coffee-break roguelike in Go on tcell. Upstream is
the `upstream` remote; this build is upstream's `v0.14.1` tag
(`686c990`, 2025-03-05). The `asciiverse` branch carries the changes
below, each small enough to offer upstream (UPSTREAM.md).

We ship the terminal version only: tcell, the default build without
tags. No Tk tiles build, no `ansi` build.

## Licence

`LICENSE` in the repository is the ISC licence ("Permission to use,
copy, modify, and distribute this software for any purpose with or
without fee..."), copyright 2017 Yon; the man page `boohu.6` carries the
same notice. The catalog lists `ISC`. The terminal build ships no tile
images.

## The game

Unchanged: same dungeon, rules, texts and screens. In a terminal of at
least 93x25 the game draws its full 100x26 layout with the side panel;
in a smaller one, down to 80x24, it switches to its compact layout by
itself, as upstream has it.

## Code

- **Windows data folder** (`io.go`, `DataDir`): the check for Windows
  read the `GOOS` environment variable, which is only set while
  building, so on Windows the save went to `$HOME/.local/share/boohu`
  (a relative path when `HOME` is unset) instead of
  `%LOCALAPPDATA%\boohu`. It now asks `runtime.GOOS`, as `main.go`
  already does.
- **State folder** (`io.go`, `DataDir`): the save, config, replay and
  statistics dump were kept in `$XDG_DATA_HOME/boohu` (default
  `~/.local/share/boohu`). They now live in `$XDG_STATE_HOME/boohu`
  (default `~/.local/state/boohu`), the XDG folder meant for per-user
  state such as saves. A folder already in the old place is moved
  across on first use, so nobody loses a game. Windows keeps
  `%LOCALAPPDATA%\boohu`.

Boohu writes no log of its own beyond those files: each game is
recorded as a replay (`boohu -r _` plays the last one back), and `#`
writes the game's statistics to `dump` in the same folder.

## Packaging

- `.github/workflows/asciiverse-release.yml` (new): av archives for Linux
  (static, x86-64 and ARM64), macOS (Intel and Apple silicon) and Windows
  (x86-64 and ARM64), plus SHA256SUMS, cross-built with cgo off on one
  runner. Each archive holds the `boohu` binary, `LICENSE`, `Changes`,
  `boohu.6`, this file and our `README.md`.
- `asciiverse/README.md` (new): the package README, shipped in the
  archives as `README.md` (upstream's README is about building from
  source and the tiles and browser versions).
- `asciiverse/record-preview.sh`, `asciiverse/drive.py` (new): the
  catalog preview.
- `asciiverse/boohu.toml`, `asciiverse/boohu.md`, `asciiverse/boohu.cast`
  (new): the catalog manifest, store page and preview, as drafts for hub.
- `CLAUDE.md`, `PATCHES.md`, `UPSTREAM.md` (new): notes.
