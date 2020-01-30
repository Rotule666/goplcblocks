package goplcblocks

import (
	"time"
)

// TOF The block's struct
type TOF struct {
	mIN bool
	mQ  bool

	mPT    int64
	mRESET bool

	mET int64

	mM  bool
	mTI bool

	mStartTime  int64
	mActualTime int64
}

// NewTOF The block's constructor
func NewTOF(pt int64) *TOF {
	o := &TOF{false, false, pt, false, 0, false, false, 0, 0}
	return o
}

// Exec The block's logic
func (o *TOF) Exec() {
	o.mActualTime = time.Now().UTC().UnixNano() / 1000 / 1000

	if o.mRESET {
		o.mStartTime = o.mActualTime
		o.mTI = false
		o.mQ = false
		o.mRESET = false
	}

	/* Falling Edge of IN should reset the TIME */
	if !(o.mIN) && o.mM {
		o.mStartTime = o.mActualTime
		o.mTI = true
	}

	/* Memory to detect the Rising Edge of IN */
	o.mM = o.mIN

	/* Count When IN is true (TOF timer) */
	if o.mTI {
		// Should manage overflow here but we need to know G_MAX_MS_COUNTER
		o.mET = o.mActualTime - o.mStartTime
	} else {
		o.mET = 0
	}

	/* Q is on after Timer is done when IN is true */
	if o.mIN || (o.mTI && (o.mET) <= (o.mPT)) {
		o.mQ = true
	} else {
		o.mQ = false
	}
}

// IN The block's input
func (o *TOF) IN(in bool) {
	o.mIN = in
}

// Q The block's output
func (o TOF) Q() bool {
	return o.mQ
}

// RESET reset command
func (o *TOF) RESET() {
	o.mRESET = true
}
