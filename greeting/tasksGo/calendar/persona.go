package calendar

import "errors"

type Persona struct {
	name string
	age  int
}

func (p *Persona) SetName(name string) error {
	if len(name) > 10 {
		return errors.New("Error")
	}
	p.name = name

	return nil

}
func (p *Persona) Name() string {
	return p.name
}
