package structs

// type VoicePacket struct {
// 	// Version+Flags 1 byte
// 	VersionFlags byte `json:"Version + Flags"`
// 	// PayloadType 1 byte
// 	PayloadType Single byte value of `0x78` `json:"Payload Type"`
// 	// Sequence 2 bytes
// 	Sequence Unsigned short (big endian) `json:"Sequence"`
// 	// Timestamp 4 bytes
// 	Timestamp Unsigned integer (big endian) `json:"Timestamp"`
// 	// SSRC 4 bytes
// 	SSRC Unsigned integer (big endian) `json:"SSRC"`
// 	// EncryptedAudio n bytes
// 	EncryptedAudio Binary data `json:"Encrypted audio"`
// }