#include <stdio.h>

void fold(char *doc,size_t doc_size, int width);
int append_char(char *base,int at, char c);
int insert_char(char *base,int len,int at, char c);


int main() {
    char doc[]="This is a folding experimentation of some kind.";
    printf("Original Test: [%s]\n", doc);
    fold(doc, sizeof(doc)/sizeof(doc[0]), 6);
    return 0;
}

void fold(char *doc,size_t doc_size, int width) {
    char c;
    int current_pos=1;
    int last_whitespace_pos=0;
    int folded_length=0;
    size_t max=doc_size+(doc_size/width)+1;
    char folded[max];

    for (int i=0;i<doc_size;i++){
        char c=doc[i];
        folded_length=append_char(folded,folded_length,c);

        if (c==' ' || c=='\t') {
            last_whitespace_pos=current_pos;
        }

        if (c=='\n' ) {
            last_whitespace_pos=0;
            current_pos=1;
            continue;
        }
        if (current_pos>width && last_whitespace_pos==current_pos) {
            folded_length=insert_char(folded,folded_length,folded_length-1,'\n');
            last_whitespace_pos=1;
            current_pos=2;
            continue;
        }

        if (current_pos>width && last_whitespace_pos>0 ) {
            folded_length=insert_char(folded,folded_length,folded_length-(width-(last_whitespace_pos-1)),'\n');
            current_pos=1+(width-(last_whitespace_pos-1));
            last_whitespace_pos=0;
            continue;
        }
        current_pos++;
    }
    folded[folded_length]='\0';
    printf("Folded text:\n[%s]\n", folded);
}

int append_char(char *doc,int at, char c) {
    doc[at]=c;
    return at+1;
}

int insert_char(char *doc,int len,int at, char c) {
    for(int i=len;i>at;i--) {
        doc[i]=doc[i-1];
    }
    doc[at]=c;
    return ++len;
}
