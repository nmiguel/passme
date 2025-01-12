package flag

type Flag struct {
	Alias    []string
	Tooltip  string
	Callback func([]string)
}
