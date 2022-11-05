package libraryfine

func LibraryFine(d1, m1, y1, d2, m2, y2 int32) int32 {
	// besaran denda
	const (
		lateDay   int32 = 15
		lateMonth int32 = 500
		lateYear  int32 = 10000
	)

	if y1 < y2 {
		return 0
	}

	// cek keterlambatan
	if y1 > y2 {
		return lateYear
	} else if m1 > m2 && y1 == y2 {
		late := m1 - m2
		return late * lateMonth
	} else if d1 > d2 && m1 == m2 {
		late := d1 - d2
		return late * lateDay
	} else {
		return 0
	}
}
