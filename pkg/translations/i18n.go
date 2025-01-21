package translations

import (
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Translation struct {
	bundle    i18n.Bundle
	localizer i18n.Localizer
}

func NewTranslation() ITranslation {
	translation := Translation{}

	bundle := i18n.NewBundle(language.Persian)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("translations/fa.json")
	bundle.LoadMessageFile("translations/en.json")

	localizer := i18n.NewLocalizer(bundle, language.Persian.String(), language.English.String())

	translation.localizer = *localizer
	translation.bundle = *bundle

	return translation
}

func (translation Translation) Message(key string) string {
	loadedMessage, _ := translation.localizer.Localize(
		&i18n.LocalizeConfig{
			MessageID: key,
		},
	)
	return loadedMessage
}

func (translation Translation) MessageWithArgs(key string, data any) string {
	loadedMessage, _ := translation.localizer.Localize(
		&i18n.LocalizeConfig{
			MessageID:    key,
			TemplateData: data,
		},
	)
	return loadedMessage
}
