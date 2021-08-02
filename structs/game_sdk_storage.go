package structs

type FileStat struct {
	// Filename the name of the file
	Filename string `json:"Filename"`
	// Size the size of the file
	Size UInt64 `json:"Size"`
	// LastModified timestamp of when the file was last modified
	LastModified UInt64 `json:"LastModified"`
}