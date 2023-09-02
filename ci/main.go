package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	var javaFile string

	for i, v := range os.Args {
		if (i == 1) {
			javaFile = v
			break
		}
	}

	if javaFile == "" {
		log.Fatalln("java entrypoint file not specified")
	}


	fmt.Println(javaFile)

	cmd := exec.Command(
		"javac",
		fmt.Sprintf("./src/%s.java", javaFile),
		"-d",
		"./out",
	) 

	err := cmd.Run()	// TODO: HANDLE THIS ERROR
	if err != nil {
		log.Fatalln(err)
	}


	err = syscall.Chdir("out")
	if err != nil {
		log.Fatalln(err)
	}

	cmd = exec.Command(
		"java",
		javaFile,
	)


    stdout, err := cmd.StdoutPipe()
    if err != nil {
        fmt.Println("Error creating stdout pipe:", err)
        return
    }
    
    stderr, err := cmd.StderrPipe()
    if err != nil {
        fmt.Println("Error creating stderr pipe:", err)
        return
    }

    if err := cmd.Start(); err != nil {
        fmt.Println("Error starting command:", err)
        return
    }

	io.Copy(os.Stdout, stdout)
	io.Copy(os.Stderr, stderr)

    if err := cmd.Wait(); err != nil {
        fmt.Println("Command finished with error:", err)
    }
}