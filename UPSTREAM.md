# Changes for upstream

Hi! The Asciiverse build of Boohu carries two small changes to
`io.go`, written so they can be taken or left as they are. PATCHES.md
lists them with the packaging files.

## Windows data folder (`io.go`)

`DataDir()` checked `os.Getenv("GOOS") == "windows"`. `GOOS` is only
set in the environment while building, so a Windows player's data went
to `$HOME/.local/share/boohu` (a relative `.local/share/boohu` when
`HOME` is unset) instead of `%LOCALAPPDATA%\boohu`. It now checks
`runtime.GOOS`, as `main.go` does for the 8-colour default.

## Save, config, replay and dump in the XDG state folder (`io.go`)

`DataDir()` used `$XDG_DATA_HOME/boohu`. The XDG base directory spec
has since added `$XDG_STATE_HOME` (default `~/.local/state`) for state
that should persist between runs but is not important data to carry
between machines: saves, history, logs. `DataDir()` now uses
`$XDG_STATE_HOME/boohu` outside Windows, and when that folder does not
exist yet but `$XDG_DATA_HOME/boohu` does, it renames the old one
across first, so existing saves, settings and replays come along.
Windows still uses `%LOCALAPPDATA%\boohu`. A small `xdgDir()` helper
reads an XDG variable with its default. Harmonist and Shamogu carry the
same change in our builds.

Thank you for Boohu!
