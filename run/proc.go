package urun

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

type CmdTry struct {
	Args []string
	Ran  *bool
}

func SetupJsonProtoPipes(bufferCapacity int, withContentLen bool, needJsonOut bool) (stdin *bufio.Scanner, rawOut *bufio.Writer, jsonOut *json.Encoder) {
	stdin = bufio.NewScanner(os.Stdin)
	stdin.Buffer(make([]byte, bufferCapacity), bufferCapacity)
	if withContentLen {
		stdin.Split(func(data []byte, ateof bool) (advance int, token []byte, err error) {
			if i_cl1 := bytes.Index(data, []byte("Content-Length: ")); i_cl1 >= 0 {
				datafromclen := data[i_cl1+16:]
				if i_cl2 := bytes.IndexAny(datafromclen, "\r\n"); i_cl2 > 0 {
					if clen, e := strconv.Atoi(string(datafromclen[:i_cl2])); e != nil {
						err = e
					} else {
						if i_js1 := bytes.Index(datafromclen, []byte("{\"")); i_js1 > i_cl2 {
							if i_js2 := i_js1 + clen; len(datafromclen) >= i_js2 {
								advance = i_cl1 + 16 + i_js2
								token = datafromclen[i_js1:i_js2]
							}
						}
					}
				}
			}
			return
		})
	}
	rawOut = bufio.NewWriterSize(os.Stdout, bufferCapacity)
	if needJsonOut {
		jsonOut = json.NewEncoder(rawOut)
		jsonOut.SetEscapeHTML(false)
		jsonOut.SetIndent("", "")
	}
	return
}

func CmdTryStart(cmdname string, cmdargs ...string) (err error) {
	cmd := exec.Command(cmdname, cmdargs...)
	err = cmd.Start()
	defer cmd.Wait()
	if cmd.Process != nil {
		cmd.Process.Kill()
	}
	return
}

func CmdsTryStart(cmds map[string]*CmdTry) {
	var w sync.WaitGroup
	run := func(cmd string, try *CmdTry) {
		defer w.Done()
		*try.Ran = nil == CmdTryStart(cmd, try.Args...)
	}
	for cmdname, cmdmore := range cmds {
		w.Add(1)
		go run(cmdname, cmdmore)
	}
	w.Wait()
}

func CmdExecStdin(stdin string, dir string, cmdname string, cmdargs ...string) (stdout string, stderr string, err error) {
	if len(cmdname) > 0 && strings.Contains(cmdname, " ") && len(cmdargs) == 0 {
		cmdargs = strings.Split(cmdname, " ")
		cmdname = cmdargs[0]
		cmdargs = cmdargs[1:]
	}
	cmd := exec.Command(cmdname, cmdargs...)
	cmd.Dir = dir
	if len(stdin) > 0 {
		cmd.Stdin = strings.NewReader(stdin)
	}
	var bufout, buferr bytes.Buffer
	cmd.Stdout = &bufout
	cmd.Stderr = &buferr
	if err = cmd.Run(); err != nil {
		if _, isexiterr := err.(*exec.ExitError); isexiterr || strings.Contains(err.Error(), "pipe has been ended") || strings.Contains(err.Error(), "pipe has been closed") {
			err = nil
		}
	}
	stdout = bufout.String()
	stderr = strings.TrimSpace(buferr.String())
	return
}

func CmdExecIn(dir string, cmdname string, cmdargs ...string) (out string, err error) {
	var output []byte
	cmd := exec.Command(cmdname, cmdargs...)
	cmd.Dir = dir
	output, err = cmd.CombinedOutput()      // wish Output() would suffice, but sadly some tools abuse stderr for all sorts of non-error 'metainfotainment' (hi godoc & gofmt!)
	out = strings.TrimSpace(string(output)) // do this regardless of err, because it might well be benign such as "exitcode 2", in which case output is still wanted
	return
}

func CmdExecInOr(def string, dir string, cmdname string, cmdargs ...string) string {
	out, err := CmdExecIn(dir, cmdname, cmdargs...)
	if err != nil {
		return def
	}
	return out
}

func CmdExec(cmdname string, cmdargs ...string) (string, error) {
	return CmdExecIn("", cmdname, cmdargs...)
}

func CmdExecOr(def string, cmdname string, cmdargs ...string) string {
	return CmdExecInOr(def, "", cmdname, cmdargs...)
}
