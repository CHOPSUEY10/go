package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
)

var path string = "./session.jsonl"

type Session struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

// FIX IT LATER AFTER ReadSession function
func checkCookies(ctx context.Context, session *Session, token string) error {
	savedSessions, err := ReadSession()
	if err != nil {
		return err
	}

	if savedSessions[token].Id == ctx.Value(session).(string) {
		return nil
	} else {
		return errors.New("invalid token")
	}
}

func InitSession(id string) *Session {
	session := &Session{
		Id:    id,
		Token: uuid.NewString(),
	}
	return session
}

func (sess *Session) WriteSession() error {

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return fmt.Errorf("Gagal membuka file : %w", err)
	}

	data, err := json.Marshal(sess)
	if err != nil {
		return fmt.Errorf("Gagal konversi ke json : %w", err)
	}

	data = append(data, '\n')

	_, err = file.Write(data)

	if err != nil {
		return fmt.Errorf("Gagal menulis ke JSON: %w", err)
	}

	return err
}

func ReadSession() (map[string]Session, error) {

	file, err := os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		return nil, errors.New("Cannot reading file")
	}

	defer file.Close()

	var jsonSlice = make(map[string]Session)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		var jsonLine Session

		line := bytes.TrimSpace(scanner.Bytes())

		// JIKA BARIS KOSONG, LEWATI (JANGAN DI-UNMARSHAL)
		if len(line) == 0 {
			continue
		}

		err := json.Unmarshal(line, &jsonLine)
		if err != nil {
			return nil, err
		}

		// var jsonLineCopy = jsonLine
		jsonSlice[jsonLine.Token] = jsonLine

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return jsonSlice, nil

}
