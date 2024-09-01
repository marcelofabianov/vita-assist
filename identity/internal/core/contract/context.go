package contract

type Context string

const (
	MainContext Context = "main"
	CoreContext Context = "core"
	AuthContext Context = "auth"
	UserContext Context = "user"
)

func (c Context) String() string {
	return string(c)
}
