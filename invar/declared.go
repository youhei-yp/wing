// -----------------------
// DECLARED METHODS
// ----------------------

package invar

// SetMimeTypesEnable set mime type map
// @declared
func SetMimeTypesEnable() {
	EnableMimeTypes()
}

// GetContentTypeHeader get mime type by file ext
// @declared
func GetContentTypeHeader(format string) *string {
	return GetContentType(format)
}

// AddAllowedDomain add web content file format based on mimeTypes
// @declared
func AddAllowedDomain(origin string) {
	PushDomain(origin)
}

// CheckRefererAllow check whether referer is allowed
// @declared
func CheckRefererAllow(referer string) bool {
	return ViaDomain(referer)
}

// SupportWebContent use to support web content
// @declared
func SupportWebContent(ext string) bool {
	return ViaWebContent(ext)
}
