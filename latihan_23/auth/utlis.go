package auth

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

func findOffset(username string) (int, error) {

	fileIndex, err := os.OpenFile(pathindex, os.O_RDWR, 0600)
	if err != nil {
		return 0, errors.New("Cannot reading file")
	}

	defer fileIndex.Close()

	idx := new(Index)

	scanner := bufio.NewScanner(fileIndex)
	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())
		if strings.Contains(string(line), username) {
			err := json.Unmarshal(line, idx)
			if err != nil {
				return 0, err
			}

		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return idx.Line, err
}

func (acc *Account) makeIndex() error {
	fileIndex, err := os.OpenFile(pathindex, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600) //
	if err != nil {
		return fmt.Errorf("Gagal membuka file : %w", err)
	}

	defer fileIndex.Close()

	idx := &Index{
		Username: acc.Username,
		Line:     lastLine(fileIndex),
	}

	dataIdx, err := json.Marshal(idx)
	if err != nil {
		return fmt.Errorf("Gagal konversi ke json : %w", err)
	}

	dataIdx = append(dataIdx, '\n')

	_, err = fileIndex.Write(dataIdx)

	if err != nil {
		return fmt.Errorf("Gagal menulis Index: %w", err)
	}

	return err

}

func (acc *Account) FindAccount(username string) (res *Account, err error) {
	acc.Mut.RLock()
	defer acc.Mut.RUnlock()

	file, err := os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		return nil, errors.New("Cannot reading file")
	}

	defer file.Close()
	offset, err := findOffset(username)
	if err != nil {
		return nil, err
	}

	_, err = file.Seek(int64(offset), 0)

	reader := bufio.NewReader(file)
	line, err := reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(line, acc)
	if err != nil {
		return nil, err
	}

	return res, err

}

// Fix Later return values
func lastLine(file *os.File) (result int) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result++
	}
	if err := scanner.Err(); err != nil {
		return 0
	}
	return result
}
