package pgm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// PGM struct represents a PGM image.
type PGM struct {
	data          [][]uint8
	width, height int
	magicNumber   string
	max           int
}

// ReadPGM lit une image PGM à partir d'un fichier et retourne une structure représentant l'image.
func ReadPGM(filename string) (*PGM, error) {
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
	data := make([][]uint8, height)
	for i := range data {
		data[i] = make([]uint8, width)
		for j := range data[i] {
			scanner.Scan()
			val, _ := strconv.Atoi(scanner.Text())
			data[i][j] = uint8(val)
		}
	}

	return &PGM{
		data:        data,
		width:       width,
		height:      height,
		magicNumber: magicNumber,
		max:         max,
	}, nil
}

// Size renvoie la largeur et la hauteur de l'image.
func (pgm *PGM) Size() (int, int) {
	return pgm.width, pgm.height
}

// At renvoie la valeur du pixel à la position (x, y).
func (pgm *PGM) At(x, y int) uint8 {
	return pgm.data[y][x]
}

// Set définit la valeur du pixel à la position (x, y).
func (pgm *PGM) Set(x, y int, value uint8) {
	pgm.data[y][x] = value
}

// Save enregistre l'image PGM dans un fichier et renvoie une erreur en cas de problème.
func (pgm *PGM) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "%s\n%d %d\n%d\n", pgm.magicNumber, pgm.width, pgm.height, pgm.max)
	for _, row := range pgm.data {
		for _, val := range row {
			fmt.Fprintf(writer, "%d ", val)
		}
		fmt.Fprintln(writer)
	}
	return writer.Flush()
}

// Invert inverse les couleurs de l'image PGM.
func (pgm *PGM) Invert() {
	for y := range pgm.data {
		for x := range pgm.data[y] {
			pgm.data[y][x] = uint8(pgm.max) - pgm.data[y][x]
		}
	}
}

// Flip retourne l'image PGM horizontalement.
func (pgm *PGM) Flip() {
	for y := range pgm.data {
		for x := 0; x < pgm.width/2; x++ {
			pgm.data[y][x], pgm.data[y][pgm.width-x-1] = pgm.data[y][pgm.width-x-1], pgm.data[y][x]
		}
	}
}

// Flop retourne l'image PGM verticalement.
func (pgm *PGM) Flop() {
	for y := 0; y < pgm.height/2; y++ {
		pgm.data[y], pgm.data[pgm.height-y-1] = pgm.data[pgm.height-y-1], pgm.data[y]
	}
}

// SetMagicNumber définit le numéro magique de l'image PGM.
func (pgm *PGM) SetMagicNumber(magicNumber string) {
	pgm.magicNumber = magicNumber
}

// SetMaxValue définit la valeur maximale de l'image PGM.
func (pgm *PGM) SetMaxValue(maxValue uint8) {
	pgm.max = int(maxValue)
}

// Rotate90CW fait pivoter l'image PGM de 90° dans le sens des aiguilles d'une montre.
func (pgm *PGM) Rotate90CW() {
	newData := make([][]uint8, pgm.width)
	for i := range newData {
		newData[i] = make([]uint8, pgm.height)
		for j := range newData[i] {
			newData[i][j] = pgm.data[pgm.height-j-1][i]
		}
	}
	pgm.data = newData
	pgm.width, pgm.height = pgm.height, pgm.width
}
