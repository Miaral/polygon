package main

type tabonode struct {
	//禁忌表的结点  局部优化
	node1  int
	node2  int
	result float32
	Next   *tabonode
}

type Rect_position struct {
	top    point
	bottom point
	left   point
	right  point
}

type Node struct {
	Data point
	Next *Node
} //结构体嵌套

type Header struct {
	// if point set
	flag     int
	pointnum int
	width    float32
	height   float32

	//当前多边形排放位置的上下左右极点坐标
	position Rect_position

	Next *Node //当前多边形的坐标点
	area float32
}

// func main() {

// 	var (
// 		packednum         = 0  //已排项数目
// 		Height    float32 = 40 // jakob1 排料带高度
// 		// Height float32 = 38 //fu  排料带高度

// 		maxx float32 = 0 //

// 		totalarea float32 //当前已排料的总面积
// 	)

// 	//读取数据
// 	header, rectnum := read()

// 	//对所有读取的项转四下  选取最宽的时候作为进去的格式
// 	bestrightchose(header, Height, rectnum)

// 	//按照宽度排序
// 	for i := 0; i < rectnum; i++ {
// 		var flag = 0
// 		for j := rectnum - 1; j > i; j-- {
// 			if header[i].width < header[j].width {
// 				flag = 1
// 				header[i], header[j] = header[j], header[i]
// 			}
// 		}
// 		if flag == 0 {
// 			break
// 		}
// 	}

// 	var xvalue []float32
// 	header[0].flag = 1
// 	p := header[0].Next

// 	for i := 0; i < header[0].pointnum; i++ {
// 		if !checkrepeat(p.Data.x, xvalue) {
// 			xvalue = append(xvalue, p.Data.x)
// 		}
// 		p = p.Next
// 	}
// 	sortxvalue(xvalue)
// 	totalarea = header[0].area
// 	packednum++

// 	fmt.Println("读取多边形数目", rectnum)

// 	maxx = BLPacking(header, xvalue, Height, rectnum, packednum)
// 	var testnum = 0

// 	// dc := gg.NewContext(1000, int(Height)*10)
// 	dc := gg.NewContext(1000, 700)
// 	dc.SetRGB(0, 0, 0)
// 	dc.Clear()
// 	for i := 0; i < rectnum; i++ {
// 		//dc.DrawLine(0, float64(50*i), 300, float64(50*i))
// 		if header[i].flag == 0 {
// 			continue
// 		} else {
// 			p := header[i].Next
// 			p1 := p.Next
// 			for j := 0; j < header[i].pointnum; j++ {
// 				if j == header[i].pointnum-1 {
// 					p1 = header[i].Next
// 				}
// 				dc.SetRGB(1, 1, 1)
// 				dc.DrawLine(float64(p.Data.x)*10, float64(p.Data.y)*10, float64(p1.Data.x)*10, float64(p1.Data.y)*10)
// 				dc.Stroke()

// 				p = p1
// 				p1 = p1.Next
// 			}
// 			testnum++
// 		}
// 	}
// 	dc.DrawLine(0, float64(10)*float64(Height), 300, float64(10)*float64(Height))
// 	dc.Stroke()
// 	dc.SavePNG("testA.png")
// 	// dc.SavePNG("fudata.png")
// 	// dc.SavePNG("normaldata.png")
// 	fmt.Println("num", testnum)

// 	fmt.Println("排料率！！！")

// 	var area float32

// 	fmt.Println(maxx, Height)
// 	totalarea = maxx * Height
// 	for i := 0; i < rectnum; i++ {
// 		if header[i].flag == 0 {
// 			continue
// 		}

// 		area += header[i].area

// 	}

// 	fmt.Println(area/totalarea, area, totalarea)

// }
