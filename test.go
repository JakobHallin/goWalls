package paths //cannot run go run test.go if i remove package main 
//https://forum.golangbridge.org/t/what-is-a-package-in-golang-official-explanation-glossary/27661/2
//so they need to be the same and when they are the same they can use the other ones function
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

func generateGridOpen(rows, columns int) [][]Door{
        grid := make([][]Door, rows) //https://go.dev/tour/moretypes/13
        for y := 0; y < rows; y++{
               grid[y] = make([]Door, columns)
               for x :=0; x< columns; x++{
                grid[y][x] = Door{
                       row: y,
                       column: x,
                       wall: false,
                       }
               }
        }
       //now i have the grid i could assign the paths
        for y:= 0; y < rows; y++{
               for x:=0; x< columns-1; x++{
                grid[y][x].right= &grid[y][x+1]
               if y < rows - 1 { grid[y][x].down= &grid[y+1][x+1]}
                if y > 0 { grid[y][x].up= &grid[y-1][x+1]}
               }
         }

        return grid
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
        }
 }
//now i have the grid i could assign the paths
for y:= 0; y < rows; y++{
        for x:=0; x< columns-1; x++{
         //grid[y][x].wall = false;
          grid[y][x].right= &grid[y][x+1] 
        if y < rows - 1 { grid[y][x].down= &grid[y+1][x+1]}
         if y > 0 { grid[y][x].up= &grid[y-1][x+1]}
        }
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
        newPath := append([]*Door{},path...) // ... inte helt hundra hur detta funkar men antar att det betyder att jag kan ha nill eller ett värde https://go.dev/ref/spec#Passing_arguments_to_..._parameters 
        //vill kunna ta in all olika vägar antar path.. gör så jag kan gå igenom slice(array) 
        //curent is the curent dor Iam on so this will make sure can later back track the paths i have gone
        newPath = append(newPath, current)
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
        //its xbyy but i start with 0 so i need to do -1

        grid:= generateGridOpen(row,cols)
        //makePaths(&grid[0][0]) 
        //now its always 0,0 need to make it for all
        //not the best in this case
        
        var allPaths[][]int //will return this
        //need to loop if i have more cols so start at 0,1,2,3 and so on if i have more rows
        for y:=0; y<row; y++ { 
          tracepathsSave(&grid[y][0], nil, &allPaths, cols-1)
        }
 return allPaths 
}

// used to creat oaths if the grid starts as walls
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

//testfunction
/*
func TestPathEnumeration(t *testing.T) {
	cases := []struct{ cols, rows, paths int }{
		{0, 0, 0},
		{1, 1, 1},
		{1, 2, 2},
		{2, 1, 1},
		{2, 2, 4},
		{3, 3, 17},
		{5, 3, 99},
		{5, 4, 178},
		{5, 5, 259},
		{6, 4, 466},
	}
	for _, test := range cases {
		var res [][]int
		res = ConnectedPaths(test.cols, test.rows)
		if len(res) != test.paths {
			t.Errorf("wrong number of paths for grid %dx%d, expected %d, got %d",
				test.cols, test.rows, test.paths, len(res))
		}
	}
}
*/
/* func main() {
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
var res [][]int
res = ConnectedPaths(2,2)
fmt.Println(len(res))
 
//TestPathEnumeration(grid)
}
*/


