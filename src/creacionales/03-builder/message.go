package builder

// Message = Producto concreto
type Message struct {
	Recipient string `json:"recipient" xml:"recipient"`
	Text      string `json:"text" xml:"text"`
}

// MessageRepresented = Representación del objeto
type MessageRepresented struct {
	Body   []byte
	Format string
}
