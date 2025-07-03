package storage

type CreateMark struct {
	Mark        string
	Description string
	Color       string
	Icon        string
}

type Mark struct {
	CreateMark
	Id        int
	Sort      int
	CreatedAt string
	ModifyAt  string
}

type FileMark struct {
	ID        int
	Dir       string
	FilePath  string
	Marks     []string
	Sha256    string
	CreatedAt string
	ModifyAt  string
}
