// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:41:00 D1890E               priveda/fixed64/[in_words_en.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"strings"
)

// InWordsEN returns the currency value as an English description in words
// This function is useful for showing amounts in invoices, etc.
//
// Uses English language names, hence the 'EN' suffix.
//
// fmt: a string specifying the currency format:
//      format_ = "<Large-Single>;<Large-Plural>;<S-Single>;<S-Plural>;<Only>"
//      Large-Single - Large Denomination, Singular. E.g. "Dollar"
//      Large-Plural - Large Denomination, Plural.   E.g. "Dollars"
//      S-Single     - Small Denomination, Singular. E.g. "Cent"
//      S-Plural     - Small Denomination, Plural.   E.g. "Cents"
//
//      All the format specifiers are optional but must follow in
//      the order specified. If a plural specifier is omitted, then
//      an "s" is added to the singular specifier for values greater
//      than 1. If the singular specifier is omitted, Then the
//      plural is used for all values (Used for some currencies).
//      If both specifiers are omitted for either denomination, the
//      denomination will not be returned. See the examples below.
//
// Returns: The amount in words as a string, including the currency
//          and the word "Only". Uses proper capitalisation.
//
// Example: PARAMETER               RESULT
//          (11.02,";;Cent;Only")   "Two Cents Only"
//          (11.02,"Dollar;;Cent")  "Eleven Dollars and Two Cents"
//          (11.02,"Euro")          "Eleven Euros"
//          (11.02,"Pound;;;Pence") "Eleven Pounds and Two Pence"
func (ob Fixed64) InWordsEN(fmt string) string {
	i := ob.i64
	if i < 0 {
		i = -ob.i64
	}
	var (
		bigUnits = i / 1E4
		smlUnits = (i - bigUnits*1E4) / 100
		hasOnly  = strings.HasSuffix(strings.ToLower(fmt), "only")
	)
	if hasOnly {
		fmt = fmt[:len(fmt)-4]
	}
	getPart := func(partNo int) string {
		parts := strings.Split(fmt, ";")
		if partNo < 0 || partNo >= len(parts) {
			return ""
		}
		return parts[partNo]
	}
	var (
		big1 = getPart(0)
		bigN = getPart(1)
		sml1 = getPart(2)
		smlN = getPart(3)
		ret  = ""
	)
	if bigUnits > 0 && (big1+bigN) != "" {
		ret += IntInWordsEN(bigUnits) + " "
		if big1 == "" && bigN != "" {
			ret += bigN
		} else if big1 != "" && bigN == "" {
			ret += big1
			if bigUnits > 1 {
				ret += "s"
			}
		} else if big1 != "" && bigN != "" {
			if bigUnits == 1 {
				ret += big1
			}
			if bigUnits > 1 {
				ret += bigN
			}
		}
	}
	if ((sml1 + smlN) != "") && smlUnits > 0 {
		if (big1+bigN != "") && bigUnits > 0 {
			ret += " and "
		}
		ret += IntInWordsEN(smlUnits) + " "
		if sml1 == "" && smlN != "" {
			ret += smlN
		} else if sml1 != "" && smlN == "" {
			ret += sml1
			if smlUnits > 1 {
				ret += "s"
			}
		} else if sml1 != "" && smlN != "" {
			if smlUnits == 1 {
				ret += sml1
			}
			if smlUnits > 1 {
				ret += smlN
			}
		}
	}
	if hasOnly && len(strings.TrimSpace(ret)) > 0 {
		ret += " Only"
	}
	return ret
}

func IntInWordsEN(i int64) string {
	return "" ////
}

//end
