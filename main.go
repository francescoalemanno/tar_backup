package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func handleErr(x error) {
	if x != nil {
		log.Fatal(x)
	}
}
func main() {
	folder_p := flag.String("f", "", "cartella da backuppare")
	out_folder_p := flag.String("o", "", "cartella in cui effettuare backups")
	delta_p := flag.Uint64("t", 2, "minuti da attendere per ogni backup, il primo è immediato.")
	flag.Parse()
	if *folder_p == "" || *out_folder_p == "" || *delta_p == 0 {
		flag.CommandLine.Usage()
		os.Exit(0)
	}
	epoch := (uint64(time.Now().Unix()) / (60 * (*delta_p)))
	fmt.Println(epoch)
	out := filepath.Join(*out_folder_p, fmt.Sprintf("backup-%020d.tar.gz", epoch))

	if _, err := os.Stat(out); err == nil {
		fmt.Println("backup pre-esistente.")
		os.Exit(0)
	}

	cmd := exec.Command("tar", "-czvpf", out, *folder_p)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	handleErr(cmd.Start())
	handleErr(cmd.Wait())
}
