package logger

var (
	defaultFilename   = "output.log"
	defaultLevel      = levelInfo
	defaultSaveToFile = false

	defaultEncoding = formatConsole
)

type OptionFn func(o *option)
type FileOptionFn func(o *fileOption)

type option struct {
	level      string
	encoding   string
	needToSave bool

	fileOption *fileOption
}

// default option
func defaultOption() *option {
	return &option{
		level:      defaultLevel,
		encoding:   defaultEncoding,
		needToSave: defaultSaveToFile,
	}
}

// apply the options
func (o *option) apply(opts ...OptionFn) {
	for _, opt := range opts {
		opt(o)
	}
}

// WithLevel setting the log level
func WithLevel(levelName string) OptionFn {
	return func(opt *option) {
		switch levelName {
		case levelDebug, levelInfo, levelWarn, levelError:
			opt.level = levelName
		default:
			opt.level = defaultLevel
		}
	}
}

// WithEncoding set the output log format, console or json
func WithEncoding(format string) OptionFn {
	return func(opt *option) {
		switch format {
		case formatJSON, formatConsole:
			opt.encoding = format
		default:
			opt.encoding = defaultEncoding
		}
	}
}

// WithSaveToFile save log to file
func WithSaveToFile(save bool, fileOpt ...FileOptionFn) OptionFn {
	return func(opt *option) {
		if save {
			opt.needToSave = true
			fo := defaultFileOption()
			fo.apply(fileOpt...)
			opt.fileOption = fo
		}
	}
}

// fileOption set the file options.
type fileOption struct {
	filename      string
	isCompression bool
}

// defaultFileOption set the default file options.
func defaultFileOption() *fileOption {
	return &fileOption{
		filename:      defaultFilename,
		isCompression: false,
	}
}

// apply file options.
func (o *fileOption) apply(opts ...FileOptionFn) {
	for _, opt := range opts {
		opt(o)
	}
}

// WithFilename set log filename
func WithFilename(filename string) FileOptionFn {
	return func(opt *fileOption) {
		opt.filename = filename
	}
}

// WithIsCompression set whether to compress log files
func WithIsCompression(isCompression bool) FileOptionFn {
	return func(opt *fileOption) {
		opt.isCompression = isCompression
	}
}
