package handler

import (
	"github.com/labstack/echo/v4"
	"microservices/app/libraries/grpc_client"
	HTMLparser "microservices/app/libraries/template_parser"
	"mime/multipart"
	"net/http"
)

type AdviceResponse struct {
	Text    string `json:"text"`
	Emotion string `json:"emotion"`
}

type ArrJSONs struct {
	PhotoDescription []*AdviceResponse
}

func (h *Handler) indexPage(ctx echo.Context) error {
	params := HTMLparser.TemplateParams{
		TemplateName: "index.html",
		Vars:         struct{}{},
	}

	data, err := HTMLparser.TemplateParser(params)
	if err != nil {
		h.logger.Errorf("error while parsing template %s: %s", params.TemplateName, err.Error())
		return err
	}
	if err := ctx.HTMLBlob(http.StatusOK, data); err != nil {
		h.logger.Errorf("error while htmlBlob : %s", err.Error())
		return err
	}
	return nil
}

func (h *Handler) getAdvices(ctx echo.Context) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		h.logger.Errorf("error while getting files from ctx.MultipartForm(): %s", err.Error())
	}

	files := form.File["data"]

	req := &grpc_client.GetAdviceRequest{}
	req.Data = make([]*multipart.FileHeader, 0, 7)

	for _, file := range files {
		req.Data = append(req.Data, file)
	}

	content, emotions, err := h.services.GetAdvise(ctx.Request().Context(), req)
	if err != nil {
		h.logger.Errorf("error while trying GetAdvice: %s", err.Error())
		err = ctx.JSON(http.StatusInternalServerError, &ArrJSONs{})
		return err
	}

	resp := &ArrJSONs{PhotoDescription: make([]*AdviceResponse, 0, len(content))}

	for i := 0; i < len(content) && i < len(emotions); i++ {
		resp.PhotoDescription = append(resp.PhotoDescription, &AdviceResponse{
			Text:    content[i],
			Emotion: emotions[i],
		})
	}

	err = ctx.JSON(http.StatusOK, resp)

	return err
}
