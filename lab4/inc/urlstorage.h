#ifndef URL_STORAGE_H__
#define URL_STORAGE_H__

#include <tidy.h>

struct url_storage {
    char **urls;
    int cap;
    int len;
};

typedef struct url_storage url_storage_t;

int init_url_storage(url_storage_t *s);
int add_url(url_storage_t *s, const char *url);
void release_url_storage(url_storage_t *s);
int read_url_storage(url_storage_t *s, const char *fname);

#define OK 0
#define MEMORY_ERROR 1
#define FILE_ERROR 2
#define NULL_POINTER_ERROR 3

#endif