package models

import "github.com/spoitler/glyph-browser/internal/models/types"

type Node struct {
	Children []Node
	NodeType types.NodeType
}
