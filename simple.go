package simple

type Person interface {
	DoThings(string, int) (int, error)
	DoNothing()
}
