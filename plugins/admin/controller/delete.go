package controller

import (
	"github.com/Yan2019120129/go-admin/context"
	"github.com/Yan2019120129/go-admin/modules/logger"
	"github.com/Yan2019120129/go-admin/plugins/admin/modules/guard"
	"github.com/Yan2019120129/go-admin/plugins/admin/modules/response"
)

// Delete delete the row from database.
func (h *Handler) Delete(ctx *context.Context) {

	param := guard.GetDeleteParam(ctx)

	//token := ctx.FormValue("_t")
	//
	//if !auth.TokenHelper.CheckToken(token) {
	//	ctx.SetStatusCode(http.StatusBadRequest)
	//	ctx.WriteString(`{"code":400, "msg":"delete fail"}`)
	//	return
	//}

	if err := h.table(param.Prefix, ctx).DeleteData(param.Id); err != nil {
		logger.ErrorCtx(ctx, "Delete error %+v", err)
		response.Error(ctx, "delete fail")
		return
	}

	response.OkWithData(ctx, map[string]interface{}{
		"token": h.authSrv().AddToken(),
	})
}
