package main

import "os"

abort := make(chan struct{})
go func() {
	os.Stdin.Read(make([]byte, 1))
	abort <- struct {}{}
}