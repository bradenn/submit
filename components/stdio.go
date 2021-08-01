package components

type Time struct {
	Elapsed string `json:"elapsed"`
	User    string `json:"user"`
	System  string `json:"system"`
}

type StdStream struct {
	Input  []string
	Output []string
	Error  []string
	Exit   int
}

type StdIOContext struct {
	Name string
	Args []string
	Env  []string
	Std  StdStream
}

type StdIOResult struct {
	Passed   bool
	Std      StdStream
	Diff     StdStream
	Duration Time
}

type StdIO struct {
}
