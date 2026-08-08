from solution import solve1, solve2


def test_solve1_1():
    text_input = "5  10  25"
    output = solve1(text_input)
    print(output)
    assert output == 0


def test_solve2_1():
    text_input = (
        "101 301 501\n102 302 502\n103 303 503\n201 401 601\n202 402 602\n203 403 603"
    )
    output = solve2(text_input)
    print(output)
    assert output == 6
