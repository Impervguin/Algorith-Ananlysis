

def SimpleMatrixMultiply(self, m1: list[list[int]], m2: list[list[int]]):
    if len(m1[0]) != len(m2):
        raise ValueError("Matrices cannot be multiplied")

    result = [[0] * len(m2[0]) for _ in range(len(m1))]

    for i in range(len(m1)):
        for j in range(len(m2[0])):
            for k in range(len(m2)):
                result[i][j] += m1[i][k] * m2[k][j]

    return result

def VinogradMatrixMultiply(self, m1: list[list[int]], m2: list[list[int]]):
    if len(m1[0]) != len(m2):
        raise ValueError("Matrices cannot be multiplied")

    result = [[0] * len(m2[0]) for _ in range(len(m1))]
    mulH = [[]]

    return result