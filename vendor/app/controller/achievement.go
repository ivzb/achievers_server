package controller

import (
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
)

// Routes
func init() {
	router.Post("/achievements", alice.
		New(auth.Handler).
		ThenFunc(AchievementOnePOST))
}

// *****************************************************************************
// Create
// *****************************************************************************
func AchievementOnePOST(w http.ResponseWriter, r *http.Request) {
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

	response.Send(w, http.StatusCreated, ItemCreated, count, nil)
}
