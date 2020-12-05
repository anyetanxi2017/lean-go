package maps

import (
	"errors"
)

type Dictionary map[string]string

// 我们通过将错误提取为变量的方式，摆脱Search中魔术错误（magic error）
// 这也会使我们获得更好的测试。
var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
	// 我们可以重用ErrNotFound 而不添加新错误，更新失败时有时更精确的错误通常更好。
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

// 我们将错误声明了常量，这需要我们创建自己的 DictionaryErr 类型来实现error接口。
// 作用就是使错误更具有可重用性和不可变性
func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	// 为了使测试通过，我们使用了一个map查找的有趣特性，它可以返回两个值。第二个值是一个布尔值
	// 表示是否成功找到key 此特性让我们区分值是不存在还是未定义
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	// 这里我们使用 switch 语句来匹配错误。如上使用switch 提供了额外的安全，以防Search返回错误而不是ErrNotFound
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	// Go的map有一个内置函数 delete。它需要两个参数，第一个是这个map，第二个要删除的键
	// delete 函数不返回任何内容，我们基于相同的概念构建Delete方法。由于删除一个不存在的值没有影响的，所以我们
	// 不需要用错误复杂化API
	delete(d, word)
}

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

// Map有一个有趣的特性，不使用指针传递就可以修改它们。因为map是引用类型。
// 这意味着它拥有对底层结构的引用，就像指针一样。它底层的数据结构是 hash table 或 hash map
// Map 作为引用类型是非常好的，因为无论map有多大，都只会有一个副本。

// 引用类型引入了nil值。如果你尝试使用一个nil的map，你会得到一个nil指针异常。
// dictionary = map[string]string{}
// dictionary = make(map[string]string)
// 这两种方法都可以创建一个空的 hash map 并指向 dictionary。这确保永远不会获得nil指针异常。
