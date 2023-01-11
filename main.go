package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/samber/lo"

	"github.com/setval/container-backup/internal/dump"
	"github.com/setval/container-backup/pkg/argument"
	"github.com/setval/container-backup/pkg/config"
	"github.com/setval/container-backup/pkg/repository"
)

func pathDate(t time.Time) string {
	return fmt.Sprintf("%d/%d/%d", t.Year(), t.Month(), t.Day())
}

func run() error {
	var configFile string
	var onlyDatabases argument.ArrayFlags
	flag.StringVar(&configFile, "config", "config.yml", "config file")
	flag.Var(&onlyDatabases, "only", "only database")
	flag.Parse()

	cfg, err := config.Parse(configFile)
	if err != nil {
		return err
	}

	if err = os.MkdirAll("data", os.ModePerm); err != nil {
		return err
	}

	var stor repository.Repository
	switch cfg.Storage.Type {
	case config.TypeStorageLocal:
		stor, err = repository.NewLocal(cfg.Storage.LocalConfig.Directory)
	case config.TypeStorageS3:
		stor, err = repository.NewS3(cfg.Storage.S3Config.Endpoint, cfg.Storage.S3Config.AccessKey, cfg.Storage.S3Config.SecretKey, cfg.Storage.S3Config.SSL)
	default:
		err = fmt.Errorf("unknown storage type: %s", cfg.Storage.Type)
	}
	if err != nil {
		return err
	}

	for _, database := range cfg.Databases {
		if len(onlyDatabases) > 0 {
			if _, ok := lo.Find(onlyDatabases, func(i string) bool {
				return i == database.Name
			}); !ok {
				continue
			}
		}
		log.Info().Msgf("dump database %s[%s:%d]", database.Name, database.Host, database.Port)
		dumpFileName := database.Name + ".sql"
		dumpFileNamePath := path.Join("data", dumpFileName)
		if err = dump.Dump(database, dumpFileNamePath); err != nil {
			return fmt.Errorf("dump: %w", err)
		}

		if err = gzip(dumpFileName); err != nil {
			return fmt.Errorf("gzip: %w", err)
		}

		log.Info().Msgf("gzip file %s", dumpFileName)
		dumpGZipFileName := database.Name + ".sql.gz"

		log.Info().Msgf("encrypt file %s", dumpGZipFileName)
		dumpGZipEncryptFile := dumpGZipFileName + ".encrypt"
		if err = encryptFile(dumpGZipFileName, dumpGZipEncryptFile, cfg.Key); err != nil {
			return fmt.Errorf("encrypt file: %w", err)
		}

		file, err := os.Open(path.Join("data", dumpGZipEncryptFile))
		if err != nil {
			return fmt.Errorf("open file to upload: %w", err)
		}
		log.Info().Msgf("upload database %s[%s:%d]", database.Name, database.Host, database.Port)
		if err = stor.Upload(fmt.Sprintf("%s/%s", pathDate(time.Now()), dumpGZipEncryptFile), file); err != nil {
			file.Close()
			return fmt.Errorf("upload to repository: %w", err)
		}
		file.Close()

		for _, deleteFileName := range []string{
			dumpGZipFileName, dumpGZipEncryptFile,
		} {
			log.Info().Msgf("remove file %s", deleteFileName)
			if err = os.Remove(path.Join("data", deleteFileName)); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func gzip(filename string) error {
	cmd := exec.Command("gzip", filename)
	cmd.Dir = "data"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func encryptFile(fileIn, fileOut, key string) error {
	cmd := exec.Command("openssl", "enc", "-aes-256-cbc", "-in", fileIn, "-out", fileOut, "-k", key)
	cmd.Dir = "data"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func decryptFile(fileIn, fileOut, key string) error {
	cmd := exec.Command("openssl", "enc", "-d", "-aes-256-cbc", "-in", fileIn, "-out", fileOut, "-k", key)
	cmd.Dir = "data"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
