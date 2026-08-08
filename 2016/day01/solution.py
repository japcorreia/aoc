from __future__ import annotations


DIRECTIONS = (
    (0, 1),  # North
    (1, 0),  # East
    (0, -1),  # South
    (-1, 0),  # West
)


def solve1(text: str) -> int:
    current_pos = (0, 0)
    direction = 0

    for instruction in text.strip().split(", "):
        turn = instruction[0]
        distance = int(instruction[1:])

        if turn == "R":
            direction = (direction + 1) % len(DIRECTIONS)
        elif turn == "L":
            direction = (direction - 1) % len(DIRECTIONS)
        else:
            raise ValueError(f"Unknown turn: {turn!r}")

        dx, dy = DIRECTIONS[direction]
        current_pos = (
            current_pos[0] + (dx * distance),
            current_pos[1] + (dy * distance),
        )

    return abs(current_pos[0]) + abs(current_pos[1])


def solve2(text: str) -> int:
    current_pos = (0, 0)
    direction = 0
    visited = set()
    visited.add(current_pos)
    for instruction in text.strip().split(", "):
        turn = instruction[0]
        distance = int(instruction[1:])

        if turn == "R":
            direction = (direction + 1) % len(DIRECTIONS)
        elif turn == "L":
            direction = (direction - 1) % len(DIRECTIONS)
        else:
            raise ValueError(f"Unknown turn: {turn!r}")

        dx, dy = DIRECTIONS[direction]
        for _ in range(distance):
            current_pos = (
                current_pos[0] + dx,
                current_pos[1] + dy,
            )
            if current_pos in visited:
                return abs(current_pos[0]) + abs(current_pos[1])

            visited.add(current_pos)

    raise ValueError("No location was visited twice!")


def main():

    with open("./input1", "r") as file:
        text = file.read()
        part1 = solve1(text)
        part2 = solve2(text)

        print(f"Solution 1: {part1}")
        print(f"Solution 2: {part2}")


if __name__ == "__main__":
    main()
