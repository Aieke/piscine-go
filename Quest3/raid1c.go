package main

import "github.com/01-edu/z01"

func main() {
	Raid1c(5,3)
}
func Raid1c(x,y int){
	if x>0 && y>0{
		if 0<y && y<3{
			for j:=0;j<y;j++{
				ThirdExample(x,y)
				z01.PrintRune(10)
			}
		}else{
			for j:=1;j<=y;j++{
				if j==1 || j==y{
					ThirdCheckY(x,j)
					z01.PrintRune(10)
				}else{
					ThirdMiddleCheckY(x)
					z01.PrintRune(10)
				}
			}
		}
	}
}
func ThirdExample(x,y int){	
	if 0<x && x<3{
		for i:=0;i<x;i++{
			z01.PrintRune('A')
		}		
	}else{
		for i:=1;i<=x;i++{
			if i==1 || i==x{
				z01.PrintRune('A')
			}else{
				z01.PrintRune('B')
			}
		}
	}
}
func ThirdMiddleCheckY(x int){
	for i:=1;i<=x;i++{
		if i==1 || i==x{
			z01.PrintRune('B')
		}else{
			z01.PrintRune(32)
		}
	}
}
func ThirdCheckY(x,y int){
	for i:=1; i<=x; i++{
		if i==1 && y!=1{
			z01.PrintRune('C')
		}else if i==x && y!=1{
			z01.PrintRune('C')
		}else if i==1 && y==1{
			z01.PrintRune('A')
		}else if i==x && y==1{
			z01.PrintRune('A')
		}else{
			z01.PrintRune('B')
		}
	}
}