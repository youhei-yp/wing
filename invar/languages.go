// Copyright (c) 2018-2019 Dunyu All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// -------------------------------------------------------------------

package invar

import (
	"strings"
)

// Lang language code type
type Lang int

const (
	Lang_ar_IL Lang = iota // Arabic (Israel)
	Lang_ar_EG             // Arabic (Egypt)
	Lang_zh_CN             // Chinese Simplified
	Lang_zh_TW             // Chinese Tradition
	Lang_zh_HK             // Chinese Hongkong
	Lang_nl_NL             // Dutch(Netherlands)
	Lang_nl_BE             // Dutch(Netherlands)
	Lang_en_US             // English(United States)
	Lang_en_AU             // English(Australia)
	Lang_en_CA             // English(Canada)
	Lang_en_IN             // English(India)
	Lang_en_IE             // English(Ireland)
	Lang_en_NZ             // English(New Zealand)
	Lang_en_SG             // English(Singapore)
	Lang_en_ZA             // English(South Africa)
	Lang_en_GB             // English(United Kingdom)
	Lang_fr_FR             // French
	Lang_fr_BE             // French
	Lang_fr_CA             // French
	Lang_fr_CH             // French
	Lang_de_DE             // German
	Lang_de_LI             // German
	Lang_de_AT             // German
	Lang_de_CH             // German
	Lang_it_IT             // Italian
	Lang_it_CH             // Italian
	Lang_pt_BR             // Portuguese
	Lang_pt_PT             // Portuguese
	Lang_es_ES             // Spanish
	Lang_es_US             // Spanish
	Lang_bn_BD             // Bengali
	Lang_bn_IN             // Bengali
	Lang_hr_HR             // Croatian
	Lang_cs_CZ             // Czech
	Lang_da_DK             // Danish
	Lang_el_GR             // Greek
	Lang_he_IL             // Hebrew
	Lang_iw_IL             // Hebrew
	Lang_hi_IN             // Hindi
	Lang_hu_HU             // Hungarian
	Lang_in_ID             // Indonesian
	Lang_ja_JP             // Japanese
	Lang_ko_KR             // Korean
	Lang_ms_MY             // Malay
	Lang_fa_IR             // Perisan
	Lang_pl_PL             // Polish
	Lang_ro_RO             // Romanian
	Lang_ru_RU             // Russian
	Lang_sr_RS             // Serbian
	Lang_sv_SE             // Swedish
	Lang_th_TH             // Thai
	Lang_tr_TR             // Turkey
	Lang_ur_PK             // Urdu
	Lang_vi_VN             // Vietnamese
	Lang_ca_ES             // Catalan
	Lang_lv_LV             // Latviesu
	Lang_lt_LT             // Lithuanian
	Lang_nb_NO             // Norwegian
	Lang_sk_SK             // slovencina
	Lang_sl_SI             // Slovenian
	Lang_bg_BG             // bulgarian
	Lang_uk_UA             // Ukrainian
	Lang_tl_PH             // Filipino
	Lang_fi_FI             // Finnish
	Lang_af_ZA             // Afrikaans
	Lang_rm_CH             // Romansh
	Lang_my_ZG             // Burmese
	Lang_my_MM             // Burmese
	Lang_km_KH             // Khmer
	Lang_am_ET             // Amharic
	Lang_be_BY             // Belarusian
	Lang_et_EE             // Estonian
	Lang_sw_TZ             // Swahili
	Lang_zu_ZA             // Zulu
	Lang_az_AZ             // Azerbaijani
	Lang_hy_AM             // Armenian
	Lang_ka_GE             // Georgian
	Lang_lo_LA             // Laotian
	Lang_mn_MN             // Mongolian
	Lang_ne_NP             // Nepali
	Lang_kk_KZ             // Kazakh
	Lang_si_LK             // Sinhala
)

// Language language information
type Language struct {
	Code   Lang
	Key    string
	EnName string
	CnName string
}

// languagesCache languages information cache
var languagesCache = make(map[Lang]*Language)

const (
	// InvalidLangCode invalid language code
	InvalidLangCode Lang = -1

	// LangsSeparator multi-langguages separator
	LangsSeparator = ","
)

func init() {
	languagesCache[Lang_ar_IL] = &Language{Lang_ar_IL, "ar_IL", "Arabic(Israel)", "阿拉伯语(以色列)"}
	languagesCache[Lang_ar_EG] = &Language{Lang_ar_EG, "ar_EG", "Arabic(Egypt)", "阿拉伯语(埃及)"}
	languagesCache[Lang_zh_CN] = &Language{Lang_zh_CN, "zh_CN", "Chinese Simplified", "中文简体"}
	languagesCache[Lang_zh_TW] = &Language{Lang_zh_TW, "zh_TW", "Chinese Tradition", "中文繁体"}
	languagesCache[Lang_zh_HK] = &Language{Lang_zh_HK, "zh_HK", "Chinese Hongkong", "中文(香港)"}
	languagesCache[Lang_nl_NL] = &Language{Lang_nl_NL, "nl_NL", "Dutch(Netherlands)", "荷兰语"}
	languagesCache[Lang_nl_BE] = &Language{Lang_nl_BE, "nl_BE", "Dutch(Netherlands)", "荷兰语(比利时)"}
	languagesCache[Lang_en_US] = &Language{Lang_en_US, "en_US", "English(United States)", "英语(美国)"}
	languagesCache[Lang_en_AU] = &Language{Lang_en_AU, "en_AU", "English(Australia)", "英语(澳大利亚)"}
	languagesCache[Lang_en_CA] = &Language{Lang_en_CA, "en_CA", "English(Canada)", "英语(加拿大)"}
	languagesCache[Lang_en_IN] = &Language{Lang_en_IN, "en_IN", "English(India)", "英语(印度)"}
	languagesCache[Lang_en_IE] = &Language{Lang_en_IE, "en_IE", "English(Ireland)", "英语(爱尔兰)"}
	languagesCache[Lang_en_NZ] = &Language{Lang_en_NZ, "en_NZ", "English(New Zealand)", "英语(新西兰)"}
	languagesCache[Lang_en_SG] = &Language{Lang_en_SG, "en_SG", "English(Singapore)", "英语(新加波)"}
	languagesCache[Lang_en_ZA] = &Language{Lang_en_ZA, "en_ZA", "English(South Africa)", "英语(南非)"}
	languagesCache[Lang_en_GB] = &Language{Lang_en_GB, "en_GB", "English(United Kingdom)", "英语(英国)"}
	languagesCache[Lang_fr_FR] = &Language{Lang_fr_FR, "fr_FR", "French", "法语"}
	languagesCache[Lang_fr_BE] = &Language{Lang_fr_BE, "fr_BE", "French", "法语(比利时)"}
	languagesCache[Lang_fr_CA] = &Language{Lang_fr_CA, "fr_CA", "French", "法语(加拿大)"}
	languagesCache[Lang_fr_CH] = &Language{Lang_fr_CH, "fr_CH", "French", "法语(瑞士)"}
	languagesCache[Lang_de_DE] = &Language{Lang_de_DE, "de_DE", "German", "德语"}
	languagesCache[Lang_de_LI] = &Language{Lang_de_LI, "de_LI", "German", "德语(列支敦斯登)"}
	languagesCache[Lang_de_AT] = &Language{Lang_de_AT, "de_AT", "German", "德语(奥地利)"}
	languagesCache[Lang_de_CH] = &Language{Lang_de_CH, "de_CH", "German", "德语(瑞士)"}
	languagesCache[Lang_it_IT] = &Language{Lang_it_IT, "it_IT", "Italian", "意大利语"}
	languagesCache[Lang_it_CH] = &Language{Lang_it_CH, "it_CH", "Italian", "意大利语(瑞士)"}
	languagesCache[Lang_pt_BR] = &Language{Lang_pt_BR, "pt_BR", "Portuguese", "葡萄牙语（巴西）"}
	languagesCache[Lang_pt_PT] = &Language{Lang_pt_PT, "pt_PT", "Portuguese", "葡萄牙语"}
	languagesCache[Lang_es_ES] = &Language{Lang_es_ES, "es_ES", "Spanish", "西班牙语"}
	languagesCache[Lang_es_US] = &Language{Lang_es_US, "es_US", "Spanish", "西班牙语(美国)"}
	languagesCache[Lang_bn_BD] = &Language{Lang_bn_BD, "bn_BD", "Bengali", "孟加拉语"}
	languagesCache[Lang_bn_IN] = &Language{Lang_bn_IN, "bn_IN", "Bengali", "孟加拉语(印度)"}
	languagesCache[Lang_hr_HR] = &Language{Lang_hr_HR, "hr_HR", "Croatian", "克罗地亚语"}
	languagesCache[Lang_cs_CZ] = &Language{Lang_cs_CZ, "cs_CZ", "Czech", "捷克语"}
	languagesCache[Lang_da_DK] = &Language{Lang_da_DK, "da_DK", "Danish", "丹麦语"}
	languagesCache[Lang_el_GR] = &Language{Lang_el_GR, "el_GR", "Greek", "希腊语"}
	languagesCache[Lang_he_IL] = &Language{Lang_he_IL, "he_IL", "Hebrew", "希伯来语(以色列)"}
	languagesCache[Lang_iw_IL] = &Language{Lang_iw_IL, "iw_IL", "Hebrew", "希伯来语(以色列)"}
	languagesCache[Lang_hi_IN] = &Language{Lang_hi_IN, "hi_IN", "Hindi", "印度语"}
	languagesCache[Lang_hu_HU] = &Language{Lang_hu_HU, "hu_HU", "Hungarian", "匈牙利语"}
	languagesCache[Lang_in_ID] = &Language{Lang_in_ID, "in_ID", "Indonesian", "印度尼西亚语"}
	languagesCache[Lang_ja_JP] = &Language{Lang_ja_JP, "ja_JP", "Japanese", "日语"}
	languagesCache[Lang_ko_KR] = &Language{Lang_ko_KR, "ko_KR", "Korean", "韩语（朝鲜语）"}
	languagesCache[Lang_ms_MY] = &Language{Lang_ms_MY, "ms_MY", "Malay", "马来语"}
	languagesCache[Lang_fa_IR] = &Language{Lang_fa_IR, "fa_IR", "Perisan", "波斯语"}
	languagesCache[Lang_pl_PL] = &Language{Lang_pl_PL, "pl_PL", "Polish", "波兰语"}
	languagesCache[Lang_ro_RO] = &Language{Lang_ro_RO, "ro_RO", "Romanian", "罗马尼亚语"}
	languagesCache[Lang_ru_RU] = &Language{Lang_ru_RU, "ru_RU", "Russian", "俄罗斯语"}
	languagesCache[Lang_sr_RS] = &Language{Lang_sr_RS, "sr_RS", "Serbian", "塞尔维亚语"}
	languagesCache[Lang_sv_SE] = &Language{Lang_sv_SE, "sv_SE", "Swedish", "瑞典语"}
	languagesCache[Lang_th_TH] = &Language{Lang_th_TH, "th_TH", "Thai", "泰语"}
	languagesCache[Lang_tr_TR] = &Language{Lang_tr_TR, "tr_TR", "Turkey", "土耳其语"}
	languagesCache[Lang_ur_PK] = &Language{Lang_ur_PK, "ur_PK", "Urdu", "乌尔都语"}
	languagesCache[Lang_vi_VN] = &Language{Lang_vi_VN, "vi_VN", "Vietnamese", "越南语"}
	languagesCache[Lang_ca_ES] = &Language{Lang_ca_ES, "ca_ES", "Catalan", "加泰隆语(西班牙)"}
	languagesCache[Lang_lv_LV] = &Language{Lang_lv_LV, "lv_LV", "Latviesu", "拉脱维亚语"}
	languagesCache[Lang_lt_LT] = &Language{Lang_lt_LT, "lt_LT", "Lithuanian", "立陶宛语"}
	languagesCache[Lang_nb_NO] = &Language{Lang_nb_NO, "nb_NO", "Norwegian", "挪威语"}
	languagesCache[Lang_sk_SK] = &Language{Lang_sk_SK, "sk_SK", "slovencina", "斯洛伐克语"}
	languagesCache[Lang_sl_SI] = &Language{Lang_sl_SI, "sl_SI", "Slovenian", "斯洛文尼亚语"}
	languagesCache[Lang_bg_BG] = &Language{Lang_bg_BG, "bg_BG", "bulgarian", "保加利亚语"}
	languagesCache[Lang_uk_UA] = &Language{Lang_uk_UA, "uk_UA", "Ukrainian", "乌克兰语"}
	languagesCache[Lang_tl_PH] = &Language{Lang_tl_PH, "tl_PH", "Filipino", "菲律宾语"}
	languagesCache[Lang_fi_FI] = &Language{Lang_fi_FI, "fi_FI", "Finnish", "芬兰语"}
	languagesCache[Lang_af_ZA] = &Language{Lang_af_ZA, "af_ZA", "Afrikaans", "南非语"}
	languagesCache[Lang_rm_CH] = &Language{Lang_rm_CH, "rm_CH", "Romansh", "罗曼什语(瑞士)"}
	languagesCache[Lang_my_ZG] = &Language{Lang_my_ZG, "my_ZG", "Burmese(Zawgyi)", "缅甸语"}
	languagesCache[Lang_my_MM] = &Language{Lang_my_MM, "my_MM", "Burmese", "缅甸语"}
	languagesCache[Lang_km_KH] = &Language{Lang_km_KH, "km_KH", "Khmer", "柬埔寨语"}
	languagesCache[Lang_am_ET] = &Language{Lang_am_ET, "am_ET", "Amharic", "阿姆哈拉语(埃塞俄比亚)"}
	languagesCache[Lang_be_BY] = &Language{Lang_be_BY, "be_BY", "Belarusian", "白俄罗斯语"}
	languagesCache[Lang_et_EE] = &Language{Lang_et_EE, "et_EE", "Estonian", "爱沙尼亚语"}
	languagesCache[Lang_sw_TZ] = &Language{Lang_sw_TZ, "sw_TZ", "Swahili", "斯瓦希里语(坦桑尼亚)"}
	languagesCache[Lang_zu_ZA] = &Language{Lang_zu_ZA, "zu_ZA", "Zulu", "祖鲁语(南非)"}
	languagesCache[Lang_az_AZ] = &Language{Lang_az_AZ, "az_AZ", "Azerbaijani", "阿塞拜疆语"}
	languagesCache[Lang_hy_AM] = &Language{Lang_hy_AM, "hy_AM", "Armenian", "亚美尼亚语(亚美尼亚)"}
	languagesCache[Lang_ka_GE] = &Language{Lang_ka_GE, "ka_GE", "Georgian", "格鲁吉亚语(格鲁吉亚)"}
	languagesCache[Lang_lo_LA] = &Language{Lang_lo_LA, "lo_LA", "Laotian", "老挝语(老挝)"}
	languagesCache[Lang_mn_MN] = &Language{Lang_mn_MN, "mn_MN", "Mongolian", "蒙古语"}
	languagesCache[Lang_ne_NP] = &Language{Lang_ne_NP, "ne_NP", "Nepali", "尼泊尔语"}
	languagesCache[Lang_kk_KZ] = &Language{Lang_kk_KZ, "kk_KZ", "Kazakh", "哈萨克语"}
	languagesCache[Lang_si_LK] = &Language{Lang_si_LK, "si_LK", "Sinhala", "僧加罗语(斯里兰卡)"}
}

// GetLanguage get language information by code
func GetLanguage(code Lang) *Language {
	return languagesCache[code]
}

// GetLangCode get language code by key
func GetLangCode(key string) Lang {
	for _, lang := range languagesCache {
		if lang.Key == key {
			return lang.Code
		}
	}
	return InvalidLangCode
}

// IsValidLang check the given language code if valid
func IsValidLang(code Lang) bool {
	return code >= Lang_ar_IL && code <= Lang_si_LK
}

// AppendLangs append a language key to multi-language string
func AppendLangs(langs string, lang Lang) string {
	langkey := GetLanguage(lang).Key
	langsarr := strings.Split(langs, LangsSeparator)
	for _, existlang := range langsarr {
		if existlang == langkey {
			return langs // already exist language
		}
	}
	langsarr = append(langsarr, langkey)
	return strings.Join(langsarr, LangsSeparator)
}

// RemoveLangs remove a language key outof multi-language string
func RemoveLangs(langs string, lang Lang) string {
	langkey := GetLanguage(lang).Key
	langsarr := strings.Split(langs, LangsSeparator)
	for i, existlang := range langsarr {
		if existlang == langkey {
			last := append(langsarr[:i], langsarr[i+1:]...)
			return strings.Join(last, LangsSeparator)
		}
	}
	return langs
}

// IsContain check the language if exist in multi-language string
func IsContain(langs string, lang Lang) bool {
	langkey := GetLanguage(lang).Key
	langsarr := strings.Split(langs, LangsSeparator)
	for _, existlang := range langsarr {
		if existlang == langkey {
			return true
		}
	}
	return false
}
