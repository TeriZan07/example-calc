package front

type Bind map[string]any

func CommonBind() Bind {
	return Bind{
		"title": "Example calc app",
	}
}

func ErrorBind() Bind {
	b := CommonBind()
	b["error_message"] = "There was an unexpected error loading the page"

	return b
}
