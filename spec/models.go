package spec

type Message struct {
	IntField int `json:"int_field""`
	StringField string `json:"string_field"`
}

type Choice string

const (
	FirstChoice Choice = "FIRST_CHOICE"
	SecondChoice Choice = "SECOND_CHOICE"
	ThirdChoice Choice = "THIRD_CHOICE"
)