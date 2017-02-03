package transform

import (
	"testing"
	"math"
)

func TestGetRotationMatrix(t *testing.T) {
	cases:=[]struct{
		x float64
		y float64
		angle float64
		scale float64
		excepted RotationMatrix
		}{
			{
				x:10,
				y:10,
				angle:45,
				scale:1.0,
				excepted:RotationMatrix([6]float64{
					0.7071067811865476, 0.7071067811865475, -4.14213562373095,
					-0.7071067811865475, 0.7071067811865476, 10,
				}),
			},
			{
				x:10,
				y:10,
				angle:-30,
				scale:2.0,
				excepted:RotationMatrix([6]float64{
					1.732050807568877, -0.9999999999999999, 2.679491924311224,
					0.9999999999999999, 1.732050807568877, -17.32050807568877,
				}),
			},
			{
				x:10,
				y:10,
				angle:70,
				scale:0.6,
				excepted:RotationMatrix([6]float64{
					0.2052120859954013, 0.563815572471545, 2.309723415330537,
					-0.563815572471545, 0.2052120859954013, 13.58603486476144,
				}),
			},
		}
	var eps =0.0000001;
	for _,c:=range cases{
		actual:=GetRotationMatrix(c.x,c.y,c.angle,c.scale)
		for i:=0;i<len(actual);i++{
			if math.Abs(actual[i]-c.excepted[i])>eps{
				t.Errorf("RotationMatrix:\nexpected: %v\nactual: %v",c.excepted,actual)
				break
			}
		}
	}
}