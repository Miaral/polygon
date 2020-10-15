package main

import "fmt"

// Uppermove Define,polygon moves from top to bottom
func Uppermove(head []Header, notpack int, packed int, height float32) (float32, float32, float32, float32) {
	var minmove = height
	var mincurrentheight float32
	var kofslope float32
	var currentx float32

	notpackrect := head[notpack]
	packedrect := head[packed]

	nodenotpack := notpackrect.Next

	//向下平移到轴外
	for nodenotpack != nil {
		nodenotpack.Data.y -= head[notpack].position.top.y
		nodenotpack = nodenotpack.Next
	}

	nodenotpack = notpackrect.Next
	nodenotpack1 := nodenotpack.Next

	nodepacked := packedrect.Next
	nodepacked1 := nodepacked.Next

	var (
		dvaluex float32
		dvaluey float32
	)

	//上面的packed对下面notpack求距离
	for i := 0; i < packedrect.pointnum; i++ {
		nodenotpack = notpackrect.Next
		nodenotpack1 = nodenotpack.Next

		for j := 0; j < notpackrect.pointnum; j++ {
			if j == notpackrect.pointnum-1 {
				nodenotpack1 = notpackrect.Next
			}
			if (nodepacked.Data.x < MinNumber(nodenotpack.Data.x, nodenotpack1.Data.x)) || (nodepacked.Data.x > MaxNumber(nodenotpack.Data.x, nodenotpack1.Data.x)) {
				//当前packed点不在notpack直线的竖向空间中 直接跳过 算下一条
				nodenotpack = nodenotpack1
				nodenotpack1 = nodenotpack1.Next
				continue
			}

			dvaluex = nodenotpack.Data.x - nodenotpack1.Data.x
			dvaluey = nodenotpack.Data.y - nodenotpack1.Data.y

			var move float32
			var k float32

			if IsEqual(dvaluex, 0) && IsEqual(nodepacked.Data.x, nodenotpack.Data.x) {
				//若dx=0 且packed点与notpack点x坐标相同 那么滑动斜率为999
				k = 999
				move = nodepacked.Data.y - MaxNumber(nodenotpack.Data.y, nodenotpack1.Data.y)

			} else {
				k = dvaluey / dvaluex
				b := nodenotpack.Data.y - k*nodenotpack.Data.x
				move = nodepacked.Data.y - (k*nodepacked.Data.x + b)

			}
			if IsEqual(move, 0) {
				move = 0
			}

			if move < minmove {
				minmove = move
				mincurrentheight = nodepacked.Data.y
				currentx = nodepacked.Data.x
				kofslope = k

			}

			nodenotpack = nodenotpack1
			nodenotpack1 = nodenotpack1.Next
		}
		nodepacked = nodepacked.Next
	}

	nodenotpack = notpackrect.Next
	nodenotpack1 = nodenotpack.Next

	for i := 0; i < notpackrect.pointnum; i++ {
		nodepacked = packedrect.Next
		nodepacked1 = nodepacked.Next

		for j := 0; j < packedrect.pointnum; j++ {
			if j == packedrect.pointnum-1 {
				nodepacked1 = packedrect.Next
			}

			if (nodenotpack.Data.x < MinNumber(nodepacked.Data.x, nodepacked1.Data.x)) || (nodenotpack.Data.x > MaxNumber(nodepacked.Data.x, nodepacked1.Data.x)) {
				//notpack上点 不再packed直线上
				nodepacked = nodepacked1
				nodepacked1 = nodepacked1.Next
				continue
			}

			dvaluex = nodepacked.Data.x - nodepacked1.Data.x
			dvaluey = nodepacked.Data.y - nodepacked1.Data.y

			var move float32
			var k float32

			if IsEqual(dvaluex, 0) && IsEqual(nodenotpack.Data.x, nodepacked.Data.x) {
				k = 999
				move = MinNumber(nodepacked.Data.y, nodepacked1.Data.y) - nodenotpack.Data.y
			} else {
				k = dvaluey / dvaluex
				b := nodepacked.Data.y - k*nodepacked.Data.x
				move = k*nodenotpack.Data.x + b - nodenotpack.Data.y
			}
			if IsEqual(move, 0) {
				move = 0
			}
			if move < minmove {
				minmove = move
				mincurrentheight = nodenotpack.Data.y + move
				currentx = nodenotpack.Data.x
				kofslope = k
			}

			nodepacked = nodepacked1
			nodepacked1 = nodepacked1.Next

		}
		nodenotpack = nodenotpack.Next

	}

	nodenotpack = notpackrect.Next

	//向下平移到轴外
	for nodenotpack != nil {
		nodenotpack.Data.y += head[notpack].position.top.y
		nodenotpack = nodenotpack.Next
	}

	//注意 这个地方 按理说应该剪掉高度  但是我们下面InsertPacking的地方需要横轴到目标点的直线距离
	//所以在需要平移距离的地方减一下
	//minmove -= head[notpack].position.top.y
	//??这个地方 原来写法是 minmove += head[notpack].position.bottom.y

	return minmove, mincurrentheight, currentx, kofslope
}

func Uppermove2(head []Header, notpack int, packed int, Height float32) (float32, float32) {
	var (
		minmove  = Height
		kofslope float32
	)

	var (
		dvaluex float32
		dvaluey float32
	)

	packedrect := head[packed]
	notpackrect := head[notpack]

	nodenotpack := notpackrect.Next
	nodenotpack1 := nodenotpack.Next

	nodepacked := packedrect.Next
	nodepacked1 := nodepacked.Next

	//上面的packed对下面notpack求距离
	for i := 0; i < packedrect.pointnum; i++ {

		nodenotpack = notpackrect.Next
		nodenotpack1 = nodenotpack.Next

		for j := 0; j < notpackrect.pointnum; j++ {
			if j == notpackrect.pointnum-1 {
				nodenotpack1 = notpackrect.Next
			}
			if (nodepacked.Data.x < MinNumber(nodenotpack.Data.x, nodenotpack1.Data.x)) || (nodepacked.Data.x >= MaxNumber(nodenotpack.Data.x, nodenotpack1.Data.x)) {
				//当前packed点不在notpack直线的竖向空间中 直接跳过 算下一条
				nodenotpack = nodenotpack1
				nodenotpack1 = nodenotpack1.Next
				continue
			}

			dvaluex = nodenotpack.Data.x - nodenotpack1.Data.x
			dvaluey = nodenotpack.Data.y - nodenotpack1.Data.y

			var move float32
			var k float32

			if IsEqual(dvaluex, 0) && IsEqual(nodepacked.Data.x, nodenotpack.Data.x) {
				//若dx=0 且packed点与notpack点x坐标相同 那么滑动斜率为999
				k = 999
				move = nodepacked.Data.y - MaxNumber(nodenotpack.Data.y, nodenotpack1.Data.y)

			} else {
				k = dvaluey / dvaluex
				b := nodenotpack.Data.y - k*nodenotpack.Data.x
				move = nodepacked.Data.y - (k*nodepacked.Data.x + b)

			}

			if IsEqual(move, 0) {

				move = 0
			} else if move < 0 {
				nodenotpack = nodenotpack1
				nodenotpack1 = nodenotpack1.Next
				continue
			}

			if move < minmove {
				minmove = move
				kofslope = k

			}

			nodenotpack = nodenotpack1
			nodenotpack1 = nodenotpack1.Next
		}
		nodepacked = nodepacked.Next
	}

	nodenotpack = notpackrect.Next

	for i := 0; i < notpackrect.pointnum; i++ {
		nodepacked = packedrect.Next
		nodepacked1 = nodepacked.Next

		for j := 0; j < packedrect.pointnum; j++ {
			if j == packedrect.pointnum-1 {
				nodepacked1 = packedrect.Next
			}

			if (nodenotpack.Data.x < MinNumber(nodepacked.Data.x, nodepacked1.Data.x)) || (nodenotpack.Data.x > MaxNumber(nodepacked.Data.x, nodepacked1.Data.x)) {
				//notpack上点 不再packed直线上
				nodepacked = nodepacked1
				nodepacked1 = nodepacked1.Next
				continue
			}

			dvaluex = nodepacked.Data.x - nodepacked1.Data.x
			dvaluey = nodepacked.Data.y - nodepacked1.Data.y

			var move float32
			var k float32

			if IsEqual(dvaluex, 0) && IsEqual(nodenotpack.Data.x, nodepacked.Data.x) {
				k = 999
				move = MinNumber(nodepacked.Data.y, nodepacked1.Data.y) - nodenotpack.Data.y
			} else {
				k = dvaluey / dvaluex
				b := nodepacked.Data.y - k*nodepacked.Data.x
				move = k*nodenotpack.Data.x + b - nodenotpack.Data.y
			}

			if IsEqual(move, 0) {
				move = 0
			} else if move < 0 {
				nodepacked = nodepacked1
				nodepacked1 = nodepacked1.Next
				continue
			}

			if move < minmove {
				minmove = move
				kofslope = k
			}

			nodepacked = nodepacked1
			nodepacked1 = nodepacked1.Next

		}
		nodenotpack = nodenotpack.Next
	}

	return minmove, kofslope
}

func Belowmove(head []Header, notpack int, packed int, height float32) (float32, float32, float32, float32) {
	var minmove = height
	var mincurrentheight float32
	var currentx float32
	var kofslope float32
	//供slope时调用 最低点接触点的直线斜率

	notpackrect := head[notpack]
	packedrect := head[packed]

	nodenotpack := notpackrect.Next

	for nodenotpack != nil {
		nodenotpack.Data.y += height
		nodenotpack = nodenotpack.Next
	}

	nodenotpack = notpackrect.Next
	nodenotpack1 := nodenotpack.Next

	nodepacked := packedrect.Next
	nodepacked1 := nodepacked.Next

	var (
		dvaluex float32
		dvaluey float32
	)

	//上面的notpackrect对下面的packedrect进行距离计算
	for i := 0; i < notpackrect.pointnum; i++ {

		nodepacked = packedrect.Next
		nodepacked1 = nodepacked.Next

		for j := 0; j < packedrect.pointnum; j++ {
			if j == packedrect.pointnum-1 {
				nodepacked1 = packedrect.Next
			}

			if (nodenotpack.Data.x < MinNumber(nodepacked.Data.x, nodepacked1.Data.x)) || (nodenotpack.Data.x > MaxNumber(nodepacked.Data.x, nodepacked1.Data.x)) {
				nodepacked = nodepacked1
				nodepacked1 = nodepacked1.Next
				continue
			}

			dvaluex = nodepacked.Data.x - nodepacked1.Data.x
			dvaluey = nodepacked.Data.y - nodepacked1.Data.y

			var move float32
			var k float32

			// var dvalueofcheck = nodenotpack.Data.x - nodenotpack1.Data.x

			if IsEqual(dvaluex, 0) && IsEqual(nodepacked.Data.x, nodenotpack.Data.x) {

				move = nodenotpack.Data.y - MaxNumber(nodepacked.Data.y, nodepacked1.Data.y)
				k = 999
			} else {
				k = dvaluey / dvaluex
				b := nodepacked.Data.y - k*nodepacked.Data.x
				move = nodenotpack.Data.y - (k*nodenotpack.Data.x + b)
			}
			if IsEqual(move, 0) {
				move = 0
			}

			if move < minmove {
				minmove = move
				kofslope = k
				mincurrentheight = nodenotpack.Data.y - move
				currentx = nodenotpack.Data.x
			}
			nodepacked = nodepacked1
			nodepacked1 = nodepacked1.Next
		}
		nodenotpack = nodenotpack.Next

	}

	nodepacked = packedrect.Next
	nodepacked1 = nodepacked.Next

	for i := 0; i < packedrect.pointnum; i++ {

		nodenotpack = notpackrect.Next
		nodenotpack1 = nodenotpack.Next

		for j := 0; j < notpackrect.pointnum; j++ {
			if j == notpackrect.pointnum-1 {
				nodenotpack1 = notpackrect.Next
			}
			if (nodepacked.Data.x < MinNumber(nodenotpack.Data.x, nodenotpack1.Data.x)) || (nodepacked.Data.x > MaxNumber(nodenotpack.Data.x, nodenotpack1.Data.x)) {
				nodenotpack = nodenotpack1
				nodenotpack1 = nodenotpack1.Next
				continue
			}

			dvaluex = nodenotpack.Data.x - nodenotpack1.Data.x
			dvaluey = nodenotpack.Data.y - nodenotpack1.Data.y

			var move float32
			var k float32

			if IsEqual(dvaluex, 0) && IsEqual(nodepacked.Data.x, nodenotpack.Data.x) {

				move = MinNumber(nodenotpack.Data.y, nodenotpack1.Data.y) - nodepacked.Data.y
				k = 999
			} else {
				k = dvaluey / dvaluex
				b := nodenotpack.Data.y - k*nodenotpack.Data.x
				move = k*nodepacked.Data.x + b - nodepacked.Data.y
			}
			if IsEqual(move, 0) {
				move = 0
			}
			if move < minmove {
				minmove = move
				kofslope = k
				mincurrentheight = nodepacked.Data.y
				currentx = nodepacked.Data.x
			}
			nodenotpack = nodenotpack1
			nodenotpack1 = nodenotpack1.Next
		}
		nodepacked = nodepacked.Next

	}

	nodenotpack = notpackrect.Next

	for nodenotpack != nil {
		nodenotpack.Data.y -= height
		nodenotpack = nodenotpack.Next
	}

	minmove = height - minmove
	//因为加了height 所以实际高度是反向的差值

	return minmove, mincurrentheight, currentx, kofslope
}

func Belowmove2(head []Header, notpack int, packed int, height float32) (float32, float32, float32, float32) {
	var minmove = height
	var mincurrentheight float32
	var currentx float32
	var kofslope float32
	//供slope时调用 最低点接触点的直线斜率

	notpackrect := head[notpack]
	packedrect := head[packed]

	nodenotpack := notpackrect.Next
	nodenotpack1 := nodenotpack.Next

	nodepacked := packedrect.Next
	nodepacked1 := nodepacked.Next

	var (
		dvaluex float32
		dvaluey float32
	)

	//上面的notpackrect对下面的packedrect进行距离计算
	for i := 0; i < notpackrect.pointnum; i++ {

		nodepacked = packedrect.Next
		nodepacked1 = nodepacked.Next

		for j := 0; j < packedrect.pointnum; j++ {
			if j == packedrect.pointnum-1 {
				nodepacked1 = packedrect.Next
			}

			if (nodenotpack.Data.x < MinNumber(nodepacked.Data.x, nodepacked1.Data.x)) || (nodenotpack.Data.x > MaxNumber(nodepacked.Data.x, nodepacked1.Data.x)) {

				nodepacked = nodepacked1
				nodepacked1 = nodepacked1.Next
				continue
			}

			dvaluex = nodepacked.Data.x - nodepacked1.Data.x
			dvaluey = nodepacked.Data.y - nodepacked1.Data.y

			var move float32
			var k float32

			if dvaluex == 0 && IsEqual(nodenotpack.Data.x, nodepacked.Data.x) {

				move = nodenotpack.Data.y - MaxNumber(nodepacked.Data.y, nodepacked1.Data.y)
				k = 999

			} else {

				k = dvaluey / dvaluex
				b := nodepacked.Data.y - k*nodepacked.Data.x
				move = nodenotpack.Data.y - (k*nodenotpack.Data.x + b)

			}
			if IsEqual(move, 0) {
				move = 0
			}

			// if move < 0 {
			// 	nodepacked = nodepacked1
			// 	nodepacked1 = nodepacked1.Next
			// 	continue
			// }

			if move < minmove {
				minmove = move
				kofslope = k
				mincurrentheight = nodenotpack.Data.y - move
				currentx = nodenotpack.Data.x
			}

			nodepacked = nodepacked1
			nodepacked1 = nodepacked1.Next

		}

		nodenotpack = nodenotpack.Next
	}

	nodepacked = packedrect.Next

	for i := 0; i < packedrect.pointnum; i++ {

		nodenotpack = notpackrect.Next
		nodenotpack1 = nodenotpack.Next

		for j := 0; j < notpackrect.pointnum; j++ {
			if j == notpackrect.pointnum-1 {
				nodenotpack1 = notpackrect.Next
			}
			if (nodepacked.Data.x < MinNumber(nodenotpack.Data.x, nodenotpack1.Data.x)) || (nodepacked.Data.x > MaxNumber(nodenotpack.Data.x, nodenotpack1.Data.x)) {
				nodenotpack = nodenotpack1
				nodenotpack1 = nodenotpack1.Next
				continue
			}

			dvaluex = nodenotpack.Data.x - nodenotpack1.Data.x
			dvaluey = nodenotpack.Data.y - nodenotpack1.Data.y

			var move float32
			var k float32

			if dvaluex == 0 && IsEqual(nodepacked.Data.x, nodenotpack.Data.x) {
				move = MinNumber(nodenotpack.Data.y, nodenotpack1.Data.y) - nodepacked.Data.y
				k = 999
			} else {
				k = dvaluey / dvaluex
				b := nodenotpack.Data.y - k*nodenotpack.Data.x
				move = k*nodepacked.Data.x + b - nodepacked.Data.y
			}

			if IsEqual(move, 0) {
				move = 0
			}
			// if move < 0 {
			// 	nodenotpack = nodenotpack1
			// 	nodenotpack1 = nodenotpack1.Next
			// 	continue
			// }

			if move < minmove {
				minmove = move
				kofslope = k
				mincurrentheight = nodepacked.Data.y
				currentx = nodepacked.Data.x
			}
			nodenotpack = nodenotpack1
			nodenotpack1 = nodenotpack1.Next
		}
		nodepacked = nodepacked.Next
	}

	nodenotpack = notpackrect.Next

	return minmove, mincurrentheight, currentx, kofslope
}

func slopevalue(head []Header, notpack int, packed int, kofslope float32, Height float32) (float32, float32, float32, float32) {
	//计算notpack相对于一个packed多边形的滑动距离

	var newkofslope float32 //考虑可能移动两次的情况 重新计算出对于新的项产生的 斜率
	var tempk float32

	nodenotpack := head[notpack].Next
	nodenotpack1 := nodenotpack.Next

	nodepacked := head[packed].Next
	nodepacked1 := nodepacked.Next

	var (
		minmove = (Height * Height)
		dvaluex float32
		dvaluey float32

		tempx float32
		tempy float32

		movex float32
		movey float32

		move float32

		k float32

		b  float32
		b1 float32
	)

	for i := 0; i < head[notpack].pointnum; i++ {
		nodepacked = head[packed].Next
		nodepacked1 = nodepacked.Next
		for j := 0; j < head[packed].pointnum; j++ {
			if j == head[packed].pointnum-1 {
				nodepacked1 = head[packed].Next
			}

			dvaluex = nodepacked.Data.x - nodepacked1.Data.x
			dvaluey = nodepacked.Data.y - nodepacked1.Data.y
			b = nodenotpack.Data.y - kofslope*nodenotpack.Data.x
			//notpack点在kofslope方向下的b  用在后面计算接触点用

			if IsEqual(dvaluey, 0) {
				//packed直线是一条横线

				tempx = (nodepacked.Data.y - b) / kofslope

				if (tempx > MaxNumber(nodepacked.Data.x, nodepacked1.Data.x)) || (tempx < MinNumber(nodepacked.Data.x, nodepacked1.Data.x)) {
					//交点不在packed 和 packed1 直线上
					nodepacked = nodepacked.Next
					nodepacked1 = nodepacked1.Next
					continue
				}
				tempx = nodenotpack.Data.x - tempx
				if IsEqual(tempx, 0) {
					tempx = 0
				}

				tempy = nodenotpack.Data.y - nodepacked.Data.y
				if IsEqual(tempy, 0) {
					tempy = 0
				}

				move = tempx*tempx + tempy*tempy

				tempk = 0

			} else if IsEqual(dvaluex, 0) {

				tempy = kofslope*nodepacked.Data.x + b
				if (tempy > MaxNumber(nodepacked.Data.y, nodepacked1.Data.y)) || (tempy < MinNumber(nodepacked.Data.y, nodepacked1.Data.y)) {
					nodepacked = nodepacked.Next
					nodepacked1 = nodepacked1.Next
					continue
				}
				tempx = nodenotpack.Data.x - nodepacked.Data.x
				if IsEqual(tempx, 0) {
					tempx = 0
				}

				tempy = nodenotpack.Data.y - tempy
				if IsEqual(tempy, 0) {
					tempy = 0
				}

				move = tempx*tempx + tempy*tempy

				tempk = 999

			} else {
				k = dvaluey / dvaluex
				b1 = nodepacked.Data.y - k*nodepacked.Data.x
				if IsEqual(k, kofslope) {
					//平行 无穷
					nodepacked = nodepacked.Next
					nodepacked1 = nodepacked1.Next
					continue
				} else {
					tempx = (b - b1) / (k - kofslope)
					tempy = kofslope*tempx + b
					if (tempx < MinNumber(nodepacked.Data.x, nodepacked1.Data.x)) || (tempx > MaxNumber(nodepacked.Data.x, nodepacked1.Data.x)) {
						//交点不在线段上 nodepacked nodepacked1 上跳过
						nodepacked = nodepacked.Next
						nodepacked1 = nodepacked1.Next
						continue
					}
					tempx = nodenotpack.Data.x - tempx
					if IsEqual(tempx, 0) {
						tempx = 0
					}

					tempy = nodenotpack.Data.y - tempy
					if IsEqual(tempy, 0) {
						tempy = 0
					}

					tempk = k

				}
				move = tempx*tempx + tempy*tempy
			}

			if tempx < 0 {
				nodepacked = nodepacked.Next
				nodepacked1 = nodepacked1.Next
				continue
			}

			if move < minmove {
				minmove = move
				movex = tempx
				movey = tempy
				newkofslope = tempk
			}

			nodepacked = nodepacked.Next
			nodepacked1 = nodepacked1.Next
		}

		nodenotpack = nodenotpack.Next
	}

	nodepacked = head[packed].Next

	for i := 0; i < head[packed].pointnum; i++ {
		nodenotpack = head[notpack].Next
		nodenotpack1 = nodenotpack.Next

		for j := 0; j < head[notpack].pointnum; j++ {
			if j == head[notpack].pointnum-1 {
				nodenotpack1 = head[notpack].Next
			}

			dvaluex = nodenotpack.Data.x - nodenotpack1.Data.x
			dvaluey = nodenotpack.Data.y - nodenotpack1.Data.y
			b = nodepacked.Data.y - kofslope*nodepacked.Data.x

			if dvaluey == 0 {

				tempx = (nodenotpack.Data.y - b) / kofslope
				if (tempx > MaxNumber(nodenotpack.Data.x, nodenotpack1.Data.x)) || (tempx < MinNumber(nodenotpack.Data.x, nodenotpack1.Data.x)) {
					//交点不在notpack直线范围内
					nodenotpack = nodenotpack1
					nodenotpack1 = nodenotpack1.Next
					continue
				}

				tempy = nodenotpack.Data.y - nodepacked.Data.y
				if IsEqual(tempy, 0) {
					tempy = 0
				}

				tempx = tempx - nodepacked.Data.x
				if IsEqual(tempx, 0) {
					tempx = 0
				}

				move = tempx*tempx + tempy*tempy

				tempk = 0

			} else if dvaluex == 0 {
				tempy = kofslope*nodenotpack.Data.x + b
				if (tempy > MaxNumber(nodenotpack.Data.y, nodenotpack1.Data.y)) || (tempy < MinNumber(nodenotpack.Data.y, nodenotpack1.Data.y)) {
					nodenotpack = nodenotpack1
					nodenotpack1 = nodenotpack1.Next
					continue
				}
				tempx = nodenotpack.Data.x - nodepacked.Data.x
				if IsEqual(tempx, 0) {
					tempx = 0
				}

				tempy = tempy - nodepacked.Data.y
				if IsEqual(tempy, 0) {
					tempy = 0
				}

				move = tempx*tempx + tempy*tempy
				tempk = 999
			} else {
				k = dvaluey / dvaluex
				if IsEqual(k, kofslope) {
					nodenotpack = nodenotpack1
					nodenotpack1 = nodenotpack1.Next
					continue
				} else {
					b1 = nodenotpack.Data.y - k*nodenotpack.Data.x
					tempx = (b - b1) / (k - kofslope)
					tempy = k*tempx + b1
					if (tempx < MinNumber(nodenotpack.Data.x, nodenotpack1.Data.x)) || (tempx > MaxNumber(nodenotpack.Data.x, nodenotpack1.Data.x)) {
						nodenotpack = nodenotpack1
						nodenotpack1 = nodenotpack1.Next
						continue
					}

					tempy = tempy - nodepacked.Data.y
					if IsEqual(tempy, 0) {
						tempy = 0
					}

					tempx = tempx - nodepacked.Data.x
					if IsEqual(tempx, 0) {
						tempx = 0
					}

					tempk = k

				}
				move = tempx*tempx + tempy*tempy
			}

			if tempx < 0 {
				nodenotpack = nodenotpack1
				nodenotpack1 = nodenotpack1.Next
				continue
			}

			if move < minmove {
				minmove = move
				movex = tempx
				movey = tempy
				newkofslope = tempk
				// fmt.Println("2", move, movex, movey)
				// fmt.Println("node value")
				// fmt.Println(nodepacked.Data.x, nodepacked.Data.y)

				// fmt.Println(nodenotpack.Data.x, nodenotpack.Data.y, nodenotpack1.Data.x, nodenotpack1.Data.y)
				// fmt.Println("node value  over !!!!!!!!!!!")
			}

			nodenotpack = nodenotpack1
			nodenotpack1 = nodenotpack1.Next
		}
		nodepacked = nodepacked.Next
	}

	return minmove, movex, movey, newkofslope
}

func MovetoLeft(head []Header, notpack int, packed int) (float32, float32) {
	var (
		minmove  = head[notpack].position.right.x //最小向左移动距离处值用当前多边形最右点x值代替
		dvaluex  float32
		dvaluey  float32
		move     float32 //记录中间计算的值 并于最小值比较
		k        float32
		b        float32
		kofslope float32
	)

	nodenotpack := head[notpack].Next
	nodenotpack1 := nodenotpack.Next

	nodepacked := head[packed].Next
	nodepacked1 := nodepacked.Next

	//notpack上每点到packed多边形上每条线段
	for i := 0; i < head[notpack].pointnum; i++ {

		nodepacked = head[packed].Next
		nodepacked1 = nodepacked.Next

		for j := 0; j < head[packed].pointnum; j++ {

			if j == head[packed].pointnum-1 {
				nodepacked1 = head[packed].Next
			}

			if (nodenotpack.Data.y < MinNumber(nodepacked.Data.y, nodepacked1.Data.y)) || (nodenotpack.Data.y > MaxNumber(nodepacked.Data.y, nodepacked1.Data.y)) {
				nodepacked = nodepacked1
				nodepacked1 = nodepacked1.Next
				continue
			}

			dvaluex = nodepacked.Data.x - nodepacked1.Data.x
			dvaluey = nodepacked.Data.y - nodepacked1.Data.y

			if IsEqual(dvaluex, 0) {
				//无法计算斜率
				move = nodenotpack.Data.x - nodepacked.Data.x
				if IsEqual(move, 0) {
					move = 0
				}
				k = 999

			} else if IsEqual(dvaluey, 0) {
				move = nodenotpack.Data.x - MinNumber(nodepacked.Data.x, nodepacked1.Data.x)
				// move = nodenotpack.Data.x - MaxNumber(nodepacked.Data.x, nodepacked1.Data.x)
				if IsEqual(move, 0) {
					move = 0
				}
				k = 0
			} else {
				k = dvaluey / dvaluex
				b = nodepacked.Data.y - k*nodepacked.Data.x

				move = nodenotpack.Data.x - (nodenotpack.Data.y-b)/k
				if IsEqual(move, 0) {
					move = 0
				}
			}

			if move < 0 {
				nodepacked = nodepacked1
				nodepacked1 = nodepacked1.Next
				continue
			}

			if move < minmove {
				minmove = move
				kofslope = k
			}
			//下一条线段
			nodepacked = nodepacked1
			nodepacked1 = nodepacked1.Next
		}

		//去下一个notpack多边形上的点
		nodenotpack = nodenotpack.Next

	}

	//计算packed多边形上的点到notpack上的距离
	nodepacked = head[packed].Next
	nodepacked1 = nodepacked.Next

	for i := 0; i < head[packed].pointnum; i++ {

		nodenotpack = head[notpack].Next
		nodenotpack1 = nodenotpack.Next

		for j := 0; j < head[notpack].pointnum; j++ {

			if j == head[notpack].pointnum-1 {
				nodenotpack1 = head[notpack].Next
			}
			if (nodepacked.Data.y < MinNumber(nodenotpack.Data.y, nodenotpack1.Data.y)) || (nodepacked.Data.y > MaxNumber(nodenotpack.Data.y, nodenotpack1.Data.y)) {
				nodenotpack = nodenotpack1
				nodenotpack1 = nodenotpack1.Next
				continue
			}

			dvaluex = nodenotpack.Data.x - nodenotpack1.Data.x
			dvaluey = nodenotpack.Data.y - nodenotpack1.Data.y

			dvalueofcheck := nodepacked.Data.x - nodepacked1.Data.x

			if IsEqual(dvaluex, 0) {

				move = nodenotpack.Data.x - nodepacked.Data.x
				if IsEqual(move, 0) {
					move = 0
				}
				k = 999
			} else if dvaluey == 0 && IsEqual(dvalueofcheck, 0) {
				move = MaxNumber(nodenotpack.Data.x, nodenotpack1.Data.x) - nodepacked.Data.x
				// move = MaxNumber(nodenotpack.Data.x, nodenotpack1.Data.x) - nodepacked.Data.x
				if IsEqual(move, 0) {
					move = 0
				}
				k = 0
			} else {
				k = dvaluey / dvaluex
				b = nodenotpack.Data.y - k*nodenotpack.Data.x
				move = (nodepacked.Data.y-b)/k - nodepacked.Data.x
				if IsEqual(move, 0) {
					move = 0
				}
			}

			if move < 0 {
				nodenotpack = nodenotpack1
				nodenotpack1 = nodenotpack1.Next
				continue
			}

			if move < minmove {
				//更新最小移动距离
				minmove = move
				kofslope = k
			}
			//下一条线段
			nodenotpack = nodenotpack1
			nodenotpack1 = nodenotpack1.Next
		}
		nodepacked = nodepacked.Next

	}
	//fmt.Println("Move to left,", minmove, " notpack ", notpack, " packed", packed)
	return minmove, kofslope
}

func FindMinmove(head []Header, movetype int, notpack int, rectnum int, kofslope float32, Height float32) (float32, float32, int, float32, float32) {

	//movetype==1 求向下的最短移动距离 Belowmove
	//movetype==2 求向上的最短移动距离 Uppermove
	//movetype==3 求向左的最短移动距离 MovetoLeft
	//movetype==4 求斜向的移动距离

	//missNO 不需要进行判断的多边形

	var (
		minmove      float32 //最小移动距离
		move         float32 //暂存minmove
		kofslope_cal float32 //滑轨的斜率
		tempofslope  float32 //滑轨暂存值

		//kforslope float32 = kofslope //在计算对missNo时可能需要用一次

		minNo int //最小移动距离项目编号
	)

	if movetype == 1 {
		for i := 0; i < rectnum; i++ {
			if head[i].flag == 0 || i == notpack {
				continue
			}

			var right = head[notpack].position.right.x
			var left = head[notpack].position.left.x
			if (right < head[i].position.left.x) || (left > head[i].position.right.x) {
				//若当前已排项与待排项及待排区域不在一个空间
				continue
			}

			move, _, _, tempofslope = Belowmove(head, notpack, i, Height)
			if move < 0 {
				continue
			}
			if move > minmove {
				minmove = move
				minNo = i
				kofslope_cal = tempofslope
			}

		}
		return minmove, kofslope_cal, minNo, 0, 0 //返回相应条件的最短移动距离  得到最小项的编号
	} else if movetype == 2 {
		var flag = 0
		minmove = Height
		for j := 0; j < rectnum; j++ {
			if head[j].flag == 0 || notpack == j {
				continue
			}

			var right = head[notpack].position.right.x
			var left = head[notpack].position.left.x
			if (right < head[j].position.left.x) || (left > head[j].position.right.x) {
				//若当前已排项与待排项及待排区域不在一个空间
				continue
			}

			move, _, _, tempofslope = Uppermove(head, notpack, j, Height)

			if move < minmove {
				flag = 1
				minmove = move
				minNo = j
				kofslope_cal = tempofslope
			}
		}
		//fmt.Println("向上", minmove, minNo)
		if flag == 0 {
			return 0, kofslope_cal, minNo, 0, 0 //返回相应条件的最短移动距离  得到最小项的编号
		}
		return minmove, kofslope_cal, minNo, 0, 0 //返回相应条件的最短移动距离  得到最小项的编号
	} else if movetype == 3 { //向左
		minmove = Height
		var flag = 0
		for j := 0; j < rectnum; j++ {
			if head[j].flag == 0 || notpack == j || head[notpack].position.bottom.y > head[j].position.top.y || head[notpack].position.top.y < head[j].position.bottom.y {
				//对比的多边形未排 或者和待排序号相同 或者不在同一横向空间  跳过
				continue
			}

			move, tempofslope = MovetoLeft(head, notpack, j)
			if move < 0 {
				continue
			}
			if head[notpack].position.left.x-move < 0 {
				continue
			}

			if move < minmove {

				flag = 1
				minmove = move
				minNo = j
				kofslope_cal = tempofslope
			}
		}
		if flag == 1 {
			return minmove, kofslope_cal, minNo, 0, 0 //返回相应条件的最短移动距离  得到最小项的编号
		}

		//无法左移
		return -1, kofslope_cal, minNo, 0, 0

	} else if movetype == 4 { //滑动
		//movetype == 4
		var (
			movex float32
			movey float32
			tempx float32
			tempy float32
			maxb  float32 = -99999
			minb  float32 = 99999
			tempb float32
			flag  float32
		)
		minmove = Height * Height
		//fmt.Println("不判断的编号", missNo)

		//计算出notpack在斜率下的两个b 可比较项应该在这个区间内
		p := head[notpack].Next
		for p != nil {
			tempb = p.Data.y - (kofslope * p.Data.x)

			if tempb > maxb {
				maxb = tempb
			}
			if tempb < minb {
				minb = tempb
			}
			p = p.Next
		}

		for i := 0; i < rectnum; i++ {
			if notpack == i || head[i].flag == 0 {
				continue
			}
			var flag2 = 0
			p := head[i].Next
			for p != nil {
				tempb = p.Data.y - (kofslope * p.Data.x)
				if tempb >= minb && tempb <= maxb {

					flag++
					flag2 = 1
					break
				}
				p = p.Next
			}

			if flag2 == 0 {
				continue
			}
			// fmt.Println("区间内计算", i)
			move, tempx, tempy, tempofslope = slopevalue(head, notpack, i, kofslope, Height)

			if move < minmove {
				minNo = i
				minmove = move
				movex = tempx
				movey = tempy
				kofslope_cal = tempofslope //新的斜率
			}
		}

		if flag == 1 {
			// fmt.Println("特殊处理")
			//若flag==1 则代表只有一个 那么 对两底边进行判断  若大于当前可移动距离 则替换
			p = head[notpack].Next
			minmove = Height * Height
			if kofslope > 0 {
				for p != nil {
					tempx = p.Data.y / kofslope
					tempy = p.Data.y
					move = tempx*tempx + tempy*tempy
					if move < minmove {
						minmove = move
						movex = tempx
						movey = tempy
					}
					p = p.Next
				}
			} else {
				for p != nil {
					tempx = (Height - p.Data.y) / kofslope
					tempy = p.Data.y - Height
					// fmt.Println("特殊处理", tempx, tempy)
					move = tempx*tempx + tempy*tempy
					if move < minmove {
						minmove = move
						movex = tempx
						movey = tempy
					}
					p = p.Next
				}
			}
		}
		return -1, kofslope_cal, minNo, movex, movey
	} else if movetype == 5 {
		//movetype ==5 直接从当前位置向上计算移动距离
		minmove = Height
		var left = head[notpack].position.left.x
		var right = head[notpack].position.right.x
		for i := 0; i < rectnum; i++ {
			if head[i].flag == 0 || i == notpack {
				continue
			}

			if head[i].position.right.x < (left+0.2) || right < head[i].position.left.x {
				continue
			}
			move, _, _, tempofslope = Belowmove2(head, notpack, i, Height)
			if move < minmove {
				minmove = move
				minNo = i
				kofslope_cal = tempofslope
			}
		}
		return minmove, kofslope_cal, minNo, 0, 0
	} else {
		//movetype ==6 直接从当前位置向下移动
		minmove = Height
		var left = head[notpack].position.left.x
		var right = head[notpack].position.right.x
		for i := 0; i < rectnum; i++ {
			if head[i].flag == 0 || i == notpack {
				continue
			}
			if head[i].position.right.x < (left+0.2) || right < head[i].position.left.x {
				continue
			}
			move, tempofslope = Uppermove2(head, notpack, i, Height)
			if move < minmove {
				minmove = move
				minNo = i
				kofslope_cal = tempofslope
			}
		}
		return minmove, kofslope_cal, minNo, 0, 0

	}
}

func Rorate(head []Header, notpack int, times int, xaxis float32, Height float32, rpointx float32, rpointy float32) {
	//该函数是多边形旋转  旋转点选择

	//旋转公式
	//x' = (x -x0)cos(b) -(y-y0)sin(b)+x0
	//y' = (x -x0)sin(b) +(y-y0)cos(b)+y0
	//P0（x0,y0)是相对旋转的点  P(x,y)是待旋转的点 b是逆时针旋转角度
	//旋转公式
	//rpointx rpointy即 P0由外部提供 从而将他转回去

	var (
		tempx float32
		tempy float32
	)

	p := head[notpack].Next

	for p != nil {
		tempx = p.Data.x
		tempy = p.Data.y
		p.Data.x = -(tempy - rpointy) + rpointx
		p.Data.y = (tempx - rpointx) + rpointy
		p = p.Next
	}

	RefreshPosition_rorate(head, notpack) //旋转之后重新确定position信息

	var (
		movex float32
		movey float32
	)

	// xaxis 是当前多边形左边的轴  旋转完之后可能相对与轴 或者 y=0/y=Height 产生变化  将他们移到轴上
	//times偶数是向下排 考虑向左移动到xaix和上方的y=0的轴

	if head[notpack].position.left.x != xaxis {
		movex = head[notpack].position.left.x - xaxis
	}

	if head[notpack].position.bottom.y != 0 {
		movey = head[notpack].position.bottom.y
	}

	RefreshPosition(head, notpack, movex, movey) //更新具体位置

}

func RefreshPosition(head []Header, rectno int, movex float32, movey float32) {

	//即movex movey  由旋转函数 或 排料的时候 提供

	p := head[rectno].Next

	//竖 和 横
	head[rectno].position.bottom.x -= movex
	head[rectno].position.top.x -= movex
	head[rectno].position.left.x -= movex
	head[rectno].position.right.x -= movex

	head[rectno].position.bottom.y -= movey
	head[rectno].position.top.y -= movey
	head[rectno].position.left.y -= movey
	head[rectno].position.right.y -= movey

	for p != nil {
		p.Data.x -= movex
		p.Data.y -= movey
		p = p.Next
	}

}

func RefreshPosition_rorate(head []Header, no int) {
	//考虑旋转之后位置信息的变化  更新多边形的position
	var (
		top  float32 = -1000
		topx float32

		bottom  float32 = 10000
		bottomx float32

		left  float32 = 10000
		lefty float32

		right  float32 = -1000
		righty float32
	)

	p := head[no].Next

	for p != nil {
		if p.Data.x < left {
			left = p.Data.x
			lefty = p.Data.y
		}
		if p.Data.x > right {
			right = p.Data.x
			righty = p.Data.y
		}

		if p.Data.y < bottom {
			bottom = p.Data.y
			bottomx = p.Data.x
		}
		if p.Data.y > top {
			top = p.Data.y
			topx = p.Data.x
		}

		p = p.Next
	}

	head[no].height = top - bottom
	head[no].width = right - left

	head[no].position.bottom.x = bottomx
	head[no].position.bottom.y = bottom

	head[no].position.left.x = left
	head[no].position.left.y = lefty

	head[no].position.right.x = right
	head[no].position.right.y = righty

	head[no].position.top.x = topx
	head[no].position.top.y = top
}

//*********************函数说明****************************//

//Belowmove 返回值 minmove最小移动距离  mincurrentheight 最低点高度   currentx
//Uppermove
//Packing 每次排一列  偶数向下  基数向上

//关于InsertPacking
/*InsertPacking函数
在if判断中加入了dvalue   即对较小的可插入项来说   有先插入的地方是空间较小的地方

所以会产生一个问题  对一个项来说 不同的转动方向后 可放置的区域是不同的 所以产生了第一列横着更好但是是竖着的情况
未处理  应该要对InsetPacking里加旋转  待处理！！！
*/
//*********************函数说明****************************//
func InsertPacking(head []Header, notpack int, Height float32, rectnum int, x float32) (float32, float32) {

	//未排多边形与待排空间中每一个求向下移动距离
	//先判断有没有在空间外
	//若没有 则从当前位置向上 看看是否可移动  若出现负值 则不可放入
	//若可移动距离 >=0 则可放入 并向上移动这个距离
	//待排空间只有一个和卒后一个的情况时进行特殊处理

	var (
		minmove float32

		backy     float32
		packednum []int
		kofslope  float32
		moveup    float32
		movedown  float32
		tempmove  float32
		tempk     float32
		tempy     float32

		flag int
	)

	backy = head[notpack].position.left.y

	for i := 0; i < rectnum; i++ {
		if head[i].flag == 0 {
			continue
		}
		if head[i].position.right.x < head[notpack].position.left.x || head[i].position.left.x > head[notpack].position.right.x {
			continue
		}

		packednum = append(packednum, i)
	}

	if len(packednum) == 1 {
		minmove, _, _, kofslope = Uppermove(head, notpack, packednum[0], Height)
		minmove -= head[notpack].position.top.y
		if minmove >= 0 {
			return minmove, kofslope
		}

		minmove, _, _, kofslope = Belowmove(head, notpack, packednum[0], Height)

		if (head[notpack].position.top.y + minmove) <= Height {
			return minmove, kofslope
		}
	} else {

		for i := 0; i < len(packednum); i++ {
			for j := i + 1; j < len(packednum); j++ {
				if head[packednum[i]].position.bottom.y > head[packednum[j]].position.bottom.y {
					packednum[i], packednum[j] = packednum[j], packednum[i]
				}
			}
		}

		// fmt.Println(packednum)

		for i := 0; i < len(packednum); i++ {

			moveup, _, _, tempk = Uppermove(head, notpack, packednum[i], Height)
			tempmove = moveup - head[notpack].position.top.y

			if head[notpack].position.bottom.y+tempmove < 0 {
				continue
			}

			RefreshPosition(head, notpack, 0, -tempmove)

			//计算向上
			flag = 0
			for j := 0; j < len(packednum); j++ {
				if j == i {
					continue
				}
				if head[notpack].position.top.y < head[packednum[j]].position.bottom.y {
					continue
				}
				movedown, _, _, _ = Belowmove2(head, notpack, packednum[j], Height)
				// fmt.Println("对应", movedown, packednum[j])
				if movedown < 0 {
					flag = 1
					break
				}

			}

			tempy = head[notpack].position.left.y - backy
			RefreshPosition(head, notpack, 0, tempy)

			if flag == 0 {

				return tempmove, tempk
			}

		}

		minmove = 0
		// fmt.Println(packednum)
		for i := 0; i < len(packednum); i++ {
			movedown, _, _, tempk = Belowmove(head, notpack, packednum[i], Height)
			// fmt.Println(movedown, packednum[i])
			if movedown > minmove {
				minmove = movedown
				kofslope = tempk
			}

		}
		if head[notpack].position.top.y+minmove <= Height {

			return minmove, kofslope
		}

	}
	return -999, 0
}

func BLPacking(head []Header, xvalue []float32, Height float32, rectnum int, packednum int) float32 {
	var (
		notpack  int
		step     float32
		minmove  float32
		kofslope float32
		movex    float32
		movey    float32

		roratenum  int
		bestrorate int

		bestright float32 = 0

		backx float32
		backy float32

		halfx float32 = 9999

		tempx float32
		tempy float32
	)

	//找到最右值作为初始最右值
	for i := 0; i < rectnum; i++ {
		if head[i].flag == 0 {
			continue
		}
		if head[i].position.right.x > bestright {
			bestright = head[i].position.right.x
		}
	}

	for notpack = 0; notpack < rectnum; notpack++ {
		if head[notpack].flag == 1 {

			continue
		}

		// if packednum == 18 {
		// 	break
		// }

		//进入旋转找最优位置的过程
		for roratenum = 0; roratenum < 4; roratenum++ {
			//定义返回初始点 方便后续的旋转
			backx = head[notpack].position.left.x
			backy = head[notpack].position.left.y

			for i := 0; i < len(xvalue); i++ {

				//特殊处理 通过差值处理接触点
				if i > 2 {
					RefreshPosition(head, notpack, 0.1, 0)
				}
				step = head[notpack].position.left.x - xvalue[i]
				if i > 1 {
					step -= 0.1
				}
				RefreshPosition(head, notpack, step, 0)

				minmove, kofslope = InsertPacking(head, notpack, Height, rectnum, xvalue[i])
				if minmove == -999 {
					continue
				}
				RefreshPosition(head, notpack, 0, -minmove)

				// fmt.Println("斜向移动的斜率", kofslope)
				//插入完毕 进入滑动过程
				if IsEqual(kofslope, 0) {
					if step != 0 {
						minmove, _, _, _, _ = FindMinmove(head, 3, notpack, rectnum, 0, Height)
						RefreshPosition(head, notpack, minmove, 0)
					}
				} else if IsEqual(kofslope, 999) {
					// fmt.Println("!!!!!!!!!kofslope===99999999")
				} else {
					if step != 0 {
						_, _, _, tempx, tempy = FindMinmove(head, 4, notpack, rectnum, kofslope, Height)
						RefreshPosition(head, notpack, tempx, tempy)
						// fmt.Println("斜向移动的距离", tempx, tempy)
						minmove, _, _, _, _ = FindMinmove(head, 3, notpack, rectnum, 0, Height)
						if minmove > 0 {
							RefreshPosition(head, notpack, minmove, 0)
						}
					}
					// fmt.Println("???????", tempx, tempy)
				}

				//能走到这 说明当前位置可插入 进入条件判断 然后break 进入旋转
				if roratenum == 0 {
					halfx = head[notpack].position.right.x
					bestrorate = roratenum
					halfx = head[notpack].position.right.x
					movex = head[notpack].position.left.x - backx
					movey = head[notpack].position.left.y - backy

				} else if head[notpack].position.right.x < halfx {
					bestrorate = roratenum
					halfx = head[notpack].position.right.x
					movex = head[notpack].position.left.x - backx
					movey = head[notpack].position.left.y - backy

				}
				break

			}

			tempx = head[notpack].position.left.x - backx
			tempy = head[notpack].position.left.y - backy
			RefreshPosition(head, notpack, tempx, tempy)

			midx := (head[notpack].position.left.x + head[notpack].position.right.x) / 2
			midy := (head[notpack].position.top.y + head[notpack].position.bottom.y) / 2
			Rorate(head, notpack, 0, 0, Height, midx, midy)
		}

		for roratenum = 0; roratenum < bestrorate; roratenum++ {
			midx := (head[notpack].position.left.x + head[notpack].position.right.x) / 2
			midy := (head[notpack].position.top.y + head[notpack].position.bottom.y) / 2
			Rorate(head, notpack, 0, 0, Height, midx, midy)
		}
		//fmt.Println("最终结果", movex, movey)
		//fmt.Println("当前已排数目", packednum, "!!!!!!!!!!!!!!!1")
		RefreshPosition(head, notpack, -movex, -movey)
		head[notpack].flag = 1

		packednum++
		xvalue = inputxvalue(head, notpack, xvalue)
		sortxvalue(xvalue)
		if head[notpack].position.right.x > bestright {
			bestright = head[notpack].position.right.x
		}
		//fmt.Println()

	}

	var max float32 = 0
	for i := 0; i < rectnum; i++ {
		if head[i].position.right.x > max {
			max = head[i].position.right.x
		}
	}
	return max
}

//判断这两个点是不是一样的
func Samenode(notpack *Node, packed *Node) bool {
	if IsEqual(notpack.Data.x, packed.Data.x) && IsEqual(notpack.Data.y, packed.Data.y) {
		return true
	}
	return false
}

func bestrightchose(head []Header, Height float32, rectnum int) {
	//对所有的多边形进行四下旋转 找到最宽的情况 将之作为最优
	var (
		bestrorate int
		bestright  float32
		midx       float32
		midy       float32
		flag       int = 0 //判断是否因旋转发生最大宽度的变化
	)

	for index := 0; index < rectnum; index++ {
		bestright = head[index].position.right.x
		bestrorate = 0
		for i := 0; i < 4; i++ {

			midx = (head[index].position.left.x + head[index].position.right.x) / 2
			midy = (head[index].position.top.y + head[index].position.bottom.y) / 2
			Rorate(head, index, 0, 0, Height, midx, midy)

			if head[index].position.right.x > bestright {
				bestright = head[index].position.right.x
				bestrorate = i
				flag = 2
			}
		}
		if flag != 0 {
			continue
		}

		for i := -1; i < bestrorate; i++ {
			midx = (head[index].position.left.x + head[index].position.right.x) / 2
			midy = (head[index].position.top.y + head[index].position.bottom.y) / 2
			Rorate(head, index, 0, 0, Height, midx, midy)
		}
	}
}

func inputxvalue(head []Header, packnum int, xvalue []float32) []float32 {
	p := head[packnum].Next
	for p != nil {
		if !checkrepeat(p.Data.x, xvalue) {
			xvalue = append(xvalue, p.Data.x)
		}
		p = p.Next
	}
	return xvalue
}

//计算任意多边形面积
func Areacal(head []Header, no int) float32 {
	var area float32

	p := head[no].Next
	p1 := p.Next

	for i := 0; i < head[no].pointnum; i++ {
		if i == head[no].pointnum-1 {
			p1 = head[no].Next
		}

		area += (p.Data.x*p1.Data.y - p1.Data.x*p.Data.y)
		p = p1
		p1 = p1.Next
	}

	area /= 2

	return area
}

func checkrepeat(x float32, xvalue []float32) bool {
	for i := 0; i < len(xvalue); i++ {
		if IsEqual(x, xvalue[i]) {
			return true
		}
	}
	return false
}

func sortxvalue(xvalue []float32) {

	for i := 0; i < len(xvalue); i++ {
		var flag = 0
		for j := len(xvalue) - 1; j > i; j-- {

			if xvalue[j-1] > xvalue[j] {

				flag = 1
				xvalue[j-1], xvalue[j] = xvalue[j], xvalue[j-1]
			}
		}

		if flag == 0 {
			break
		}
	}

}

func sameTaboNode(a int, b int, list *tabonode) bool {
	p := list
	for p != nil {
		if (a == p.node1 && b == p.node2) || (a == p.node2 && b == p.node1) {
			return true
		}
		p = p.Next
	}
	return false
}

func firstresult(Height float32, num int) float32 {
	header, rectnum := read()
	//将第一个与第num个交换
	header[0], header[num] = header[num], header[0]

	var packednum = 0
	var result float32
	var xvalue []float32
	header[0].flag = 1
	p := header[0].Next

	for i := 0; i < header[0].pointnum; i++ {
		if !checkrepeat(p.Data.x, xvalue) {

			xvalue = append(xvalue, p.Data.x)
		}
		p = p.Next
	}
	sortxvalue(xvalue)

	packednum++
	result = BLPacking(header, xvalue, Height, rectnum, packednum)
	return result
}

func taboRecursive(num int, Height float32) float32 {
	var (
		result    float32
		packednum int
		minresult float32 = 999
	)

	//初始解
	result = firstresult(Height, num)

	header, rectnum := read()
	tempheader, _ := read()

	//对所有读取的项转四下  选取最宽的时候作为进去的格式
	bestrightchose(header, Height, rectnum)

	//按照宽度排序
	for i := 0; i < rectnum; i++ {
		var flag = 0
		for j := rectnum - 1; j > i; j-- {
			if header[i].width < header[j].width {
				flag = 1
				header[i], header[j] = header[j], header[i]
				tempheader[i], tempheader[j] = tempheader[j], tempheader[i]
			}
		}
		if flag == 0 {
			break
		}
	}

	//将第一个与第num个交换
	header[0], header[num] = header[num], header[0]
	tempheader[0], tempheader[num] = tempheader[num], tempheader[0]

	var xvalue []float32
	var tempxvalue []float32
	header[0].flag = 1
	tempheader[0].flag = 1

	p := header[0].Next

	for i := 0; i < header[0].pointnum; i++ {
		if !checkrepeat(p.Data.x, xvalue) {
			xvalue = append(xvalue, p.Data.x)
			tempxvalue = append(tempxvalue, p.Data.x)
		}
		p = p.Next
	}
	sortxvalue(xvalue)
	sortxvalue(tempxvalue)

	packednum++

	//第一个排完  交换  加入禁忌表

	//这里记录初始状态的值 最为最后禁忌表中筛选值
	var tabolist tabonode
	tabolist.node1 = 0
	tabolist.node2 = num
	tabolist.result = result
	tabolist.Next = nil

	var times int = 0

	var (
		flag       int
		mini       int
		minj       int
		bestresult float32 = 999
	)

	for i := 1; i < (rectnum - 1); i++ {
		for j := 1; j < rectnum; j++ {
			if i == j {
				// fmt.Println("!!")
				continue

			}
			if sameTaboNode(i, j, &tabolist) {
				// fmt.Println("!!!!")
				continue

			}
			header[i], header[j] = header[j], header[i]

			result = BLPacking(header, xvalue, Height, rectnum, packednum)

			if result < bestresult {
				mini = i
				minj = j
				bestresult = result
			}
			packednum = 1
			header = tempheader

			xvalue = tempxvalue
			header[i], header[j] = header[j], header[i]
			times++

		}
		packednum = 1
		tempheader[mini], tempheader[minj] = tempheader[minj], tempheader[mini]
		header = tempheader
		xvalue = tempxvalue
		inserttotabolist(mini, minj, bestresult, &tabolist)
		bestresult = 999
	}

	//找到第一个固定的情况下 后面一次全互换的最优解 并加入禁忌表
	p1 := &tabolist
	minresult = 999
	flag = 0
	minflag := 0
	for p1 != nil {
		if p1.result < minresult {
			minflag = flag
			minresult = p1.result
		}
		flag++
		p1 = p1.Next
	}
	fmt.Println("禁忌表最小值", minflag)
	fmt.Println("次数", times)
	return minresult
}

func inserttotabolist(i int, j int, result float32, list *tabonode) bool {
	var p tabonode

	p.node1 = i
	p.node2 = j
	p.result = result

	if list.Next == nil {
		list.Next = &p
		return true
	}

	q := list
	for q.Next != nil {
		q = q.Next
	}
	q.Next = &p
	return true
}
