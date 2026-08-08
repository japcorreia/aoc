from __future__ import annotations


def solve1(text: str) -> int:
    return 0


def solve2(text: str) -> int:
    return 0


def main():

    with open("./input1", "r") as file:
        text = file.read()
        part1 = solve1(text)
        part2 = solve2(text)

        print(f"Solution 1: {part1}")
        print(f"Solution 2: {part2}")


if __name__ == "__main__":
    main()
