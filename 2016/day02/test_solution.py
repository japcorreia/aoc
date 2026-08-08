from solution import solve1, solve2


def test_solve1_1():
    text_input = "ULL\nRRDDD\nLURDL\nUUUUD"
    output = solve1(text_input)
    print(output)
    assert output == 1985


def test_solve2_1():
    text_input = "ULL\nRRDDD\nLURDL\nUUUUD"
    output = solve2(text_input)
    print(output)
    assert output == "5DB3"
