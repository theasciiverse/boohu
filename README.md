Break Out Of Hareka's Underground (Boohu) is a roguelike game mainly inspired
from DCSS and its tavern, with some ideas from Brogue, but aiming for very
short games, almost no character building, and a simplified inventory.

*Every year, the elders send someone to collect medicinal simella plants in the
Underground.  This year, the honor fell upon you, and so here you are.
According to the elders, deep in the Underground, a magical monolith will lead you
back to your village.  Along the way, you will collect simellas, as well as
various items that will help you deal with monsters, which you may
fight or flee...*

Website
-------

![Boohu introduction screen](https://anaseto.codeberg.page/games/boohu/intro-screen-tiles.png)

You can visit the [game's
website](https://anaseto.codeberg.page/games/boohu)
for more information and tips about the game.

Install from Sources
--------------------

First, you need to install the [go compiler](https://golang.org/), so that the
`go` command is available.

### ASCII

You can build a native ASCII version from source by using this command:

    go build

You should now have a working `boohu` executable in the current directory. Use
the `--help` option for command-line usage.
  
The only dependency outside of the go standard library is the
curses-like library [tcell](https://github.com/gdamore/tcell), which is
installed automatically by the previous `go build` command.

*Portability note.* If you want to avoid a dependency, try adding `--tags ansi`
to the `go build` command. It will work on POSIX systems with a `stty` command,
but has more limited functionality.

### Tiles

You can build a graphical version depending on Tcl/Tk (8.6) using this command:

	go build --tags tk

This will install the [gothic](https://codeberg.org/anaseto/gothic) Go bindings
for Tcl/Tk. You need to install Tcl/Tk first.

### Browser (Tiles or ASCII)

You can also build a WebAssembly version with:

    GOOS=js GOARCH=wasm go build --tags js -o boohu.wasm

You can then play by serving a directory containing the wasm file via http. The
directory should contain some other files that you can find in the main
website instance.

Colors
------

If the default colors do not display nicely on your terminal emulator, you can
use the `-s` option: `boohu -s` to use the 16-color palette, which
will display nicely if the [solarized](http://ethanschoonover.com/solarized)
palette is used. Configurations are available for most terminal emulators,
otherwise, colors may have to be configured manually to one's liking in
the terminal emulator options.

Documentation
-------------

See the man page boohu(6) for more information on command line options and use
of the replay file. For example:

    boohu -r _

launches an auto-replay of your last game.
