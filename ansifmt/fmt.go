package ansifmt

import (
  "os"
	"fmt"
	"strings"
)

// Wrap returns a string with all provided formatting codes applied to the input string. These
// formatting codes are reset at the returned string's end. 
// NOTE: wrap does not implicitly append "\n" to its strings. Use ansifmt.Wrapln() for this.
func Wrap(str string, codes ...Code) string {
  fmtr := NewFormatter()
  return fmtr.Set(codes...).Append(str).Reset().String()
}

// Wrapln returns a string with all provided formatting codes applied to the input string. 
// These codes are reset at its end, and the returned string includes a newline character.
func Wrapln(str string, codes ...Code) string {
  fmtr := NewFormatter()
  return fmtr.Set(codes...).Append(str).Reset().String() + "\n"
}

// Formatter implements functions to be used to build and write
// ansi-formatted strings to stdout or as returned string values.
//
// Basic usage:
//   formatter := ansifmt.NewFormatter()
//   formatter.Set(ansifmt.BOLD, ansifmt.GREEN, ansifmt.UNDERLINE)
//   formatter.Append("Now in color!")
//   formatter.Println()
type Formatter struct {
  fmtrSlice  []string
}

// NewFormatter creates a new pointer to a Formatter
func NewFormatter() *Formatter {
  return &Formatter{fmtrSlice: make([]string, 0)}
}


// String returns the held internal string slice, with all formatting codes reset at its end.
func (fmtr *Formatter) String() string {
  return fmtr.Join("")
}

// Join returns the held internal string slice, joined together with the provided separator,
// and with all formatting codes reset at its end.
func (fmtr *Formatter) Join(sep string) string {
  fmtr = fmtr.Reset()
  return strings.Join(fmtr.fmtrSlice, sep)
}

// Printf takes inputs to fill any go string formatting directives held in the
// built string, similar to fmt.Printf() and prints the string to stdout. It will not
// contain a newline unless one was included within the input string set.
func (fmtr *Formatter) Printf(a ...any) (n int, err error) {
  return fmt.Fprintf(os.Stdout, fmtr.String(), a...)
}

// Println prints the input string set to stdout, appending a newline character. It does
// not take arguments for formatting. See ansifmt.Printf()
func (fmtr *Formatter) Println() (n int, err error) {
  return fmtr.Printf(fmtr.String()+"\n")
}

func (fmtr *Formatter) ansiOp(op func(Code) string, codes ...Code) *Formatter {
  fmtStrSlice := make([]string, len(codes))
  for idx, code := range codes {
    fmtStrSlice[idx] = op(code)
  }

  fmtStr := strings.Join(fmtStrSlice, ";")

  fmtr.fmtrSlice = append(fmtr.fmtrSlice, ESCAPE+fmtStr+"m")
  
  return fmtr
}

// internal function for DRY purposes
func set(code Code) string {
  return code.String()
}

// internal function for DRY purposes
func unset(code Code) string {
  return code.Reset()
}
// Set adds any number of ansi formatting codes to the string slice, without resetting.
// For automatic reset, use ansifmt.Wrap()
func (fmtr *Formatter) Set(codes ...Code) *Formatter {
  return fmtr.ansiOp(set, codes...)
}

// Unsets any current formatting codes. Primarily used for graphics modes. When used
// with a color code, ansifmt.DEFAULT is set.
func (fmtr *Formatter) Unset(codes ...Code) *Formatter {
  return fmtr.ansiOp(unset, codes...)
}

// Add an input string to the stringset for formatted output.
func (fmtr *Formatter) Append(str string) *Formatter {
  fmtr.fmtrSlice = append(fmtr.fmtrSlice, str)
  return fmtr
}

// Reset all formatting codes previously set by the Formatter. ansifmt.Set may be called
// after Reset to continue string building with new formatting codes.
func (fmtr *Formatter) Reset() *Formatter {
  return fmtr.Set(RESET)
}
