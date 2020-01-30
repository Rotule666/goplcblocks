package goplcblocks

// SFCSTEP The block's struct
type SFCSTEP struct {
	mX       bool
	mTON     *TON
	mRTRIG   *RTRIG
	mFTRIG   *FTRIG
	mYbefore []*bool
	mYafter  []*bool
}

// NewSFCSTEP The block's constructor
func NewSFCSTEP(i bool, pt int64) *SFCSTEP {
	o := &SFCSTEP{i, NewTON(pt), NewRTRIG(), NewFTRIG(), []*bool{}, []*bool{}}
	return o
}

// Exec The block's logic
func (o *SFCSTEP) Exec() {

	var yb bool = false
	var ya bool = false

	// Or transitions Before
	yb = false
	for _, y := range o.mYbefore {
		yb = yb || *y
	}

	// Or transitions after
	ya = false
	for _, y := range o.mYafter {
		ya = ya || *y
	}

	o.mX = yb || (o.mX && !ya)

	o.mTON.IN(o.mX)
	o.mTON.Exec()

	o.mRTRIG.CLK(o.mX)
	o.mRTRIG.Exec()

	o.mFTRIG.CLK(o.mX)
	o.mFTRIG.Exec()
}

// X The block's output
func (o SFCSTEP) X() bool {
	return o.mX
}

// RTRIG The block's output
func (o SFCSTEP) RTRIG() bool {
	return o.mRTRIG.Q()
}

// FTRIG The block's output
func (o SFCSTEP) FTRIG() bool {
	return o.mFTRIG.Q()
}

// TON The block's output
func (o SFCSTEP) TON() bool {
	return o.mTON.Q()
}

// AddYbefore ...
func (o *SFCSTEP) AddYbefore(y *bool) {
	o.mYbefore = append(o.mYbefore, y)
}

// AddYafter ...
func (o *SFCSTEP) AddYafter(y *bool) {
	o.mYafter = append(o.mYafter, y)
}
