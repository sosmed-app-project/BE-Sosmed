package handler

import "hris-app-golang/feature/divisions"

type DivisionsResponseAll struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type DivisionsResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func DivisionCoreToResponseAll(input divisions.DivisionCore) DivisionsResponseAll {
	var divResp = DivisionsResponseAll{
		ID:   input.ID,
		Name: input.Name,
	}
	return divResp
}

