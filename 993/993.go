package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func find(root *TreeNode, x, deep, par int) (bool, int, int) {
	if x == 1 {
		return true, 1, 0
	}
	if root.Val == x {
		return true, deep, par
	}
	var a, b int
	var flag bool
	if root.Left != nil {
		flag, a, b = find(root.Left, x, deep+1, root.Val)
	}
	// 找到后提前退出
	if flag {
		return true, a, b
	}
	if root.Right != nil {
		flag, a, b = find(root.Right, x, deep+1, root.Val)
	}
	// 找到后提前退出
	if flag {
		return true, a, b
	}
	return false, a, b
}

func isCousins(root *TreeNode, x int, y int) bool {
	_, deepX, parX := find(root, x, 1, 0)
	_, deepY, parY := find(root, y, 1, 0)
	if deepX == deepY && parX != parY {
		return true
	}

	return false
}

func main() {

}
