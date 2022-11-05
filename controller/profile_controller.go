package controller

import (
	"github.com/FTN-TwitterClone/profile/service"
	"go.opentelemetry.io/otel/trace"
)

type ProfileController struct {
	tracer         trace.Tracer
	profileService *service.ProfileService
}

func NewProfileController(tracer trace.Tracer, profileService *service.ProfileService) *ProfileController {
	return &ProfileController{
		tracer,
		profileService,
	}
}

//func (c *AuthController) RegisterUser(w http.ResponseWriter, req *http.Request) {
//	ctx, span := c.tracer.Start(req.Context(), "AuthController.RegisterUser")
//	defer span.End()
//
//	userForm, err := json.DecodeJson[model.RegisterUser](req.Body)
//
//	if err != nil {
//		span.SetStatus(codes.Error, err.Error())
//		http.Error(w, err.Error(), 500)
//		return
//	}
//
//	appErr := c.authService.RegisterUser(ctx, userForm)
//	if appErr != nil {
//		span.SetStatus(codes.Error, appErr.Error())
//		http.Error(w, appErr.Message, appErr.Code)
//		return
//	}
//}
