from solution import solve1, solve2


def test_solve1_1():
    text_input = "R2, L3"
    output = solve1(text_input)
    print(output)
    assert output == 5


def test_solve1_2():
    text_input = "R2, R2, R2"
    output = solve1(text_input)
    print(output)
    assert output == 2


def test_solve1_3():
    text_input = "R5, L5, R5, R3"
    output = solve1(text_input)
    print(output)
    assert output == 12


def test_solve2_1():
    text_input = "R8, R4, R4, R8"
    output = solve2(text_input)
    print(output)
    assert output == 4
