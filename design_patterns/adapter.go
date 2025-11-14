package designpatterns

import (
	"fmt"
	"io"
)

// ---------- ADAPTEE (Interface B / existing object) ----------
type LegacyPrinter struct{} // <-- this is the "existing object" (implements API B)

func (LegacyPrinter) Print(s string) { // <-- API B: Print(string)
	fmt.Println("legacy:", s)
}

// ---------- ADAPTER (implements Interface A) ----------
type PrinterAdapter struct {
	p LegacyPrinter // <-- adapter *wraps* the existing object (composition)
}

// Write implements io.Writer (Interface A) by translating bytes->string
// and delegating to the LegacyPrinter.Print method.
func (a PrinterAdapter) Write(b []byte) (int, error) {
	// translation step: convert io.Writer's bytes into the string the adaptee expects
	s := string(b)

	// delegation step: forward the translated call to the adaptee
	a.p.Print(s)

	// satisfy io.Writer contract: return number of bytes "written" and nil error
	return len(b), nil
}

// ---------- CONSUMER (expects Interface A) ----------
func LogSomething(w io.Writer, msg string) {
	// This function is written to use any io.Writer (Interface A).
	// It does not (and should not) know about LegacyPrinter or Print.
	fmt.Fprintf(w, "log: %s\n", msg) // fmt.Fprintf calls w.Write(...) internally
}

func Adapter() {
	legacy := LegacyPrinter{}            // create the adaptee (existing object)
	adapter := PrinterAdapter{p: legacy} // wrap it in the adapter that implements Interface A

	// The consumer only needs an io.Writer (Interface A). We pass the adapter.
	LogSomething(adapter, "hello adapter")
}
