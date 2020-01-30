package goplcblocks

// RTRIG ... The block's struct
type RTRIG struct {
	mCLK bool
	mQ   bool
	mM   bool
}

// NewRTRIG The block's constructor
func NewRTRIG() *RTRIG {
	o := &RTRIG{false, false, false}
	return o
}

// Exec The block's logic
func (o *RTRIG) Exec() {
	o.mQ = o.mCLK && !o.mM
	o.mM = o.mCLK
}

// CLK The block's input
func (o *RTRIG) CLK(clk bool) {
	o.mCLK = clk
}

// Q The block's output
func (o RTRIG) Q() bool {
	return o.mQ
}
