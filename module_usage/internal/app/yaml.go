package app

type File struct {
	Filename  string `yaml:"path"`
	Substring string `yaml:"substring"`
}
