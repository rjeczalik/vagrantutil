// generated by stringer -type=BoxSubcommand,Status -output=stringer.go; DO NOT EDIT

package vagrantutil

import "fmt"

const _BoxSubcommand_name = "AddListOutdatedRemoveRepackageUpdate"

var _BoxSubcommand_index = [...]uint8{0, 3, 7, 15, 21, 30, 36}

func (i BoxSubcommand) String() string {
	if i < 0 || i >= BoxSubcommand(len(_BoxSubcommand_index)-1) {
		return fmt.Sprintf("BoxSubcommand(%d)", i)
	}
	return _BoxSubcommand_name[_BoxSubcommand_index[i]:_BoxSubcommand_index[i+1]]
}

const _Status_name = "UnknownNotCreatedRunning"

var _Status_index = [...]uint8{0, 7, 17, 24}

func (i Status) String() string {
	if i < 0 || i >= Status(len(_Status_index)-1) {
		return fmt.Sprintf("Status(%d)", i)
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}
