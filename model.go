package main

type Model struct {
	contents []string
}

func InitModel() *Model {
	return &Model{
		contents: []string{
			"1:hogehoge",
			"2:hogehoge",
			"3:hogehoge",
			"4:hogehoge",
			"5:hogehoge",
			"6:hogehoge",
		},
	}
}
