from __future__ import annotations


def is_triangle(sides: tuple[int, int, int]) -> bool:
    (
        a,
        b,
        c,
    ) = sorted(sides)
    return a + b > c


def parse(text: str) -> list[tuple[int, int, int]]:
    triangles = []

    for line in text.splitlines():
        sides = tuple(map(int, line.split()))

        if len(sides) != 3:
            raise ValueError(f"Expected 3 sides, got {len(sides)}")
        triangles.append(sides)

    return triangles


def solve1(text: str) -> int:
    triangles = parse(text)
    return sum(is_triangle(t) for t in triangles)


def solve2(text: str) -> int:
    rows = parse(text)

    if len(rows) % 3 != 0:
        raise ValueError("Number of rows must be a multiple of 3")

    result = 0

    for i in range(0, len(rows), 3):
        group = rows[i : i + 3]

        for triangle in zip(*group):
            result += is_triangle(triangle)

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
