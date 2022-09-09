package items

type Item struct {
	Id                string      `json:"id"`
	Seller            int64       `json:"seller"`
	Title             string      `json:"title"`
	Description       Description `json:"description"`
	Pictures          []Picture   `json:"pictures"`
	Video             string      `json:"video"`
	Price             int64       `json:"price"`
	AvailableQuantity int64       `json:"available_quantity"`
	SoldQuantity      int64       `json:"sold_quantity"`
	Status            string      `json:"status"`
}

type Description struct {
	PlainText string `json:"plaintext"`
	Html      string `json:"html"`
}
type Picture struct {
	Id  int64  `json:"pid"`
	Url string `json:"url"`
}
