package day10

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Machine struct {
	Lights  []int
	Buttons [][]int
	Joltage []int
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

		lights := make([]int, 0, len(lightsRes[1]))
		for _, light := range lightsRes[1] {
			if light == '#' {
				lights = append(lights, 1)
			} else {
				lights = append(lights, 0)
			}
		}

		buttonsRes := buttonPattern.FindAllStringSubmatch(line, -1)
		buttons := make([][]int, 0, len(buttonsRes))
		for _, btnGroup := range buttonsRes {
			individualButtons := strings.Split(btnGroup[1], ",")
			buttonGroup := make([]int, 0, len(individualButtons))
			for _, flag := range individualButtons {
				val, _ := strconv.Atoi(flag)
				buttonGroup = append(buttonGroup, val)
			}
			buttons = append(buttons, buttonGroup)
		}

		joltageRes := joltagePattern.FindStringSubmatch(line)
		joltage := make([]int, 0)
		if len(joltageRes) > 1 {
			jolts := strings.Split(joltageRes[1], ",")
			for _, jolt := range jolts {
				val, _ := strconv.Atoi(jolt)
				joltage = append(joltage, val)
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
