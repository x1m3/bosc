package bosc

import (
	"errors"
	"fmt"
	"strings"
)

type binaryTreeFactory func() (BinarySearchTree)

var knownTypeTrees map[string]binaryTreeFactory = make(map[string]binaryTreeFactory)

func NewTree(name string) (BinarySearchTree, error) {
	if factory, found := knownTypeTrees[name]; found {
		return factory(), nil
	}else {
		return nil, errors.New(fmt.Sprintf("Unknown tree type <%s>. Valid types are <%s>", name, strings.Join(TreeTypes(),", ")))
	}
}

func register(name string, factory binaryTreeFactory) {
	knownTypeTrees[name] = factory
}

func TreeTypes() []string {
	types := make([]string,0,1)
	for t := range knownTypeTrees {
		types = append(types, t)
	}
	return types
}