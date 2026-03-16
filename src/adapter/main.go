package main

import "fmt"

type Banner struct {
	string string
}

func (b *Banner) ShowWithParen() {
	fmt.Println("(" + b.string + ")")}

func (b *Banner) ShowWithAster() {
	fmt.Println("*" + b.string + "*")
}

type Print interface {
	PrintWeak()
	PrintStrong()
}

type PrintBanner struct {
	banner *Banner
}

func (p *PrintBanner) PrintWeak() {
	p.banner.ShowWithParen()
}

func (p *PrintBanner) PrintStrong() {
	p.banner.ShowWithAster()
}

func main() {
	var p Print
	p = &PrintBanner{&Banner{"Hello"}}
	p.PrintWeak()
	p.PrintStrong()
}
