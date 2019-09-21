#include <stdio.h>

#define TAB_LENGTH 4
#define MAX_LENGTH 1000

int append_space(char *base, int at,int tab_size);
int append_char(char *base,int at, char c);

int main() {
    char c;
    char line[MAX_LENGTH];
    int pos=0;

    while ((c=getchar())!=EOF) {
        if (c!='\t') {
            pos=append_char(line,pos,c);
        } else {
            pos=append_space(line,pos, TAB_LENGTH);
        }
    }
    append_char(line,pos,'\0');
    printf("%s", line);
    return 0;
}

int append_space(char *line, int at,int tab_size) {
    for (int i=0;i<tab_size;i++){
        at=append_char(line,at,' ');
    }
    return at;
}

int append_char(char *line,int at, char c) {
    line[at]=c;
    return at+1;
}