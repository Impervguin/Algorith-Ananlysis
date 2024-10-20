#ifndef RECIPE_PARSER_H__
#define RECIPE_PARSER_H__

#include <tidy.h>

int download_article(const char *url, const char *fname);
#define RECIPE_SUCCESS 0
#define RECIPE_NOT_FOUND 1

#endif