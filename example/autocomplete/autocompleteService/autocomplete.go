package autocompleteService

import "context"
import "strings"

//AutocompleteRequest contains request parameter of autocomplete service
type AutocompleteRequest struct {
	Keyword string
}

//AutocompleteSuggestionRespond contains all responds field of autocomplete suggestion
type AutocompleteSuggestionRespond struct {
	Suggestions []string `json:"suggestion"`
	TookTime    int      `json:"took_time"`
}

//AutocompletePopularRespond contains popular keywords
type AutocompletePopularRespond struct {
	Popular  []string `json:"popular"`
	TookTime int      `json:"took_time"`
}

//AutocompleteService give keyword suggestion based on given keyword
type AutocompleteService interface {
	GetSuggestion(context.Context, AutocompleteRequest) (AutocompleteSuggestionRespond, error)
	GetPopular(context.Context) (AutocompletePopularRespond, error)
}

type autocompleteService struct{}

//GetSuggestion returns autocomplete suggestion based on given keyword
func (autocompleteService) GetSuggestion(ctx context.Context, request AutocompleteRequest) (respond AutocompleteSuggestionRespond, err error) {
	if request.Keyword == "" {
		return
	}
	for _, keyword := range dataSet {
		if strings.Contains(keyword, request.Keyword) {
			respond.Suggestions = append(respond.Suggestions, keyword)
		}
	}
	return
}

//GetPopular returns current popular keyword
func (autocompleteService) GetPopular(ctx context.Context) (respond AutocompletePopularRespond, err error) {
	respond.Popular = []string{"baju", "baju bayi"}
	return
}

//New create new autocomplete service instance
func New() AutocompleteService {
	return autocompleteService{}
}

var (
	dataSet = []string{"baju", "baju bayi", "celana", "celana bayi"}
)
