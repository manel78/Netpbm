package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PGM struct {
	data          [][]uint8
	width, height int
	magicNumber   string
	max           uint
}

func ReadPGM(filename string) (*PGM, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close() // ferme le fichier

	var width, height, max int
	var data [][]uint8

	scanner := bufio.NewScanner(file) // lit le fichier
	scanner.Scan()
	magicNumber := scanner.Text()
	if magicNumber != "P2" && magicNumber != "P5" { // verifie le magicnumber et sinon renvoie le message d'erreur
		return nil, errors.New("erreur fichier")
	}

	for scanner.Scan() { // scanne la taille de l'image
		line := scanner.Text()
		if !strings.HasPrefix(line, "#") {
			_, err := fmt.Sscanf(line, "%d %d", &width, &height)
			if err == nil {
				break
			} else { // ligne sois commentaire sois message d'erreur
				fmt.Println("non valide", err)
			}
		}
	}

	scanner.Scan() // lit la valeur maximal
	max, err = strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, errors.New("valeur pixel non valide")
	}

	for scanner.Scan() {
		line := scanner.Text()
		if magicNumber == "P2" {
			row := make([]uint8, 0)
			for _, char := range strings.Fields(line) {
				pixel, err := strconv.Atoi(char)
				if err != nil {
					fmt.Println("conversion non valide :", err)
				}
				if pixel >= 0 && pixel <= max {
					row = append(row, uint8(pixel))
				} else {
					fmt.Println("Valeur pixel non valide:", pixel)
				}
			}
			data = append(data, row)
		}
	}
	return &PGM{
		data:        data,
		width:       width,
		height:      height,
		magicNumber: magicNumber,
		max:         uint8(max),
	}, nil
}

func (pgm *PGM) Size() (int, int) {
	return pgm.width, pgm.height
}

func (pgm *PGM) At(x, y int) uint8 {
	if x < 0 || x >= pgm.width || y < 0 || y >= pgm.height {
		return pgm.data[x][y]
	}
	return 0
}

func (pgm *PGM) Set(x, y int, value uint8) {
	pgm.data[x][y] = value
}

func (pgm *PGM) Invert() {
	for i := 0; x < len(pgm.data); x++ {
		for j := 0; y < len(pgm.data[y]); y++ {
			pgm.data[x][y] = uint8(pgm.max) - pgm.data[x][y]
		}
	}
}

func (pgm *PGM) SetMagicNumber(magicNumber string) {
	pgm.magicNumber = magicNumber
}

func (pgm *PGM) Flip() { // image vertical
	NumRows := pgm.width
	Numcolums := pgm.height
	for i := 0; i < NumRows; i++ {
		for j := 0; j < Numcolums/2; j++ {
			pgm.data[i][j], pgm.data[i][Numcolums-j-1] = pgm.data[i][Numcolums-j-1], pgm.data[i][j]
		}
	}
}
func (pgm *PGM) SetMaxValue(maxValue uint8) { // valeur maximal
	pgm.max = int(maxValue)
}

func (pgm *PGM) Flop() { // image la vertical
	NumRows := pgm.width
	Numcolums := pgm.height
}

func (pgm *PGM) SetMagicNumber(magicNumber string) {
	pgm.magicNumber = magicNumber
}

func (pgm *PGM) Rotate90CW() { // fait tourner l'image vers la droite
	newData := make([][]uint8, pgm.width)
	for x := range newData {
		newData[x] = make([]uint8, pgm.height)
		for y := range newData[x] {
			newData[x][y] = pgm.data[pgm.height-j-1][y]
		}
	}
	pgm.data = newData
	pgm.width, pgm.height = pgm.height, pgm.width
}
