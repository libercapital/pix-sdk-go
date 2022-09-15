package services

type Violations struct {
	Reason      string      `json:"razao"`
	Property    string      `json:"propriedade"`
}

type PixError struct {
	Type        string      `json:"type,omitempty"`
	Title       string      `json:"title,omitempty"`
	Status      int8        `json:"status,omitempty"`
	Detail      string      `json:"detail,omitempty"`
}

type Error struct {
	Name        string      `json:"nome,omitempty"`
	Message     string      `json:"mensagem,omitempty"`
	Errors      []*PixError `json:"errors,omitempty"`
}


func (p *Error) Error() string {
	return p.Name + p.Message
}