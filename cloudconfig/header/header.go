package header

type Header struct {
	content string
}

func New() Header {
	h := new(Header)
	h.content = "#cloud-config"

	return *h
}

func (h Header) Content() string {
	return h.content
}
