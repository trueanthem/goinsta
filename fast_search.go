package goinsta

func (search *Search) FastSearchUser(query string) (*SearchResult, error) {
	return search.user(query)
}

func (search *Search) FastSearchHashtag(query string) (*SearchResult, error) {
	return search.tags(query)
}
