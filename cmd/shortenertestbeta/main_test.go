package main

//go:generate go test -c -o=../../bin/shortenertest

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestMain(m *testing.M) {
	// Основной тест, запускает все остальные тесты
	os.Exit(m.Run())
}

func TestIteration1(t *testing.T) {
	// Запускает тест-сьют для первой итерации
	suite.Run(t, new(Iteration1Suite))
}

func TestIteration2(t *testing.T) {
	// Запускает тест-сьют для второй итерации
	suite.Run(t, new(Iteration2Suite))
}

func TestIteration3(t *testing.T) {
	// Запускает тест-сьют для третьей итерации
	suite.Run(t, new(Iteration3Suite))
}

func TestIteration4(t *testing.T) {
	// Запускает тест-сьют для четвертой итерации
	suite.Run(t, new(Iteration4Suite))
}

func TestIteration5(t *testing.T) {
	// Запускает тест-сьют для пятой итерации
	suite.Run(t, new(Iteration5Suite))
}

func TestIteration6(t *testing.T) {
	// Запускает тест-сьют для шестой итерации
	suite.Run(t, new(Iteration6Suite))
}

func TestIteration7(t *testing.T) {
	// Запускает тест-сьют для седьмой итерации
	suite.Run(t, new(Iteration7Suite))
}

func TestIteration8(t *testing.T) {
	// Запускает тест-сьют для восьмой итерации
	suite.Run(t, new(Iteration8Suite))
}

func TestIteration9(t *testing.T) {
	// Запускает тест-сьют для восьмой итерации
	suite.Run(t, new(Iteration9Suite))
}
