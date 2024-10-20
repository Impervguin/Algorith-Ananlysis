#ifndef HTML_NODE_H__
#define HTML_NODE_H__

#include <stdio.h>
#include <tidy.h>

#define PARSE_OK 0
#define DOC_INIT_ERROR 1
#define DOC_PARSE_ERROR 2

# define DUMP_OK 0
# define DUMP_FILE_ERROR 1
# define DUMP_PRINT_ERROR 2

int dump_node(TidyDoc doc, TidyNode node, int indent, FILE *f);
int parse_html_from_buffer(TidyDoc *doc, TidyBuffer *buf);
void release_doc(TidyDoc *doc, TidyBuffer *buf);

#endif