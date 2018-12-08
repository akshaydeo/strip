package main

type arrayFlags []string

func (i *arrayFlags) String() string {
	str := ""
	for _, v := range *i {
		str += v
	}
	return str
}

func (i *arrayFlags) Get() []string {
	return *i
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}
