package main

type snippetCmd struct {
	Add    addSnippetCmd
	Get    getSnippetCmd
	Update updateSnippetCmd
	Delete deleteSnippetCmd
	List   listSnippetCmd
}

type addSnippetCmd struct{}
type getSnippetCmd struct{}
type updateSnippetCmd struct{}
type deleteSnippetCmd struct{}
type listSnippetCmd struct{}

func (a *addSnippetCmd) Execute([]string) error {
	return nil
}
func (g *getSnippetCmd) Execute([]string) error {
	return nil
}
func (u *updateSnippetCmd) Execute([]string) error {
	return nil
}
func (d *deleteSnippetCmd) Execute([]string) error {
	return nil
}
func (l *listSnippetCmd) Execute([]string) error {
	return nil
}
