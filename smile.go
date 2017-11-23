package smile

import "time"

// Smile smiles... duh...
func Smile() string {
	return ":)"
}

// LongRunningSmile sleeps for a period and
// then smiles
func LongRunningSmile(sleep string) string {
	dur, err := time.ParseDuration(sleep)
	if err != nil {
		return err.Error()
	}
	time.Sleep(dur)
	return ":)"
}
