package main

type person struct {
	first string
}

func changeFisrtP(p *person, s string) {
	p.first = s
}

func changeFisrt(p person, s string) person {
	p.first = s
	return p
}
func main() {
	p := person{
		first: "jaja",
	}
	changeFisrtP(&p, "papa")
	p = changeFisrt(p, "nana")

}
