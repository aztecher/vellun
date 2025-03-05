package cmd

import (
	"github.com/containernetworking/cni/pkg/skel"
)

// type Option func(cmd *Cmd)

// func NewCmd(opts ...Option) *Cmd {
// }

type Cmd struct{}

// NewCmd creates a new Cmd instance with Add, Del and Check methods
func NewCmd() *Cmd {
	return &Cmd{}
}

// CNIFuncs returns the CNI functions
func (cmd *Cmd) CNIFuncs() skel.CNIFuncs {
	return skel.CNIFuncs{
		Add:    cmd.Add,
		Del:    cmd.Del,
		Check:  cmd.Check,
		Status: cmd.Status,
	}
}

func (cmd *Cmd) Add(args *skel.CmdArgs) (err error) {
	// cniTypes.PrintResult()
	return err
}

func (cmd *Cmd) Del(args *skel.CmdArgs) (err error) {
	return nil
}

func (cmd *Cmd) Check(args *skel.CmdArgs) (err error) {
	return nil
}

func (cmd *Cmd) Status(args *skel.CmdArgs) (err error) {
	return nil
}
