from __future__ import annotations


PAD = (
    (1, 2, 3),
    (4, 5, 6),
    (7, 8, 9),
)

DIRECTIONS = {
    "U": (-1, 0),
    "D": (1, 0),
    "R": (0, 1),
    "L": (0, -1),
}

PAD_2 = (
    ("", "", "1", "", ""),
    ("", "2", "3", "4", ""),
    ("5", "6", "7", "8", "9"),
    ("", "A", "B", "C", ""),
    ("", "", "D", "", ""),
)


def clamp(x: int, a: int, b: int) -> int:
    return max(a, min(x, b))


def solve1(text: str) -> int:
    result = ""
    current_pos = (1, 1)
    for line in text.strip().split("\n"):
        for direction in line:
            row, col = current_pos
            dr, dc = DIRECTIONS[direction]

            row = clamp(row + dr, 0, len(PAD) - 1)
            col = clamp(col + dc, 0, len(PAD[0]) - 1)
            current_pos = row, col

        result += str(PAD[current_pos[0]][current_pos[1]])

    return int(result)


def solve2(text: str) -> str:
    result = ""
    current_pos = (2, 0)
    for line in text.strip().split("\n"):
        for direction in line:
            dr, dc = DIRECTIONS[direction]
            row, col = current_pos

            row = clamp(row + dr, 0, len(PAD_2) - 1)
            col = clamp(col + dc, 0, len(PAD_2[0]) - 1)
            next_pos = row, col

            if PAD_2[next_pos[0]][next_pos[1]] != "":
                current_pos = next_pos

        result += PAD_2[current_pos[0]][current_pos[1]]
    return result


def main():

    with open("./input1", "r") as file:
        text = file.read()
        part1 = solve1(text)
        part2 = solve2(text)

        print(f"Solution 1: {part1}")
        print(f"Solution 2: {part2}")


if __name__ == "__main__":
    main()
