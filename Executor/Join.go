package Executor

import (
	"fmt"
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/Split"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionJoin(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanJoinNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{&enode.LeftInput, &enode.RightInput}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunJoin() (err error) {
	defer self.Clear()
	writer := self.Writers[0]
	enode := self.EPlanNode.(*EPlan.EPlanJoinNode)

	//read md
	if len(self.Readers) != 2 {
		return fmt.Errorf("join readers number %v <> 2", len(self.Readers))
	}

	mds := make([]*Metadata.Metadata, 2)
	if len(self.Readers) != 2 {
		return fmt.Errorf("join input number error")
	}
	for i, reader := range self.Readers {
		mds[i] = &Metadata.Metadata{}
		if err = Util.ReadObject(reader, mds[i]); err != nil {
			return err
		}
	}
	leftReader, rightReader := self.Readers[0], self.Readers[1]
	leftMd, rightMd := mds[0], mds[1]

	//write md
	if err = Util.WriteObject(writer, enode.Metadata); err != nil {
		return err
	}

	leftRbReader, rightRbReader := Split.NewSplitBuffer(leftMd, leftReader, nil), Split.NewSplitBuffer(rightMd, rightReader, nil)
	rbWriter := Split.NewSplitBuffer(enode.Metadata, nil, writer)

	defer func() {
		rbWriter.Flush()
	}()

	//write
	var sp *Split.Split
	rightSp := Split.NewSplit(rightMd)

	switch enode.JoinType {
	case Plan.INNERJOIN:
		fallthrough
	case Plan.LEFTJOIN:
		for {
			sp, err = rightRbReader.ReadSplit()
			if err == io.EOF {
				err = nil
				break
			}
			if err != nil {
				return err
			}
			rightSp.Append(sp)

		}

		for {
			sp, err = leftRbReader.ReadSplit()
			if err == io.EOF {
				err = nil
				break
			}
			if err != nil {
				return err
			}

			joinNum := 0
			for i := 0; i < sp.GetRowsNumber(); i++ {
				for j := 0; j < rightSp.GetRowsNumber(); j++ {
					joinSp := Split.NewSplit(enode.Metadata)
					vals := sp.GetValues(i)
					vals = append(vals, rightSp.GetValues(j)...)
					joinSp.AppendValues(vals)

					if ok, err := enode.JoinCriteria.Result(joinSp, 0); ok && err == nil {
						if err = rbWriter.Write(joinSp, 0); err != nil {
							return err
						}
						joinNum++
					} else if err != nil {
						return err
					}
				}

				if enode.JoinType == Plan.LEFTJOIN && joinNum == 0 {
					joinSp := Split.NewSplit(enode.Metadata)
					vals := sp.GetValues(i)
					vals = append(vals, make([]interface{}, rightSp.GetColumnNumber())...)
					joinSp.AppendValues(vals)

					if err = rbWriter.Write(joinSp, 0); err != nil {
						return err
					}
				}

			}

		}

	case Plan.RIGHTJOIN:
	}

	Logger.Infof("RunJoin finished")
	return err
}
