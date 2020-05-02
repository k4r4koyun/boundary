package iam

// GetOpts - iterate the inbound Options and return a struct
func GetOpts(opt ...Option) Options {
	opts := getDefaultOptions()
	for _, o := range opt {
		o(&opts)
	}
	return opts
}

// Option - how Options are passed as arguments
type Option func(*Options)

// Options = how options are represented
type Options struct {
	withPublicId     string
	withFriendlyName string
	withScope        *Scope
	withDescription  string
	withGroupGrants  bool
}

func getDefaultOptions() Options {
	return Options{
		withPublicId:     "",
		withFriendlyName: "",
		withScope:        nil,
		withDescription:  "",
		withGroupGrants:  false,
	}
}

// WitPublicId provides an optional public id
func WitPublicId(id string) Option {
	return func(o *Options) {
		o.withPublicId = id
	}
}

// WithDescription provides an optional description
func WithDescription(desc string) Option {
	return func(o *Options) {
		o.withDescription = desc
	}
}

// WithScope provides an optional scope
func WithScope(s *Scope) Option {
	return func(o *Options) {
		o.withScope = s
	}
}

// WithFriendlyName provides an option to search by a friendly name
func WithFriendlyName(name string) Option {
	return func(o *Options) {
		o.withFriendlyName = name
	}
}

// WithGroupGrants provides an option to include group grants
func WithGroupGrants(include bool) Option {
	return func(o *Options) {
		o.withGroupGrants = include
	}
}