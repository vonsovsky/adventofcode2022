def dec_to_snafu(num: int):
    pre_snafu = ""
    while num > 0:
        rest = num % 5
        num //= 5
        pre_snafu += str(rest)

    num_builder = ""
    add = 0
    for i in range(len(pre_snafu)):
        curr_num = int(pre_snafu[i]) + add
        if curr_num <= 2:
            num_builder += str(curr_num)
            add = 0
            continue

        add = 1
        if curr_num == 3:
            num_builder += "="
        if curr_num == 4:
            num_builder += "-"
        if curr_num == 5:
            num_builder += "0"

    if add == 1:
        num_builder += "1"

    return num_builder[::-1]


def snafu_to_dec(num: str):
    dec = 0
    rev_num = reversed(num)
    for i, char in enumerate(rev_num):
        exp = 5 ** i
        if char == '2':
            dec += 2 * exp
        if char == '1':
            dec += 1 * exp
        if char == '-':
            dec += -1 * exp
        if char == '=':
            dec += -2 * exp
    return dec


def read_input(filename):
    snafu = []
    with open(filename, 'r') as fr:
        for line in fr.readlines():
            if line[-1] == "\n":
                line = line[:-1]

            snafu.append(line)

    return snafu


snafu = read_input('input2')
suma = 0
for num in snafu:
    suma += snafu_to_dec(num)
print(suma)
print(dec_to_snafu(suma))

#print(snafu_to_dec("1=="))
#print(dec_to_snafu(13))
#print(dec_to_snafu(11))
#print(dec_to_snafu(201))
#print(dec_to_snafu(906))
#print(dec_to_snafu(1747))
