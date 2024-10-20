#include <stdlib.h>
#include <stdio.h>
#include "urlstorage.h"

#define INIT_LEN 10
#define CAP_MULTIPLIER 2

int init_url_storage(url_storage_t *s) {
    if (s == NULL) return 0;
    
    char **urls = malloc(INIT_LEN * sizeof(char *));
    if (urls == NULL) return MEMORY_ERROR;

    s->cap = INIT_LEN;
    s->len = 0;
    s->urls = urls;
    return OK;
}

void release_url_storage(url_storage_t *s) {
    if (s) {
        for (int i = 0; i < s->len; i++) {
            if (s->urls[i] != NULL) {
                free(s->urls[i]);
                s->urls[i] = NULL;
            }
        }
        if (s->urls) {
            free(s->urls);
            s->urls = NULL;
        }
        s->cap = 0;
        s->len = 0;
    }
}

int add_url(url_storage_t *s, const char *url) {
    if (s == NULL || url == NULL) return NULL_POINTER_ERROR;

    if (s->len >= s->cap) {
        char **new_urls = realloc(s->urls, s->cap * CAP_MULTIPLIER * sizeof(char *));
        if (new_urls == NULL) return MEMORY_ERROR;

        s->urls = new_urls;
        s->cap *= CAP_MULTIPLIER;
    }

    s->urls[s->len] = strdup(url);
    if (s->urls[s->len] == NULL) return MEMORY_ERROR;

    s->len++;
    return OK;
}


int read_url_storage(url_storage_t *s, const char *fname) {
    url_storage_t tmp;
    int err = init_url_storage(&tmp);
    if (err) return err;

    FILE *f = fopen(fname, "r");
    if (f == NULL) return FILE_ERROR;

    char *line = NULL;
    size_t slen = 0;
    ssize_t len;

    while ((len = getline(&line, &len, f)) != -1) {
        line[len - 1] = '\0';  // remove trailing newline
        err = add_url(&tmp, line);
        if (err) {
            release_url_storage(&tmp);
            return err;
        }
    }
    fclose(f);
    if (line) free(line);
    release_url_storage(s);
    *s = tmp;
    return OK;
}