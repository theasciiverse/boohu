# Boohu (Asciiverse packaging)

anaseto's Boohu: Break Out Of Hareka's Underground (Go, tcell, ISC per
the repository's LICENSE), packaged for the Asciiverse catalog (id
`boohu`, kind `game`, genre `roguelike`). Not our code: the game stays
as its author made it; list every change in PATCHES.md and keep each one
offerable upstream (UPSTREAM.md). Its line: Boohu (2017), then Harmonist
(2019, grown from Boohu's code, same Hareka's Underground) and Shamogu
(2025), all three in the catalog.

- `upstream` is codeberg.org/anaseto/boohu (branch `master`); ours is
  branch `asciiverse` (from the `v0.14.1` tag), `origin`
  github.com/theasciiverse/boohu (private).
- Build: `CGO_ENABLED=0 go build -trimpath .` (terminal version, tcell;
  never `-tags tk` or `-tags ansi`). go.mod says 1.17; the workflow
  builds with Go 1.26.
- Release: tag `v<upstream version>-av.<n>` (first `v0.14.1-av.1`) and
  push it; `.github/workflows/asciiverse-release.yml` cross-builds
  Linux, macOS and Windows on one runner.
- Save, config, replay and statistics dump live in
  `$XDG_STATE_HOME/boohu`. Tests and scripted runs always set scratch
  `HOME`, `XDG_STATE_HOME` and `XDG_DATA_HOME` in the same command (the
  old data folder is moved on first use), under
  `systemd-run --user --scope -p MemoryMax=4G` with `GOMAXPROCS=4`.
- The full layout is 100x26 (map and side panel); below 93x25 the game
  switches to its compact 80x24 layout by itself.
- Preview: `asciiverse/record-preview.sh ./boohu`.
- Catalog drafts: `asciiverse/boohu.toml` and `.md`; package README
  `asciiverse/README.md` (shipped as README.md in the archives).
