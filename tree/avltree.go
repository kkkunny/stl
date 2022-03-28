package tree

import . "github.com/kkkunny/stl/types"

// AVL节点
type AVLTreeNode[T Comparator[T]] struct {
	Value  T               // 值
	Times  Usize           // 值出现的次数
	Height Usize           // 该节点作为树根节点，树的高度，方便计算平衡因子
	Left   *AVLTreeNode[T] // 左子树
	Right  *AVLTreeNode[T] // 右字树
}

// 更新节点的树高度
func (node *AVLTreeNode[T]) UpdateHeight() {
	if node == nil {
		return
	}

	var leftHeight, rightHeight Usize
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}
	// 哪个子树高算哪棵的
	maxHeight := leftHeight
	if rightHeight > maxHeight {
		maxHeight = rightHeight
	}
	// 高度加上自己那一层
	node.Height = maxHeight + 1
}

// 计算平衡因子
func (node *AVLTreeNode[T]) BalanceFactor() Isize {
	var leftHeight, rightHeight Usize
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}
	return Isize(leftHeight) - Isize(rightHeight)
}

func (node *AVLTreeNode[T]) Add(value T) *AVLTreeNode[T] {
	// 添加值到根节点node，如果node为空，那么让值成为新的根节点，树的高度为1
	if node == nil {
		return &AVLTreeNode[T]{Value: value, Height: 1}
	}

	// 如果值重复，什么都不用做，直接更新次数
	if value.Compare(node.Value) == 0 {
		node.Times = node.Times + 1
		return node
	}

	// 辅助变量
	var newTreeNode *AVLTreeNode[T]

	o := value.Compare(node.Value)
	if o > 0 {
		// 插入的值大于节点值，要从右子树继续插入
		node.Right = node.Right.Add(value)
		// 平衡因子，插入右子树后，要确保树根左子树的高度不能比右子树低一层。
		factor := node.BalanceFactor()
		// 右子树的高度变高了，导致左子树-右子树的高度从-1变成了-2。
		if factor == -2 {
			if value.Compare(node.Right.Value) > 0 {
				// 表示在右子树上插上右儿子导致失衡，需要单左旋：
				newTreeNode = LeftRotation(node)
			} else {
				//表示在右子树上插上左儿子导致失衡，先右后左旋：
				newTreeNode = RightLeftRotation(node)
			}
		}
	} else {
		// 插入的值小于节点值，要从左子树继续插入
		node.Left = node.Left.Add(value)
		// 平衡因子，插入左子树后，要确保树根左子树的高度不能比右子树高一层。
		factor := node.BalanceFactor()
		// 左子树的高度变高了，导致左子树-右子树的高度从1变成了2。
		if factor == 2 {
			if value.Compare(node.Left.Value) < 0 {
				// 表示在左子树上插上左儿子导致失衡，需要单右旋：
				newTreeNode = RightRotation(node)
			} else {
				//表示在左子树上插上右儿子导致失衡，先左后右旋：
				newTreeNode = LeftRightRotation(node)
			}
		}
	}

	if newTreeNode == nil {
		// 表示什么旋转都没有，根节点没变，直接刷新树高度
		node.UpdateHeight()
		return node
	} else {
		// 旋转了，树根节点变了，需要刷新新的树根高度
		newTreeNode.UpdateHeight()
		return newTreeNode
	}
}

func (node *AVLTreeNode[T]) FindMinValue() *AVLTreeNode[T] {
	// 左子树为空，表面已经是最左的节点了，该值就是最小值
	if node.Left == nil {
		return node
	}

	// 一直左子树递归
	return node.Left.FindMinValue()
}

func (node *AVLTreeNode[T]) FindMaxValue() *AVLTreeNode[T] {
	// 右子树为空，表面已经是最右的节点了，该值就是最大值
	if node.Right == nil {
		return node
	}

	// 一直右子树递归
	return node.Right.FindMaxValue()
}

func (node *AVLTreeNode[T]) Find(value T) *AVLTreeNode[T] {
	o := value.Compare(node.Value)
	if o > 0 {
		// 如果查找的值大于节点值，从节点的右子树开始找
		if node.Right == nil {
			// 右子树为空，表示找不到该值了，返回nil
			return nil
		}
		return node.Right.Find(value)
	} else if o < 0 {
		// 如果查找的值小于节点值，从节点的左子树开始找
		if node.Left == nil {
			// 左子树为空，表示找不到该值了，返回nil
			return nil
		}
		return node.Left.Find(value)
	} else {
		// 如果该节点刚刚等于该值，那么返回该节点
		return node
	}
}

func (node *AVLTreeNode[T]) Delete(value T) *AVLTreeNode[T] {
	if node == nil {
		// 如果是空树，直接返回
		return nil
	}
	o := value.Compare(node.Value)
	if o < 0 {
		// 从左子树开始删除
		node.Left = node.Left.Delete(value)
		// 删除后要更新该子树高度
		node.Left.UpdateHeight()
	} else if o > 0 {
		// 从右子树开始删除
		node.Right = node.Right.Delete(value)
		// 删除后要更新该子树高度
		node.Right.UpdateHeight()
	} else {
		// 找到该值对应的节点
		// 该节点没有左右子树
		// 第一种情况，删除的节点没有儿子，直接删除即可。
		if node.Left == nil && node.Right == nil {
			return nil // 直接返回nil，表示直接该值删除
		}

		// 该节点有两棵子树，选择更高的哪个来替换
		// 第二种情况，删除的节点下有两个子树，选择高度更高的子树下的节点来替换被删除的节点，如果左子树更高，选择左子树中最大的节点，也就是左子树最右边的叶子节点，如果右子树更高，选择右子树中最小的节点，也就是右子树最左边的叶子节点。最后，删除这个叶子节点。
		if node.Left != nil && node.Right != nil {
			// 左子树更高，拿左子树中最大值的节点替换
			if node.Left.Height > node.Right.Height {
				maxNode := node.Left
				for maxNode.Right != nil {
					maxNode = maxNode.Right
				}

				// 最大值的节点替换被删除节点
				node.Value = maxNode.Value
				node.Times = maxNode.Times

				// 把最大的节点删掉
				node.Left = node.Left.Delete(maxNode.Value)
				// 删除后要更新该子树高度
				node.Left.UpdateHeight()
			} else {
				// 右子树更高，拿右子树中最小值的节点替换
				minNode := node.Right
				for minNode.Left != nil {
					minNode = minNode.Left
				}

				// 最小值的节点替换被删除节点
				node.Value = minNode.Value
				node.Times = minNode.Times

				// 把最小的节点删掉
				node.Right = node.Right.Delete(minNode.Value)
				// 删除后要更新该子树高度
				node.Right.UpdateHeight()
			}
		} else {
			// 只有左子树或只有右子树
			// 只有一个子树，该子树也只是一个节点，将该节点替换被删除的节点，然后置子树为空
			if node.Left != nil {
				//第三种情况，删除的节点只有左子树，因为树的特征，可以知道左子树其实就只有一个节点，它本身，否则高度差就等于2了。
				node.Value = node.Left.Value
				node.Times = node.Left.Times
				node.Height = 1
				node.Left = nil
			} else if node.Right != nil {
				//第四种情况，删除的节点只有右子树，因为树的特征，可以知道右子树其实就只有一个节点，它本身，否则高度差就等于2了。
				node.Value = node.Right.Value
				node.Times = node.Right.Times
				node.Height = 1
				node.Right = nil
			}
		}

		// 找到值后，进行替换删除后，直接返回该节点
		return node
	}

	// 左右子树递归删除节点后需要平衡
	var newNode *AVLTreeNode[T]
	// 相当删除了右子树的节点，左边比右边高了，不平衡
	if node.BalanceFactor() == 2 {
		if node.Left.BalanceFactor() >= 0 {
			newNode = RightRotation(node)
		} else {
			newNode = LeftRightRotation(node)
		}
		//  相当删除了左子树的节点，右边比左边高了，不平衡
	} else if node.BalanceFactor() == -2 {
		if node.Right.BalanceFactor() <= 0 {
			newNode = LeftRotation(node)
		} else {
			newNode = RightLeftRotation(node)
		}
	}

	if newNode == nil {
		node.UpdateHeight()
		return node
	} else {
		newNode.UpdateHeight()
		return newNode
	}
}

// AVL树
type AVLTree[T Comparator[T]] struct {
	Root *AVLTreeNode[T] // 树根节点
}

// 初始化一个AVL树
func NewAVLTree[T Comparator[T]]() *AVLTree[T] {
	return new(AVLTree[T])
}

// 添加元素
func (tree *AVLTree[T]) Add(value T) {
	// 往树根添加元素，会返回新的树根
	tree.Root = tree.Root.Add(value)
}

// 找出最小值的节点
func (tree *AVLTree[T]) FindMinValue() *AVLTreeNode[T] {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}

	return tree.Root.FindMinValue()
}

// 找出最大值的节点
func (tree *AVLTree[T]) FindMaxValue() *AVLTreeNode[T] {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}

	return tree.Root.FindMaxValue()
}

// 查找指定节点
func (tree *AVLTree[T]) Find(value T) *AVLTreeNode[T] {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}

	return tree.Root.Find(value)
}

// 删除指定的元素
func (tree *AVLTree[T]) Delete(value T) {
	if tree.Root == nil {
		// 如果是空树，直接返回
		return
	}

	tree.Root = tree.Root.Delete(value)
}

// 单右旋操作，看图说话
func RightRotation[T Comparator[T]](Root *AVLTreeNode[T]) *AVLTreeNode[T] {
	// 只有Pivot和B，Root位置变了
	Pivot := Root.Left
	B := Pivot.Right
	Pivot.Right = Root
	Root.Left = B

	// 只有Root和Pivot变化了高度
	Root.UpdateHeight()
	Pivot.UpdateHeight()
	return Pivot
}

// 单左旋操作，看图说话
func LeftRotation[T Comparator[T]](Root *AVLTreeNode[T]) *AVLTreeNode[T] {
	// 只有Pivot和B，Root位置变了
	Pivot := Root.Right
	B := Pivot.Left
	Pivot.Left = Root
	Root.Right = B

	// 只有Root和Pivot变化了高度
	Root.UpdateHeight()
	Pivot.UpdateHeight()
	return Pivot
}

// 先左后右旋操作，看图说话
func LeftRightRotation[T Comparator[T]](node *AVLTreeNode[T]) *AVLTreeNode[T] {
	node.Left = LeftRotation(node.Left)
	return RightRotation(node)
}

// 先右后左旋操作，看图说话
func RightLeftRotation[T Comparator[T]](node *AVLTreeNode[T]) *AVLTreeNode[T] {
	node.Right = RightRotation(node.Right)
	return LeftRotation(node)
}
