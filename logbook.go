package main

type Logbook struct {
	name string
}

func LoadDiary(path string) (*Logbook, error) {
	return &Logbook{"test"}, nil
}
