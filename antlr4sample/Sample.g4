grammar Sample;

/* 
install https://marketplace.visualstudio.com/items?itemName=mike-lischke.vscode-antlr4 

then use this VSCode settings:

"antlr4.generation": {
    "mode": "external",
    "language": "Go",
    "listeners": true,
    "visitors": false,
    "outputDir": "parser"
},

*/

SAMPLE: 'Sample';
INT: [1-9] DIG*;
fragment DIG: [0-9];
DOT: '.';
WS: [ \t\r\n\f] -> channel(HIDDEN);

main: 
    SAMPLE
    samplenum=INT 
    DOT 
    EOF;
