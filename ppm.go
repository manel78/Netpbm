package main

type PPM struct {
	Data          [][]Pixel
	width, height int
	magicNumber   string
	max           int
}

type Pixel struct {
	R, G, B uint8
}

func (ppm *PPM) Size() (int, int) {
	return ppm.width, ppm.height
}

func (ppm *PPM) At(x, y int) Pixel {
	return ppm.Data[y][x]
}

func (ppm *PPM) Set(x, y int, value Pixel) {
	ppm.Data[y][x] = value
}

func (ppm *PPM) SetMagicNumber(magicNumber string) {
	ppm.magicNumber = magicNumber
}

func (ppm *PPM) Invert() {
	for i := 0; i < len(ppm.Data); i++ { // parcours la longueur de pixels
		for j := 0; j < len(ppm.Data[0]); j++ {
			ppm.Data[i][j].R = ppm.Max - ppm.Data[i][j].R // inverse les valeurs
			ppm.Data[i][j].G = ppm.Max - ppm.Data[i][j].G
			ppm.Data[i][j].B = ppm.Max - ppm.Data[i][j].B
		}
	}
}

func (ppm *PPM) DrawPolygon(points []Point, color Pixel) {
	for i := 0; i < len(points); i++ {
		nextIndex := (i + 1) % len(points)
		ppm.DrawLine(points[i], points[nextIndex], color)
	}
}

func (ppm *PPM) SetMaxValue(maxValue uint8) {
	ppm.max = uint8(maxValue)
}

func NewPPM(width, height int) *PPM {
	data := make([][]Pixel, height)
	for i := range data {
		data[i] = make([]Pixel, width)
	}

	return &PPM{
		data:        data,
		width:       width,
		height:      height,
		magicNumber: "P3",
		max:         255,
	}
}
