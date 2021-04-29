package smtp

// Envelopes have their own pool

type EnvelopePool struct {
	// envelopes that are ready to be borrowed
	pool chan *Envelope
	// semaphore to control number of maximum borrowed envelopes
	sem chan bool
}

func NewEnvelopePool(poolSize int) *EnvelopePool {
	return &EnvelopePool{
		pool: make(chan *Envelope, poolSize),
		sem:  make(chan bool, poolSize),
	}
}

func (p *EnvelopePool) Borrow(remoteAddr string, clientID uint64) *Envelope {
	var e *Envelope
	p.sem <- true // block the envelope until more room
	select {
	case e = <-p.pool:
		e.Reseed(remoteAddr, clientID)
	default:
		e = NewEnvelope(remoteAddr, clientID)
	}
	return e
}

// Return returns an envelope back to the envelope pool
// Make sure that envelope finished processing before calling this
func (p *EnvelopePool) Return(e *Envelope) {
	select {
	case p.pool <- e:
		//placed envelope back in pool
	default:
		// pool is full, discard it
	}
	// take a value off the semaphore to make room for more envelopes
	<-p.sem
}
