package phonenumber

import "errors"

func Number(phoneNumber string) (string, error) {
    num := ""

    // keep only digits
    for i := 0; i < len(phoneNumber); i++ {
        ch := phoneNumber[i]
        if ch >= '0' && ch <= '9' {
            num += string(ch)
        }
    }

    // handle country code
    if len(num) == 11 && num[0] == '1' {
        num = num[1:]
    }

    // length check
    if len(num) != 10 {
        return "", errors.New("invalid")
    }

    // basic validation
    if num[0] < '2' || num[3] < '2' {
        return "", errors.New("invalid")
    }

    return num, nil
}

func AreaCode(phoneNumber string) (string, error) {
    num, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return num[0:3], nil
}

func Format(phoneNumber string) (string, error) {
    num, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }

    return "(" + num[0:3] + ") " + num[3:6] + "-" + num[6:], nil
}