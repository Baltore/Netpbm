package pbm

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// PBM représente une image PBM.
type PBM struct {
	data          [][]bool
	width, height int
	magicNumber   string
}

// ReadPBM lit une image PBM à partir d'un fichier et renvoie une structure qui représente l'image.
func ReadPBM(filename string) (*PBM, error) {
	// Ouvrir le fichier
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Initialiser le scanner pour lire le fichier ligne par ligne
	scanner := bufio.NewScanner(file)

	// Ignorer les commentaires au début du fichier
	for scanner.Scan() {
		if line := scanner.Text(); !strings.HasPrefix(line, "#") {
			break
		}
	}

	// Lire le nombre magique (P1 ou P4)
	magicNumber := scanner.Text()
	if magicNumber != "P1" && magicNumber != "P4" {
		return nil, errors.New("type de fichier non pris en charge")
	}

	// Lire les dimensions de l'image
	scanner.Scan()
	dimensions := strings.Fields(scanner.Text())
	if len(dimensions) != 2 {
		return nil, errors.New("dimensions d'image non valides")
	}

	// Convertir les dimensions en entiers (largeur et hauteur)
	width, _ := strconv.Atoi(dimensions[0])
	height, _ := strconv.Atoi(dimensions[1])

	// Initialiser une matrice pour stocker les valeurs des pixels
	var data [][]bool
	for scanner.Scan() {
		// Ignorer les lignes de commentaire
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		// Initialiser une ligne pour stocker les valeurs des pixels
		row := make([]bool, width)
		for i, char := range strings.Fields(line) {
			// Convertir chaque caractère en un booléen et le stocker dans la ligne
			row[i] = char == "1"
		}
		// Ajouter la ligne à la matrice
		data = append(data, row)
	}

	// Vérifier s'il y a des erreurs lors de la lecture du fichier
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Créer et renvoyer une nouvelle instance de la structure PBM avec les données lues
	return &PBM{
		data:        data,
		width:       width,
		height:      height,
		magicNumber: magicNumber,
	}, nil
}

// Size renvoie la largeur et la hauteur de l'image.
func (pbm *PBM) Size() (int, int) {
	return pbm.width, pbm.height
}

// At renvoie la valeur du pixel en (x, y).
func (pbm *PBM) At(x, y int) bool {
	if len(pbm.data) == 0 || x < 0 || y < 0 || x >= pbm.width || y >= pbm.height {
		// Les coordonnées sont hors de la plage valide ou le tableau est vide.
		// Vous pouvez renvoyer une valeur par défaut ou gérer l'erreur de la manière qui vous convient.
		return false
	}

	return pbm.data[y][x]
}

// Set définit la valeur du pixel à (x, y).
func (pbm *PBM) Set(x, y int, value bool) {
	pbm.data[y][x] = value
}

// Save enregistre l'image PBM dans un fichier et renvoie une erreur en cas de problème.
func (pbm *PBM) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Écrire le nombre magique et les dimensions
	fmt.Fprintf(file, "%s\n%d %d\n", pbm.magicNumber, pbm.width, pbm.height)

	// Écrire les pixels
	for _, row := range pbm.data {
		for _, pixel := range row {
			if pixel {
				fmt.Fprint(file, "1 ")
			} else {
				fmt.Fprint(file, "0 ")
			}
		}
		fmt.Fprintln(file)
	}

	return nil
}

// Inverser inverse les couleurs de l'image PBM.
func (pbm *PBM) Invert() {
	for i := 0; i < pbm.height; i++ {
		for j := 0; j < pbm.width; j++ {
			pbm.data[i][j] = !pbm.data[i][j]
		}
	}
}

// Flip retourne l'image PBM horizontalement.
func (pbm *PBM) Flip() {
	for i := 0; i < pbm.height; i++ {
		for j := 0; j < pbm.width/2; j++ {
			pbm.data[i][j], pbm.data[i][pbm.width-j-1] = pbm.data[i][pbm.width-j-1], pbm.data[i][j]
		}
	}
}

// Flop floppe l'image PBM verticalement.
func (pbm *PBM) Flop() {
	for i := 0; i < pbm.height/2; i++ {
		pbm.data[i], pbm.data[pbm.height-i-1] = pbm.data[pbm.height-i-1], pbm.data[i]
	}
}

// SetMagicNumber définit le nombre magique de l'image PBM.
func (pbm *PBM) SetMagicNumber(magicNumber string) {
	pbm.magicNumber = magicNumber
}
