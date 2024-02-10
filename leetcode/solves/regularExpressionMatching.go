package solves

func IsRegularExpressionMatching(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}

	if len(s) != 0 && len(p) == 0 {
		return false
	}

	if len(p) > 1 && p[1] == '*' {
		if len(s) != 0 && (s[0] == p[0] || p[0] == '.') {
			return IsRegularExpressionMatching(s[1:], p) || IsRegularExpressionMatching(s, p[2:])
		} else {
			return IsRegularExpressionMatching(s, p[2:])
		}
	} else {
		if len(s) != 0 && (s[0] == p[0] || p[0] == '.') {
			return IsRegularExpressionMatching(s[1:], p[1:])
		} else {
			return false
		}
	}
}
