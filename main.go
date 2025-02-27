package main

import (
    "flag"
    "os"
    "github.com/gen2brain/beeep"
)

var (
    title string
    message string
)

func init() {
    // FIXME: take stdin
    flag.StringVar(&title, "title", "", "The notification title")
    flag.StringVar(&message, "message", "", "The message text")
    flag.Parse()

    if title == "" || message == "" {
        flag.PrintDefaults()
        os.Exit(1)
    }
}

func main() {
    err := beeep.Notify(title, message, "")
    if err != nil {
        panic(err)
    }
}
