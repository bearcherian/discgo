package structs

type ImageType string

const ImageTypeUser ImageType = "User"

type ImageDimensions struct {
	// Width the width of the image
	Width uint32 `json:"Width"`
	// Height the height of the image
	Height uint32 `json:"Height"`
}

type ImageHandle struct {
	// Type the source of the image
	Type ImageType `json:"Type"`
	// Id the id of the user whose avatar you want to get
	Id int64 `json:"Id"`
	// Size the resolution at which you want the image
	Size uint32 `json:"Size"`
}
