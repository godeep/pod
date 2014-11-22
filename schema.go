package pod

// Schema is use to create some predefine middleware stack.
type Schema []MiddleWare

// NewScheme generate a New Schema with the provided Middleware.
func NewSchema(m ...interface{}) *Schema {
	if len(m) > 0 {
		sche := &Schema{}
		stack := toMiddleware(m)
		for _, s := range stack {
			*sche = append(*sche, s)
		}
		return sche
	}
	return nil
}
