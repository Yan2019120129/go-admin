package display

import (
	"strconv"

	"ProgrammerYan/go-admin/context"
	"ProgrammerYan/go-admin/modules/utils"
	"ProgrammerYan/go-admin/template/types"
)

type FileSize struct {
	types.BaseDisplayFnGenerator
}

func init() {
	types.RegisterDisplayFnGenerator("filesize", new(FileSize))
}

func (f *FileSize) Get(ctx *context.Context, args ...interface{}) types.FieldFilterFn {
	return func(value types.FieldModel) interface{} {
		size, _ := strconv.ParseUint(value.Value, 10, 64)
		return utils.FileSize(size)
	}
}
