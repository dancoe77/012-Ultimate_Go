package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
)

func main() {
	// GoLand: kill_server/server.pid -- caveat have to use the content root path in GoLand
	err := KillServer("02-Data_Structures_and_REST_APIs/06-Working_with_files_using_defer_error_handling/kill_server/server.pid")
	if err != nil {
		fmt.Println("ERROR:", err)
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("not found")
		}
		for e := err; e != nil; e = errors.Unwrap(e) {
			fmt.Printf("> %s\n", e)
		}
	}
}

func KillServer(pidFile string) error {
	file, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	/*
		- defer happens when function exits, no matter what (panic)
		- defer works at the function level
		- multiple defers are executed in reverse order (stack, LIFO)
		- IDIOM: try to acquire a resource, check for an error, defer release
	*/
	defer func() {
		if err := file.Close(); err != nil {
			slog.Warn("close", "file", pidFile, "error", err)
		}
	}()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return fmt.Errorf("%q - bad pid: %w\n", pidFile, err)
	}

	slog.Info("killing", "pid", pid)

	if err := os.Remove(pidFile); err != nil {
		// We are not failing, only warning.
		slog.Warn("delete", "file", pidFile, "error", err)
	}

	return nil
}
