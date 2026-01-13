def addition(a, b):
    if isinstance(a, (int, float)) and isinstance(b, (int, float)):
        return a + b
    else:
        raise TypeError("Unsupported operand type(s) for +: 'int' and 'str'")
