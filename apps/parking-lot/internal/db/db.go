package db

import (
	"fmt"
	"github.com/harshabangi/LLD/apps/parking-lot/pkg"
)

type db struct {
	floors   []pkg.Floor
	floorMap map[string]struct{}
}

type DB interface {
	AddFloor(floor pkg.Floor) error
	AddSpot(floorID string, spot pkg.Spot) error
}

func NewDB() DB {
	return &db{
		floors:   []pkg.Floor{},
		floorMap: map[string]struct{}{},
	}
}

func (d *db) AddFloor(floor pkg.Floor) error {

	// check if there is any other floor create with same id
	if _, ok := d.floorMap[floor.ID]; ok {
		return fmt.Errorf("floor with id %s already exists", floor.ID)
	}

	d.floors = append(d.floors, floor)
	d.floorMap[floor.ID] = struct{}{}

	return nil
}

func (d *db) GetFloor(floorID string) (floor *pkg.Floor, err error) {

	if _, ok := d.floorMap[floorID]; !ok {
		err = fmt.Errorf("no such floor with id: %s", floorID)
		return
	}

	for _, v := range d.floors {
		if v.ID == floorID {
			floor = &v
			return
		}
	}
	return
}

func (d *db) AddSpot(floorID string, spot pkg.Spot) error {

	floor, err := d.GetFloor(floorID)
	if err != nil {
		return err
	}

	if _, ok := floor.SpotMap[spot.ID]; ok {
		return fmt.Errorf("spot with id %s already exists in floor with id %s",
			spot.ID, floor.ID)
	}

	// add spot in a floor
	floor.Spots = append(floor.Spots, spot)
	floor.SpotMap[spot.ID] = len(floor.Spots) - 1

	return nil
}

func remove(s []pkg.Spot, i int) []pkg.Spot {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (d *db) VacateSpot(floorID string, spotID string) error {

	floor, err := d.GetFloor(floorID)
	if err != nil {
		return err
	}

	idx, ok := floor.SpotMap[spotID]
	if !ok {
		return fmt.Errorf("no such spot %s in the floor %s", spotID, floorID)
	}

	delete(floor.SpotMap, spotID)
	floor.Spots = remove(floor.Spots, idx)

	return nil
}
