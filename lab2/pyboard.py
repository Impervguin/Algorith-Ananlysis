import urandom as r
import utime as t

def SimpleMatrixMultiply(m1: list[list[int]], m2: list[list[int]]):
    if len(m1[0]) != len(m2):
        raise ValueError("Matrices cannot be multiplied")

    result = [[0] * len(m2[0]) for _ in range(len(m1))]

    for i in range(len(m1)):
        for j in range(len(m2[0])):
            for k in range(len(m2)):
                result[i][j] += m1[i][k] * m2[k][j]

    return result

def VinogradMatrixMultiply(m1: list[list[int]], m2: list[list[int]]):
    if (len(m1) == 0 or len(m2) == 0):
        raise ValueError("Empty matrix")

    if len(m1[0]) != len(m2):
        raise ValueError("Matrices cannot be multiplied")

    M = len(m1)
    N = len(m1[0]) # == len(m2)
    Q = len(m2[0])
    result = [[0] * Q for _ in range(M)]

    mulH = [0] * (M)
    for i in range(M):
        for j in range(N // 2):
            mulH[i] = mulH[i] + m1[i][2*j] * m1[i][2*j + 1]

    mulV = [0] * (Q)
    for i in range(Q):
        for j in range(N // 2):
            mulV[i] = mulV[i] + m2[2*j][i] * m2[2*j + 1][i]

    for i in range(M):
        for j in range(Q):
            result[i][j] = -mulH[i] -mulV[j]
            for k in range(N // 2):
                result[i][j] = result[i][j] + (m1[i][2*k]+m2[2*k + 1][j]) * (m1[i][2*k + 1]+m2[2*k][j])

    if (N % 2 != 0):
        for i in range(M):
            for j in range(Q):
                result[i][j] = result[i][j] + m1[i][-1] * m2[-1][j]

    return result

def OptimizedVinogradMatrixMultiply(m1: list[list[int]], m2: list[list[int]]):
    if (len(m1) == 0 or len(m2) == 0):
        raise ValueError("Empty matrix")

    if len(m1[0]) != len(m2):
        raise ValueError("Matrices cannot be multiplied")

    M = len(m1)
    N = len(m1[0]) # == len(m2)
    Q = len(m2[0])
    result = [[0] * Q for _ in range(M)]

    mulH = [0] * (M)
    for i in range(M):
        mulH[i] = m1[i][0] * m1[i][1]
        for j in range(2, N - 1, 2):
            mulH[i] += m1[i][j] * m1[i][j + 1]

    mulV = [0] * (Q)
    for i in range(Q):
        mulV[i] = m2[0][i] * m2[1][i]
        for j in range(2, N - 1, 2):
            mulV[i] += m2[j][i] * m2[j + 1][i]

    for i in range(M):
        for j in range(Q):
            result[i][j] = -mulH[i] -mulV[j] + (m1[i][0]+m2[1][j]) * (m1[i][1]+m2[0][j])
            for k in range(2, N - 1, 2):
                result[i][j] += (m1[i][k]+m2[k + 1][j]) * (m1[i][k + 1]+m2[k][j])

    if (N % 2 != 0):
        for i in range(M):
            for j in range(Q):
                result[i][j]  += m1[i][-1] * m2[-1][j]

    return result



RANDMAX = 10000
RANDMIN = -10000
def RandomMatrix(n, m):
    return [[r.randint(-10000, 10000)] * m for _ in range(n)]

# REPEATCOUNT = 20
# REPEATCOUNT = 10
REPEATCOUNT = 3
# Памяти на все тесты за раз на плате не хватает, выдаёт ошибку
# Поэтому разобьём тесты на несколь запусков
# SIZES = [5, 10, 15, 20, 25]
# SIZES = [30, 35]
# SIZES = [40]
# SIZES = [45]
SIZES = [50]


print("size\tstd\tvin\topt\trepeat")
for size in SIZES:
    tstd = 0
    tvin = 0
    topt = 0
    for _ in range(REPEATCOUNT):
        mat1 = RandomMatrix(size, size)
        mat2 = RandomMatrix(size, size)

        start = t.ticks_ms()
        res = SimpleMatrixMultiply(mat1, mat2)
        end = t.ticks_ms()
        tstd += t.ticks_diff(end, start)

        start = t.ticks_ms()
        res = VinogradMatrixMultiply(mat1, mat2)
        end = t.ticks_ms()
        tvin += t.ticks_diff(end, start)

        start = t.ticks_ms()
        res = OptimizedVinogradMatrixMultiply(mat1, mat2)
        end = t.ticks_ms()
        topt += t.ticks_diff(end, start)
    tstd /= REPEATCOUNT
    # tstd //= 1_000_000
    tvin /= REPEATCOUNT
    # tvin //= 1_000_000
    topt /= REPEATCOUNT
    # topt //= 1_000_000
    print(f"{size}\t{tstd:.3g}\t{tvin:.3g}\t{topt:.3g}\t{REPEATCOUNT}")
