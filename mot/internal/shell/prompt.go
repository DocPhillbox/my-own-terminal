package shell

import "fmt"

func (s *Shell) prompt() string {
	user := currentUser()
	host := hostname()
	cwd := currentDir()

	statusColor := colorGreen
	if s.lastStatus != 0 {
		statusColor = colorRed
	}

	return fmt.Sprintf(
		"%s%s@%s%s:%s%s%s %s> %s",
		colorCyan,
		user,
		host,
		colorReset,
		colorBlue,
		cwd,
		colorReset,
		statusColor,
		colorReset,
	)
}
