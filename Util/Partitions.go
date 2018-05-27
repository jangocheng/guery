package Util

import ()

type PartitionInfo struct {
	Metadata  *Metadata
	Rows      []*Row
	Locations []string
	FileTypes []string
}

func NewPartitionInfo(md *Metadata) *PartitionInfo {
	return &PartitionInfo{
		Metadata: md,
		Rows:     []*Row{},
	}
}

func (self *PartitionInfo) GetPartitionNum() int {
	return len(self.Rows)
}

func (self *PartitionInfo) GetPartition(i int) *RowsBuffer {
	if i >= len(self.Rows) {
		return nil
	}
	rowsBuffer := NewRowsBuffer(self.Metadata)
	for _, row := range self.Rows {
		rowsBuffer.Write(row)
	}
	return rowsBuffer
}

func (self *PartitionInfo) GetLocation(i int) string {
	if i >= len(self.Locations) {
		return ""
	}
	return self.Locations[i]
}

func (self *PartitionInfo) GetFileType(i int) string {
	if i >= len(self.FileTypes) {
		return ""
	}
	return self.FileTypes[i]
}

func (self *PartitionInfo) Write(row *Row) {
	self.Rows = append(self.Rows, row)
}