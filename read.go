package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Point Define point
type Point struct {
	x int
	y int
}

// Piece Define Piece
type Piece struct {
	PointNum int
	Points   []Point
}

// Obejct Define
type Obejct struct {
	Hight, Weight int
	PiecesNum     int
	Pieces        []*Piece
}

// GetObjects returns the object from dataset file
func GetObjects(filename string) ([]*Obejct, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines, err := contentStr2Arry(string(bytes))
	if err != nil {
		return nil, err
	}

	var objs []*Obejct
	readLineNum := 2
	// mutile objects
	if len(lines[0]) != 1 {
		for i := 1; i <= lines[0][0]; i++ {
			objs = append(objs, getObject(lines, lines[0][i], readLineNum, readLineNum+lines[0][i]))
			readLineNum += lines[0][i]
		}
	} else {
		objs = append(objs, getObject(lines, lines[0][0], readLineNum, readLineNum+lines[0][0]))
	}
	return objs, err
}

// getObject
func getObject(lines [][]int, piecesNum, start, end int) (obj *Obejct) {
	obj = &Obejct{}
	obj.Hight, obj.Weight = lines[1][0], lines[1][1]
	obj.PiecesNum = piecesNum
	for _, line := range lines[start:end] {
		tempPiece := Piece{}
		tempPiece.PointNum = line[0]
		for i := 1; i < len(line); {
			tempPoint := Point{
				x: line[i],
				y: line[i+1],
			}
			tempPiece.Points = append(tempPiece.Points, tempPoint)
			i += 2
		}
		obj.Pieces = append(obj.Pieces, &tempPiece)
	}
	return
}

// contentStr2Arry txt文本字符串转为二维数组
func contentStr2Arry(conStr string) (lines [][]int, err error) {
	lineStrs := strings.Split(strings.TrimSpace(conStr), "\n")
	for _, lineStr := range lineStrs {
		lineStr = strings.TrimSpace(lineStr)
		strs := strings.Split(lineStr, " ")
		line := make([]int, len(strs))
		for i, str := range strs {
			line[i], err = strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
		}
		lines = append(lines, line)
	}
	return
}

func main() {
	objs, err := GetObjects("dataset/OpTA001C5.txt")
	if err != nil {
		log.Println(err)
	}
	PrintObjs(objs)
}

// PrintObjs p
func PrintObjs(objs []*Obejct) {
	for _, obj := range objs {
		log.Println("--------------\n--------------\n")
		log.Println("Hight:", obj.Hight, "Weight:", obj.Weight, "PiecesNum:", obj.PiecesNum)
		for _, piece := range obj.Pieces {
			log.Println("\nPointNum:", piece.PointNum)
			for _, point := range piece.Points {
				log.Println("point:", point)
			}
		}
	}
}
