package model

type Video struct {
	Aid   int64
	BV    string
	Title string
	Desc  string
	Up    Up
	Parts []VideoPart
}

type VideoPart struct {
	Cid   int64
	Title string
}

type Up struct {
	Mid  int64
	Name string
	Face string
}
