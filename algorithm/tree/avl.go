package main

import (
	"fmt"
)

// todo 平衡树
func main() {
	arr := []int {3,2,1,4,5,6,7,16,15,14,13,12,11,10,8,9}

	tree := &avlTree{}
	for _, key := range arr {
		tree.insert(key)
	}

	/**

	依次添加: 3 2 1 4 5 6 7 16 15 14 13 12 11 10 8 9

	前序遍历: 7 4 2 1 3 6 5 13 11 9 8 10 12 15 14 16
	中序遍历: 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16
	后序遍历: 1 3 2 5 6 4 8 10 9 12 11 14 16 15 13 7
	高度: 5
	最小值: 1
	最大值: 16


	删除节点8之后，再打印该AVL树的信息。
	高度: 5
	中序遍历: 1 2 3 4 5 6 7 9 10 11 12 13 14 15 16
	 */

	fmt.Printf("\n== 前序遍历: ")
	tree.preOrder()

	fmt.Printf("\n== 中序遍历: ")
	tree.middleOrder()

	fmt.Printf("\n== 后序遍历: ")
	tree.lastOrder()
	fmt.Printf("\n")

	fmt.Printf("== 高度: %d\n", tree.height())
	fmt.Printf("== 最小值: %d\n", tree.minimum().key)
	fmt.Printf("== 最大值: %d\n", tree.maximum().key)
	fmt.Printf("== 树的详细信息: \n")
	tree.print()

	removeKey := 8
	fmt.Printf("\n== 删除根节点: %d", removeKey)
	tree.remove(removeKey)

	fmt.Printf("\n== 高度: %d", tree.height())
	fmt.Printf("\n== 中序遍历: ")
	tree.middleOrder()
	fmt.Printf("\n== 树的详细信息: \n")
	tree.print()

	// 销毁二叉树
	tree.destroy()
}

func maxFn(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type avlNode struct {
	key, height int // 关键字, 节点高度

	left, right *avlNode
}

func heightAVLnode(node *avlNode) int {
	if nil != node {
		return node.height
	}
	return -1 // 空节点的 高度为 -1, 有一个节点的树 高度为1, 有一个或者两个 子节点的node高度为 1
}

func preOrderAVLnode(root *avlNode) {
	if nil != root {
		fmt.Print(root.key, " ") // 先打印,  再递归
		preOrderAVLnode(root.left)
		preOrderAVLnode(root.right)
	}
}

func middleOrderAVLnode(root *avlNode) {
	if nil != root {
		middleOrderAVLnode(root.left)
		fmt.Print(root.key, " ") // 中间打印
		middleOrderAVLnode(root.right)
	}
}

func lastOrderAVLnode(root *avlNode) {
	if nil != root {
		lastOrderAVLnode(root.left)
		lastOrderAVLnode(root.right)
		fmt.Print(root.key, " ") // 最后 打印
	}
}

func searchAVLnode(root *avlNode, key int) *avlNode {
	if nil == root {
		return nil
	}

	if key < root.key {
		return searchAVLnode(root.left, key)
	} else if key > root.key {
		return searchAVLnode(root.right, key)
	} else {
		return root
	}
}

func minimumAVLnode(root *avlNode) *avlNode {
	if nil == root {
		return nil
	}

	for nil != root.left { // 只要 还有 左节点 就一直往左找
		root = root.left
	}
	return root
}

func maximumAVLnode(root *avlNode) *avlNode {
	if nil == root {
		return nil
	}

	for nil != root.right { // 只要 还有 右节点 就一直往右找
		root = root.right
	}
	return root
}

func destroyAVLnode(root *avlNode) {

	if nil == root {
		return
	}
	// 销毁左子树
	if nil != root.left {
		destroyAVLnode(root.left)
	}
	// 销毁右子树
	if nil != root.right {
		destroyAVLnode(root.right)
	}
	root = nil
}

func printAVLnode (root *avlNode, parentKey, direction int) {
    if nil != root {
        if direction == 0 { // tree是根节点
			fmt.Printf("%d is root\n", root.key)
		} else if direction == 1{ // tree是分支节点
			fmt.Printf("%d is %d's right child\n", root.key, parentKey)
		} else {
			fmt.Printf("%d is %d's left child\n", root.key, parentKey)
		}
		printAVLnode(root.left, root.key, -1)
		printAVLnode(root.right,root.key,  1)
    }
}


// LL 则 右单旋
func leftLeftRotation(root *avlNode) *avlNode {

	var newRoot *avlNode

	newRoot = root.left
	root.left = newRoot.right
	newRoot.right = root

	// 重新计算高度
	root.height = maxFn(heightAVLnode(root.left), heightAVLnode(root.right)) + 1 // 每个节点的高度 依赖于自己的 left 和 right 的最大值高度 + 1 (之所以 +1 因为自己是人家的 父节点啊)
	newRoot.height = maxFn(heightAVLnode(newRoot.left), root.height) + 1         // 新的 root 依赖于 自己的 left 和 新的right也就是原来的root的 高度最大值 + 1 (道理 同上)  todo  这一句 在 insert 方法中 计算重复了 (不影响 因为计算结果一样), 但是 remove 中却依赖着

	return newRoot
}

// RR 则 左单旋
func rightRightRotation(root *avlNode) *avlNode {

	var newRoot *avlNode

	newRoot = root.right
	root.right = newRoot.left
	newRoot.left = root

	// 重新计算高度
	root.height = maxFn(heightAVLnode(root.left), heightAVLnode(root.right)) + 1
	newRoot.height = maxFn(root.height, heightAVLnode(newRoot.right)) + 1   	// todo  这一句 在 insert 方法中 计算重复了 (不影响 因为计算结果一样), 但是 remove 中却依赖着

	return newRoot
}

// LR 双旋
//		其实是 第一次旋转是围绕 "儿子" 进行的"RR旋转 [左单旋]"，第二次是围绕 "自己" 进行的"LL旋转 [右单旋]"
func leftRightRotation(root *avlNode) *avlNode {

	// 先将 自己的 儿子和孙子 做旋转 todo RR 旋转 [左单旋]
	root.left = rightRightRotation(root.left)

	// 再到 自己和新儿子做旋转  todo LL 旋转 [右单旋]
	return leftLeftRotation(root)
}

// RL 双旋
//		其实是 第一次旋转是围绕 "儿子" 进行的"LL旋转 [右单旋]"，第二次是围绕 "自己" 进行的"RR旋转 [左单旋]"
func rightLeftRotation(root *avlNode) *avlNode {

	// 儿子和孙子 做旋转
	root.right = leftLeftRotation(root.right)

	// 自己和儿子 做旋转
	return rightRightRotation(root)
}

//
func insertAVLnode(root, node *avlNode) *avlNode {

	if nil == root {
		root = node // 这里的 root heightAVLnode 为 0
	} else {

		// 二叉树插入, 然后旋转
		if node.key < root.key {

			root.left = insertAVLnode(root.left, node)

			// 往左插, 如果左子 节点高度大于右子 节点

			// 插入节点后，若AVL树失去平衡，则进行相应的调节
			if heightAVLnode(root.left)-heightAVLnode(root.right) == 2 {

				// 需要旋转, 到底是  LL 还是 LR 看情况
				if node.key < root.left.key {
					root = leftLeftRotation(root) // todo LL 旋转
				} else {
					root = leftRightRotation(root) // todo LR 旋转
				}
			}
		} else if node.key > root.key {
			root.right = insertAVLnode(root.right, node) // 往右边 插入
			if heightAVLnode(root.right)-heightAVLnode(root.left) == 2 { // 判断高度, 是否需要旋转
				if node.key > root.right.key {
					root = rightRightRotation(root) // RR 单旋
				} else {
					root = rightLeftRotation(root) // RL 双旋
				}
			}
		}
	}

	// todo 这一句加了  不是和  旋转里面的重复了么?
	// todo 其实可以理解为, 这一句是 计算 不需要旋转时, key 插入的node的 height 的
	// todo 而需要旋转时, 虽然和旋转里面的 计算 height 重复了, 但是不影响结果
	root.height = maxFn(heightAVLnode(root.left), heightAVLnode(root.right)) + 1
	return root
}

//
func removeAVLnode(root, node *avlNode) *avlNode {

	if nil == root || nil == node {
		return nil
	}

	if node.key < root.key {
		root.left = removeAVLnode(root.left, node)

		// 删除完后, 判断是否需要做旋转
		if heightAVLnode(root.right)-heightAVLnode(root.left) == 2 { // 当 右节点的高度 > 左节点的高度时, 需要对 右子树做旋转

			if heightAVLnode(root.right.left) > heightAVLnode(root.right.right) { // 处理 右节点 和 右节点的儿子
				root = rightLeftRotation(root) // 如果是 右节点的左儿子高度大, 则需要 EL
			} else {
				root = rightRightRotation(root) // 否则 RR
			}
		}
	} else if node.key > root.key {
		root.right = removeAVLnode(root.right, node)

		if heightAVLnode(root.left)-heightAVLnode(root.right) == 2 { // 需要处理 左子树
			if heightAVLnode(root.left.left) > heightAVLnode(root.left.right) {
				root = leftLeftRotation(root)
			} else {
				root = leftRightRotation(root)
			}
		}
	} else { // todo 找到 需要被删除的  节点了

		// todo #############################
		// todo #############################
		// todo 主要逻辑在这里
		// todo #############################
		// todo #############################

		// root 存在 左右孩子 时
		if nil != root.left && nil != root.right {

			if  heightAVLnode(root.left) > heightAVLnode(root.right) {
				// todo 非常重要  (有点 类似 Treap 中删除节点时的其中一种处理方式)
				//
				// 如果 root 的左子树比右子树高；
				// 则(01)找出 root 的左子树中的最大节点
				//   (02)将该最大节点的值赋值给 root
				//   (03)删除该最大节点
				// 这类似于用 "root 的左子树中最大节点"  做  "root" 的替身
				// 采用这种方式的好处是：删除 "root 的左子树中最大节点" 之后，AVL树仍然是平衡的
				//

				max := maximumAVLnode(root.left)
				root.key = max.key
				root.left = removeAVLnode(root.left, max)

			} else {

				// todo 同理

				// 如果 root 的左子树不比右子树高(即它们相等，或右子树比左子树高1)
				// 则(01)找出 root 的右子树中的最小节点
				//   (02)将该最小节点的值赋值给 root
				//   (03)删除该最小节点。
				// 这类似于用 "root 的右子树中最小节点" 做 "root" 的替身
				// 采用这种方式的好处是：删除 " root的右子树中最小节点" 之后，AVL树仍然是平衡的
				//

				min := maximumAVLnode(root.right)
				root.key = min.key
				root.right = removeAVLnode(root.right, min)
			}

		// TODO 但是如果 缺少一边子树, 或者不存在子树时
		// 剩余  左子树使用左子树接到 root位置, 反之用右子树, 否则为 nil
		} else {

			if nil != root.left {
				root = root.left
			} else {
				root = root.right
			}
		}
	}
	return root
}

type avlTree struct {
	root *avlNode
}

// todo 树的高度
func (t *avlTree) height() int {
	return heightAVLnode(t.root)
}

/**
todo AVL树的 插入/删除, 导致失去平衡时的情况一定是【 LL、LR、RL、RR】 这 4 种 之一，它们都由各自的定义：

(1) LL：LeftLeft，也称为"左左"。插入或删除一个节点后，根节点的左子树的左子树还有非空子节点，导致"根的左子树的高度"比"根的右子树的高度"大2，导致AVL树失去了平衡。  	todo 单旋
     例如，在上面LL情况中，由于"根节点(8)的左子树(4)的左子树(2)还有非空子节点"，而"根节点(8)的右子树(12)没有子节点"；导致"根节点(8)的左子树(4)高度"比"根节点(8)的右子树(12)"高2。



(2) LR：LeftRight，也称为"左右"。插入或删除一个节点后，根节点的左子树的右子树还有非空子节点，导致"根的左子树的高度"比"根的右子树的高度"大2，导致AVL树失去了平衡。 	todo 双旋
     例如，在上面LR情况中，由于"根节点(8)的左子树(4)的左子树(6)还有非空子节点"，而"根节点(8)的右子树(12)没有子节点"；导致"根节点(8)的左子树(4)高度"比"根节点(8)的右子树(12)"高2。



(3) RL：RightLeft，称为"右左"。插入或删除一个节点后，根节点的右子树的左子树还有非空子节点，导致"根的右子树的高度"比"根的左子树的高度"大2，导致AVL树失去了平衡。   	todo 双旋
     例如，在上面RL情况中，由于"根节点(8)的右子树(12)的左子树(10)还有非空子节点"，而"根节点(8)的左子树(4)没有子节点"；导致"根节点(8)的右子树(12)高度"比"根节点(8)的左子树(4)"高2。



(4) RR：RightRight，称为"右右"。插入或删除一个节点后，根节点的右子树的右子树还有非空子节点，导致"根的右子树的高度"比"根的左子树的高度"大2，导致AVL树失去了平衡。	todo 单旋
     例如，在上面RR情况中，由于"根节点(8)的右子树(12)的右子树(14)还有非空子节点"，而"根节点(8)的左子树(4)没有子节点"；导致"根节点(8)的右子树(12)高度"比"根节点(8)的左子树(4)"高2。


如果在AVL树中进行插入或删除节点后，可能导致AVL树失去平衡。AVL失去平衡之后，可以通过旋转使其恢复平衡，下面分别介绍"LL(左左)，LR(左右)，RR(右右)和RL(右左)"这4种情况对应的旋转方法。
 */

// todo 插入
func (t *avlTree) insert(key int) {
	node := &avlNode{
		key:    key,
		height: 0,
	}
	t.root = insertAVLnode(t.root, node)
}

// todo 删除
func (t *avlTree) remove(key int) {
	if node := t.search(key); nil != node {
		t.root = removeAVLnode(t.root, node)
	}
}

// todo 前序遍历
func (t *avlTree) preOrder() {
	preOrderAVLnode(t.root)
}

// todo 中序遍历  (层序遍历)
func (t *avlTree) middleOrder() {
	middleOrderAVLnode(t.root)
}

// todo 后序遍历
func (t *avlTree) lastOrder() {
	lastOrderAVLnode(t.root)
}

// todo 二叉树查找
func (t *avlTree) search(key int) *avlNode {
	return searchAVLnode(t.root, key)
}

// todo 最小值
func (t *avlTree) minimum() *avlNode {
	return minimumAVLnode(t.root)
}

// todo 最大值
func (t *avlTree) maximum() *avlNode {
	return maximumAVLnode(t.root)
}

// todo 销毁树
func (t *avlTree) destroy () {
	destroyAVLnode(t.root)
}

// todo 打印树
func (t *avlTree) print () {
	printAVLnode(t.root,0, 0)
}