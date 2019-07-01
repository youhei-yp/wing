// Copyright (c) 2018-2019 WING All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// 00002       2019/06/30   zhaixing       Add function from godfs
// -------------------------------------------------------------------

package invar

import (
	"container/list"
	"strings"
)

// mime types for browser download header 'Content-Type'
var (
	mimeTypes       = make(map[string]string)
	webMimeTypes    = make(map[string]string)
	allowedDomains  = list.New()
	enable          = false
	defaultMimeType = "application/octet-stream"
)

// SetMimeTypesEnable set mime type map
// if enable it will auto add download header 'Content-Type' before download,
// if it is disabled(by default), download header 'Content-Type' will set to 'application/octet-stream'.
func SetMimeTypesEnable() {

	enable = true

	mimeTypes["html"] = "text/html"                                                                 // file ext is html
	mimeTypes["shtml"] = "text/html"                                                                // file ext is shtml
	mimeTypes["css"] = "text/css"                                                                   // file ext is css
	mimeTypes["xml"] = "text/xml"                                                                   // file ext is xml
	mimeTypes["gif"] = "image/gif"                                                                  // file ext is gif
	mimeTypes["jpeg"] = "image/jpeg"                                                                // file ext is jpeg
	mimeTypes["jpg"] = "image/jpeg"                                                                 // file ext is jpg
	mimeTypes["js"] = "application/javascript"                                                      // file ext is js
	mimeTypes["atom"] = "application/atom+xml"                                                      // file ext is atom
	mimeTypes["rss"] = "application/rss+xml"                                                        // file ext is rss
	mimeTypes["mml"] = "text/mathml"                                                                // file ext is mml
	mimeTypes["txt"] = "text/plain"                                                                 // file ext is txt
	mimeTypes["jad"] = "text/vnd.sun.j2me.app-descriptor"                                           // file ext is jad
	mimeTypes["wml"] = "text/vnd.wap.wml"                                                           // file ext is wml
	mimeTypes["htc"] = "text/x-component"                                                           // file ext is htc
	mimeTypes["png"] = "image/png"                                                                  // file ext is png
	mimeTypes["tif"] = "image/tiff"                                                                 // file ext is tif
	mimeTypes["tiff"] = "image/tiff"                                                                // file ext is tiff
	mimeTypes["wbmp"] = "image/vnd.wap.wbmp"                                                        // file ext is wbmp
	mimeTypes["ico"] = "image/x-icon"                                                               // file ext is ico
	mimeTypes["jng"] = "image/x-jng"                                                                // file ext is jng
	mimeTypes["bmp"] = "image/x-ms-bmp"                                                             // file ext is bmp
	mimeTypes["svg"] = "image/svg+xml"                                                              // file ext is svg
	mimeTypes["svgz"] = "image/svg+xml"                                                             // file ext is svgz
	mimeTypes["webp"] = "image/webp"                                                                // file ext is webp
	mimeTypes["woff"] = "application/font-woff"                                                     // file ext is woff
	mimeTypes["jar"] = "application/java-archive"                                                   // file ext is jar
	mimeTypes["war"] = "application/java-archive"                                                   // file ext is war
	mimeTypes["ear"] = "application/java-archive"                                                   // file ext is ear
	mimeTypes["json"] = "application/json"                                                          // file ext is json
	mimeTypes["hqx"] = "application/mac-binhex40"                                                   // file ext is hqx
	mimeTypes["doc"] = "application/msword"                                                         // file ext is doc
	mimeTypes["pdf"] = "application/pdf"                                                            // file ext is pdf
	mimeTypes["ps"] = "application/postscript"                                                      // file ext is ps
	mimeTypes["eps"] = "application/postscript"                                                     // file ext is eps
	mimeTypes["ai"] = "application/postscript"                                                      // file ext is ai
	mimeTypes["rtf"] = "application/rtf"                                                            // file ext is rtf
	mimeTypes["m3u8"] = "application/vnd.apple.mpegurl"                                             // file ext is m3u8
	mimeTypes["xls"] = "application/vnd.ms-excel"                                                   // file ext is xls
	mimeTypes["eot"] = "application/vnd.ms-fontobject"                                              // file ext is eot
	mimeTypes["ppt"] = "application/vnd.ms-powerpoint"                                              // file ext is ppt
	mimeTypes["wmlc"] = "application/vnd.wap.wmlc"                                                  // file ext is wmlc
	mimeTypes["kml"] = "application/vnd.google-earth.kml+xml"                                       // file ext is kml
	mimeTypes["kmz"] = "application/vnd.google-earth.kmz"                                           // file ext is kmz
	mimeTypes["7z"] = "application/x-7z-compressed"                                                 // file ext is 7z
	mimeTypes["cco"] = "application/x-cocoa"                                                        // file ext is cco
	mimeTypes["jardiff"] = "application/x-java-archive-diff"                                        // file ext is jardiff
	mimeTypes["jnlp"] = "application/x-java-jnlp-file"                                              // file ext is jnlp
	mimeTypes["run"] = "application/x-makeself"                                                     // file ext is run
	mimeTypes["pl"] = "application/x-perl"                                                          // file ext is pl
	mimeTypes["pm"] = "application/x-perl"                                                          // file ext is pm
	mimeTypes["prc"] = "application/x-pilot"                                                        // file ext is prc
	mimeTypes["pdb"] = "application/x-pilot"                                                        // file ext is pdb
	mimeTypes["rar"] = "application/x-rar-compressed"                                               // file ext is rar
	mimeTypes["rpm"] = "application/x-redhat-package-manager"                                       // file ext is rpm
	mimeTypes["sea"] = "application/x-sea"                                                          // file ext is sea
	mimeTypes["swf"] = "application/x-shockwave-flash"                                              // file ext is swf
	mimeTypes["sit"] = "application/x-stuffit"                                                      // file ext is sit
	mimeTypes["tcl"] = "application/x-tcl"                                                          // file ext is tcl
	mimeTypes["tk"] = "application/x-tcl"                                                           // file ext is tk
	mimeTypes["der"] = "application/x-x509-ca-cert"                                                 // file ext is der
	mimeTypes["pem"] = "application/x-x509-ca-cert"                                                 // file ext is pem
	mimeTypes["crt"] = "application/x-x509-ca-cert"                                                 // file ext is crt
	mimeTypes["xpi"] = "application/x-xpinstall"                                                    // file ext is xpi
	mimeTypes["xhtml"] = "application/xhtml+xml"                                                    // file ext is xhtml
	mimeTypes["xspf"] = "application/xspf+xml"                                                      // file ext is xspf
	mimeTypes["zip"] = "application/zip"                                                            // file ext is zip
	mimeTypes["docx"] = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"   // file ext is docx
	mimeTypes["xlsx"] = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"         // file ext is xlsx
	mimeTypes["pptx"] = "application/vnd.openxmlformats-officedocument.presentationml.presentation" // file ext is pptx
	mimeTypes["mid"] = "audio/midi"                                                                 // file ext is mid
	mimeTypes["midi"] = "audio/midi"                                                                // file ext is midi
	mimeTypes["kar"] = "audio/midi"                                                                 // file ext is kar
	mimeTypes["mp3"] = "audio/mpeg"                                                                 // file ext is mp3
	mimeTypes["ogg"] = "audio/ogg"                                                                  // file ext is ogg
	mimeTypes["m4a"] = "audio/x-m4a"                                                                // file ext is m4a
	mimeTypes["ra"] = "audio/x-realaudio"                                                           // file ext is ra
	mimeTypes["3gpp"] = "video/3gpp"                                                                // file ext is 3gpp
	mimeTypes["3gp"] = "video/3gpp"                                                                 // file ext is 3gp
	mimeTypes["ts"] = "video/mp2t"                                                                  // file ext is ts
	mimeTypes["mp4"] = "video/mp4"                                                                  // file ext is mp4
	mimeTypes["mpeg"] = "video/mpeg"                                                                // file ext is mpeg
	mimeTypes["mpg"] = "video/mpeg"                                                                 // file ext is mpg
	mimeTypes["mov"] = "video/quicktime"                                                            // file ext is mov
	mimeTypes["webm"] = "video/webm"                                                                // file ext is webm
	mimeTypes["flv"] = "video/x-flv"                                                                // file ext is flv
	mimeTypes["m4v"] = "video/x-m4v"                                                                // file ext is m4v
	mimeTypes["mng"] = "video/x-mng"                                                                // file ext is mng
	mimeTypes["asx"] = "video/x-ms-asf"                                                             // file ext is asx
	mimeTypes["asf"] = "video/x-ms-asf"                                                             // file ext is asf
	mimeTypes["wmv"] = "video/x-ms-wmv"                                                             // file ext is wmv
	mimeTypes["avi"] = "video/x-msvideo"                                                            // file ext is avi
}

// GetContentTypeHeader get mime type by file ext
func GetContentTypeHeader(fileFormat string) *string {
	if enable {
		format := mimeTypes[strings.TrimLeft(fileFormat, ".")]
		if format == "" {
			return &defaultMimeType
		}
		return &format
	}
	return &defaultMimeType
}

// AddWebMimeType add web content file format based on mimeTypes
func AddWebMimeType(format string) {
	webMimeTypes[format] = mimeTypes[format]
}

// AddAllowedDomain add web content file format based on mimeTypes
func AddAllowedDomain(origin string) {
	for ele := allowedDomains.Front(); ele != nil; ele = ele.Next() {
		if ele.Value.(string) == origin {
			return
		}
	}
	allowedDomains.PushBack(origin)
}

// CheckRefererAllow check whether referer is allowed
func CheckRefererAllow(referer string) bool {
	if allowedDomains.Len() == 0 {
		return true
	}
	for ele := allowedDomains.Front(); ele != nil; ele = ele.Next() {
		if strings.HasPrefix(referer, ele.Value.(string)) {
			return true
		}
	}
	return false
}

// SupportWebContent use to support web content
func SupportWebContent(ext string) bool {
	ret := webMimeTypes[strings.TrimLeft(ext, ".")]
	if ret == "" {
		return false
	}
	return true
}
