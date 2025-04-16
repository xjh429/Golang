type Calendar struct {
	start,end   int
}
type MyCalendar []Calendar

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	for _, p := range *this {
		if p.start < endTime && p.end > startTime {
			return false
		}
	}
	*this = append(*this, Calendar{startTime, endTime})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
