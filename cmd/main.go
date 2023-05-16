package main
type options struct {
	week string
	help bool
	version bool
}
func buildOptions(args []string) (*options, *flag.FlagSet){
	opts := &options{}
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() {fmt.Println(helpMessage(args[0]))}
	flags.StringVarP(&opts.week, "week", "w", "",  "このオプションはその週の天気予報を出力します．")
	flags.BoolVarP(&opts.help, "help", "h", false, "このメッセージを出力して終了する")
	flags.BoolVarP(&opts.version, "version", "v", false, "バージョンを出力して終了する")
	return opts, flags
}

func perform(opts *options, args []string) *cwfError { 
	fmt.Println("Hello World")
	return nil
}
func parseOptions(args []string) (*options, []string, *cwfError) { 
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	if opts.help {
		fmt.Println(helpMessage(args[0]))
		return nil, nil, &cwfError{statusCode: 0, message: ""}
	}
	if opts.token == "" {
		return nil, nil,
		&cwfError{statusCode: 3, message: "no token was given"} 
	}
	return opts, flags.Args(), nil
}
func goMain(args []string){
	opts, args, err := parseOptions(args)
	if err != nil {
		if err.statusCode != 0{
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
	