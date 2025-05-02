package engine

import "fmt"

var (
	LabelPrefix = "wp.autoscaler/"
	LabelPool   = fmt.Sprintf("%spool", LabelPrefix)
	LabelImage  = fmt.Sprintf("%simage", LabelPrefix)

	CloudInitUserDataUbuntuDefault = `
#cloud-config

final_message: "The system is finally up, after $UPTIME seconds"
`
)
