package main

// 0ms
// 2.5MB
func addStrings(num1 string, num2 string) string {
    i1 := len(num1) - 1
    i2 := len(num2) - 1
    if i1 < i2 {
        // ensure num1 longer
        num1, i1, num2, i2 = num2, i2, num1, i1
    }
    buf := make([]byte, i1+2)

    var carry byte = 0
    for i2 >= 0 {
        carry = num1[i1] - '0' + num2[i2] - '0' + carry
        buf[i1+1] = '0' + carry % 10
        carry /= 10
        i1--
        i2--
    }
    
    for i1 >= 0 && carry > 0 {
        carry = num1[i1] - '0' + carry
        buf[i1+1] = '0' + carry % 10
        carry /= 10
        i1--
    }

    if carry > 0 {
        buf[0] = '1'
        return string(buf)
    }

    if i1 >= 0 {
        copy(buf[1:], num1[:i1+1])
    }
    return string(buf[1:])
}