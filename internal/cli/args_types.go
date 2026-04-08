package cli

type Dest struct {
	Loc string
}

type Source struct {
	Loc   string
	Name  string
	IsDir bool
}

type Paths struct {
	Src  []Source
	Dest Dest
}
