package controller

import (
	"github.com/FTN-TwitterClone/profile/controller/json"
	"github.com/FTN-TwitterClone/profile/model"
	"github.com/FTN-TwitterClone/profile/service"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"net/http"
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

func (c *ProfileController) GetUser(w http.ResponseWriter, req *http.Request) {
	_, span := c.tracer.Start(req.Context(), "AuthController.RegisterUser")
	defer span.End()
}

func (c *ProfileController) SaveUser(w http.ResponseWriter, req *http.Request) {
	ctx, span := c.tracer.Start(req.Context(), "ProfileController.SaveUser")
	defer span.End()

	user, err := json.DecodeJson[*model.ProfileUser](req.Body)

	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	newUser, appErr := c.profileService.SaveUser(ctx, user)
	if appErr != nil {
		span.SetStatus(codes.Error, appErr.Error())
		http.Error(w, appErr.Message, appErr.Code)
		return
	}
	responseErr := json.EncodeJson(w, newUser)
	if responseErr != nil {
		span.SetStatus(codes.Error, responseErr.Error())
		http.Error(w, responseErr.Error(), 500)
		return
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
