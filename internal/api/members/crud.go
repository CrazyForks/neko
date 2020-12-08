package members

import (
	"net/http"

	"demodesk/neko/internal/utils"
	"demodesk/neko/internal/types"
)

type MemberCreatePayload struct {
	ID string `json:"id"`
	*types.MemberProfile
}

func (h *MembersHandler) membersCreate(w http.ResponseWriter, r *http.Request) {
	data := &MemberCreatePayload{
		MemberProfile: &types.MemberProfile{
			IsAdmin: false,
			CanLogin: true,
			CanConnect: true,
			CanWatch: true,
			CanHost: true,
			CanAccessClipboard: true,
		},
	}

	if !utils.HttpJsonRequest(w, r, data) {
		return
	}

	if data.Secret == "" {
		utils.HttpBadRequest(w, "Secret cannot be empty.")
		return
	}

	if data.Name == "" {
		utils.HttpBadRequest(w, "Name cannot be empty.")
		return
	}

	if data.ID == "" {
		var err error
		if data.ID, err = utils.NewUID(32); err != nil {
			utils.HttpInternalServerError(w, err)
			return
		}
	} else {
		if _, ok := h.sessions.Get(data.ID); ok {
			utils.HttpBadRequest(w, "Member ID already exists.")
			return
		}
	}

	session, err := h.sessions.Create(data.ID, *data.MemberProfile)
	if err != nil {
		utils.HttpInternalServerError(w, err)
		return
	}

	utils.HttpSuccess(w, MemberCreatePayload{
		ID: session.ID(),
	})
}

func (h *MembersHandler) membersRead(w http.ResponseWriter, r *http.Request) {
	member := GetMember(r)

	// TODO: Get whole profile from session.
	utils.HttpSuccess(w, types.MemberProfile{
		Name: member.Name(),
		IsAdmin: member.IsAdmin(),
		CanLogin: member.CanLogin(),
		CanConnect: member.CanConnect(),
		CanWatch: member.CanWatch(),
		CanHost: member.CanHost(),
		CanAccessClipboard: member.CanAccessClipboard(),
	})
}

func (h *MembersHandler) membersUpdate(w http.ResponseWriter, r *http.Request) {
	member := GetMember(r)

	// TODO: Get whole profile from session.
	profile := types.MemberProfile{
		Name: member.Name(),
		IsAdmin: member.IsAdmin(),
		CanLogin: member.CanLogin(),
		CanConnect: member.CanConnect(),
		CanWatch: member.CanWatch(),
		CanHost: member.CanHost(),
		CanAccessClipboard: member.CanAccessClipboard(),
	}

	if !utils.HttpJsonRequest(w, r, &profile) {
		return
	}

	if err := h.sessions.Update(member.ID(), profile); err != nil {
		utils.HttpInternalServerError(w, err)
		return
	}

	utils.HttpSuccess(w)
}

func (h *MembersHandler) membersDelete(w http.ResponseWriter, r *http.Request) {
	member := GetMember(r)

	if err := h.sessions.Delete(member.ID()); err != nil {
		utils.HttpInternalServerError(w, err)
		return
	}

	utils.HttpSuccess(w)
}
