package logger

import (
	"fmt"
	"time"
)

type ColorCode string

const (
	Reset      ColorCode = "\033[0m"
	Bold       ColorCode = "\033[1m"
	Red        ColorCode = "\033[31m"
	Green      ColorCode = "\033[32m"
	Yellow     ColorCode = "\033[33m"
	Blue       ColorCode = "\033[34m"
	Magenta    ColorCode = "\033[35m"
	Cyan       ColorCode = "\033[36m"
	White      ColorCode = "\033[37m"
	BoldRed    ColorCode = "\033[1;31m"
	BoldGreen  ColorCode = "\033[1;32m"
	BoldYellow ColorCode = "\033[1;33m"
	BoldBlue   ColorCode = "\033[1;34m"
	BoldPurple ColorCode = "\033[1;35m"
	BoldCyan   ColorCode = "\033[1;36m"
	BoldWhite  ColorCode = "\033[1;37m"
	BlackBg    ColorCode = "\033[40m"
	RedBg      ColorCode = "\033[41m"
	GreenBg    ColorCode = "\033[42m"
	YellowBg   ColorCode = "\033[43m"
	BlueBg     ColorCode = "\033[44m"
	MagentaBg  ColorCode = "\033[45m"
	CyanBg     ColorCode = "\033[46m"
	WhiteBg    ColorCode = "\033[47m"
)

func ColorTest() {
	// Test all colors
	fmt.Printf("%sRed%s ", Red, Reset)
	fmt.Printf("%sGreen%s ", Green, Reset)
	fmt.Printf("%sYellow%s ", Yellow, Reset)
	fmt.Printf("%sBlue%s ", Blue, Reset)
	fmt.Printf("%sMagenta%s ", Magenta, Reset)
	fmt.Printf("%sCyan%s ", Cyan, Reset)
	fmt.Printf("%sWhite%s\n", White, Reset)
	fmt.Printf("%sRed%s ", BoldRed, Reset)
	fmt.Printf("%sGreen%s ", BoldGreen, Reset)
	fmt.Printf("%sYellow%s ", BoldYellow, Reset)
	fmt.Printf("%sBlue%s ", BoldBlue, Reset)
	fmt.Printf("%sMagenta%s ", BoldPurple, Reset)
	fmt.Printf("%sCyan%s ", BoldCyan, Reset)
	fmt.Printf("%sWhite%s\n", BoldWhite, Reset)
	fmt.Printf("%sRed%s ", RedBg, Reset)
	fmt.Printf("%sGreen%s ", GreenBg, Reset)
	fmt.Printf("%sYellow%s ", YellowBg, Reset)
	fmt.Printf("%sBlue%s ", BlueBg, Reset)
	fmt.Printf("%sMagenta%s ", MagentaBg, Reset)
	fmt.Printf("%sCyan%s ", CyanBg, Reset)
	fmt.Printf("%sWhite%s ", WhiteBg, Reset)
	fmt.Printf("%sBlack%s\n", BlackBg, Reset)
}

/**
 * Logger which makes things easier
 * since these are pretty functions
 * for repetitive logging tasks
 */

type Logger struct {
	hasPrefix, includeTimestamp bool
	prefix                      string
	color                       ColorCode
}

func NewLogger() *Logger {
	return &Logger{
		hasPrefix:        false,
		includeTimestamp: false,
		prefix:           "",
		color:            White,
	}
}

func (l *Logger) SetPrefix(prefix string, color ColorCode) *Logger {
	l.hasPrefix = true
	l.prefix = prefix
	l.color = color
	return l
}

func (l *Logger) ResetPrefix() *Logger {
	l.hasPrefix = false
	l.prefix = ""
	l.color = White
	return l
}

func (l *Logger) IncludeTimestamp() *Logger {
	l.includeTimestamp = true
	return l
}

func (l *Logger) ExcludeTimestamp() *Logger {
	l.includeTimestamp = false
	return l
}

func (l *Logger) timestamp() string {
	return time.Now().Format("01/02 15:04")
}

func (l *Logger) printTimestamp() {
	if l.includeTimestamp {
		fmt.Printf("[%s%s%s]", White, l.timestamp(), Reset)
	}
}

func (l *Logger) buildOutput(format string, args ...any) {
	if l.hasPrefix {
		format = "%s%s%s " + format
		args = append([]any{l.color, l.prefix, Reset}, args...)
	}

	if l.includeTimestamp {
		format = "[%s%s%s] " + format
		args = append([]any{White, l.timestamp(), Reset}, args...)
	}

	fmt.Printf(format, args...)
}

func (l *Logger) Basicf(format string, args ...any) {
	// if l.hasPrefix {
	// 	fmt.Printf("%s%s%s ", l.color, l.prefix, Reset)
	// }

	// l.printTimestamp()

	// fmt.Printf("%s[>]%s %s", Cyan, Reset, fmt.Sprintf(format, args...))
	l.buildOutput("%s[>]%s %s", Cyan, Reset, fmt.Sprintf(format, args...))
}

func (l *Logger) Basic(message string) {
	l.Basicf("%s\n", message)
}

func (l *Logger) Statusf(format string, args ...any) {
	// if l.hasPrefix {
	// 	fmt.Printf("%s%s%s ", l.color, l.prefix, Reset)
	// }

	// l.printTimestamp()

	// fmt.Printf("%s[@]%s %s", Magenta, Reset, fmt.Sprintf(format, args...))
	l.buildOutput("%s[@]%s %s", Magenta, Reset, fmt.Sprintf(format, args...))
}

func (l *Logger) Status(message string) {
	l.Statusf("%s\n", message)
}

func (l *Logger) Errorf(format string, args ...any) {
	// if l.hasPrefix {
	// 	fmt.Printf("%s%s%s ", l.color, l.prefix, Reset)
	// }

	// l.printTimestamp()

	// fmt.Printf("%s[!]%s %s", BoldRed, Reset, fmt.Sprintf(format, args...))
	l.buildOutput("%s[!]%s %s", BoldRed, Reset, fmt.Sprintf(format, args...))
}

func (l *Logger) Error(message string) {
	l.Errorf("%s\n", message)
}

func (l *Logger) Importantf(format string, args ...any) {
	// if l.hasPrefix {
	// 	fmt.Printf("%s%s%s ", l.color, l.prefix, Reset)
	// }

	// l.printTimestamp()

	// fmt.Printf("%s[#]%s %s%s%s", BoldRed, Reset, Bold, fmt.Sprintf(format, args...), Reset)
	l.buildOutput("%s[#]%s %s%s%s", BoldRed, Reset, Bold, fmt.Sprintf(format, args...), Reset)
}

func (l *Logger) Important(message string) {
	l.Importantf("%s\n", message)
}

func (l *Logger) Successf(format string, args ...any) {
	// if l.hasPrefix {
	// 	fmt.Printf("%s%s%s ", l.color, l.prefix, Reset)
	// }

	// l.printTimestamp()

	// fmt.Printf("%s[+]%s %s", Green, Reset, fmt.Sprintf(format, args...))
	l.buildOutput("%s[+]%s %s", Green, Reset, fmt.Sprintf(format, args...))
}

func (l *Logger) Success(message string) {
	l.Successf("%s\n", message)
}

func (l *Logger) Warningf(format string, args ...any) {
	// if l.hasPrefix {
	// 	fmt.Printf("%s%s%s ", l.color, l.prefix, Reset)
	// }

	// l.printTimestamp()

	// fmt.Printf("%s[-]%s %s", Yellow, Reset, fmt.Sprintf(format, args...))
	l.buildOutput("%s[-]%s %s", Yellow, Reset, fmt.Sprintf(format, args...))
}

func (l *Logger) Warning(message string) {
	l.Warningf("%s\n", message)
}

func (l *Logger) Queryf(format string, args ...any) {
	// if l.hasPrefix {
	// 	fmt.Printf("%s%s%s ", l.color, l.prefix, Reset)
	// }

	// l.printTimestamp()

	// fmt.Printf("%s[?]%s %s ", Blue, Reset, fmt.Sprintf(format, args...))
	l.buildOutput("%s[?]%s %s ", Blue, Reset, fmt.Sprintf(format, args...))
}

func (l *Logger) Query(message string) {
	l.Queryf("%s", message)
}

func (l *Logger) Customf(format string, color ColorCode, args ...any) {
	// if l.hasPrefix {
	// 	fmt.Printf("%s%s%s ", l.color, l.prefix, Reset)
	// }

	// l.printTimestamp()

	// fmt.Printf("%s%s%s", color, Reset, fmt.Sprintf(format, args...))
	l.buildOutput("%s%s%s", color, Reset, fmt.Sprintf(format, args...))
}

func (l *Logger) Custom(message string, color ColorCode) {
	l.Customf("%s\n", color, message)
}
