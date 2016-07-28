//print index and value of arguments, one per line
package main

import (
	"fmt"
)

type panel struct{
	x,y,width int;
	h int; // where the first vertical 
	e bool; //whether width-2 is even (extra column) 
	expanded,assigned bool;
}

var list []panel 

func main() {
	//list := []panel{panel{0,0,1024,false,false}}
	list = append(list,panel{0,0,9,0,false,false,false})

	// add kids to the list
	var j int;
	for {
		j = addKids()
		if j<0 { break }
	}

	for i:=0; i<len(list); i++{
		assign(i)	
	}

	for i:=0; i<len(list); i++{
		//fmt.Println("i=",i,"w=",list[i].width,"x=",list[i].x,"y=",list[i].y,"vert=",list[i].h,"extra?",list[i].e)
	}
}


func assign(i int){
	p := list[i]
	if !list[i].assigned { list[i].assigned=true }
	if list[i].width==1 {fmt.Println("i=",i,"paint: (",p.x,",",p.y,")"); return}
	if list[i].width==2 {fmt.Println("i=",i,"paint: (",p.x,",",p.y,")..4");return}
	if p.e { fmt.Println("i=",i,"paint: vert at x=",p.x+p.h," and x=",p.x+2*p.h)
	} else { fmt.Println("i=",i,"paint: vert at x=",p.x+p.h) }
}

func addKids() int {
	j := 0; L := len(list)
	for j<L && list[j].expanded { j++ }
	if j==L {return -1}
	// add kids to the list
	kids := quad(j)
	list[j].expanded=true
	if kids[0].width > 0  {
		for i:=0; i<4; i++{
			list = append(list,kids[i])
		}
	}
	return j	
}

func quad(j int) [4]panel{
	var out [4]panel
	p := list[j] // convenience
	if p.expanded || p.width<3 { 
		//fmt.Println("done width=",p.width," expanded=",p.expanded)
		return out
	}
	w := p.width
	if w%2==0 { w = (w-2)/2; list[j].e=true } else { w = (w-1)/2 }
	list[j].h = w+1 // place the vertical
	for i:=0; i<4; i++{
		out[i].x = p.x; out[i].y = p.y
		if i%2==1 { out[i].x = p.x + w + 1 } 
		if i>1 { out[i].y = p.y + w + 1 }
		out[i].width = w;
	}
	return out	
}
