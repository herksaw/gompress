package format

import (
	"archive/zip"
	"gompress/helper"
	"io"
	"os"
)

func CreateZip(filename string, files []string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	zipper := zip.NewWriter(file)

	defer zipper.Close()

	for _, name := range files {
		if err := writeFileToZip(zipper, name); err != nil {
			return err
		}
	}
	return nil
}

func writeFileToZip(zipper *zip.Writer, filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	info, err := file.Stat()

	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)

	if err != nil {
		return err
	}

	header.Name = helper.SanitizedName(filename)

	writer, err := zipper.CreateHeader(header)

	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)

	return err
}
