package leap


func IsLeapYear(year int) bool {
	/* IsLeapYear returns if an year is leap year or not */
    if year % 4 == 0 {
        if year % 100 == 0 {
            if year % 400 == 0 {
                return true
            }
            return false
        }
        return true
    }
    return false
}
