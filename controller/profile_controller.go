package controller

import (
	"github.com/FTN-TwitterClone/profile/controller/json"
	"github.com/FTN-TwitterClone/profile/service"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"strconv"
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
	ctx, span := c.tracer.Start(req.Context(), "ProfileController.GetUser")
	defer span.End()

	username := mux.Vars(req)["username"]

	user, appErr := c.profileService.GetUser(ctx, username)
	if appErr != nil {
		span.SetStatus(codes.Error, appErr.Error())
		http.Error(w, appErr.Message, appErr.Code)
		return
	}

	json.EncodeJson(w, user)
}

func (c *ProfileController) UpdateMyDetails(w http.ResponseWriter, req *http.Request) {
	ctx, span := c.tracer.Start(req.Context(), "ProfileController.UpdateMyDetails")
	defer span.End()

	username := mux.Vars(req)["username"]

	user, appErr := c.profileService.GetUser(ctx, username)
	if appErr != nil {
		span.SetStatus(codes.Error, appErr.Error())
		http.Error(w, appErr.Message, appErr.Code)
		return
	}
	changedValue, err := strconv.ParseBool(mux.Vars(req)["private"])
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, "Wrong value.", 500)
		return
	}
	user.Private = changedValue

}
