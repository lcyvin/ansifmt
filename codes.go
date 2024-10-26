package ansifmt

import (
  "strconv"
)

type Code interface {
  String() string
  Reset()  string
}

type Escape string

func (eac Escape) String() string {
  return "\\"+string(eac)+"["
}

const (
  ESCAPE_OCT Escape = "033"
  ESCAPE_HEX Escape = "\x1b"
)

var ESCAPE string = ESCAPE_OCT.String()

type Reset int

func (arc Reset) String() string {
  return strconv.Itoa(int(arc))
}

func (arc Reset) Reset() string {
  return strconv.Itoa(int(arc))
}

const RESET Reset = 0

type Color int
func (acc Color) String() string {
  return strconv.Itoa(int(acc))
}

func (acc Color) Reset() string {
  return strconv.Itoa(int(DEFAULT))
}

func (acc Color) Background() Background {
  return Background(acc)
}

const (
  BLACK Color = iota + 30
  RED
  GREEN
  YELLOW
  BLUE
  MAGENTA
  CYAN
  WHITE
  DEFAULT Color = 39
)

type Background int

func (abcc Background) String() string {
  return strconv.Itoa(int(abcc)+10)
}

func (abcc Background) Reset() string {
  return strconv.Itoa(int(DEFAULT)+10)
}

type Graphics int

func (agc Graphics) String() string {
  return strconv.Itoa(int(agc))
}


func (agc Graphics) Reset() string {
  if agc == 1 {
    return "22"
  }

  return strconv.Itoa(int(agc) + 20)
}

const (
  BOLD Graphics = iota + 1
  DIM
  ITALIC
  UNDERLINE
  BLINK
  _
  INVERSE
  INVISIBLE
  STRIKETHROUGH
)
