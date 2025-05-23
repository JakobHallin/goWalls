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
  //grid[0][x].wall = false;
  grid[0][x].right= &grid[0][x+1]
  grid[0][x].down= &grid[1][x+1] 
  //ther is no up on first row
 }
 for y:= 1; y < rows - 1; y++{
        for x:=y-1; x< columns-1; x++{
         //grid[y][x].wall = false;
         grid[y][x].right= &grid[y][x+1]
         grid[y][x].down= &grid[y+1][x+1]
         grid[y][x].up= &grid[y-1][x+1]
        }
  }
//last row
  lastrow := rows - 1
  for x:= 0; x< columns - 1; x++{
  	//grid[lastrow][x].wall = false;
         grid[lastrow][x].right= &grid[lastrow][x+1]
         grid[lastrow][x].up= &grid[lastrow-1][x+1]
 }

 return grid
}
//start  point could be ?,? in the grid
//proposed idea is to use dfs
//func findPath(startX int , startyY int,grid [][]Door){
//need to know the path number and the que and the priority is up then right then down
func tracepaths(curent *Door, counter *int, column int){
	//if curent door is nill or a wall
	if curent == nil || curent.wall {
        return
	}
	//add curent path
        //path = append(path, current)
	if curent.column == column{
		*counter = *counter + 1	
		curent.paths = append(curent.paths, *counter)
		return
	}

 tracepaths(curent.up, counter, column)
 tracepaths(curent.right, counter, column)
 tracepaths(curent.down, counter, column)
}
//will use kinda same logic but not with counter
func tracepathsSave(current *Door, path []*Door, allPaths *[][]int, column int){
	//if curent door is nill or a wall
	if current == nil || current.wall {
        return
	}
        //
        newPath := append([]*Door{},path...) // ... inte helt hundra hur detta funkar men antar att det betyder att jag kan ha nill eller ett värde https://go.dev/ref/spec#Passing_arguments_to_..._parameters 
        //vill kunna ta in all olika vägar antar path.. gör så jag kan gå igenom slice(array) 
        //curent is the curent dor Iam on so this will make sure can later back track the paths i have gone
        newPath = append(newPath, current)

	//add curent path
        //path = append(path, current)
	if current.column == column{
                //now im on the last cant go more right
                var pathNumber []int
                for _, test := range newPath {
                        pathNumber = append(pathNumber , test.row*column+test.column)
                }
		*allPaths = append(*allPaths, pathNumber)
                return
	}

        tracepathsSave(current.up, path, allPaths, column)
        tracepathsSave(current.right, path, allPaths, column)
        tracepathsSave(current.down, path, allPaths, column)

}
func ConnectedPaths(cols int , row int)[][]int{ 
        //empty
        if cols == 0 || row == 0 {
		return [][]int{}
	}
        grid:= generateGrid(row,cols)
        makePaths(&grid[0][0])
        
        var allPaths[][]int //will return this
        tracepathsSave(&grid[0][0], nil, &allPaths, cols)

 return allPaths 
}

//}
func makePaths(start *Door){

	if start == nil {
	return
	}
	start.wall= false
	makePaths(start.right)
	makePaths(start.down)
        makePaths(start.up)
}



func printGrid(grid [][]Door){
 for y:= 0; y<len(grid); y++ { //loop over rows
    for x:=0; x<len(grid[0]); x++ { //lop over columns
     	if grid[y][x].wall {
	 fmt.Print("X ")
    	} else {
	 fmt.Print("0 ")
	}
    }
    fmt.Println("")
  }
}

func main() {
    fmt.Println("hello world")
    grid:= generateGrid(4,4)
    fmt.Println(grid[1][1].down.wall)
    fmt.Println(grid[1][1].down)

    fmt.Println(grid[2][2].down)
    fmt.Println(grid[2][2].down)

    fmt.Println(grid[2][2].down.wall)
fmt.Println(grid[0][2].down)
fmt.Println(grid[3][2].down)
fmt.Println(grid[0][0].down)
fmt.Println(grid[0][0].up)


    fmt.Println(" ")
 makePaths(&grid[0][0])
  printGrid(grid)
//works but counter is amount of steps
counter := 0
tracepaths(&grid[0][0], &counter, 2)
fmt.Println(grid[0][1])

    fmt.Println(" ")

counter2 := 0
tracepaths(&grid[0][0], &counter2, 1)
fmt.Println(grid[0][2])
fmt.Println(grid[0][1])
    fmt.Println(" ")

counter3 := 0
tracepaths(&grid[0][0], &counter3, 3)
fmt.Println(grid[0][3])
fmt.Println(grid[0][2])
fmt.Println(grid[0][1])
fmt.Println(grid[0][0])
 ConnectedPaths(1,4)
}



