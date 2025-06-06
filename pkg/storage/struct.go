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
