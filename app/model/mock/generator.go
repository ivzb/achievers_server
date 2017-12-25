package mock

type Generator struct {
}

func New() *Generator {
	return &Generator{}
}

// func (g *Generator) Exists(string, string, string) (bool, error) {
// 	return mock.ExistsMock.B, mock.ExistsMock.E
// }
