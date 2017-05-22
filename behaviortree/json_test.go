package behaviortree

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

type funcCreateColor func() ColoredThing // 声明了一个函数类型
var gColorCreator map[string]funcCreateColor

func init() {
	gColorCreator = make(map[string]funcCreateColor, 0)
	AppendColorCreator("plant", func() ColoredThing { return new(Plant) })
	AppendColorCreator("animal", func() ColoredThing { return new(Animal) })
}

func AppendColorCreator(name string, f func() ColoredThing) {
	gColorCreator[name] = f
}

type HeheColor struct {
	Colors ColorfulEcosystem
}
type ColorfulEcosystem struct {
	Name   string
	Things []ColoredThing `json:"things"`
}

type ColoredThing interface {
	Color() string
	Name() string
}

type Plant struct {
	MyColor string `json:"color"`
	Ani     Animal
}

type Animal struct {
	MyColor string `json:"color"`
}

func (p *Plant) Color() string {
	return p.MyColor
}
func (p *Plant) Name() string {
	return "plant"
}

// func (p *Plant) MarshalJSON() (b []byte, e error) {
// 	return json.Marshal(map[string]string{
// 		"type":  "plant",
// 		"color": p.Color(),
// 	})
// }

func (a *Animal) Color() string {
	return a.MyColor
}
func (a *Animal) Name() string {
	return "animal"
}

// func (a *Animal) MarshalJSON() (b []byte, e error) {
// 	return json.Marshal(map[string]string{
// 		"type":  "animal",
// 		"color": a.Color(),
// 	})
// }

func (ce *ColorfulEcosystem) MarshalJSON() (b []byte, e error) {
	// ce.Things
	newThings := make([]string, 0)
	for k := range ce.Things {
		xb, _ := json.Marshal(ce.Things[k])
		xbs := strings.Replace(string(xb), "{", "{\"type\":\""+ce.Things[k].Name()+"\",", -1)
		newThings = append(newThings, xbs)
	}

	things := "["
	for k := range newThings {
		if k < len(newThings) {
			things += newThings[k] + ","
		} else {
			things += newThings[k]
		}
	}

	return json.Marshal(map[string]string{
		"things": things,
	})
}

func (ce *ColorfulEcosystem) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	var rawMessagesForColoredThings []*json.RawMessage
	err = json.Unmarshal(*objMap["things"], &rawMessagesForColoredThings)
	if err != nil {
		return err
	}

	ce.Things = make([]ColoredThing, len(rawMessagesForColoredThings))

	var m map[string]string
	for index, rawMessage := range rawMessagesForColoredThings {
		err = json.Unmarshal(*rawMessage, &m)
		if err != nil {
			return err
		}

		p := gColorCreator[m["type"]]()
		err := json.Unmarshal(*rawMessage, &p)
		if err != nil {
			return err
		}

		ce.Things[index] = p
	}

	return nil
}

func TestCustomJSON(t *testing.T) {
	// First let's create some things to live in the ecosystem
	fern := &Plant{MyColor: "green"}
	flower := &Plant{MyColor: "purple"}

	panther := &Animal{MyColor: "black"}
	lizard := &Animal{MyColor: "green"}

	// Then let's create a ColorfulEcosystem
	colorfulEcosystem := ColorfulEcosystem{
		Things: []ColoredThing{
			fern,
			flower,
			panther,
			lizard,
		},
	}

	hehe := &HeheColor{Colors: colorfulEcosystem}

	// prints:
	// {"things":[{"color":"green","type":"plant"},{"color":"purple","type":"plant"},{"color":"black","type":"animal"},{"color":"green","type":"animal"}]}
	byteSlice, _ := json.Marshal(hehe)
	fmt.Println(string(byteSlice))

	// // Now let's try deserializing the JSON back to a new struct
	// newCE := ColorfulEcosystem{}
	// err := json.Unmarshal(byteSlice, &newCE)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, colorfulThing := range newCE.Things {
	// 	fmt.Println(colorfulThing.Color())
	// }

}
