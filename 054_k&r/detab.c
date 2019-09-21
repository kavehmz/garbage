#include <stdio.h>

void process_input(int max, int tab_stop);
int append_space(char *base, int at,int tab_size);
int append_char(char *base,int at, char c);
short append_fits(int pos,int max, int tab_top, char c);


int main() {
    process_input(7,4);
    return 0;
}

void process_input(int max, int tab_stop) {
    char c;
    char line[max];
    int pos=0;

    while ((c=getchar())!=EOF && append_fits(pos,max,tab_stop,c) ) {
        if (c=='\t') {
            pos=append_space(line,pos, tab_stop);
        } else {
            pos=append_char(line,pos,c);
        }
    }
    append_char(line,pos,'\0');
    printf("[%s]", line);
}

int append_space(char *line, int at, int tab_size) {
    for (int i=0;i<tab_size;i++){
        at=append_char(line,at,' ');
    }
    return at;
}

int append_char(char *line,int at, char c) {
    line[at]=c;
    return at+1;
}

short append_fits(int pos, int max, int tab_stop, char c) {
    max=max-1; // save room for \0
    int size=1;
    if (c=='\t') {
        size=tab_stop;
    }
    if (pos>=max || c=='\t' && pos+size>=max) {
        return 0;
    }
    return 1;
}

