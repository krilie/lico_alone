package ginutil

func (g *GinWrap) PathParamOrDefault(key, def string) string {
	var val = g.Context.Param(key)
	if val == "" {
		return def
	} else {
		return val
	}
}

func (g *GinWrap) QueryParamOrDefault(key, def string) string {
	val, ok := g.Context.GetQuery(key)
	if !ok {
		return def
	}
	if val == "" {
		return def
	}
	return val
}

func (g *GinWrap) FormParamOrDefault(key, def string) string {
	val, ok := g.Context.GetPostForm(key)
	if !ok {
		return def
	}
	if val == "" {
		return def
	}
	return val
}
