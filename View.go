package main

import (
	"text/template"
	"container/list"
	"bytes"
	//"fmt"
)

const (
	CLIENT = "Client"
	SERVER = "Server"
)

type HandshakeProtocolStep struct {
	Actor       string
	Description string
	Success     bool
	Optional    bool
	//Other attributes
}

type StepGroup struct {
	Actor       string
	Steps       []HandshakeProtocolStep
}


func createStepGroups(steps []HandshakeProtocolStep)([]StepGroup) {
	groups := make([]StepGroup, 0)
	var pre HandshakeProtocolStep = steps[0]
	var currGroup StepGroup = StepGroup{Actor: pre.Actor, Steps: make([]HandshakeProtocolStep,0)}
	for _, step := range steps {
		if (step.Actor != pre.Actor) {
			groups = append(groups, currGroup)
			currGroup = StepGroup{Actor: step.Actor, Steps: make([]HandshakeProtocolStep,0)}
		}
		currGroup.Steps = append(currGroup.Steps, step)
		pre = step
	}
	groups = append(groups, currGroup)
	return groups
}

func Visualize(data list.List, format string)([]byte) {
	groups := getViewDataModel()
	//var result []byte
	output := new(bytes.Buffer)
	switch (format) {
		case "txt": {
			tmpl, err := template.ParseFiles("template/txt/HandshakeProtocolDetails.txt")
			if err != nil { panic(err) }
			err = tmpl.Execute(output, groups)
			if err != nil { panic(err) }
		}
	}
	return output.Bytes()
}

//For now only test data
func getViewDataModel()([]StepGroup) {
	steps := make([]HandshakeProtocolStep,0)
	step1 := HandshakeProtocolStep { Actor: CLIENT, Description: "Client Hello"}
	step2 := HandshakeProtocolStep { Actor: SERVER, Description: "Server Hello"}
	step3 := HandshakeProtocolStep { Actor: SERVER, Description: "Server Certificate"}
	step4 := HandshakeProtocolStep { Actor: SERVER, Description: "Server Hello Done"}
	steps = append(steps,step1)
	steps = append(steps,step2)
	steps = append(steps,step3)
	steps = append(steps,step4)


	groups := createStepGroups(steps)
	return groups
}

//test only
//func main() {
//	bytes := Visualize(list.List{}, "txt")
//	fmt.Println(string(bytes))
//}

//func main() {
//	steps := make([]HandshakeProtocolStep,0)
//	step1 := HandshakeProtocolStep { Actor: CLIENT, Description: "Client Hello"}
//	step2 := HandshakeProtocolStep { Actor: SERVER, Description: "Server Hello"}
//	step3 := HandshakeProtocolStep { Actor: SERVER, Description: "Server Certificate"}
//	step4 := HandshakeProtocolStep { Actor: SERVER, Description: "Server Hello Done"}
//	steps = append(steps,step1)
//	steps = append(steps,step2)
//	steps = append(steps,step3)
//	steps = append(steps,step4)
//
//
//	groups := createStepGroups(steps)
//
//	//fmt.Println("There was an error:", steps)
//	tmpl, err := template.ParseFiles("template/txt/HandshakeProtocolDetails.txt")
//	if err != nil {
//		panic(err)
//	}
//	//tmpl, _ := template.New("detailsteps").Parse("{{.Actor}} items are made of {{.Description}}")
//	err = tmpl.Execute(os.Stdout, groups)
//	if err != nil {
//		panic(err)
//	}
//}