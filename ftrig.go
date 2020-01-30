package goplcblocks

// FTRIG ... The block's struct
type FTRIG struct {
	mCLK bool
	mQ   bool
	mM   bool
}

// NewFTRIG The block's constructor
func NewFTRIG() *FTRIG {
	o := &FTRIG{false, false, false}
	return o
}

// Exec The block's logic
func (o *FTRIG) Exec() {
	o.mQ = !o.mCLK && o.mM
	o.mM = o.mCLK
}

// CLK The block's input
func (o *FTRIG) CLK(clk bool) {
	o.mCLK = clk
}

// Q The block's output
func (o FTRIG) Q() bool {
	return o.mQ
}
