package controllers

import (
	"net/http"
)

func (c *{{ .FileNameCamelCase }}Controller) Get{{ .FileNameCamelCase }}(wctx restserver.WebContext) {
	var {{ .FileNameCamelCaseTitleLower }} models.{{ .FileNameCamelCase }}

	if err := wctx.DecodeQueryParams(&{{ .FileNameCamelCaseTitleLower }}); err != nil {
		wctx.ErrorResponse(http.StatusBadRequest, err)
		return
	}

	result, err := c.Usecase.Execute(wctx.Context(), {{ .FileNameCamelCaseTitleLower }})

	if err != nil {
		wctx.ErrorResponse(http.StatusInternalServerError, err)
		return
	}

	wctx.JsonResponse(http.StatusOK, result)
}
