package rect

type Rect struct {
	width  float64
	lenght float64
}

func CreateRect(w, l float64) Rect {
	return Rect{
		width:  w,
		lenght: l,
	}

}
func (r *Rect) CalcArea() float64 {
	return r.lenght * r.width
}
func (r *Rect) Double() {
	r.lenght = r.lenght * 2
	r.width = r.width * 2
}
