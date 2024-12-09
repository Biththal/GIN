package services

type NotesService struct {
}

type Note struct {
	Id   int
	Name string
}

func (n *NotesService) GetNotesService() []Note {
	data := []Note{
		{
			Id:   1,
			Name: "Biththal",
		},
		{
			Id:   2,
			Name: "Success",
		},
	}
	return data
}

func (n *NotesService) CreatePostService() []Note {
	data := []Note{
		{
			Id:   1,
			Name: "Biththal",
		},
	}
	return data
}
