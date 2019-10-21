package main

import "github.com/01-edu/z01"

func main() {
	Raid1b(1,3)
}
func Raid1b(x,y int){
	if x>0 && y>0{
		if 0<y && y<3{
			for j:=0;j<y;j++{
				SecondExample(x,y)
				z01.PrintRune(10)			
			}
		}else{
			for j:=1;j<=y;j++{
				if j==1 || j==y{
					SecondLastCheckY(x,j)
					z01.PrintRune(10)		
				}else{
					MiddleCheckY(x)
					z01.PrintRune(10)	
				}
			}
		}
	}
}
func SecondExample(x,y int){
	if 0<x && x<3{
		for i:=0;i<x;i++ {
		z01.PrintRune(47)
		}
	}else{
		for i:=1;i<=x;i++ {
			if i==1 && i==y{
			z01.PrintRune(47)
			} else if i==x {
			z01.PrintRune(92)
			}else if i==1 && i==x {
			z01.PrintRune(92)
			}else if i==1{
			z01.PrintRune(47)
			}else{
			z01.PrintRune(42)
			}
		}
	}
}
func MiddleCheckY(x int){
	for i:=1; i<=x;i++{
		if i==1 || i==x{
			z01.PrintRune(42)
		}else{
			z01.PrintRune(32)
		}
	}
}
func SecondLastCheckY(x,y int){
	for i:=1; i<=x;i++{
		if i==1 && y!=1{
			z01.PrintRune(92)
		}else if i==x && y!=1{
			z01.PrintRune(47)
		}else if i==1 && y==1{
			z01.PrintRune(47)
		}else if i==x && y==1{
			z01.PrintRune(92)
		}else{
			z01.PrintRune(42)
		}
	}
}

