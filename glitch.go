package main

import (
        "os"
        "math/rand"
        "time"
        "encoding/base64"
        "strings"
        "log"
)

var Rand = rand.New(rand.NewSource(time.Now().Unix()))
var Characters = "abcdefghigklmnopqrstuvwxyzABCDEFGHIGKLMNOPQRSTUVWXYZ0123456789+/"

type Glitch struct {
        ImageFile string
}

func (g *Glitch) glitch() {
        src, _ := os.Open(g.ImageFile)
        defer src.Close()
        stat, _ := src.Stat()

        img := make([]byte, stat.Size())
        src.Read(img)
        img64 := base64.StdEncoding.EncodeToString(img)

        slice := make([]string, len(img64))
        for i, c := range img64 {
                slice[i] = string(c)
        }

        try := Rand.Intn(10) + 1
        for i := 0; i < try; i++ {
                spot := Rand.Intn(len(img64))
                slice[spot] = string(Characters[Rand.Intn(len(Characters))])
        }

        newImg, _ := base64.StdEncoding.DecodeString(strings.Join(slice, ""))

        dst, _ := os.Create("_" + g.ImageFile)
        defer dst.Close()
        dst.Write(newImg)
}

func NewGlitch(image string) *Glitch {
        return &Glitch{image}
}

func main() {
        if len(os.Args) < 2 {
                log.Fatal("must specified JPEG file")
        }
        glitch := NewGlitch(os.Args[1])
        glitch.glitch()
}
