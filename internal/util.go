package internal

import (
	"errors"

	"github.com/YoshihikoAbe/avsproperty"
)

func NewUniqueNode(parent *avsproperty.Node, name string) *avsproperty.Node {
	node := parent.SearchChild(name)
	if node != nil {
		return node
	}
	node, err := parent.NewNode(name)
	if err != nil {
		panic(err)
	}
	return node
}

func SetChildValue(node *avsproperty.Node, name string, v any) {
	if child := node.SearchChild(name); child != nil {
		child.SetValue(v)
	} else {
		node.NewNodeWithValue(name, v)
	}
}

func CompareNodeName(want string, node *avsproperty.Node) error {
	wantName, err := avsproperty.NewNodeName(want)
	if err != nil {
		return err
	}
	if !node.Name().Equals(wantName) {
		return errors.New("invalid node name: " + node.Name().String() + " != " + want)
	}
	return nil
}
