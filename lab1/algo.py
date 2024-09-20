def RecursiveLevenshtein(s1, s2):
    if len(s1) == 0:
        return len(s2)
    if len(s2) == 0:
        return len(s1)
    if s1[0] == s2[0]:
        return RecursiveLevenshtein(s1[1:], s2[1:])
    return 1 + min(
        RecursiveLevenshtein(s1[1:], s2),
        RecursiveLevenshtein(s1[1:], s2[1:]),
        RecursiveLevenshtein(s1, s2[1:])
    )

def CacheLevenshtein(s1, s2):
    cacheRows = len(s1) + 1
    cacheCols = len(s2) + 1
    cache = [[0] * cacheCols for _ in range(cacheRows)]
    for i in range(1, cacheCols):
        cache[0][i] = i
    for i in range(1, cacheRows):
        cache[i][0] = i

    for i in range(1, cacheRows):
        for j in range(1, cacheCols):
                cache[i][j] = min(cache[i - 1][j] + 1, cache[i - 1][j - 1]  + (0 if s1[i - 1] == s2[j - 1] else 1), cache[i][j - 1] + 1)

    return cache[cacheRows - 1][cacheCols - 1]

def CacheDamerauLevenshtein(s1, s2):
    cacheRows = len(s1) + 1
    cacheCols = len(s2) + 1
    cache = [[0] * cacheCols for _ in range(cacheRows)]
    for i in range(1, cacheCols):
        cache[0][i] = i
    for i in range(1, cacheRows):
        cache[i][0] = i
    for i in range(1, cacheRows):
        for j in range(1, cacheCols):
                if i >= 2 and j >= 2 and s1[i - 1] == s2[j - 2] and s1[i - 2] == s2[j - 1]:
                    cache[i][j] = min(
                        cache[i - 1][j] + 1,
                        cache[i - 1][j - 1] + (0 if s1[i - 1] == s2[j - 1] else 1),
                        cache[i - 2][j - 2] + 1,
                        cache[i][j - 1] + 1,
                    )
                else:
                    cache[i][j] = min(
                        cache[i - 1][j] + 1,
                        cache[i - 1][j - 1] + (0 if s1[i - 1] == s2[j - 1] else 1),
                        cache[i][j - 1] + 1,
                    )

    return cache[cacheRows - 1][cacheCols - 1]








