package modelTypes

// define all the structures used to build the chair model variants
type chairLower struct {
	Particle string `json:"particle" binding:"required"`
	Top      string `json:"top" binding:"required"`
	Front    string `json:"front" binding:"required"`
	Side     string `json:"side" binding:"required"`
	Bottom   string `json:"bottom" binding:"required"`
}

type chairUpper struct {
	Particle string `json:"particle" binding:"required"`
	Top      string `json:"top" binding:"required"`
	Front    string `json:"front" binding:"required"`
	Side     string `json:"side" binding:"required"`
}

type chairBackLower struct {
	Particle string `json:"particle" binding:"required"`
	Top      string `json:"top" binding:"required"`
	Front    string `json:"front" binding:"required"`
	Side     string `json:"side" binding:"required"`
	Back     string `json:"back" binding:"required"`
	Bottom   string `json:"bottom" binding:"required"`
}

type chairBackUpper struct {
	Particle string `json:"particle" binding:"required"`
	Top      string `json:"top" binding:"required"`
	Front    string `json:"front" binding:"required"`
	Side     string `json:"side" binding:"required"`
	Back     string `json:"back" binding:"required"`
}

// define all structures used to build the chair couchState variants
type chairState struct {
	Model string `json:"model" binding:"required"`
	Y     int    `json:"y,omitempty"`
}

//func MakeChairs(parentNameBase string, texturePath string, fileName string, mainDir string) {
//	space := "_"
//	condition := strings.ToTitle(parser.GetCondition())
//	for chairType := 0; chairType < 2; chairType++ {
//		switch chairType {
//		case 0:
//			if condition == "NOBACK" {
//				upper := chairUpper{
//					Particle: "#top",
//					Top:      "",
//					Front:    "",
//					Side:     "",
//				}
//				msg := helpers.GenericParent[chairUpper]{Parent: parentName, Textures: upper}
//				helpers.WriteModelFile(msg, mainDir, fileName+".json")
//			} else {
//				upper := chairBackUpper{
//					Particle: "#front",
//					Top:      "",
//					Front:    "",
//					Side:     "",
//					Back:     "",
//				}
//				msg := helpers.GenericParent[chairBackUpper]{Parent: parentName, Textures: upper}
//				helpers.WriteModelFile(msg, mainDir, fileName+".json")
//			}
//			break
//		case 1:
//
//			break
//		default:
//			return
//		}
//	}
//}
