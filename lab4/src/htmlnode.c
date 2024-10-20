#include <stdio.h>
#include <tidy.h>
#include <tidybuffio.h>
#include "htmlnode.h"

void fill_indent_string(char *str, int indent) {
    for (int i = 0; i < indent; i++) {
        str[i] = ' ';
    }
    str[indent] = '\0';
}

int dump_node(TidyDoc doc, TidyNode node, int indent, FILE *f) {
    if (f == NULL)
        return DUMP_FILE_ERROR;
    if (node == NULL) return DUMP_OK;
    char space_buffer[indent + 1];
    
    fill_indent_string(space_buffer, indent);
    int err = DUMP_OK;
    for( TidyNode child = tidyGetChild(node); child; child = tidyGetNext(child) ) {
        ctmbstr name = tidyNodeGetName(child);
        if (name) {
            TidyAttr attr;
            fprintf(f, "%s%s%s ", space_buffer, "<", name);
            /* walk the attribute list */
            for(attr = tidyAttrFirst(child); attr; attr = tidyAttrNext(attr) ) {
                fprintf(f, "%s", tidyAttrName(attr));
                tidyAttrValue(attr) ? fprintf(f, "=\"%s\" ", tidyAttrValue(attr)) : fprintf(f, " ");
            }
            fprintf(f, ">\n");
            err = dump_node(doc, child, indent + 4, f);
            if (err!= DUMP_OK)
                return err;
            fprintf(f, "%s%s%s%s\n", space_buffer, "</", name, ">");
        } else {
            TidyBuffer buf;
            tidyBufInit(&buf);
            tidyNodeGetText(doc, child, &buf);
            fprintf(f, "%s%s\n", space_buffer, buf.bp ? (char *)buf.bp : "");
            tidyBufFree(&buf);
        }
    }
    return DUMP_OK;
}

int parse_html_from_buffer(TidyDoc *doc, TidyBuffer *buf) {
    TidyDoc tmp_doc = tidyCreate();
    TidyBuffer err_buf;
    tidyBufInit(&err_buf);

    int err = 0;

    err = tidySetErrorBuffer(tmp_doc, &err_buf);

    if (err != 0) 
        return DOC_INIT_ERROR;
    
    tidyParseBuffer(tmp_doc, buf);
    tidyBufFree(&err_buf);
    
    *doc = tmp_doc;
    return PARSE_OK;
}

void release_doc(TidyDoc *doc, TidyBuffer *buf) {
    tidyBufFree(buf);
    tidyRelease(*doc);
}