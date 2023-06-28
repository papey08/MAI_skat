package chich_err

type Bookmark struct {
	Offset int
	LineNo int
	pos    int
	Line   string
}

func NewBookmark(fileCnt string, filePos, offset int) Bookmark {
	bline, blineNo, relative := line(fileCnt, filePos)

	return Bookmark{
		Offset: offset,
		Line:   bline,
		LineNo: blineNo,
		pos:    relative,
	}
}
