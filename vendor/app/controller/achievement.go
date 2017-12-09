package controller

import (
	"strconv"
	"net/http"

	"app/model/achievement"
	"app/model/involvement"

	"app/route/middleware/auth"

	"app/shared/form"
	"app/shared/response"
	"app/shared/router"

	"github.com/gorilla/context"
	"github.com/justinas/alice"
)

const (
	InvolvementNotFound = "involvement not found"
	InvalidPage = "invalid page"
)

// Routes
func init() {
	router.Post("/achievement", alice.
		New(auth.Handler).
		ThenFunc(PostAchievement))

	router.Get("/achievement/:id", alice.
		New(auth.Handler).
		ThenFunc(GetAchievementById))

	router.Get("/achievements/:page", alice.
		New(auth.Handler).
		ThenFunc(GetAchievementsByPage))
}

// *****************************************************************************
// Read
// *****************************************************************************

func GetAchievementById(w http.ResponseWriter, r *http.Request) {
	// Get the parameter id
	params := router.Params(r)
	ID := params.ByName("id")

	// Get an item
	entity, err := achievement.Get(ID)
	if err == achievement.ErrNoResult {
		response.Send(w, http.StatusOK, ItemNotFound, 0, nil)
		return
	}

	if err != nil {
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	response.Send(w, http.StatusOK, ItemFound, 1, entity)
}

func GetAchievementsByPage(w http.ResponseWriter, r *http.Request) {
	// Get the parameter id
	params := router.Params(r)

	page64, err := strconv.ParseInt(params.ByName("page"), 0, 32)
	
	if err != nil || page64 < 0 {
		response.SendError(w, http.StatusBadRequest, InvalidPage)
		return
	}

	page := int(page64)

	// Get all items
	group, err := achievement.Load(page)
	if err != nil {
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	if len(group) < 1 {
		response.Send(w, http.StatusNotFound, ItemsFindEmpty, len(group), nil)
		return
	}

	response.Send(w, http.StatusOK, ItemsFound, len(group), group)
}

// *****************************************************************************
// Create
// *****************************************************************************
func PostAchievement(w http.ResponseWriter, r *http.Request) {
	m, err := achievement.New()
	if err != nil {
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	// Validate the required fields are present
	err, errMsg := form.Validate(r, m)
	if err == form.ErrRequiredMissing || err == form.ErrWrongContentType {
		response.SendError(w, http.StatusBadRequest, errMsg)
		return
	}

	if err == form.ErrBadStruct || err == form.ErrNotStruct {
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	// Validate value types and copy values to struct
	err, errMsg = form.StructCopy(r, m)
	if err == form.ErrWrongType {
		response.SendError(w, http.StatusBadRequest, errMsg)
		return
	}

	if err == form.ErrNotSupported || err == form.ErrNotStruct {
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	i, err := involvement.Exist(m.InvolvementID)

	if err != nil || i != true {
		response.SendError(w, http.StatusBadRequest, InvolvementNotFound)
		return
	}

	// Create item
	uID := context.Get(r, "userID").(string)

	count, err := m.Create(uID)
	if err != nil {
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	response.Send(w, http.StatusCreated, ItemCreated, count, m.ID)
}
