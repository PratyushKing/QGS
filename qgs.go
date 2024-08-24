package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Reset string = "\033[0m"; // Reset Color of Console
const Underline string = "\033[4m"; // Underline Console Text

func main() {
	defer func() {
		fmt.Println("cleanup");
	}()
	fmt.Println(retCellStr("Hello!", true, true));
	fmt.Println(strconv.Itoa(getCursorX()) + " " + strconv.Itoa(getCursorY()));
	return;
}

func moveCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col);
}

func getCursorX() int {
	// Request cursor position
	fmt.Print("\033[6n")

	// Read the response from stdin
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('R')

	// Parse the response
	response = strings.TrimPrefix(response, "\033[")
	response = strings.TrimSuffix(response, "R")
	parts := strings.Split(response, ";")

	x, _ := strconv.Atoi(parts[1])
	return x
}

func getCursorY() int {
	// Request cursor position
	fmt.Print("\033[6n")

	// Read the response from stdin
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('R')

	// Parse the response
	response = strings.TrimPrefix(response, "\033[")
	response = strings.TrimSuffix(response, "R")
	parts := strings.Split(response, ";")

	y, _ := strconv.Atoi(parts[0])
	return y
}

func getLenStr(times int) int {
	return (times + 4);
}

func retCellStr(cellText string, isCellAbove bool, isCellLeft bool) string {
	formedStr := " ";
	if (isCellAbove) {
		formedStr += Underline;
	}
	for i := 1; i < getLenStr(len(cellText)) - 1; i++ {
		formedStr += " ";
	}
	formedStr += Reset + "\n";
	if (isCellLeft) {
		formedStr += "|";
	} else { formedStr += " "; }
	formedStr += Underline + " " + cellText + " " + Reset + "|";
	return formedStr;
}
