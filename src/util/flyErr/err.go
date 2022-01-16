package flyErr

type Error struct {
	Text string
}

func (err Error) Error() string {
	return err.Text
}
