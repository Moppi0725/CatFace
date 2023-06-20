package main

import (
	"fmt"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
	//"github.com/moppi0725/cwf"
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

/*
func (opts *options) mode(args []string) cwf.Mode {
	switch {
	case opts.runOpt.qrcode != "":
		return urleap.QRCode
	default:
		return urleap.Shorten
	}
}
*/
//適宜修正↑

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

// ↓適宜修正
// func shortenEach(bitly *urleap.Bitly, config *urleap.Config, url string) error {
// 	result, err := bitly.Shorten(config, url)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(result)
// 	return nil
// }

// func deleteEach(bitly *urleap.Bitly, config *urleap.Config, url string) error {
// 	return bitly.Delete(config, url)
// }

// func listUrls(bitly *urleap.Bitly, config *urleap.Config) error {
// 	urls, err := bitly.List(config)
// 	if err != nil {
// 		return err
// 	}
// 	for _, url := range urls {
// 		fmt.Println(url)
// 	}
// 	return nil
// }

// func listGroups(bitly *urleap.Bitly, config *urleap.Config) error {
// 	groups, err := bitly.Groups(config)
// 	if err != nil {
// 		return err
// 	}
// 	for i, group := range groups {
// 		fmt.Printf("GUID[%d] %s\n", i, group.Guid)
// 	}
// 	return nil
// }

// func performImpl(args []string, executor func(url string) error) *UrleapError {
// 	for _, url := range args {
// 		err := executor(url)
// 		if err != nil {
// 			return makeError(err, 3)
// 		}
// 	}
// 	return nil
// }

// 5/30ここの修正
// func perform(opts *options, args []string) *UrleapError {
// 	bitly := urleap.NewBitly(opts.runOpt.group)
// 	config := urleap.NewConfig(opts.runOpt.config, opts.mode(args))
// 	config.Token = opts.runOpt.token
// 	switch config.RunMode {
// 	case urleap.List:
// 		err := listUrls(bitly, config)
// 		return makeError(err, 1)
// 	case urleap.ListGroup:
// 		err := listGroups(bitly, config)
// 		return makeError(err, 2)
// 	case urleap.Delete:
// 		return performImpl(args, func(url string) error {
// 			return deleteEach(bitly, config, url)
// 		})
// 	case urleap.Shorten:
// 		return performImpl(args, func(url string) error {
// 			return shortenEach(bitly, config, url)
// 		})
// 	}
// 	return nil
// }

//↑適宜修正

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
	//opts, args, err := parseOptions(args)
	_, args, err := parseOptions(args)

	if err != nil {
		if err.statusCode != 0 {
			fmt.Println(err.Error())
		}
		return err.statusCode
	}
	// if err := perform(opts, args); err != nil {
	// 	fmt.Println(err.Error())
	// 	return err.statusCode
	// }
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
