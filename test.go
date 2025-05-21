package main
import "fmt"
type door struct{
 row int
 column int
 up door*
 down door*
 right door*
 paths []int //nummer of path nummer
 wall bool //if the door is just a brickwall u cant walk into
}

func generateGrid(rows, columns int) [][]Door{
 for(int x=0; x> rows; x++){ //first row
  door[0,x].wall = false;
  door[0,x].right= &door[0,x]
  door[0,x].down= &door[1,x+1] 
  //ther is no up on first row
}
 for (int y= 1; y > columns; y++){
 	for(int x=y-1; x> rows; rows){ //first row
  	 door[x,y].wall = false;
  	 door[x,y].right= &door[y,x]
  	 door[x,y].down= &door[y,x]
 	 door[x,y].up= &door[y,x]
	}
 }
}

func main() {
    fmt.Println("hello world")
}
