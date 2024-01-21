package main

import (
	"Netpbm/pbm"
	"Netpbm/pgm"
	"Netpbm/ppm"
	"fmt"
	"log"
)

func main() {
	// Lire une image PBM depuis un fichier
	imagePBM, err := pbm.ReadPBM("canard.pbm")
	if err != nil {
		log.Fatal("Erreur canard PBM :", err)
	}

	// Obtenir la largeur et la hauteur de l'image PBM
	width, height := imagePBM.Size()
	fmt.Printf("Largeur : %d, Hauteur : %d\n", width, height)

	// Obtenir la valeur du pixel en (0, 0)
	value := imagePBM.At(0, 0)
	fmt.Println("Valeur du pixel en (0, 0) :", value)

	// Enregistrer l'image PBM modifiée dans un fichier
	err = imagePBM.Save("savedPBM.pbm")
	if err != nil {
		log.Fatal("Erreur savedPBM :", err)
	}

	// Modifier la valeur du pixel en (1, 1)
	imagePBM.Set(1, 1, true)

	// Inverser les couleurs de l'image PBM
	imagePBM.Invert()
	err = imagePBM.Save("invertedPBM.pbm")
	if err != nil {
		log.Fatal("Erreur invertedPBM :", err)
	}

	// Retourner horizontalement l'image PBM
	imagePBM.Flip()
	err = imagePBM.Save("flippedPBM.pbm")
	if err != nil {
		log.Fatal("Erreur flippedPBM :", err)
	}

	// Retourner verticalement l'image PBM
	imagePBM.Flop()
	err = imagePBM.Save("floppedPBM.pbm")
	if err != nil {
		log.Fatal("Erreur floppedPBM :", err)
	}

	// Changer le numéro magique de l'image PBM
	imagePBM.SetMagicNumber("P1")
	err = imagePBM.Save("MagicNumberPBM.pbm")
	if err != nil {
		log.Fatal("Erreur MagicNumberPBM :", err)
	}

	// Lire une image PGM depuis un fichier
	imagePGM, err := pgm.ReadPGM("canard.pgm")
	if err != nil {
		log.Fatal("Erreur canard PGM :", err)
	}

	// Obtenir la largeur et la hauteur de l'image PGM
	width, height = imagePGM.Size()
	fmt.Printf("Largeur : %d, Hauteur : %d\n", width, height)

	// Enregistrer l'image PGM modifiée dans un fichier
	err = imagePGM.Save("savedPGM.pgm")
	if err != nil {
		log.Fatal("Erreur savedPGM :", err)
	}

	// Obtenir la valeur du pixel en (2, 3)
	pixelValue := imagePGM.At(2, 3)
	fmt.Println("Valeur du pixel en (2, 3) :", pixelValue)

	// Inverser les couleurs de l'image PGM
	imagePGM.Invert()
	err = imagePGM.Save("invertedPGM.pgm")
	if err != nil {
		log.Fatal("Erreur invertedPGM :", err)
	}

	// Retourner horizontalement l'image PGM
	imagePGM.Flip()
	err = imagePGM.Save("flippedPGM.pgm")
	if err != nil {
		log.Fatal("Erreur flippedPGM :", err)
	}

	// Retourner verticalement l'image PGM
	imagePGM.Flop()
	err = imagePGM.Save("floppedPGM.pgm")
	if err != nil {
		log.Fatal("Erreur floppedPGM :", err)
	}

	// Changer le numéro magique de l'image PGM
	imagePGM.SetMagicNumber("P2")

	// Lire une image PPM depuis un fichier
	imagePPM, err := ppm.ReadPPM("canard.ppm")
	if err != nil {
		log.Fatal("Erreur canard PPM :", err)
	}

	// Obtenir la largeur et la hauteur de l'image PPM
	width, height = imagePPM.Size()
	fmt.Printf("Largeur : %d, Hauteur : %d\n", width, height)

	// Obtenir la valeur du pixel en (0, 0)
	Vpixel := imagePPM.At(0, 0)
	fmt.Printf("Valeur du pixel en (0, 0) : %+v\n", Vpixel)

	// Enregistrer l'image PPM modifiée dans un fichier
	err = imagePPM.Save("savedPPM.ppm")
	if err != nil {
		log.Fatal("Erreur savedPPM :", err)
	}

	// Définir un nouveau pixel à la position (1, 1)
	newPixel := ppm.Pixel{R: 100, G: 150, B: 200}
	imagePPM.Set(1, 1, newPixel)

	// Inverser les couleurs de l'image
	imagePPM.Invert()
	err = imagePPM.Save("invertedPPM.ppm")
	if err != nil {
		log.Fatal("Erreur invertedPPM :", err)
	}

	// Retourner l'image horizontalement
	imagePPM.Flip()
	err = imagePPM.Save("flippedPPM.ppm")
	if err != nil {
		log.Fatal("Erreur flippedPPM :", err)
	}

	// Retourner l'image verticalement
	imagePPM.Flop()
	err = imagePPM.Save("floppedPPM.ppm")
	if err != nil {
		log.Fatal("Erreur floppedPPM :", err)
	}

	// Changer le numéro magique de l'image en "P3"
	imagePPM.SetMagicNumber("P3")
	err = imagePPM.Save("MagicNumberPPM.ppm")
	if err != nil {
		log.Fatal("Erreur MagicNumberPPM :", err)
	}

	// Définir la valeur maximale de couleur à 255
	imagePPM.SetMaxValue(255)

	// Faire une rotation de 90 degrés dans le sens des aiguilles d'une montre
	imagePPM.Rotate90CW()
	err = imagePPM.Save("rotatedPPM.ppm")
	if err != nil {
		log.Fatal("Erreur rotatedPPM :", err)
	}

	// Créer une nouvelle image PPM de 100x100 et dessiner une ligne rouge au centre
	imagePPM = ppm.NewPPM(100, 100)
	imagePPM.DrawLine(ppm.Point{0, 50}, ppm.Point{99, 50}, ppm.Pixel{255, 0, 0})
	err = imagePPM.Save("Line.ppm")
	if err != nil {
		log.Fatal("Erreur Line :", err)
	}

	// Créer une nouvelle image PPM de 100x100 et dessiner un rectangle blanc
	imagePPM = ppm.NewPPM(100, 100)
	imagePPM.DrawRectangle(ppm.Point{20, 30}, 60, 40, ppm.Pixel{255, 255, 255})
	err = imagePPM.Save("Rectangle.ppm")
	if err != nil {
		log.Fatal("Erreur Rectangle :", err)
	}

	// Créer une nouvelle image PPM de 100x100 et dessiner un rectangle blanc rempli
	imagePPM = ppm.NewPPM(100, 100)
	imagePPM.DrawFilledRectangle(ppm.Point{20, 30}, 60, 40, ppm.Pixel{255, 255, 255})
	err = imagePPM.Save("FilledRectangle.ppm")
	if err != nil {
		log.Fatal("Erreur FilledRectangle :", err)
	}

	// Créer une nouvelle image PPM de 100x100 et dessiner un cercle blanc
	imagePPM = ppm.NewPPM(100, 100)
	imagePPM.DrawCircle(ppm.Point{50, 50}, 20, ppm.Pixel{R: 255, G: 255, B: 255})
	err = imagePPM.Save("circle.ppm")
	if err != nil {
		log.Fatal("Erreur circle :", err)
	}

	// Créer une nouvelle image PPM de 100x100 et dessiner un cercle blanc rempli
	imagePPM = ppm.NewPPM(100, 100)
	imagePPM.DrawFilledCircle(ppm.Point{50, 50}, 20, ppm.Pixel{R: 255, G: 255, B: 255})
	err = imagePPM.Save("FilledCircle.ppm")
	if err != nil {
		log.Fatal("Erreur FilledCircle :", err)
	}

	// Créer une nouvelle image PPM de 100x100 et dessiner un triangle blanc
	imagePPM = ppm.NewPPM(100, 100)
	imagePPM.DrawTriangle(ppm.Point{20, 80}, ppm.Point{80, 80}, ppm.Point{50, 20}, ppm.Pixel{R: 255, G: 255, B: 255})
	err = imagePPM.Save("Triangle.ppm")
	if err != nil {
		log.Fatal("Erreur Triangle :", err)
	}

	// Créer une nouvelle image PPM de 100x100 et dessiner un triangle blanc rempli
	imagePPM = ppm.NewPPM(100, 100)
	imagePPM.DrawFilledTriangle(ppm.Point{20, 80}, ppm.Point{80, 80}, ppm.Point{50, 20}, ppm.Pixel{R: 255, G: 255, B: 255})
	err = imagePPM.Save("FilledTriangle.ppm")
	if err != nil {
		log.Fatal("Erreur FilledTriangle :", err)
	}

	// Créer une nouvelle image PPM de 100x100 et dessiner un polygone blanc
	imagePPM = ppm.NewPPM(100, 100)
	polygonPoints := []ppm.Point{{20, 80}, {80, 80}, {50, 20}, {30, 40}}
	imagePPM.DrawPolygon(polygonPoints, ppm.Pixel{R: 255, G: 255, B: 255})
	err = imagePPM.Save("Polygon.ppm")
	if err != nil {
		log.Fatal("Erreur Polygon :", err)
	}

	// Créez une autre image PPM et le polygone rempli
	imagePPM = ppm.NewPPM(100, 100)
	imagePPM.DrawFilledPolygon(polygonPoints, ppm.Pixel{R: 255, G: 255, B: 255})
	imagePPM.Save("FilledPolygon.ppm")

	// Créez une image PPM et Dessine le flocon de neige de Koch
	imagePPM = ppm.NewPPM(400, 400)
	startPoint := ppm.Point{X: 50, Y: 350}
	imagePPM.DrawKochSnowflake(3, startPoint, 300, ppm.Pixel{R: 255, G: 255, B: 255})
	err = imagePPM.Save("KochSnowflake.ppm")
	if err != nil {
		log.Fatal("Erreur KochSnowflake :", err)
	}

	// Créez une image PPM et Dessine le triangle de Sierpinski
	imagePPM = ppm.NewPPM(100, 100)
	imagePPM.DrawSierpinskiTriangle(4, ppm.Point{25, 75}, 50, ppm.Pixel{255, 0, 0})
	imagePPM.Save("SierpinskiTriangle.ppm")

}
