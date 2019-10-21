package main

import "github.com/01-edu/z01"

func main() {
	Raid1a(1,2)
	z01.PrintRune(10)
}
func Raid1a(x,y int){
	if x>0 && y>0{
		if 0<y && y<3{
			for j:=0; j<y;j++{
				FirstExample(x,y)
			}
		}else{
			for j:=1; j<=y;j++{
				if j==1 || j==y{
				FirstExample(x,y)
				z01.PrintRune(10)
				}else{
					MiddleCheck(x)
					z01.PrintRune(10)
				}
				
			}
		}
	}
}
func FirstExample(x,y int){
	if 0<x && x<3{
		for i:=0;i<x;i++{
			z01.PrintRune('o')
		}
	}else{
		for i:=1;i<=x;i++{
			if i==1 || i==x{
			z01.PrintRune('o')
			}else if i==1{ 
			z01.PrintRune('o')
			}else{
			z01.PrintRune('-')
			}
		}
	}
}
func MiddleCheck(x int){
	for i:=1;i<=x;i++{
		if i==1||i==x{
			z01.PrintRune('|')
		}else{
			z01.PrintRune(32)
		}
	}
}
