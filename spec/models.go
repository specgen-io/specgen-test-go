package spec

type Message struct {
	IntField int `json:"int_field"`
	StringField string `json:"string_field"`
}

type Choice string

const (
	ChoiceFirst Choice = "FIRST_CHOICE"
	ChoiceSecond Choice = "SECOND_CHOICE"
	ChoiceThird Choice = "THIRD_CHOICE"
)

var ChoiceValuesStrings = []string{string(ChoiceFirst), string(ChoiceSecond), string(ChoiceThird)}
var ChoiceValues = []Choice{ChoiceFirst, ChoiceSecond, ChoiceThird}

func ChoiceNullable(s *string) *Choice {
	if s != nil {
		v := Choice(*s)
		return &v
	} else {
		return nil
	}
}

func ChoiceArray(s []string) []Choice {
	values := []Choice{}
	for _, s := range s {
		values = append(values, Choice(s))
	}
	return values
}