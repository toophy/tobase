package skills

import "testing"
import "encoding/json"

func TestSpaceTime(t *testing.T) {
	gWorld.Strings["作者"] = "toophy"
	gWorld.Ints["年龄"] = 37
	gWorld.Floats["坐标X"] = 1.01
	gWorld.Floats["坐标Y"] = 1.01
	gWorld.Floats["坐标Z"] = 1.01
	u := NewUniverse("第一宇宙", gWorld)
	scene := NewScene("地球", u)
	outJSON := gWorld.OutputJSON()
	println(outJSON)

	realScene := NewRealScene("火之地球", scene)
	outRealScene, err := json.MarshalIndent(realScene, "", "  ")
	if err != nil {
		println("%s", err.Error())
	} else {
		println(string(outRealScene))
	}
}

func TestLoadWorld(t *testing.T) {
	data := `
{
  "Name": "世界",
  "Parent": "",
  "Time": 0,
  "Strings": {
    "作者": "toophy"
  },
  "Ints": {
    "年龄": 37
  },
  "Floats": {
    "坐标X": 1.01,
    "坐标Y": 1.01,
    "坐标Z": 1.01
  },
  "Universes": {
    "第一宇宙": {
      "Name": "第一宇宙",
      "Parent": "",
      "Time": 0,
      "Strings": {
        "作者": "toophy"
      },
      "Ints": {
        "年龄": 37
      },
      "Floats": {
        "坐标X": 1.01,
        "坐标Y": 1.01,
        "坐标Z": 1.01
      },
      "Scenes": {
        "地球": {
          "Name": "地球",
          "Parent": "",
          "Time": 0,
          "Strings": {
            "作者": "toophy"
          },
          "Ints": {
            "年龄": 37
          },
          "Floats": {
            "坐标X": 1.01,
            "坐标Y": 1.01,
            "坐标Z": 1.01
          }
        }
      }
    }
  }
}`

	newWorld := new(World)
	newWorld.InputJSON(data)
	outJSON := newWorld.OutputJSON()
	println(outJSON)
}
