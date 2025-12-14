package day10

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Machine struct {
	Lights  uint16
	Buttons [][]byte
	Joltage []uint16
}

type Input []Machine

var lightsPattern *regexp.Regexp = regexp.MustCompile(`\[([.#]+)\]`)
var buttonPattern *regexp.Regexp = regexp.MustCompile(`\(((?:\d+,?)+)\)`)
var joltagePattern *regexp.Regexp = regexp.MustCompile(`\{((?:\d+,?)+)\}`)

func ParseInput(input io.Reader) Input {
	scanner := bufio.NewScanner(input)

	machines := make([]Machine, 0)
	for scanner.Scan() {
		line := scanner.Text()

		lightsRes := lightsPattern.FindStringSubmatch(line)

		var lights uint16 = 0
		for i, light := range strings.Split(lightsRes[1], "") {
			if light == "#" {
				lights |= 1 << i
			}
		}

		buttonsRes := buttonPattern.FindAllStringSubmatch(line, -1)
		buttons := make([][]byte, 0, len(buttonsRes))
		for _, btnGroup := range buttonsRes {
			individualButtons := strings.Split(btnGroup[1], ",")
			buttonGroup := make([]byte, 0, len(individualButtons))
			for _, flag := range individualButtons {
				val, _ := strconv.ParseUint(flag, 10, 8)
				buttonGroup = append(buttonGroup, byte(val))
			}
			buttons = append(buttons, buttonGroup)
		}

		joltageRes := joltagePattern.FindStringSubmatch(line)
		joltage := make([]uint16, 0)
		if len(joltageRes) > 1 {
			jolts := strings.Split(joltageRes[1], ",")
			for _, jolt := range jolts {
				val, _ := strconv.ParseUint(jolt, 10, 16)
				joltage = append(joltage, uint16(val))
			}
		}

		machine := Machine{
			Lights:  lights,
			Buttons: buttons,
			Joltage: joltage,
		}

		machines = append(machines, machine)
	}

	return machines
}
