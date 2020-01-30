package goplcblocks

import (
	"time"
)

// TON The block's struct
type TON struct {
	mIN bool
	mQ  bool

	mPT    int64
	mRESET bool

	mET int64

	mM bool

	mStartTime  int64
	mActualTime int64
}

// NewTON The block's constructor
func NewTON(pt int64) *TON {
	o := &TON{false, false, pt, false, 0, false, 0, 0}
	return o
}

// Exec The block's logic
func (o *TON) Exec() {
	o.mActualTime = time.Now().UTC().UnixNano() / 1000 / 1000

	if o.mRESET && o.mIN {
		o.mStartTime = o.mActualTime
		o.mRESET = false
	}

	/* Rising Edge of IN should reset the TIME */
	if o.mIN && !(o.mM) {
		o.mStartTime = o.mActualTime
	}

	/* Memory to detect the Rising Edge of IN */
	o.mM = o.mIN

	/* Count When IN is true (TON timer) */
	if o.mIN {
		// Should manage overflow here but we need to know G_MAX_MS_COUNTER
		o.mET = o.mActualTime - o.mStartTime
	} else {
		o.mET = 0
	}

	/* Q is on after Timer is done when IN is true */
	if o.mIN && (((o.mET) >= (o.mPT)) || o.mQ) {
		o.mQ = true
	} else {
		o.mQ = false
	}
}

// IN The block's input
func (o *TON) IN(in bool) {
	o.mIN = in
}

// Q The block's output
func (o TON) Q() bool {
	return o.mQ
}

// RESET reset command
func (o *TON) RESET() {
	o.mRESET = true
}
