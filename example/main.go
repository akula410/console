package main

import (
	"console"
	"fmt"
)

var testAttr = []*console.Attr{
	&console.Attr{
		Name:"-param1",
		Default:"1",
	},
	&console.Attr{
		Name:"-param2",
		Default:"2",
	},
	&console.Attr{
		Name:"-param3",
		Default:"false",
	},
}

var testCallback = func(attr map[string]string){
	for param, value := range attr{
		fmt.Println("param = ", param, "; value = ", value)
	}
}

func main() {
	var cmd = console.Exec{
		Command: map[string]*console.Cmd{
			"test":&console.Cmd{
				Callback:testCallback,
				Attr: testAttr,
				Async:true,
			},
		},
		Sleep:3,
	}
	cmd.Response()
	/** Console
	$ test -param1="Hello world" -param3=true
		param =  -param1 ; value =  Hello world
		param =  -param2 ; value =  2
		param =  -param3 ; value =  true
	*/
}
