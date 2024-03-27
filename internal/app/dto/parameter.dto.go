package dto

// ParameterCreateDTO is used to capture the data required to create a new parameter.
type ParameterCreateDTO struct {
	Name       string `json:"name" binding:"required"`
	HostDevice int    `json:"hostDevice"`
	Device     int    `json:"device"`
	Log        int    `json:"log"`
	Point      string `json:"point" binding:"required"`
	Unit       string `json:"unit" binding:"required"`
}
