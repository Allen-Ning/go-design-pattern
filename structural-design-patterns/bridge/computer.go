package bridge

type computer interface {
	setPrinter(*printer)
	print()
}

type mac struct {
	printer printer
}

func (mac *mac) setPrinter(printer printer) {
	mac.printer = printer
}

func (mac *mac) print() {
	mac.printer.printFile()
}

type windows struct {
	printer printer
}

func (windows *windows) setPrinter(printer printer) {
	windows.printer = printer
}

func (windows *windows) print() {
	windows.printer.printFile()
}

type linux struct {
	printer printer
}

func (linux *linux) setPrinter(printer printer) {
	linux.printer = printer
}

func (linux *linux) print() {
	linux.printer.printFile()
}
