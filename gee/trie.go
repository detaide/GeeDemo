package gee

import (
	"fmt"
	"strings"
)

//isPattern : 是否开启*匹配
type Node struct {
	path string
	children []*Node
	isEnd bool
	sign string
	Params []string
}


type Stack struct{
	length  int
	nodeList []*Node
}

func (stack *Stack) Push(newNode *Node){
	stack.length += 1
	stack.nodeList = append(stack.nodeList, newNode)
}

func (stack *Stack) Pop() *Node {
	if stack.length == 0 {
		return nil
	}
	stack.length -= 1
	return stack.nodeList[stack.length]
}

func (stack *Stack) getLength() int {
	return stack.length
}

func NewStack() *Stack {
	return &Stack{
		length: 0,
		nodeList: []*Node{},
	}
}

func NewHeadNode() *Node {
	return &Node{
		path: "/",
		children: []*Node{},
		isEnd: true,
	}
}

func (node *Node) insertNode(path string) *Node {
	newNode := &Node{
		path : path,
		children: []*Node{},
		Params: []string{},
	}

	node.children = append(node.children, newNode)

	return newNode
}

//输出isEnd结点
func (node *Node) GetTreePath(path string) {
	if node == nil {
		fmt.Println(path)
		return
	}
	levelNode := node

	if levelNode.path == "/" && levelNode.isEnd {
		fmt.Println("/")
	}

	if len(levelNode.children) == 0 || levelNode.isEnd {
		newPath := path

		if levelNode.sign == ":" && levelNode.isEnd{
			for _, param := range levelNode.Params {
				newPath += "/:" + param
			}
		}

		if levelNode.sign == "*" && levelNode.isEnd {
			newPath += "/*"
		}
		if newPath != "" {
			fmt.Println(newPath)
		}
		
		
	}
	
	for _, childrenNode := range levelNode.children {
		childrenNode.GetTreePath(path + "/" + childrenNode.path)
	}
}

//后续使用反射进行
// func eraserEmpty[T int | string | nil | []interface{}](arr []interface{}, filter T) []interface{} {
// 	newArr := []interface{}{}
// 	for _, s := range arr {
// 		if filter 
// 		newArr := append()
// 	}
// }

func eraserEmpty(arr []string) []string {
	newArr := []string{}
	for _, s := range arr {
		if s != "" {
			newArr = append(newArr, s)
		}
	}
	return newArr
}

func (node *Node) MatchPattern(path string) (string, map[string]string) {

	if path == "/" {
		return path, nil
	}

	relativePath := ""
	if node == nil  {
		return "", nil
	}
	params := make(map[string]string)

	parts := strings.Split(path, "/")
	if len(parts) ==0 {
		return "", nil
	}
	//去除空白符
	// parts = parts[1:]
	parts = eraserEmpty(parts)
	levelNode := node
	endIndex := 0
	outer:
	for index, part := range parts {
		if part == "" {
			continue
		}

		for _, childNode := range levelNode.children {
			//路径解析
			if childNode.path == part {
				levelNode = childNode
				relativePath += "/" + levelNode.path
				endIndex = index + 1
				continue outer
			}
		}

		//到这儿说明全部子节点路径都不匹配，假如是终结符号，就可以匹配参数
		if !levelNode.isEnd {
			return "", nil
		}
		
	}

	if endIndex == 0 {
		return relativePath, nil
	}

	if !levelNode.isEnd  {
		return "", nil
	}

	if levelNode.sign == "*" {
		for endIndex < len(parts) {
			params["*"] += "/" + parts[endIndex]
			endIndex += 1
		}

		return relativePath, params
	}

	for index, paramName := range levelNode.Params {		
		if endIndex + index >= len(parts) {
			break
		}
		params[paramName] = parts[endIndex  + index]
		// fmt.Println(paramName)
	}

	return relativePath, params

}

//插入非空才会
func (node *Node) InsertPath(path string) bool {
	//插入路径，先对路径进行解析，/分割
	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		return false
	}
	//遍历node，查看node的每一个结点是否有前缀
	levelNode := node
	for _, part := range parts {
		if part == "" {
			continue
		}

		//特殊情况下进行匹配,*仅仅匹配一次
		if  part[0] == '*' {
			levelNode.isEnd = true
			levelNode.sign = string(part[0])
			break
		}

		if part[0] == ':' {
			levelNode.isEnd = true
			levelNode.sign = string(part[0])
			levelNode.Params = append(levelNode.Params, part[1:])
			continue
		}

		//如果子节点为空，就添加到该层的空
		if len(levelNode.children ) == 0 {
			levelNode = levelNode.insertNode(part)
			continue
		}

		

		isFind := false
		//遍历children，此时子节点不为空,找到了子节点
		for _, childNode := range levelNode.children {
			if childNode.path == part {
				levelNode = childNode
				isFind = true
				break
			}
		}

		if !isFind {
			//遍历完还没有匹配到，就增结点
			levelNode = levelNode.insertNode(part)
		}
	}
	levelNode.isEnd = true

	return true
}



