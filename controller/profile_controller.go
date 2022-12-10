package controller

import (
	"github.com/FTN-TwitterClone/profile/controller/json"
	"github.com/FTN-TwitterClone/profile/model"
	"github.com/FTN-TwitterClone/profile/service"
	"github.com/gorilla/mux"
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

	updateForm, err := json.DecodeJson[model.UpdateProfile](req.Body)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, "Error decoding from form json", 500)
		return
	}

	authUser := ctx.Value("authUser").(model.AuthUser)

	appErr := c.profileService.UpdateUser(ctx, &updateForm, &authUser)
	if appErr != nil {
		span.SetStatus(codes.Error, appErr.Error())
		http.Error(w, appErr.Message, appErr.Code)
		return
	}

}

func (c *ProfileController) GetPrivacy(w http.ResponseWriter, req *http.Request) {
	ctx, span := c.tracer.Start(req.Context(), "ProfileController.GetPrivacy")
	defer span.End()
	username := mux.Vars(req)["username"]

	user, appErr := c.profileService.GetUser(ctx, username)
	if appErr != nil {
		span.SetStatus(codes.Error, appErr.Error())
		http.Error(w, appErr.Message, appErr.Code)
		return
	}

	json.EncodeJson(w, user.Private)
}
