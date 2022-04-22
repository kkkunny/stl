package tree

import . "github.com/kkkunny/stl/types"

// 定义颜色
const (
	RED   = true
	BLACK = false
)

// 普通红黑树节点
type RBTreeNode[T Comparator[T]] struct {
	Value  T              // 值
	Times  int            // 值出现的次数
	Left   *RBTreeNode[T] // 左子树
	Right  *RBTreeNode[T] // 右子树
	Parent *RBTreeNode[T] // 父节点
	Color  bool           // 父亲指向该节点的链接颜色
}

func (node *RBTreeNode[T]) FindMinValue() *RBTreeNode[T] {
	// 左子树为空，表面已经是最左的节点了，该值就是最小值
	if node.Left == nil {
		return node
	}

	// 一直左子树递归
	return node.Left.FindMinValue()
}

func (node *RBTreeNode[T]) FindMaxValue() *RBTreeNode[T] {
	// 右子树为空，表面已经是最右的节点了，该值就是最大值
	if node.Right == nil {
		return node
	}

	// 一直右子树递归
	return node.Right.FindMaxValue()
}

func (node *RBTreeNode[T]) Find(value T) *RBTreeNode[T] {
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

// 普通红黑树
type RBTree[T Comparator[T]] struct {
	Root *RBTreeNode[T] // 树根节点
}

// 新建一棵普通红黑树
func NewRBTree[T Comparator[T]]() *RBTree[T] {
	return new(RBTree[T])
}

// 对某节点左旋转
func (tree *RBTree[T]) RotateLeft(h *RBTreeNode[T]) {
	if h != nil {

		// 看图理解
		x := h.Right
		h.Right = x.Left

		if x.Left != nil {
			x.Left.Parent = h
		}

		x.Parent = h.Parent
		if h.Parent == nil {
			tree.Root = x
		} else if h.Parent.Left == h {
			h.Parent.Left = x
		} else {
			h.Parent.Right = x
		}
		x.Left = h
		h.Parent = x
	}
}

// 对某节点右旋转
func (tree *RBTree[T]) RotateRight(h *RBTreeNode[T]) {
	if h != nil {

		// 看图理解
		x := h.Left
		h.Left = x.Right

		if x.Right != nil {
			x.Right.Parent = h
		}

		x.Parent = h.Parent
		if h.Parent == nil {
			tree.Root = x
		} else if h.Parent.Right == h {
			h.Parent.Right = x
		} else {
			h.Parent.Left = x
		}
		x.Right = h
		h.Parent = x
	}
}

// 普通红黑树添加元素 O(logN)
func (tree *RBTree[T]) Add(value T) {
	// 根节点为空
	if tree.Root == nil {
		// 根节点都是黑色
		tree.Root = &RBTreeNode[T]{
			Value: value,
			Color: BLACK,
		}
		return
	}

	// 辅助变量 t，表示新元素要插入到该子树，t是该子树的根节点
	t := tree.Root

	// 插入元素后，插入元素的父亲节点
	var parent *RBTreeNode[T]

	// 辅助变量，为了知道元素最后要插到左边还是右边
	var cmp int

	for {
		parent = t

		cmp = value.Compare(t.Value)
		if cmp < 0 {
			// 比当前节点小，往左子树插入
			t = t.Left
		} else if cmp > 0 {
			// 比当前节点节点大，往右子树插入
			t = t.Right
		} else {
			// 已经存在值了，更新出现的次数
			t.Times = t.Times + 1
			return
		}

		// 终于找到要插入的位置了
		if t == nil {
			break // 这时叶子节点是 parent，要插入到 parent 的下面，跳到外层去
		}
	}

	// 新节点，它要插入到 parent下面
	newNode := &RBTreeNode[T]{
		Value:  value,
		Parent: parent,
	}
	if cmp < 0 {
		// 知道要从左边插进去
		parent.Left = newNode
	} else {
		// 知道要从右边插进去
		parent.Right = newNode
	}

	// 插入新节点后，可能破坏了红黑树特征，需要修复，核心函数
	tree.fixAfterInsertion(newNode)
}

// 调整新插入的节点，自底而上
// 可以看图理解
func (tree *RBTree[T]) fixAfterInsertion(node *RBTreeNode[T]) {
	// 插入的新节点一定要是红色
	node.Color = RED

	// 节点不能是空，不能是根节点，父亲的颜色必须为红色（如果是黑色，那么直接插入不破坏平衡，不需要调整了）
	for node != nil && node != tree.Root && node.Parent.Color == RED {
		// 父亲在祖父的左边
		if ParentOf(node) == LeftOf(ParentOf(ParentOf(node))) {
			// 叔叔节点
			uncle := RightOf(ParentOf(ParentOf(node)))

			// 图例3左边部分，叔叔是红节点，祖父变色，也就是父亲和叔叔变黑，祖父变红
			if IsRed(uncle) {
				SetColor(ParentOf(node), BLACK)
				SetColor(uncle, BLACK)
				SetColor(ParentOf(ParentOf(node)), RED)
				// 还要向上递归
				node = ParentOf(ParentOf(node))
			} else {
				// 图例4左边部分，叔叔是黑节点，并且插入的节点在父亲的右边，需要对父亲左旋
				if node == RightOf(ParentOf(node)) {
					node = ParentOf(node)
					tree.RotateLeft(node)
				}

				// 变色，并对祖父进行右旋
				SetColor(ParentOf(node), BLACK)
				SetColor(ParentOf(ParentOf(node)), RED)
				tree.RotateRight(ParentOf(ParentOf(node)))
			}
		} else {
			// 父亲在祖父的右边，与父亲在祖父的左边相似
			// 叔叔节点
			uncle := LeftOf(ParentOf(ParentOf(node)))

			// 图例3右边部分，叔叔是红节点，祖父变色，也就是父亲和叔叔变黑，祖父变红
			if IsRed(uncle) {
				SetColor(ParentOf(node), BLACK)
				SetColor(uncle, BLACK)
				SetColor(ParentOf(ParentOf(node)), RED)
				// 还要向上递归
				node = ParentOf(ParentOf(node))
			} else {
				// 图例4右边部分，叔叔是黑节点，并且插入的节点在父亲的左边，需要对父亲右旋
				if node == LeftOf(ParentOf(node)) {
					node = ParentOf(node)
					tree.RotateRight(node)
				}

				// 变色，并对祖父进行左旋
				SetColor(ParentOf(node), BLACK)
				SetColor(ParentOf(ParentOf(node)), RED)
				tree.RotateLeft(ParentOf(ParentOf(node)))
			}
		}
	}

	// 根节点永远为黑
	tree.Root.Color = BLACK
}

// 普通红黑树删除元素 O(logN)
func (tree *RBTree[T]) Delete(value T) {
	// 查找元素是否存在，不存在则退出
	p := tree.Find(value)
	if p == nil {
		return
	}

	// 删除该节点
	tree.delete(p)
}

// 删除节点核心函数
// 找最小后驱节点来补位，删除内部节点转为删除叶子节点
func (tree *RBTree[T]) delete(node *RBTreeNode[T]) {
	// 如果左右子树都存在，那么从右子树的左边一直找一直找，就找能到最小后驱节点
	if node.Left != nil && node.Right != nil {
		s := node.Right
		for s.Left != nil {
			s = s.Left
		}

		// 删除的叶子节点找到了，删除内部节点转为删除叶子节点
		node.Value = s.Value
		node.Times = s.Times
		node = s // 可能存在右儿子
	}

	if node.Left == nil && node.Right == nil {
		// 没有子树，要删除的节点就是叶子节点。
	} else {
		// 只有一棵子树，因为红黑树的特征，该子树就只有一个节点
		// 找到该唯一节点
		replacement := node.Left
		if node.Left == nil {
			replacement = node.Right
		}

		// 替换开始，子树的唯一节点替代被删除的内部节点
		replacement.Parent = node.Parent
		if node.Parent == nil {
			// 要删除的节点的父亲为空，表示要删除的节点为根节点，唯一子节点成为树根
			tree.Root = replacement
		} else if node == node.Parent.Left {
			// 子树的唯一节点替代被删除的内部节点
			node.Parent.Left = replacement
		} else {
			// 子树的唯一节点替代被删除的内部节点
			node.Parent.Right = replacement
		}

		// delete this node
		node.Parent = nil
		node.Right = nil
		node.Left = nil

		//  case 1: not enter this logic
		//      R(del)
		//    B   B
		//
		//  case 2: node's color must be black, and it's son must be red
		//    B(del)     B(del)
		//  R  O       O   R
		//
		// 单子树时删除的节点绝对是黑色的，而其唯一子节点必然是红色的
		// 现在唯一子节点替换了被删除节点，该节点要变为黑色
		// now son replace it's father, just change color to black
		replacement.Color = BLACK
		return
	}

	// 要删除的叶子节点没有父亲，那么它是根节点，直接置空，返回
	if node.Parent == nil {
		tree.Root = nil
		return
	}

	// 要删除的叶子节点，是一个黑节点，删除后会破坏平衡，需要进行调整，调整成可以删除的状态
	if !IsRed(node) {
		// 核心函数
		tree.fixAfterDeletion(node)
	}

	// 现在可以删除叶子节点了
	if node == node.Parent.Left {
		node.Parent.Left = nil
	} else if node == node.Parent.Right {
		node.Parent.Right = nil
	}

	node.Parent = nil
}

// 调整删除的叶子节点，自底向上
// 可以看图理解
func (tree *RBTree[T]) fixAfterDeletion(node *RBTreeNode[T]) {
	// 如果不是递归到根节点，且节点是黑节点，那么继续递归
	for tree.Root != node && !IsRed(node) {
		// 要删除的节点在父亲左边，对应图例1，2
		if node == LeftOf(ParentOf(node)) {
			// 找出兄弟
			brother := RightOf(ParentOf(node))

			// 兄弟是红色的，对应图例1，那么兄弟变黑，父亲变红，然后对父亲左旋，进入图例21,22,23
			if IsRed(brother) {
				SetColor(brother, BLACK)
				SetColor(ParentOf(node), RED)
				tree.RotateLeft(ParentOf(node))
				brother = RightOf(ParentOf(node)) // 图例1调整后进入图例21,22,23，兄弟此时变了
			}

			// 兄弟是黑色的，对应图例21，22，23
			// 兄弟的左右儿子都是黑色，进入图例23，将兄弟设为红色，父亲所在的子树作为整体，当作删除的节点，继续向上递归
			if !IsRed(LeftOf(brother)) && !IsRed(RightOf(brother)) {
				SetColor(brother, RED)
				node = ParentOf(node)
			} else {
				// 兄弟的右儿子是黑色，进入图例22，将兄弟设为红色，兄弟的左儿子设为黑色，对兄弟右旋，进入图例21
				if !IsRed(RightOf(brother)) {
					SetColor(LeftOf(brother), BLACK)
					SetColor(brother, RED)
					tree.RotateRight(brother)
					brother = RightOf(ParentOf(node)) // 图例22调整后进入图例21，兄弟此时变了
				}

				// 兄弟的右儿子是红色，进入图例21，将兄弟设置为父亲的颜色，兄弟的右儿子以及父亲变黑，对父亲左旋
				SetColor(brother, ParentOf(node).Color)
				SetColor(ParentOf(node), BLACK)
				SetColor(RightOf(brother), BLACK)
				tree.RotateLeft(ParentOf(node))

				node = tree.Root
			}
		} else {
			// 要删除的节点在父亲右边，对应图例3，4
			// 找出兄弟
			brother := RightOf(ParentOf(node))

			// 兄弟是红色的，对应图例3，那么兄弟变黑，父亲变红，然后对父亲右旋，进入图例41,42,43
			if IsRed(brother) {
				SetColor(brother, BLACK)
				SetColor(ParentOf(node), RED)
				tree.RotateRight(ParentOf(node))
				brother = LeftOf(ParentOf(node)) // 图例3调整后进入图例41,42,43，兄弟此时变了
			}

			// 兄弟是黑色的，对应图例41，42，43
			// 兄弟的左右儿子都是黑色，进入图例43，将兄弟设为红色，父亲所在的子树作为整体，当作删除的节点，继续向上递归
			if !IsRed(LeftOf(brother)) && !IsRed(RightOf(brother)) {
				SetColor(brother, RED)
				node = ParentOf(node)
			} else {
				// 兄弟的左儿子是黑色，进入图例42，将兄弟设为红色，兄弟的右儿子设为黑色，对兄弟左旋，进入图例41
				if !IsRed(LeftOf(brother)) {
					SetColor(RightOf(brother), BLACK)
					SetColor(brother, RED)
					tree.RotateLeft(brother)
					brother = LeftOf(ParentOf(node)) // 图例42调整后进入图例41，兄弟此时变了
				}

				// 兄弟的左儿子是红色，进入图例41，将兄弟设置为父亲的颜色，兄弟的左儿子以及父亲变黑，对父亲右旋
				SetColor(brother, ParentOf(node).Color)
				SetColor(ParentOf(node), BLACK)
				SetColor(LeftOf(brother), BLACK)
				tree.RotateRight(ParentOf(node))

				node = tree.Root
			}
		}
	}

	// this node always black
	SetColor(node, BLACK)
}

// 找出最小值的节点
func (tree *RBTree[T]) FindMinValue() *RBTreeNode[T] {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}

	return tree.Root.FindMinValue()
}

// 找出最大值的节点
func (tree *RBTree[T]) FindMaxValue() *RBTreeNode[T] {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}

	return tree.Root.FindMaxValue()
}

// 查找指定节点 O(logN)
func (tree *RBTree[T]) Find(value T) *RBTreeNode[T] {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}

	return tree.Root.Find(value)
}

// 节点的颜色
func IsRed[T Comparator[T]](node *RBTreeNode[T]) bool {
	if node == nil {
		return false
	}
	return node.Color == RED
}

// 返回节点的父亲节点
func ParentOf[T Comparator[T]](node *RBTreeNode[T]) *RBTreeNode[T] {
	if node == nil {
		return nil
	}

	return node.Parent
}

// 返回节点的左子节点
func LeftOf[T Comparator[T]](node *RBTreeNode[T]) *RBTreeNode[T] {
	if node == nil {
		return nil
	}

	return node.Left
}

// 返回节点的右子节点
func RightOf[T Comparator[T]](node *RBTreeNode[T]) *RBTreeNode[T] {
	if node == nil {
		return nil
	}

	return node.Right
}

// 设置节点颜色
func SetColor[T Comparator[T]](node *RBTreeNode[T], color bool) {
	if node != nil {
		node.Color = color
	}
}
