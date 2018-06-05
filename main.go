package main

import (
	"fmt"
	"github.com/intdxdt/geom"
	"github.com/gopherjs/gopherjs/js"
	"github.com/TopoSimplify/opts"
	"encoding/json"
)

func main() {
	js.Global.Set("_go", map[string]interface{}{
		"simplx": map[string]interface{}{
			"Options": pointArray,
		},
	})
}

type SimplxObj struct {
	pln         string
	constraints []string
}

func Simplify(siplxobj_json, options_json string) *js.Object {
	var options = opts.Opts{}
	var err = json.Unmarshal([]byte(options_json), &options)
	if err != nil {
		panic(err)
	}
	fmt.Println(options)
	var obj = SimplxObj{}
	err = json.Unmarshal([]byte(siplxobj_json), &obj)
	if err != nil {
		panic(err)
	}
	return js.MakeWrapper(options)
}

func pointArray(points [][]float64) *js.Object {
	var pts = make([]*geom.Point, 0)
	for _, pt := range points {
		pts = append(pts, geom.NewPoint(pt))
	}
	var ln = geom.NewLineString(pts)
	fmt.Println(ln.WKT())

	//options := &opts.Opts{
	//	Threshold:              50.0,
	//	MinDist:                20.0,
	//	RelaxDist:              30.0,
	//	PlanarSelf:             false,
	//	AvoidNewSelfIntersects: true,
	//	GeomRelation:           true,
	//	DirRelation:            false,
	//	DistRelation:           false,
	//}

	return js.MakeWrapper(pts)
}
