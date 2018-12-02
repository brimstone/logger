package logger

import "io"

//Method Fatalf[string ...interface {}] is missing
func (l *Logger) Fatalf(s string, args ...interface{}) {
	l.prepFields(1).Fatalf(s, args...)
}
func (l *Logger) Fatal(args ...interface{}) {
	l.prepFields(1).Fatal(args...)
}

//Method Fatalln[...interface {}] is missing
func (l *Logger) Fatalln(msg ...interface{}) {
	l.prepFields(1).Fatalln(msg...)
}

// Panic
func (l *Logger) Panic(args ...interface{}) {
	l.prepFields(1).Panic(args...)
}
func (l *Logger) Panicf(s string, args ...interface{}) {
	l.prepFields(1).Panicf(s, args...)
}
func (l *Logger) Panicln(msg ...interface{}) {
	l.prepFields(1).Panicln(msg...)
}

// Print
func (l *Logger) Print(args ...interface{}) {
	l.prepFields(1).Print(args...)
}
func (l *Logger) Printf(s string, args ...interface{}) {
	l.prepFields(1).Printf(s, args...)
}
func (l *Logger) Println(msg ...interface{}) {
	l.prepFields(1).Println(msg...)
}

// Flags
func (l *Logger) Flags() {
	l.prepFields(1).Debug("unsupported method Flags called")
}
func (l *Logger) SetFlags(i int) {
	l.prepFields(1).Debug("unsupported method SetFlags called")
}

// Output
func (l *Logger) Output(i int, s string) error {
	l.prepFields(1).Debug("unsupported method Output called")
	return nil
}
func (l *Logger) SetOutput(w io.Writer) {
	l.prepFields(1).Debug("unsupported method SetOutput called")
}

// Prefix
func (l *Logger) Prefix() {
	l.prepFields(1).Debug("unsupported method Prefix called")
}
func (l *Logger) SetPrefix(s string) {
	l.prepFields(1).Debug("unsupported method SetPrefix called")
}
