package main

import (
	"fmt"
	"image/png"
	"os"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	flag "github.com/spf13/pflag"
)

/*
VERSION of qrg
*/
const VERSION = "1.0.0"

type options struct {
	dest        string
	level       string
	size        int
	helpFlag    bool
	versionFlag bool
	args        []string
}

func helpMessage() string {
	return `qrg [OPTIONS] <CONTENT>
OPTIONS
    -d, --dest <DEST>      specifies destination.
                           if this option did not specified,
                           it generates to 'qrcode.png'.
    -l, --level <LEVEL>    specifies error correction level.
                           available values: L, M, Q, H.  default is M.
                           L: 7%, M: 15%, Q: 25%, H: 30%.
    -s, --size <SIZE>      specifies pixel size of QR code.  default is 256.
    -h, --help             print this message.
    -v, --version          print version.
CONTENT
    specifies content of QR code.`
}

func version() string {
	return fmt.Sprintf("qrg version %s", VERSION)
}

func buildFlagSet(args []string) (*flag.FlagSet, *options) {
	var flags = flag.NewFlagSet("qrg", flag.ContinueOnError)
	var opts = options{}
	flags.Usage = func() { fmt.Println(helpMessage()) }
	flags.StringVarP(&opts.dest, "dest", "d", "qrcode.png", "specifies the destination")
	flags.StringVarP(&opts.level, "level", "l", "M", "specifies error correction level.")
	flags.IntVarP(&opts.size, "size", "s", 256, "specifies pixel size of QR code.")
	flags.BoolVarP(&opts.helpFlag, "help", "h", false, "print this messgae.")
	flags.BoolVarP(&opts.versionFlag, "version", "v", false, "print version.")
	return flags, &opts
}

func parseOptions(args []string) (*options, error) {
	var flag, opts = buildFlagSet(args)
	if err := flag.Parse(args); err != nil {
		return nil, err
	}
	opts.args = flag.Args()
	return opts, nil
}

func printError(err error) {
	fmt.Println(err.Error())
}

func validateLevel(level string) error {
	var template = "LMQH"
	var levelUpper = strings.ToUpper(level)
	if len(levelUpper) != 1 || !strings.Contains(template, levelUpper) {
		return fmt.Errorf("%s: not level character", level)
	}
	return nil
}

func validate(opts *options) error {
	if err := validateLevel(opts.level); err != nil {
		return err
	}
	if len(opts.args) == 1 {
		return fmt.Errorf("no arguments are specified")
	}
	return nil
}

func concat(args []string) string {
	var builder = new(strings.Builder)
	for _, arg := range args {
		fmt.Fprintf(builder, "%s ", arg)
	}
	return strings.TrimSpace(builder.String())
}

func toCorrectionLevel(level string) qr.ErrorCorrectionLevel {
	if level == "L" {
		return qr.L
	} else if level == "M" {
		return qr.M
	} else if level == "Q" {
		return qr.Q
	}
	return qr.H
}

func generateCode(content string, opts *options) (barcode.Barcode, error) {
	var level = toCorrectionLevel(opts.level)
	var qrcode, err = qr.Encode(content, level, qr.Auto)
	if err != nil {
		return nil, err
	}
	return barcode.Scale(qrcode, opts.size, opts.size)
}

func writeImage(qrcode barcode.Barcode, opts *options) error {
	var file, err = os.OpenFile(opts.dest, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	png.Encode(file, qrcode)
	return nil
}

func helpOrVersion(opts *options) bool {
	if opts.versionFlag {
		fmt.Printf("qrg version %s\n", VERSION)
	}
	if opts.helpFlag {
		fmt.Println(helpMessage())
	}
	return opts.versionFlag || opts.helpFlag
}

func perform(opts *options) error {
	if helpOrVersion(opts) {
		return nil
	}
	if err := validate(opts); err != nil {
		return err
	}
	var content = concat(opts.args)
	var qrcode, err = generateCode(content, opts)
	if err != nil {
		return err
	}
	return writeImage(qrcode, opts)
}

func goMain(args []string) int {
	var opts, err = parseOptions(args)
	if err != nil {
		printError(err)
		return 1
	}
	var err2 = perform(opts)
	if err2 != nil {
		printError(err2)
		return 1
	}
	return 0
}

func main() {
	var status = goMain(os.Args)
	os.Exit(status)
}
