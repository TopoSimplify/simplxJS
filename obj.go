package main

import (
	"github.com/intdxdt/geom"
	"github.com/gopherjs/gopherjs/js"
	"github.com/TopoSimplify/offset"
	"strings"
	"github.com/TopoSimplify/constdp"
	"github.com/TopoSimplify/opts"
)

var offsetFuncs = map[string]func([]*geom.Point) (int, float64){
	"dp":  offset.MaxOffset,
	"sed": offset.MaxSEDOffset,
}

type SplxObj struct {
	Polyline               *geom.LineString `json:"polyline"`    //array of geom.LineString
	Constraints            []string         `json:"constraints"` //WKTs
	OffsetFunc             string           `json:"offset"`      //offset func : dp, sed
	Threshold              float64          `json:"threshold"`
	MinDist                float64          `json:"mindist"`
	RelaxDist              float64          `json:"relaxdist"`
	PlanarSelf             bool             `json:"planarself"`
	NonPlanarSelf          bool             `json:"nonplanarself"`
	AvoidNewSelfIntersects bool             `json:"avoidself"`
	GeomRelation           bool             `json:"geomrelate"`
	DistRelation           bool             `json:"distrelate"`
	DirRelation            bool             `json:"dirrelate"`
}

func (o *SplxObj) SetPolyline(v [][]float64) {
	var coordinates []*geom.Point
	for _, c := range v {
		coordinates = append(coordinates, geom.NewPoint(c))
	}
	o.Polyline = geom.NewLineString(coordinates)
}
func (o *SplxObj) GetPolyline() *js.Object {
	var obj = js.Global.Get("Object").New()
	var coords [][]float64
	for _, o := range o.Polyline.Coordinates() {
		coords = append(coords, o[:])
	}
	obj.Set("value", coords)
	return obj
}

func (o *SplxObj) SetConstraints(v []string) {
	o.Constraints = v
}
func (o *SplxObj) GetConstraints(v float64) *js.Object {
	var obj = js.Global.Get("Object").New()
	obj.Set("value", o.Constraints)
	return obj
}

func (o *SplxObj) SetOffsetFunc(key string) {
	key = strings.ToLower(key)
	var fn = offsetFuncs[key]
	if fn == nil {
		panic("not implemented, expects : dp | sed as key")
	}
	o.OffsetFunc = key
}
func (o *SplxObj) GetOffsetFunc(v float64) *js.Object {
	var obj = js.Global.Get("Object").New()
	obj.Set("value", o.OffsetFunc)
	return obj
}

func (o *SplxObj) SetThreshold(v float64) {
	o.Threshold = v
}
func (o *SplxObj) GetThreshold(v float64) *js.Object {
	var obj = js.Global.Get("Object").New()
	obj.Set("value", o.Threshold)
	return obj
}

func (o *SplxObj) SetMinDist(v float64) {
	o.MinDist = v
}
func (o *SplxObj) GetMinDist(v float64) *js.Object {
	var obj = js.Global.Get("Object").New()
	obj.Set("value", o.MinDist)
	return obj
}

func (o *SplxObj) SetRelaxDist(v float64) {
	o.RelaxDist = v
}
func (o *SplxObj) GetRelaxDist(v float64) *js.Object {
	var obj = js.Global.Get("Object").New()
	obj.Set("value", o.RelaxDist)
	return obj
}

func (o *SplxObj) SetPlanarSelf(v bool) {
	o.PlanarSelf = v
}
func (o *SplxObj) GetPlanarSelf() *js.Object {
	var obj = js.Global.Get("Object").New()
	obj.Set("value", o.PlanarSelf)
	return obj
}

func (o *SplxObj) SetNonPlanarSelf(v bool) {
	o.NonPlanarSelf = v
}
func (o *SplxObj) GetNonPlanarSelf() *js.Object {
	var obj = js.Global.Get("Object").New()
	obj.Set("value", o.NonPlanarSelf)
	return obj
}

func (o *SplxObj) SetAvoidNewSelfIntersects(v bool) {
	o.AvoidNewSelfIntersects = v
}
func (o *SplxObj) GetAvoidNewSelfIntersects() *js.Object {
	var obj = js.Global.Get("Object").New()
	obj.Set("value", o.AvoidNewSelfIntersects)
	return obj
}

func (o *SplxObj) SetGeomRelation(v bool) {
	o.GeomRelation = v
}
func (o *SplxObj) GetGeomRelation() *js.Object {
	var obj = js.Global.Get("Object").New()
	obj.Set("value", o.GeomRelation)
	return obj
}

func (o *SplxObj) SetDistRelation(v bool) {
	o.DistRelation = v
}
func (o *SplxObj) GetDistRelation() *js.Object {
	var obj = js.Global.Get("Object").New()
	obj.Set("value", o.DistRelation)
	return obj
}

func (o *SplxObj) SetDirRelation(v bool) {
	o.DirRelation = v
}
func (o *SplxObj) GetDirRelation() *js.Object {
	var obj = js.Global.Get("Object").New()
	obj.Set("value", marshal(o.DirRelation))
	return obj
}

func (o *SplxObj) Simplify() *js.Object {
	options := &opts.Opts{
		Threshold:              o.Threshold,
		MinDist:                o.MinDist,
		RelaxDist:              o.RelaxDist,
		PlanarSelf:             o.PlanarSelf,
		AvoidNewSelfIntersects: o.AvoidNewSelfIntersects,
		GeomRelation:           o.GeomRelation,
		DirRelation:            o.DirRelation,
		DistRelation:           o.DistRelation,
	}

	var constraints []geom.Geometry
	for _, wkt := range o.Constraints {
		constraints = append(constraints, geom.NewPolygonFromWKT(wkt))
	}

	var fn = offsetFuncs[o.OffsetFunc]
	var coords = o.Polyline.Coordinates()
	var homo = constdp.NewConstDP(coords, constraints, options, fn)

	var ptset = homo.Simplify().SimpleSet

	var indices []int
	var simple [][]float64
	for _, i := range ptset.Values() {
		indices = append(indices, i.(int))
	}
	for _, i := range indices {
		simple = append(simple, coords[i][:])
	}

	var obj = js.Global.Get("Object").New()
	obj.Set("simple", simple)
	obj.Set("indices", indices)
	return obj

}
