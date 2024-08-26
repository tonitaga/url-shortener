package server

type JsonBody struct {
	Alias       string `json:"alias"`
	OriginalUrl string `json:"original_url"`
}
// обычно все стурктурки складываем в папке моделей, что бы в одном месте смотерть
// но смотря по какой архетекруте делаешь, можем созвониться покажу приемры
// можно и тут оставить, совет назваит файл model
