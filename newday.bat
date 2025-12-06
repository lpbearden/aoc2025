@echo off
setlocal enabledelayedexpansion
set day=%1
if "%day:~1%" == "" set day=0%day%

mkdir day%day% 2>nul
cd day%day%

(
echo package main
echo.
echo import ^(
echo 	"fmt"
echo 	"aoc2025/util"
echo ^)
echo.
echo func part1^(input string^) int {
echo 	lines := util.Lines^(input^)
echo 	_ = lines
echo 	return 0
echo }
echo.
echo func part2^(input string^) int {
echo 	lines := util.Lines^(input^)
echo 	_ = lines
echo 	return 0
echo }
echo.
echo func main^(^) {
echo 	input := util.ReadInput^(^)
echo 	fmt.Println^("Part 1:", part1^(input^)^)
echo 	fmt.Println^("Part 2:", part2^(input^)^)
echo }
) > main.go

(
echo package main
echo.
echo import "testing"
echo.
echo const sample = ``
echo.
echo func TestPart1^(t *testing.T^) {
echo 	got := part1^(sample^)
echo 	want := 0
echo 	if got != want {
echo 		t.Errorf^("part1^(^) = %%v, want %%v", got, want^)
echo 	}
echo }
echo.
echo func TestPart2^(t *testing.T^) {
echo 	got := part2^(sample^)
echo 	want := 0
echo 	if got != want {
echo 		t.Errorf^("part2^(^) = %%v, want %%v", got, want^)
echo 	}
echo }
) > main_test.go

type nul > input.txt
cd ..
echo Created day%day%