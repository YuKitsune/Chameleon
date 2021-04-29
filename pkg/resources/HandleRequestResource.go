package resources

// Todo: Allow for multiple recipients

type HandleRequestResource struct {
	Recipient string `json:"recipient"`
	RawData   []byte `json:"rawData"`
}
