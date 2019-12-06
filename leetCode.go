package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 稀疏数组概念
// 当一个数组中大部分元素为0，或者为同一个值的数组时，可以使用稀疏数组来保存该数组
// 处理方法
// 1.记录数组一共有几行几列 有多少个不同的值
// 2. 把具有不通值的元素的行列及值记录在一个小规模的数组中 从而缩小程序的规模

/*
0	0	0	0	0	0	0	0	0	0	0
0	0	1	0	0	0	0	0	0	0	0
0	0	0	2	0	0	0	0	0	0	0
0	0	0	0	0	0	0	0	0	0	0
0	0	0	0	0	0	0	0	0	0	0
0	0	0	0	0	0	0	0	0	0	0
0	0	0	0	0	0	0	0	0	0	0
0	0	0	0	0	0	0	0	0	0	0
0	0	0	0	0	0	0	0	0	0	0
0	0	0	0	0	0	0	0	0	0	0
0	0	0	0	0	0	0	0	0	0	0
*/
type VolNode struct {
	row int
	col int
	val int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sparseArray() error {
	// 先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	// 看chessMap 值 不为0 则放到新的切片中
	var sparseArr []VolNode
	valNode := VolNode{
		row: 11,
		col: 11,
		val: 0,
	}

	sparseArr = append(sparseArr, valNode)
	for i, v := range chessMap {
		for j, v2 := range v {
			//fmt.Printf("%d\t",v2)
			if v2 != 0 {
				// 创建一个节点
				valNode := VolNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	var filename = "sparse.data"
	var priFile *os.File
	var err1 error

	if checkFileIsExist(filename) { //如果文件存在
		priFile, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		priFile, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	check(err1)
	defer priFile.Close()
	// 存储数据
	writer := bufio.NewWriter(priFile)
	for _, valNode := range sparseArr {
		// 先转成字符串
		str1 := strconv.Itoa(valNode.row)
		str2 := strconv.Itoa(valNode.col)
		str3 := strconv.Itoa(valNode.val)
		strNode := str1 + " " + str2 + " " + str3 + "\n"
		//fmt.Println(strNode)
		//fmt.Printf("%d: %d %d %d \n",i,valNode.row,valNode.col,valNode.val)
		n2, err := writer.WriteString(strNode)
		if err != nil {
			return err
		}
		writer.Flush()
		fmt.Printf("wrote %d bytes\n", n2)
	}
	// 取数据并还原数组
	var chessMap2 [11][11]int

	if checkFileIsExist(filename) { //如果文件存在
		priFile, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		priFile, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	check(err1)
	defer priFile.Close()
	// 读取文件   newreader 会读取到换行  数据分割后会导致其他问题    推荐用NewScanner
	inputReader := bufio.NewScanner(priFile)
	for inputReader.Scan() {
		res := strings.Split(inputReader.Text(), " ")
		s1, err1 := strconv.Atoi(res[0])
		check(err1)
		if s1 != 11 {
			s2, err1 := strconv.Atoi(res[1])
			s3, err1 := strconv.Atoi(res[2])
			check(err1)
			chessMap2[s1][s2] = s3
		}
	}
	for _, v4 := range chessMap2 {
		for _, v5 := range v4 {
			fmt.Printf("%d\t", v5)
		}
	}

	return nil

}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
