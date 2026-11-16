//go:build !js
// +build !js

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

func Replay(file string) error {
	ui := &gameui{}
	g := &game{}
	ui.g = g
	g.ui = ui
	err := g.LoadReplay(file)
	if err != nil {
		return fmt.Errorf("loading replay: %v", err)
	}
	err = ui.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "boohu: %v\n", err)
		os.Exit(1)
	}
	defer ui.Close()
	ui.DrawBufferInit()
	ui.Replay()
	return nil
}

// DataDir returns the game's data directory location, which holds the
// save, config, replay and dump. On Windows it is %LOCALAPPDATA%\boohu.
// Elsewhere it is the XDG state folder, $XDG_STATE_HOME/boohu (default
// ~/.local/state/boohu), where per-user state such as saves belongs; a
// folder left in the data folder ($XDG_DATA_HOME/boohu) by older versions
// is moved there on first use.
func (g *game) DataDir() (string, error) {
	var dataDir string
	if runtime.GOOS == "windows" {
		xdg := os.Getenv("LOCALAPPDATA")
		if xdg == "" {
			xdg = filepath.Join(os.Getenv("HOME"), ".local", "share")
		}
		dataDir = filepath.Join(xdg, "boohu")
	} else {
		dataDir = filepath.Join(xdgDir("XDG_STATE_HOME", ".local", "state"), "boohu")
		oldDir := filepath.Join(xdgDir("XDG_DATA_HOME", ".local", "share"), "boohu")
		if _, err := os.Stat(dataDir); err != nil {
			if _, err := os.Stat(oldDir); err == nil {
				if err := os.MkdirAll(filepath.Dir(dataDir), 0755); err == nil {
					// on failure the old folder stays where it is and
					// the game starts afresh in the new one
					os.Rename(oldDir, dataDir)
				}
			}
		}
	}
	_, err := os.Stat(dataDir)
	if err != nil {
		err = os.MkdirAll(dataDir, 0755)
		if err != nil {
			return "", fmt.Errorf("%v\n", err)
		}
	}
	return dataDir, nil
}

// xdgDir returns the XDG base directory named by env, or its default
// under the home directory.
func xdgDir(env string, def ...string) string {
	if dir := os.Getenv(env); dir != "" {
		return dir
	}
	return filepath.Join(append([]string{os.Getenv("HOME")}, def...)...)
}

func (g *game) Save() error {
	dataDir, err := g.DataDir()
	if err != nil {
		g.Print(err.Error())
		return err
	}
	saveFile := filepath.Join(dataDir, "save")
	data, err := g.GameSave()
	if err != nil {
		g.Print(err.Error())
		return err
	}
	err = ioutil.WriteFile(saveFile, data, 0644)
	if err != nil {
		g.Print(err.Error())
		return err
	}
	return nil
}

func (g *game) RemoveSaveFile() error {
	return g.RemoveDataFile("save")
}

func (g *game) Load() (bool, error) {
	dataDir, err := g.DataDir()
	if err != nil {
		return false, err
	}
	saveFile := filepath.Join(dataDir, "save")
	_, err = os.Stat(saveFile)
	if err != nil {
		// no save file, new game
		return false, err
	}
	data, err := ioutil.ReadFile(saveFile)
	if err != nil {
		return true, err
	}
	lg, err := g.DecodeGameSave(data)
	if err != nil {
		return true, err
	}
	if lg.Version != Version {
		return true, fmt.Errorf("saved game for previous version %s.", lg.Version)
	}
	*g = *lg
	return true, nil
}

func (g *game) SaveConfig() error {
	dataDir, err := g.DataDir()
	if err != nil {
		g.Print(err.Error())
		return err
	}
	saveFile := filepath.Join(dataDir, "config.gob")
	data, err := GameConfig.ConfigSave()
	if err != nil {
		g.Print(err.Error())
		return err
	}
	err = ioutil.WriteFile(saveFile, data, 0644)
	if err != nil {
		g.Print(err.Error())
		return err
	}
	return nil
}

func (g *game) LoadConfig() (bool, error) {
	dataDir, err := g.DataDir()
	if err != nil {
		return false, err
	}
	saveFile := filepath.Join(dataDir, "config.gob")
	_, err = os.Stat(saveFile)
	if err != nil {
		// no save file, new game
		return false, err
	}
	data, err := ioutil.ReadFile(saveFile)
	if err != nil {
		return true, err
	}
	c, err := g.DecodeConfigSave(data)
	if err != nil {
		return true, err
	}
	GameConfig = *c
	return true, nil
}

func (g *game) RemoveDataFile(file string) error {
	dataDir, err := g.DataDir()
	if err != nil {
		return err
	}
	dataFile := filepath.Join(dataDir, file)
	_, err = os.Stat(dataFile)
	if err == nil {
		err := os.Remove(dataFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *game) SaveReplay() error {
	dataDir, err := g.DataDir()
	if err != nil {
		g.Print(err.Error())
		return err
	}
	saveFile := filepath.Join(dataDir, "replay")
	data, err := g.EncodeDrawLog()
	if err != nil {
		g.Print(err.Error())
		return err
	}
	err = ioutil.WriteFile(saveFile, data, 0644)
	if err != nil {
		g.Print(err.Error())
		return err
	}
	return nil
}

func (g *game) LoadReplay(file string) error {
	dataDir, err := g.DataDir()
	if err != nil {
		return err
	}
	replayFile := filepath.Join(dataDir, "replay")
	if file != "_" {
		replayFile = file
	}
	_, err = os.Stat(replayFile)
	if err != nil {
		// no save file, new game
		return err
	}
	data, err := ioutil.ReadFile(replayFile)
	if err != nil {
		return err
	}
	dl, err := g.DecodeDrawLog(data)
	if err != nil {
		return err
	}
	g.DrawLog = dl
	return nil
}

func (g *game) WriteDump() error {
	dataDir, err := g.DataDir()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(dataDir, "dump"), []byte(g.Dump()), 0644)
	if err != nil {
		return fmt.Errorf("writing game statistics: %v", err)
	}
	err = g.SaveReplay()
	if err != nil {
		return fmt.Errorf("writing replay: %v", err)
	}
	return nil
}
