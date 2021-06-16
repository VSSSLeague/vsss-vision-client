package client

import (
	"math"
	"sort"
	"strconv"

	"github.com/VSSSLeague/vsss-vision-client/pkg/vision"
)

var white = "white"
var orange = "orange"
var yellow = "yellow"
var blue = "blue"
var black = "black"
var lineWidth = float32(10)
var ballRadius = float32(21)
var botLength = float64(80)
var center2Dribbler = float64(75)
var noFill = float32(0)
var botStrokeWidth = float32(10)
var ballStrokeWidth = float32(0)
var zero = float32(0)
var fbRadius = float32(10)

type ShapesByOrderNumber []Shape

func (a ShapesByOrderNumber) Len() int           { return len(a) }
func (a ShapesByOrderNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ShapesByOrderNumber) Less(i, j int) bool { return a[i].OrderNumber < a[j].OrderNumber }

func (p *Package) AddDetectionFrame(frame *vision.Frame) {
	p.Shapes = append(p.Shapes, Shape{OrderNumber: 3, Circle: createBallShape(float32(*&frame.Ball.X), float32(*&frame.Ball.Y), 0)})

	for _, bot := range frame.RobotsBlue {
		p.Shapes = append(p.Shapes, Shape{OrderNumber: 1, Rect: createBotRect(float32(*&bot.X), float32(*&bot.Y), float32(*&bot.Orientation), blue)})
		p.Shapes = append(p.Shapes, Shape{OrderNumber: 2, Text: createBotId(*&bot.RobotId, float32(*&bot.X), float32(*&bot.Y), white)})
	}

	for _, bot := range frame.RobotsYellow {
		p.Shapes = append(p.Shapes, Shape{OrderNumber: 1, Rect: createBotRect(float32(*&bot.X), float32(*&bot.Y), float32(*&bot.Orientation), yellow)})
		p.Shapes = append(p.Shapes, Shape{OrderNumber: 2, Text: createBotId(*&bot.RobotId, float32(*&bot.X), float32(*&bot.Y), black)})
	}
}

func createBallShape(x, y, z float32) *Circle {
	heightFactor := 0.01*math.Abs(float64(z)) + 1
	return &Circle{
		Center: Point{x * 1000.0, -y * 1000.0},
		Radius: float32(heightFactor) * ballRadius,
		Style: Style{
			StrokeWidth: &ballStrokeWidth,
			Fill:        &orange,
		},
	}
}

func createBotRect(posX, posY, orientation float32, fillColor string) *Rect {
	return &Rect{
		X:      (posX * 1000.0) - float32(botLength/2),
		Y:      -(posY * 1000.0) - float32(botLength/2),
		Width:  float32(botLength),
		Height: float32(botLength),
		Ori:    float32(orientation),
		Style: Style{
			Fill:        &fillColor,
			Stroke:      &black,
			StrokeWidth: &botStrokeWidth,
		},
	}
}

func createBotId(id uint32, x, y float32, strokeColor string) *Text {
	return &Text{
		Text: strconv.Itoa(int(id)),
		P:    Point{x * 1000.0, -y * 1000.0},
		Style: Style{
			Fill: &strokeColor,
		},
	}
}

func (p *Package) AddGeometryShapes(geometry *vision.Field) {
	p.FieldWidth = float32(*&geometry.Width) * 1000.0
	p.FieldLength = float32(*&geometry.Length) * 1000.0
	p.BoundaryWidth = float32(150) /// TODO: receive this by protobuf later
	p.GoalWidth = float32(*&geometry.GoalWidth) * 1000.0
	p.GoalDepth = float32(*&geometry.GoalDepth) * 1000.0
	p.PenaltyWidth = float32(*&geometry.PenaltyWidth) * 1000.0
	p.PenaltyDepth = float32(*&geometry.PenaltyDepth) * 1000.0
	p.PenaltyPoint = float32(*&geometry.PenaltyPoint) * 1000.0
	p.CenterRadius = float32(*&geometry.CenterRadius) * 1000.0

	p.Shapes = append(p.Shapes, goalLinesPositive(geometry)...)
	p.Shapes = append(p.Shapes, goalLinesNegative(geometry)...)
	p.Shapes = append(p.Shapes, fieldExternalLines(geometry)...)
	p.Shapes = append(p.Shapes, fieldPenaltyAreas(geometry)...)
	p.Shapes = append(p.Shapes, fieldCircles(geometry)...)
}

func fieldPenaltyAreas(geometry *vision.Field) (rects []Shape) {
	len := float32(float32(*&geometry.Length) * 1000.0)
	penWid := float32(float32(*&geometry.PenaltyWidth) * 1000.0)
	penDepth := float32(float32(*&geometry.PenaltyDepth) * 1000.0)

	// Left penalty area
	rects = append(rects, Shape{Rect: &Rect{X: -len / 2, Y: -penWid / 2, Width: penDepth, Height: penWid, Ori: 0,
		Style: Style{Stroke: &white, FillOpacity: &zero, StrokeWidth: &lineWidth}}})

	// Right penalty area
	rects = append(rects, Shape{Rect: &Rect{X: len/2 - penDepth, Y: -penWid / 2, Width: penDepth, Height: penWid, Ori: 0,
		Style: Style{Stroke: &white, FillOpacity: &zero, StrokeWidth: &lineWidth}}})

	/// TODO: check how to do this in an better way
	if len == 2200 {
		// Take vars
		goalWid := float32(float32(*&geometry.GoalWidth * 1000.0))

		// Left great area
		rects = append(rects, Shape{Rect: &Rect{X: -len / 2, Y: -goalWid, Width: 350, Height: goalWid * 2, Ori: 0,
			Style: Style{Stroke: &white, FillOpacity: &zero, StrokeWidth: &lineWidth}}})

		// Right great area
		rects = append(rects, Shape{Rect: &Rect{X: len/2 - 350, Y: -goalWid, Width: 350, Height: goalWid * 2, Ori: 0,
			Style: Style{Stroke: &white, FillOpacity: &zero, StrokeWidth: &lineWidth}}})
	}

	return
}

func fieldCircles(geometry *vision.Field) (circles []Shape) {
	len := float32(float32(*&geometry.Length) * 1000.0)
	wid := float32(float32(*&geometry.Width) * 1000.0)
	centerRad := float32(float32(*&geometry.CenterRadius) * 1000.0)
	penPoint := float32(float32(*&geometry.PenaltyPoint) * 1000.0)

	// Center circle
	circles = append(circles, Shape{Circle: &Circle{Center: Point{0, 0}, Radius: centerRad,
		Style: Style{StrokeWidth: &lineWidth, FillOpacity: &zero}}})

	// FB Quadrant Top-Right
	circles = append(circles, Shape{Circle: &Circle{Center: Point{len / 4, (wid / 2) - (centerRad + 50)}, Radius: fbRadius,
		Style: Style{StrokeWidth: &lineWidth, Fill: &white}}})

	// FB Quadrant Top-Left
	circles = append(circles, Shape{Circle: &Circle{Center: Point{-len / 4, (wid / 2) - (centerRad + 50)}, Radius: fbRadius,
		Style: Style{StrokeWidth: &lineWidth, Fill: &white}}})

	// FB Quadrant Bottom-Left
	circles = append(circles, Shape{Circle: &Circle{Center: Point{-len / 4, -(wid / 2) + (centerRad + 50)}, Radius: fbRadius,
		Style: Style{StrokeWidth: &lineWidth, Fill: &white}}})

	// FB Quadrant Bottom-Right
	circles = append(circles, Shape{Circle: &Circle{Center: Point{len / 4, -(wid / 2) + (centerRad + 50)}, Radius: fbRadius,
		Style: Style{StrokeWidth: &lineWidth, Fill: &white}}})

	// Penalty Left Mark
	circles = append(circles, Shape{Circle: &Circle{Center: Point{-(len / 2) + penPoint, 0}, Radius: fbRadius,
		Style: Style{StrokeWidth: &lineWidth, Fill: &white}}})

	// Penalty Right Mark
	circles = append(circles, Shape{Circle: &Circle{Center: Point{(len / 2) - penPoint, 0}, Radius: fbRadius,
		Style: Style{StrokeWidth: &lineWidth, Fill: &white}}})

	return
}

func fieldExternalLines(geometry *vision.Field) (lines []Shape) {
	len := float32(float32(*&geometry.Length) * 1000.0)
	wid := float32(float32(*&geometry.Width) * 1000.0)
	tsz := float32(70)

	// Left
	lines = append(lines, Shape{Line: &Line{P1: Point{-len / 2, -wid/2 + tsz}, P2: Point{-len / 2, wid/2 - tsz},
		Style: Style{Stroke: &white, StrokeWidth: &lineWidth}}})

	// Left-Bottom
	lines = append(lines, Shape{Line: &Line{P1: Point{-len / 2, wid/2 - tsz}, P2: Point{-len/2 + tsz, wid / 2},
		Style: Style{Stroke: &white, StrokeWidth: &lineWidth}}})

	// Bottom
	lines = append(lines, Shape{Line: &Line{P1: Point{-len/2 + tsz, wid / 2}, P2: Point{len/2 - tsz, wid / 2},
		Style: Style{Stroke: &white, StrokeWidth: &lineWidth}}})

	// Bottom-Right
	lines = append(lines, Shape{Line: &Line{P1: Point{len/2 - tsz, wid / 2}, P2: Point{len / 2, wid/2 - tsz},
		Style: Style{Stroke: &white, StrokeWidth: &lineWidth}}})

	// Right
	lines = append(lines, Shape{Line: &Line{P1: Point{len / 2, wid/2 - tsz}, P2: Point{len / 2, -wid/2 + tsz},
		Style: Style{Stroke: &white, StrokeWidth: &lineWidth}}})

	// Right-Top
	lines = append(lines, Shape{Line: &Line{P1: Point{len / 2, -wid/2 + tsz}, P2: Point{len/2 - tsz, -wid / 2},
		Style: Style{Stroke: &white, StrokeWidth: &lineWidth}}})

	// Top
	lines = append(lines, Shape{Line: &Line{P1: Point{len/2 - tsz, -wid / 2}, P2: Point{-len/2 + tsz, -wid / 2},
		Style: Style{Stroke: &white, StrokeWidth: &lineWidth}}})

	// Top-Left
	lines = append(lines, Shape{Line: &Line{P1: Point{-len/2 + tsz, -wid / 2}, P2: Point{-len / 2, -wid/2 + tsz},
		Style: Style{Stroke: &white, StrokeWidth: &lineWidth}}})

	// MidLine
	lines = append(lines, Shape{Line: &Line{P1: Point{0, wid / 2}, P2: Point{0, -wid / 2},
		Style: Style{Stroke: &white, StrokeWidth: &lineWidth}}})

	return
}

func goalLinesNegative(geometry *vision.Field) (lines []Shape) {
	lines = goalLinesPositive(geometry)
	for i := range lines {
		lines[i].Line.P1.X *= -1
		lines[i].Line.P2.X *= -1
	}
	return
}

func goalLinesPositive(geometry *vision.Field) (lines []Shape) {
	flh := float32(float32(*&geometry.Length) * 1000.0 / 2)
	gwh := float32(float32(*&geometry.GoalWidth) * 1000.0 / 2)
	gd := float32(float32(*&geometry.GoalDepth) * 1000.0)

	lines = append(lines, Shape{Line: &Line{P1: Point{flh, -gwh}, P2: Point{flh + gd, -gwh},
		Style: Style{Stroke: &black, StrokeWidth: &lineWidth}}})
	lines = append(lines, Shape{Line: &Line{P1: Point{flh, gwh}, P2: Point{flh + gd, gwh},
		Style: Style{Stroke: &black, StrokeWidth: &lineWidth}}})
	lines = append(lines, Shape{Line: &Line{P1: Point{flh + gd, -gwh}, P2: Point{flh + gd, gwh},
		Style: Style{Stroke: &black, StrokeWidth: &lineWidth}}})

	return
}

func (p *Package) SortShapes() {
	sort.Sort(ShapesByOrderNumber(p.Shapes))
}
