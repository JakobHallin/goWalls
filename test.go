package main
import "fmt"
type Door struct{
 row int
 column int
 up *Door
 down *Door
 right *Door
 paths []int //nummer of path nummer
 wall bool //if the door is just a brickwall u cant walk into
}
func generateGrid(rows, columns int) [][]Door{
 grid := make([][]Door, rows) //https://go.dev/tour/moretypes/13
 for y := 0; y < rows; y++{
	grid[y] = make([]Door, columns)
        for x :=0; x< columns; x++{
	 grid[y][x] = Door{
		row: y,
		column: x,
		wall: true,
		}
        //cant assign beofre made the grid 
	//door[x,y].right= &door[y,x]
        // door[x,y].down= &door[y,x]
        // door[x,y].up= &door[y,x]
        }
 }
//now i have the grid i could assign the paths
 for x :=0; x< columns-1; x++{ //first row
  grid[0][x].wall = false;
  grid[0][x].right= &grid[0][x+1]
  grid[0][x].down= &grid[1][x+1] 
  //ther is no up on first row
}
 for y:= 1; y < rows-1; y++{
        for x:=y-1; x< columns-1; x++{
         grid[x][y].wall = false;
         grid[x][y].right= &grid[y][x+1]
         grid[x][y].down= &grid[y+1][x+1]
         grid[x][y].up= &grid[y-1][x+1]
        }
 }
return grid
}

func main() {
    fmt.Println("hello world")
    grid:= generateGrid(3,3)
    fmt.Println(grid[1][1].down.wall)
}

