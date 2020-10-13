package main

func Abs(a float32) float32 {

	if a < 0 {
		a = -a
	}
	return a
}

//判断两个浮点数是否相等
func IsEqual(a float32, b float32) bool {
	if ((a - b) < 0.01) && ((a - b) > -0.01) {
		return true
	}
	return false
}

//返回两个值的最小值
func MinNumber(a float32, b float32) float32 {
	if a > b {
		return b
	}
	return a
}

//返回两个值的最大值
func MaxNumber(a float32, b float32) float32 {
	if a < b {
		return b
	}
	return a
}
func Insert(h  , d *Node) bool {
	//h the headpoint of p,   d want to insert  p,insert position

	if h.Next == nil {
		h.Next = d
		return true
	}
	//fmt.Println(h)
	n := h //*Node 
	for n.Next != nil {
		n = n.Next
	}
	d.Next = n.Next
	n.Next = d

	return true
}

func rectset(value1 float32, value2 float32) Node { //位置摆放
	var h Node
	h.Next = nil

	var d Node
	d.Data.x = value1
	d.Data.y = value2
	Insert(&h, &d)

	return h
}
