from pythonds.basic.stack import Stack


def rev_string(mystr):
    s = Stack()
    for char in mystr:
        s.push(char)

    reversed = ""
    while s.size() > 0:
        reversed += s.pop()

    return reversed


print rev_string("Hello")
