package main

import (
	"fmt"
	"io/ioutil"
)

// Define point
type point struct {
	x int
	y int
}

// Piece Define Piece
type Piece struct {
	PointNum int
	points   []point
}

// Obejct Define
type Obejct struct {
	Hight, weight int
	PiecesNum     int
	Pieces        []*Piece
}

// GetObjects returns the object from dataset file
func GetObjects(filename string) ([]*Obejct, error) {
	// file, err := os.Open(filename)
	// defer file.Close()
	// if err != nil {
	// 	return nil, err
	// }
	// //创建一个从fin读取数据的Reader
	// reader := bufio.NewReader(file)
	ioutil.ReadFile
	bytes, err := ioutil.ReadFile(filename)
	fmt.Println(string(bytes))
	// io
	// var lines []string
	// for line, err := buf.ReadString(' '); ; {
	// 	if err == nil {
	// 		if err == io.EOF { // EOF break loop
	// 			break
	// 		}
	// 		return nil, err
	// 	}
	// 	lines = append(lines, line)
	// }
	return nil, err

}
func main() {
	fmt.Println(GetObjects("dataset/OpTA001C5.txt"))
}

// func read() ([]Header, int) {
// 	//注意这个read是读  /n  需要在data最后一行后面加回车

// 	//var max = 20
// 	//filename := "/mnt/wwwroot/gofile/src/packing/jakobsdata.txt"jakobsdata
// 	filename := "./OpTA005C5.txt"
// 	//filename := "./fudata.txt"
// 	// filename := "/mnt/wwwroot/gofile/src/packing/fudata.txt"
// 	fin, err := os.Open(filename)
// 	defer fin.Close() //quit and close退出并关闭
// 	if err != nil {
// 		fmt.Println(filename, err)
// 		return nil, -1
// 	}
// 	buf := bufio.NewReader(fin) //创建一个从fin读取数据的Reader

// 	header := make([]Header, 30) //读取长度
// 	//	fmt.Println(header)
// 	//将txt文件读取 并将文件中数据进行存储
// 	rectnum := 0

// 	for {

// 		//定义母版的大小
// 		var (
// 			top  float32 = -1
// 			topx float32

// 			bottom  float32 = 10000
// 			bottomx float32

// 			left  float32 = 10000
// 			lefty float32

// 			right  float32 = -1
// 			righty float32
// 		)

// 		v := strings.Fields(line) //读取
// 		var h Header
// 		var p Node
// 		h.Next = &p
// 		//header[rectnum] = h
// 		//fmt.Println(v)

// 		//可在这里对想要的多边形信息进行添加
// 		var pointnum = 1
// 		t, _ := strconv.Atoi(v[0]) //字符串转换整型
// 		p.Data.x = float32(t)      //转换成浮点型

// 		t, _ = strconv.Atoi(v[1])
// 		//p.Data.y = float32(t)

// 		//		fmt.Println("===========")
// 		if p.Data.x < left {
// 			left = p.Data.x
// 			lefty = p.Data.y
// 		}
// 		if p.Data.x > right {
// 			right = p.Data.x
// 			righty = p.Data.y
// 		}

// 		if p.Data.y < bottom {
// 			bottom = p.Data.y
// 			bottomx = p.Data.x
// 		}
// 		if p.Data.y > top {
// 			top = p.Data.y
// 			topx = p.Data.x
// 		}

// 		for i := 2; i < len(v); i += 2 {
// 			var p1 Node
// 			t, _ = strconv.Atoi(v[i])
// 			p1.Data.x = float32(t)

// 			t, _ = strconv.Atoi(v[i+1])
// 			p1.Data.y = float32(t)

// 			if p1.Data.x < left {
// 				left = p1.Data.x
// 				lefty = p1.Data.y
// 			}
// 			if p1.Data.x > right {
// 				right = p1.Data.x
// 				righty = p1.Data.y
// 			}

// 			if p1.Data.y < bottom {
// 				bottom = p1.Data.y
// 				bottomx = p1.Data.x
// 			}
// 			if p1.Data.y > top {
// 				top = p1.Data.y
// 				topx = p1.Data.x
// 			}

// 			Insert(&p, &p1)
// 			pointnum++
// 		}
// 		h.pointnum = pointnum
// 		h.height = top - bottom
// 		h.width = right - left

// 		h.position.bottom.x = bottomx
// 		h.position.bottom.y = bottom

// 		h.position.left.x = left
// 		h.position.left.y = lefty

// 		h.position.right.x = right
// 		h.position.right.y = righty

// 		h.position.top.x = topx
// 		h.position.top.y = top

// 		header[rectnum] = h
// 		header[rectnum].area = Areacal(header, rectnum)
// 		rectnum++

// 	}

// }
