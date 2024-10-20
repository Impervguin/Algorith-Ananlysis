#include <string.h>
#include <stdio.h>

char *get_url_basename(char *url) {
    char *start = strrchr(url, '/');
    if (start == NULL) return strdup(url);
    if (start[1] == '\0') {
        // printf("start1");
        start[0] = '\0';
        char *start2 = strrchr(url, '/');
        if (start2 == NULL) {
            start[0] = '\0';
            return strdup(url);
        }
        char *res = strdup(start2 + 1);
        start[0] = '/';
        return res;
    }
    return strdup(start + 1);
}