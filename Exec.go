package console

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

type Exec struct {
	Command map[string]*Cmd
	Attr map[string]string
	Sleep time.Duration
}

type Cmd struct {
	Async bool
	Callback func(attr map[string]string)
	Attr []*Attr
}

type Attr struct {
	Name string
	Default string
}

func (e *Exec) Response(){
	flag.Parse()
	if cmd := e.getCmd(); cmd != nil {
		var attrs = e.getAttr(cmd)
		if cmd.Async {
			go cmd.Callback(attrs)
			time.Sleep(time.Second*e.Sleep)
		}else{
			cmd.Callback(attrs)
		}
	}
}

func (e *Exec) getCmd() *Cmd{
	var cmd *Cmd

	if e.Command != nil{
		if flag.NArg() > 0 {
			for i:=0;i<flag.NArg();i++{
				f := flag.Arg(i)
				for n, v := range e.Command{
					if n == f {
						cmd = v
					}
				}
				e.setAttr(f)
			}
		}else{
			fmt.Println("Not enough arguments")
		}
	}else{
		fmt.Println("Are you not added command (Exec.Command)")
	}

	return cmd
}

func (e *Exec) setAttr(str string){
	var attrs = strings.Split(str, "=")
	var name string
	var value string

	for i, v := range attrs{
		if i == 0 {
			name = v
		}else{
			value += v
		}
	}

	if e.Attr == nil {
		e.Attr = make(map[string]string)
	}
	e.Attr[name] = value
}

func (e *Exec) getAttr(cmd *Cmd)map[string]string{
	var attrs = make(map[string]string)
	if cmd.Attr != nil {
		for _, attr := range cmd.Attr{
			if result, ok := e.Attr[attr.Name]; ok {
				attrs[attr.Name] = result
			}else{
				attrs[attr.Name] = attr.Default
			}
		}
	}
	return attrs
}
