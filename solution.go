package square

import "math"
// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#


type sides int

const SidesTriangle sides = 3
const SidesSquare sides = 4
const SidesCircle sides = 0

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum sides) float64 {
     if sidesNum == 0{
          return (math.Pi)*(math.Pow(sideLen,2))
     } else if sidesNum == 3{
          return (math.Pow(sideLen,2)/4)*(math.Sqrt(3))
     } else if sidesNum ==4 {
         return math.Pow(sideLen,2)
     } else{
         return 0
     }
}
