//Cоздать программу (простую) в которой будешь заполнять данными структурус 2мя полями:
//	address (country, city, street), info (fio, job, age, birthday).
//	Поле заполнение должно быть через 3 функции:
//	первая заполняет всю структуру,
//    внутри вызывает 2 функции - одна заполняет address, вторая поле info.

package main

import "fmt"

type Persona struct {
	address Address
	info    Info
}
type Address struct {
	country string
	city    string
	street  string
}
type Info struct {
	fio      string
	job      string
	age      int
	birthday int
}

func fullPersona(p *Persona) {
	infoAddress(&p.address)
	infoInfo(&p.info)
}
func infoInfo(inf *Info) {
	inf.fio = ""
	inf.job = ""
	inf.age = 0
	inf.birthday = 0

}
func infoAddress(addr *Address) {
	addr.country = ""
	addr.city = ""
	addr.street = ""
}

func main() {
	var serghei Persona
	fullPersona(&serghei)
	infoAddress(&serghei.address)
	fmt.Println(serghei)
}
