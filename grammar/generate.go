package grammar

//go:generate java -Xmx500M -cp "./antlr-4.13.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool -Dlanguage=Go -visitor -package parsing -o ../parsing ./BoardGameLang.g4
