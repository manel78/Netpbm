package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type PBM struct {
	data          [][]bool
	width, height int
	magicNumber   string
}

func ReadPBM(filename string) (*PBM, error) {
	file, err := os.Open("image.pbm")
	if err != nil {
		return nil, err
	}
	defer file.Close() // ferme le fichier

	scanner := bufio.NewScanner(file) // lire le fichier mot par mot
	scanner.Split(bufio.ScanWords)

	var magicNumber string // lire le magicnumber
	scanner.Scan()
	magicNumber = scanner.Text()

	scanner.Scan() // lire la largeur
	width, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, err
	}

	scanner.Scan() // lire la hauteur
	height, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, err
	}

	data := make([][]bool, height) // stockers les valeurs de pixels
	for i := 0; i < height; i++ {
		data[i] = make([]bool, width)
		for j := 0; j < width; j++ {
			scanner.Scan()
			pixelValue := scanner.Text()
			data[i][j] = pixelValue == "1"
		}
	}

	return &PBM{data, width, height, magicNumber}, nil
}

func (pbm *PBM) Size() (int, int) { // la largeur et la hauteur sont retourner
	return pbm.width, pbm.height
}

func (pbm *PBM) At(x, y int) bool { // retourne la valeur du pixel x et y
	return pbm.data[y][x]
}

func (pbm *PBM) Set(x, y int, value bool) { // remet la valeur d'origine de x et y
	pbm.data[y][x] = value
}

func (pbm *PBM) Save(filename string) error { // enregistre l image pbm
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close() // ferme le fichier lorsque l'on a fini

	writer := bufio.NewWriter(file) // comme un editeur de fichier

	fmt.Fprintf(writer, "%s\n%d %d\n", pbm.magicNumber, pbm.width, pbm.height) // la hauteur et largeur  dans le fichier

	for _, row := range pbm.data { // parcours les 0 et les 1 pour connaitre les pixels blancs et noirs
		for _, val := range row {
			if val {
				fmt.Fprintf(writer, "1 ")
			} else {
				fmt.Fprintf(writer, "0 ")
			}
		}
		fmt.Fprintln(writer) // nouvelle ligne pour nouvelle image
	}

	return writer.Flush()
}

func (pbm *PBM) Invert() { // inverse les pixels i et j
	for i := 0; i < pbm.height; i++ {
		for j := 0; j < pbm.width; j++ {
			pbm.data[i][j] = !pbm.data[i][j]
		}
	}
}

func (pbm *PBM) Flip() { // inverse l'ordre des pixels horizontalement
	for i := 0; i < pbm.height; i++ {
		for j, k := 0, pbm.width-1; j < k; j, k = j+1, k-1 {
			pbm.data[i][j], pbm.data[i][k] = pbm.data[i][k], pbm.data[i][j]
		}
	}
}

func (pbm *PBM) Flop() { // inversant l'ordre des pixels verticalement
	for i := 0; i < pbm.height/2; i++ {
		pbm.data[i], pbm.data[pbm.height-i-1] = pbm.data[pbm.height-i-1], pbm.data[i]
	}
}

func (pbm *PBM) SetMagicNumber(magicNumber string) { // met a jour le numero magique
	pbm.magicNumber = magicNumber
}
