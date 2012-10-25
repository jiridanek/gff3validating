package services

import (
	"bytes"
	"io/ioutil"
	"os/exec"
)

import (
	"../view"
)

var (
	GT = "/home/jirka/Downloads/gt-1.4.2-Linux_x86_64-64bit/bin/gt"
)

func Validate(fname string) (view.Results, error) {
	var results view.Results
	valid, stderr, err := RunValidator(fname)
	if err != nil {
		return results, err
	}
	if valid {
		results.Result = "The GFF3 file is valid"
	} else {
		results.Result = "The GFF3 file is invalid"
	}
	errors, warnings := parseStderr(stderr)
	results.NError = len(errors)
	upto := 10
	if (results.NError < 10) {
		upto = results.NError
	}
	results.Errors = errors[0:upto]
	results.NWarning = len(warnings)
	results.Warnings = warnings
	
	return results, nil
}

func RunValidator(fname string) (bool, *bytes.Buffer, error) {
	
// 	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	cmd := exec.Command(GT, "gff3validator", fname)
// 	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Start()
	if err != nil {
		return false, nil, err
	}
	err = cmd.Wait()
	if err != nil {
		if e, ok := err.(*exec.ExitError); ok {
			if e.ProcessState.Exited() && !e.ProcessState.Success() {
				return false, stderr, nil
			}
		}
		return false, stderr, err
	}
	return true, stderr, nil
}

func RunTidy(fname string) (string, error) {
	GT := "/home/jirka/Downloads/gt-1.4.2-Linux_x86_64-64bit/bin/gt"
	r, err := ioutil.TempFile("", "gfftidy")
	rname := r.Name()
	r.Close()
	cmd := exec.Command(GT, "gff3", "-o " + rname, fname)
	err = cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	if err != nil {
		return "", err
	}
	return rname, nil
}