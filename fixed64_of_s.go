// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:42:36 D4B3B0              priveda/fixed64/[fixed64_of_s.go]
// -----------------------------------------------------------------------------

package fixed64

// Fixed64OfS converts a numeric string to a Fixed64.
// If the string is not numeric, logs an error and sets the Fixed64 to zero.
func Fixed64OfS(s string) Fixed64 {
	var (
		minus bool
		fract bool
		dp    int
		ret   Fixed64
	)
	for _, r := range s {
		if r == '-' {
			minus = true
		} else if r == '.' {
			fract = true
		} else if r >= '0' && r <= '9' {
			if fract {
				dp++
				if dp > 4 {
					break
				}
			}
			ret.i64 *= 10
			ret.i64 += int64(r - '0')
		} else if r != ',' &&
			r != ' ' && r != '\a' && r != '\b' && r != '\f' &&
			r != '\n' && r != '\r' && r != '\t' && r != '\v' {
			mod.Error("Non-numeric string:^", s)
			break
		}
	}
	for dp < 4 {
		dp++
		ret.i64 *= 10
	}
	if minus {
		ret.i64 = -ret.i64
	}
	return ret
}

//end
