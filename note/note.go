package note

type Note struct {
	title   string
	content string
}

func New(title, content string) *Note {
	return &Note{
		title,
		content,
	}
}
