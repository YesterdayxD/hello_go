package main

import "fmt"
//属性
type Attribution struct{
	attack int
	defence int

}
//定义英雄

type Hero struct {
	attribution Attribution
	name string
	skill Skill

}

//
type Skill interface {
	SkillQ()
	SkillW()
}
func (h *Hero)SkillQ(){
	fmt.Println("use Q")
}
func (h *Hero)SkillW(){
	fmt.Println("use W")
}

func main()  {
	var _ Skill=&Hero{}
}
