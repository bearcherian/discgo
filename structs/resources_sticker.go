package structs

type Sticker struct {
	// Id [id of the sticker](#DOCS_REFERENCE/image-formatting)
	Id string `json:"id"`
	// PackId for standard stickers, id of the pack the sticker is from
	PackId string `json:"pack_id"`
	// Name name of the sticker
	Name string `json:"name"`
	// Description description of the sticker
	Description string `json:"description"`
	// Tags for guild stickers, the Discord name of a unicode emoji representing the sticker's expression. for standard stickers, a comma-separated list of related expressions.
	Tags string `json:"tags"`
	// Asset **Deprecated** previously the sticker asset hash, now an empty string
	Asset string `json:"asset"`
	// Type [type of sticker](#DOCS_RESOURCES_STICKER/sticker-object-sticker-types)
	Type int `json:"type"`
	// FormatType [type of sticker format](#DOCS_RESOURCES_STICKER/sticker-object-sticker-format-types)
	FormatType int `json:"format_type"`
	// Available whether this guild sticker can be used, may be false due to loss of Server Boosts
	Available bool `json:"available"`
	// GuildId id of the guild that owns this sticker
	GuildId string `json:"guild_id"`
	// User the user that uploaded the guild sticker
	User User `json:"user"`
	// SortValue the standard sticker's sort order within its pack
	SortValue int `json:"sort_value"`
}

type StickerItem struct {
	// Id id of the sticker
	Id string `json:"id"`
	// Name name of the sticker
	Name string `json:"name"`
	// FormatType [type of sticker format](#DOCS_RESOURCES_STICKER/sticker-object-sticker-format-types)
	FormatType int `json:"format_type"`
}

type StickerPack struct {
	// Id id of the sticker pack
	Id string `json:"id"`
	// Stickers the stickers in the pack
	Stickers []Sticker `json:"stickers"`
	// Name name of the sticker pack
	Name string `json:"name"`
	// SkuId id of the pack's SKU
	SkuId string `json:"sku_id"`
	// CoverStickerId id of a sticker in the pack which is shown as the pack's icon
	CoverStickerId string `json:"cover_sticker_id"`
	// Description description of the sticker pack
	Description string `json:"description"`
	// BannerAssetId id of the sticker pack's [banner image](#DOCS_REFERENCE/image-formatting)
	BannerAssetId string `json:"banner_asset_id"`
}

type Response struct {
	// StickerPacks
	StickerPacks []StickerPack `json:"sticker_packs"`
}
