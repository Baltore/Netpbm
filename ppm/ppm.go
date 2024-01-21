package ppm

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Pixel struct represents a pixel with red, green, and blue values.
type Pixel struct {
	R, G, B uint8
}

type Point struct {
	X, Y int
}

// PPM struct represents a PPM image.
type PPM struct {
	data          [][]Pixel
	width, height int
	magicNumber   string
	max           uint
}

// ReadPPM lit une image PPM à partir d'un fichier et retourne une structure représentant l'image.
func ReadPPM(filename string) (*PPM, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	// Lire le numéro magique
	scanner.Scan()
	magicNumber := scanner.Text()

	// Lire la largeur et la hauteur
	scanner.Scan()
	width, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	height, _ := strconv.Atoi(scanner.Text())

	// Lire la valeur maximale
	scanner.Scan()
	max, _ := strconv.Atoi(scanner.Text())

	// Lire les données de l'image
	data := make([][]Pixel, height)
	for i := range data {
		data[i] = make([]Pixel, width)
		for j := range data[i] {
			var pixel Pixel

			// Lire la composante rouge
			scanner.Scan()
			red, err := strconv.ParseUint(scanner.Text(), 10, 8)
			if err != nil {
				return nil, err
			}
			pixel.R = uint8(red)

			// Lire la composante verte
			scanner.Scan()
			green, err := strconv.ParseUint(scanner.Text(), 10, 8)
			if err != nil {
				return nil, err
			}
			pixel.G = uint8(green)

			// Lire la composante bleue
			scanner.Scan()
			blue, err := strconv.ParseUint(scanner.Text(), 10, 8)
			if err != nil {
				return nil, err
			}
			pixel.B = uint8(blue)

			data[i][j] = pixel
		}
	}

	return &PPM{
		data:        data,
		width:       width,
		height:      height,
		magicNumber: magicNumber,
		max:         uint(max),
	}, nil
}

// Size renvoie la largeur et la hauteur de l'image.
func (ppm *PPM) Size() (int, int) {
	return ppm.width, ppm.height
}

// At renvoie la valeur du pixel à la position (x, y).
func (ppm *PPM) At(x, y int) Pixel {
	return ppm.data[y][x]
}

// Set définit la valeur du pixel à la position (x, y).
func (ppm *PPM) Set(x, y int, value Pixel) {
	ppm.data[y][x] = value
}

// Save enregistre l'image PPM dans un fichier et renvoie une erreur en cas de problème.
func (ppm *PPM) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "%s\n%d %d\n%d\n", ppm.magicNumber, ppm.width, ppm.height, ppm.max)
	for _, row := range ppm.data {
		for _, pixel := range row {
			fmt.Fprintf(writer, "%d %d %d ", pixel.R, pixel.G, pixel.B)
		}
		fmt.Fprintln(writer)
	}
	return writer.Flush()
}

// Invert inverse les couleurs de l'image PPM.
func (ppm *PPM) Invert() {
	for y := range ppm.data {
		for x := range ppm.data[y] {
			ppm.data[y][x].R = uint8(ppm.max) - ppm.data[y][x].R
			ppm.data[y][x].G = uint8(ppm.max) - ppm.data[y][x].G
			ppm.data[y][x].B = uint8(ppm.max) - ppm.data[y][x].B
		}
	}
}

// Flip retourne l'image PPM horizontalement.
func (ppm *PPM) Flip() {
	for y := range ppm.data {
		for x := 0; x < ppm.width/2; x++ {
			ppm.data[y][x], ppm.data[y][ppm.width-x-1] = ppm.data[y][ppm.width-x-1], ppm.data[y][x]
		}
	}
}

// Flop retourne l'image PPM verticalement.
func (ppm *PPM) Flop() {
	for y := 0; y < ppm.height/2; y++ {
		ppm.data[y], ppm.data[ppm.height-y-1] = ppm.data[ppm.height-y-1], ppm.data[y]
	}
}

// SetMagicNumber définit le numéro magique de l'image PPM.
func (ppm *PPM) SetMagicNumber(magicNumber string) {
	ppm.magicNumber = magicNumber
}

// SetMaxValue définit la valeur maximale de l'image PPM.
func (ppm *PPM) SetMaxValue(maxValue uint8) {
	ppm.max = uint(maxValue)
}

// Rotate90CW fait pivoter l'image PPM de 90° dans le sens des aiguilles d'une montre.
func (ppm *PPM) Rotate90CW() {
	newData := make([][]Pixel, ppm.width)
	for i := range newData {
		newData[i] = make([]Pixel, ppm.height)
		for j := range newData[i] {
			newData[i][j] = ppm.data[ppm.height-j-1][i]
		}
	}
	ppm.data = newData
	ppm.width, ppm.height = ppm.height, ppm.width
}

// maxAbs retourne la valeur absolue maximale entre deux nombres flottants.
func maxAbs(a, b float64) float64 {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	if a > b {
		return a
	}
	return b
}

// DrawLine dessine une ligne entre deux points.
func (ppm *PPM) DrawLine(p1, p2 Point, color Pixel) {
	dx := float64(p2.X - p1.X)
	dy := float64(p2.Y - p1.Y)

	steps := int(maxAbs(dx, dy))

	xIncrement := dx / float64(steps)
	yIncrement := dy / float64(steps)

	x, y := float64(p1.X), float64(p1.Y)

	for i := 0; i <= steps; i++ {
		ppm.Set(int(x+0.5), int(y+0.5), color)
		x += xIncrement
		y += yIncrement
	}
}

// DrawRectangle dessine un rectangle.
func (ppm *PPM) DrawRectangle(p1 Point, width, height int, color Pixel) {
	p2 := Point{X: p1.X + width, Y: p1.Y}
	p3 := Point{X: p1.X, Y: p1.Y + height}
	p4 := Point{X: p1.X + width, Y: p1.Y + height}

	ppm.DrawLine(p1, p2, color)
	ppm.DrawLine(p2, p4, color)
	ppm.DrawLine(p4, p3, color)
	ppm.DrawLine(p3, p1, color)
}

// DrawFilledRectangle dessine un rectangle rempli.
func (ppm *PPM) DrawFilledRectangle(p1 Point, width, height int, color Pixel) {
	for y := p1.Y; y < p1.Y+height; y++ {
		for x := p1.X; x < p1.X+width; x++ {
			ppm.Set(x, y, color)
		}
	}
}

// DrawCircle dessine un cercle.
func (ppm *PPM) DrawCircle(center Point, radius int, color Pixel) {
	for y := 0; y < ppm.height; y++ {
		for x := 0; x < ppm.width; x++ {
			distance := math.Sqrt(math.Pow(float64(x-center.X), 2) + math.Pow(float64(y-center.Y), 2))
			if math.Abs(distance-float64(radius)) < 1.0 {
				ppm.Set(x, y, color)
			}
		}
	}
}

// DrawFilledCircle dessine un cercle rempli.
func (ppm *PPM) DrawFilledCircle(center Point, radius int, color Pixel) {
	for y := 0; y < ppm.height; y++ {
		for x := 0; x < ppm.width; x++ {
			distance := math.Sqrt(math.Pow(float64(x-center.X), 2) + math.Pow(float64(y-center.Y), 2))
			if distance < float64(radius) {
				ppm.Set(x, y, color)
			}
		}
	}
}

// DrawTriangle dessine un triangle.
func (ppm *PPM) DrawTriangle(p1, p2, p3 Point, color Pixel) {
	// Utilisez une fonction de dessin de ligne pour dessiner les trois côtés du triangle
	ppm.DrawLine(p1, p2, color)
	ppm.DrawLine(p2, p3, color)
	ppm.DrawLine(p3, p1, color)
}

// DrawFilledTriangle dessine un triangle rempli.
func (ppm *PPM) DrawFilledTriangle(p1, p2, p3 Point, color Pixel) {
	// Triez les points par ordre croissant de Y
	points := []Point{p1, p2, p3}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2-i; j++ {
			if points[j].Y > points[j+1].Y {
				points[j], points[j+1] = points[j+1], points[j]
			}
		}
	}

	// Calculez les pentes inverses des côtés du triangle
	invSlope1 := float64(points[1].X-points[0].X) / float64(points[1].Y-points[0].Y)
	invSlope2 := float64(points[2].X-points[0].X) / float64(points[2].Y-points[0].Y)

	// Initialisez les valeurs de départ des bords gauche et droit du triangle
	xLeft := float64(points[0].X)
	xRight := float64(points[0].X)

	// Parcourez les lignes du triangle et dessinez-les
	for y := points[0].Y; y <= points[2].Y; y++ {
		// Utilisez la fonction DrawLine pour dessiner chaque ligne horizontale du triangle
		ppm.DrawLine(Point{int(xLeft), y}, Point{int(xRight), y}, color)

		// Mettez à jour les bords gauche et droit
		xLeft += invSlope1
		xRight += invSlope2
	}
}

// DrawPolygon dessine un polygone.
func (ppm *PPM) DrawPolygon(points []Point, color Pixel) {
	for i := 0; i < len(points); i++ {
		nextIndex := (i + 1) % len(points)
		ppm.DrawLine(points[i], points[nextIndex], color)
	}
}

// NewPPM crée une nouvelle image PPM avec la largeur et la hauteur spécifiées.
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
