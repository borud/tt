package main

type projectCmd struct {
	Add    addProjectCmd
	Get    getProjectCmd
	Update updateProjectCmd
	Delete deleteProjectCmd
	List   listProjectCmd
}

type addProjectCmd struct{}
type getProjectCmd struct{}
type updateProjectCmd struct{}
type deleteProjectCmd struct{}
type listProjectCmd struct{}

func (a *addProjectCmd) Execute([]string) error {
	return nil
}
func (g *getProjectCmd) Execute([]string) error {
	return nil
}
func (u *updateProjectCmd) Execute([]string) error {
	return nil
}
func (d *deleteProjectCmd) Execute([]string) error {
	return nil
}
func (l *listProjectCmd) Execute([]string) error {
	return nil
}
