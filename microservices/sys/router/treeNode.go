package router

import (
	"strings"
)

var systreeNode *treeNode

type treeNode struct {
	name       string
	children   []*treeNode
	routerName string
	isEnd      bool
}

func (t *treeNode) Put(path string) {
	root := t
	///user/add
	//获取url路径分割/数组["","user","add"]
	strs := strings.Split(path, "/")

	for index, name := range strs {
		if index == 0 {
			continue
		}

		children := t.children //子层级
		isMatch := false       //是否存在
		for _, node := range children {
			//判断子层级是否存在node
			if node.name == name {
				isMatch = true //存在
				t = node       //置换t,进行下一次循环的父层级
				break
			}
		}
		if !isMatch {
			isEnd := false //是否是结尾
			//判断当前循环是否已经是url切割数组的最后一个
			if index == len(strs)-1 {
				isEnd = true
			}
			node := &treeNode{name: name, children: make([]*treeNode, 0), isEnd: isEnd}
			children = append(children, node)
			t.children = children
			t = node
		}
	}
	t = root
}

//get path: /user/get/1
func (t *treeNode) Get(path string) *treeNode {
	strs := strings.Split(path, "/")
	routerName := ""
	for index, name := range strs {
		if index == 0 {
			continue
		}
		children := t.children
		isMatch := false
		for _, node := range children {
			//匹配node 并获取拼接routerName
			if node.name == name || node.name == "*" || strings.Contains(node.name, ":") {
				isMatch = true
				routerName += "/" + node.name
				node.routerName = routerName
				t = node
				if index == len(strs)-1 {
					if node.isEnd {
						return node
					} else {
						return nil
					}

				}
				break
			}
		}
		if !isMatch {
			for _, node := range children {
				// /user/**
				// /user/get/userInfo
				// /user/aa/bb
				if node.name == "**" {
					return node
				}
			}

		}
	}
	return nil
}

//func main() {
//	systreeNode := &treeNode{
//		"/",
//		make([]*treeNode, 0),
//		"sys",
//		false,
//	}
//	systreeNode.Put("/user/add")
//	systreeNode.Put("/user/info")
//
//	systreeNode.Put("/order/add")
//	systreeNode.Put("/order/info/{id}")
//	fmt.Printf("treenode:%+v", systreeNode)
//
//}
