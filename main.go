package main

import (
    "bufio"
    "fmt"
	"os"
	"strings"
)
type workspace struct {
	s map[string]string
	i map[string]int
	d map[string]float64
}
func init_workspace() workspace {
	s0 := make(map[string]string)
	i0 := make(map[string]int)
	d0 := make(map[string]float64)
	var ws = workspace{
		s: s0,
		i: i0,
		d: d0,
	}
	return ws
}
func declaration(name, val string, ws workspace) {
	ws.s[name] = val
}
func handle_input(input string, ws workspace) (string, workspace) {
	var ret string
	switch {
		case input == "q":
			os.Exit(3)
		case strings.Contains(input, "="):
			proposed := strings.Split(input, "=")
			if len(proposed) > 2{
				ret = "Bad variable declaration!"
			} else {
				declaration(proposed[0], proposed[1], ws)
			}
		default:
			ret = ws.s[input]
		}
	return ret, ws
}
func main(){
	ws := init_workspace()
    reader := bufio.NewReader(os.Stdin)
    for{
        fmt.Print(">> ")
        input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\n")

		output, ws1 := handle_input(input, ws)
		fmt.Println(output)

		ws = ws1
    }

}