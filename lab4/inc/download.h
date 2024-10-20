#ifndef DOWNLOAD_H__
#define DOWNLOAD_H__

#include <tidy.h>

#define DOWNLOAD_OK 0
#define DOWNLOAD_ERROR 1

typedef unsigned int uint;

int download_page_to_docbuf(const char * url, TidyBuffer *buf);

#endif