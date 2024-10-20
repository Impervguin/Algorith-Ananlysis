#include <tidy.h>
#include <tidybuffio.h>
#include <curl/curl.h>
#include "download.h"
// #include <types.h>

/* curl write callback, to fill tidy's input buffer...  */
uint write_cb(char *in, uint size, uint nmemb, TidyBuffer *out)
{
  uint r;
  r = size * nmemb;
  tidyBufAppend(out, in, r);
  return r;
}

int download_page_to_docbuf(const char * url, TidyBuffer *buf) {
    tidyBufInit(buf);
    CURL *curl;
    int err;

    curl = curl_easy_init();
    curl_easy_setopt(curl, CURLOPT_URL, url);
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_cb);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, buf);
    err = curl_easy_perform(curl);
    curl_easy_cleanup(curl);

    if (err != CURLE_OK) {
        tidyBufFree(buf);
        return DOWNLOAD_ERROR;
    }
        
    return DOWNLOAD_OK;
}