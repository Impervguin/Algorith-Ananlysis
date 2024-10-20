#include <tidy.h>
#include <tidybuffio.h>
#include <stdbool.h>
#include <stdio.h>
#include "download.h"
#include "recipeparser.h"
#include "htmlnode.h"


bool check_recipe_article(TidyNode node) {
    if (node == NULL)
        return false;

    ctmbstr name = tidyNodeGetName(node);
    if (name && strcmp(name, "article") == 0) {
        TidyAttr attr;
        for(attr = tidyAttrFirst(node); attr; attr = tidyAttrNext(attr) ) {
            if (strcmp(tidyAttrName(attr), "class") == 0 && (strcmp(tidyAttrValue(attr), "recipe") == 0 || strcmp(tidyAttrValue(attr), "post-default") == 0)) {
                return true;
            }
        }
    }

    return false;
}

TidyNode get_recipe_article(TidyDoc doc, TidyNode current) {
    if (!current) return current;
    for (TidyNode child = tidyGetChild(current); child; child = tidyGetNext(child)) {
        if (check_recipe_article(child)) {
            return child;
        }
        TidyNode found = get_recipe_article(doc, child);
        if (found) return found;
    }

    return NULL;
}

int download_article(const char *url, const char *fname) {
    TidyDoc doc;
    TidyBuffer buf;
    int err;
    FILE *file = fopen(fname, "w");
    if (file == NULL) {
        return DUMP_FILE_ERROR;
    }
    // printf("File created: %s\n", url);

    err = download_page_to_docbuf(url, &buf);
    if (err!= DOWNLOAD_OK) {
        fclose(file);
        return err;
    }
    // printf("File downloaded: %s\n", url);

    err = parse_html_from_buffer(&doc, &buf);
    if (err!= PARSE_OK) {
        fclose(file);
        release_doc(&doc, &buf);
        return err;
    }
    // printf("HTML parsed: %s\n", url);

    TidyNode article = get_recipe_article(doc, tidyGetRoot(doc));
    if (article == NULL) {
        fclose(file);
        release_doc(&doc, &buf);
        return RECIPE_NOT_FOUND;
    }
    // printf("Recipe found: %s\n", url);

    err = dump_node(doc, article, 0, file);
    fclose(file);
    release_doc(&doc, &buf);
    return err;
}