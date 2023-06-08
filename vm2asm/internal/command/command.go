package command

type Command interface {
	Out() ([]string, error)
}
