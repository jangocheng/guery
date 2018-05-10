package Plan

import (
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromSampleRelation(t parser.ISampledRelationContext) PlanNode {
	tt := t.(*parser.SampledRelationContext)
	res := NewPlanNodeFromRelationPrimary(tt.RelationPrimary())
	if id := tt.Identifier(); id != nil {
		idNode := NewIdentifierNode(id)
		rename := idNode.GetText()
		renameNode := NewPlanRenameNode(rename, res)
		res.SetOutput(renameNode)
		res = renameNode
	}
	return res
}
