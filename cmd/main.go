package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/moppi0725/cwf"
	flag "github.com/spf13/pflag"
)

const VERSION = "0.0.1"

func versionString(args []string) string {
	prog := "cwf"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf("%s version %s", prog, VERSION)
}

func helpMessage(args []string) string {
	prog := "cwf"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf(`%s cwf [オプション] <県名>

	オプションをつけずに実行した場合は，現在の天気を出力する
	
	オプション
		-w, --week    このオプションはその週の天気予報を出力します．
		-h, --help    このメッセージを出力して終了する
		-v, --version バージョンを出力して終了する
	変数
		県名    調べたい都道府県を指定する．`, prog)
}

// ↓適宜修正↓
type CwfError struct {
	statusCode int
	message    string
}

func (e CwfError) Error() string {
	return e.message
}

type flags struct {
	helpFlag    bool
	versionFlag bool
}

type runOpts struct {
	week string
}

/*
This struct holds the values of the options.
*/
type options struct {
	runOpt  *runOpts
	flagSet *flags
}

func newOptions() *options {
	return &options{runOpt: &runOpts{}, flagSet: &flags{}}
}

func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := newOptions()
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args)) }
	flags.StringVarP(&opts.runOpt.week, "week", "w", "", "このオプションはその週の天気予報を出力します．")
	flags.BoolVarP(&opts.flagSet.helpFlag, "help", "h", false, "このメッセージを出力して終了する")
	flags.BoolVarP(&opts.flagSet.versionFlag, "version", "v", false, "バージョンを出力して終了する")
	return opts, flags
}

func parseOptions(args []string) (*options, []string, *CwfError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	if opts.flagSet.helpFlag {
		fmt.Println(helpMessage(args))
		return nil, nil, &CwfError{statusCode: 0, message: ""}
	}
	if opts.flagSet.versionFlag {
		fmt.Println(versionString(args))
		return nil, nil, &CwfError{statusCode: 0, message: ""}
	}
	return opts, flags.Args(), nil
}

func perform(opts *options, args []string) *CwfError {
	switch {
	case opts.runOpt.week != "":
		city, err := cwf.GetCityInfo(opts.runOpt.week)
		fmt.Println(city, err)
		date, weathercode, err := cwf.MakeWeekUrl(city, "w")
		for i := range date {
			fmt.Printf("%s  %s\n", date[i], weathercode[i])
		}
		return nil

	default:
		city, err := cwf.GetCityInfo(args[0])
		fmt.Println(city, err)
		date, weathercode, err := cwf.MakedayUrl(city, "d")
		for i := range date {
			fmt.Printf("%s  %s\n", date[i], weathercode[i])
		}
		return nil
	}
}

func makeError(err error, status int) *CwfError {
	if err == nil {
		return nil
	}
	ue, ok := err.(*CwfError)
	if ok {
		return ue
	}
	return &CwfError{statusCode: status, message: err.Error()}
}

func goMain(args []string) int {
	opts, args, err := parseOptions(args)

	if err != nil {
		if err.statusCode != 0 {
			fmt.Println(err.Error())
		}
		return err.statusCode
	}
	if err := perform(opts, args); err != nil {
		fmt.Println(err.Error())
		return err.statusCode
	}
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
