package service

import (
	"github.com/harshabangi/LLD/apps/parking-lot/internal/db"
	"github.com/harshabangi/LLD/apps/parking-lot/pkg"
)

type Handler struct {
	db db.DB
}

func NewHandler(db db.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) AddFloor(floor pkg.Floor) error {
	return h.db.AddFloor(floor)
}

func (h *Handler) AddSpot(floorID string, spot pkg.Spot) error {
	return h.db.AddSpot(floorID, spot)
}
