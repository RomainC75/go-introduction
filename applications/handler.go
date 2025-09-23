package applications

import "sandbox/repos"

type Handler struct {
	repos []repos.Repos
}

func NewHandler(repos ...repos.Repos) *Handler {
	return &Handler{
		repos: repos,
	}
}

func (h *Handler) GetAllIds() []int {
	allIds := []int{}
	for _, r := range h.repos {
		ids := r.GetIds()
		allIds = append(allIds, ids...)
	}
	return allIds
}
