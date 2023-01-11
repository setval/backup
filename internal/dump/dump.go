package dump

import (
	"io"
	"os"
	"os/exec"
	"strconv"

	"github.com/setval/container-backup/pkg/config"
)

// function run cmd and wait result
func runCmd(cmd *exec.Cmd, out io.Writer) error {
	cmd.Stdout = out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func createDumpCMD(cfg config.Database) *exec.Cmd {
	args := []string{
		cfg.Name, "--host", cfg.Host, "-u", cfg.Login, "--port", strconv.Itoa(cfg.Port),
		"--column-statistics=0",
	}
	args = append(args, cfg.Tables...)
	// TODO: select dump program
	cmd := exec.Command("mysqldump", args...)
	cmd.Env = append(cmd.Env, "MYSQL_PWD="+cfg.Password)
	return cmd
}

func Dump(cfg config.Database, out string) error {
	outFile, err := os.Create(out)
	if err != nil {
		return err
	}
	defer outFile.Close()

	cmd := createDumpCMD(cfg)

	return runCmd(
		cmd,
		outFile,
	)
}
