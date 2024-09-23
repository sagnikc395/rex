package main

type Regex struct {
	subpatterns []string
	capturing   bool
}

func NewRegex() *Regex {
	return &Regex{
		subpatterns: make([]string, 0),
		capturing:   true,
	}
}

func (r *Regex) isCapturingGroup() bool {
	return r.capturing
}

func (r *Regex) nonCapturing() *Regex {
	r.capturing = false
	return r
}
