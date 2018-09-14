/*
Package parser implements a parser for SteamKit's SteamLanguage.

Grammar (EBNF)

The grammar was inferred from SteamKit's source code and is not official.

	letter = ? unicode-letter ? ;
	digit = ? unicode-digit ? ;
	character = ? unicode-character ? ;
	string = ? quote-delimited-string ? ;
	integer = ? integer ? | ? hex-integer ? | ? octal-integer ? ;
	float = ? float ? ;
	number = integer | float ;
	identifier = letter, { "_" | letter | digit | ":" | "." } ;

	type-ref = identifier | number ;
	type-param = "<", type-ref, ">" ;

	class-comment = "removed" ;
	class = "class", identifier, [ type-param ], [ class-comment ], "{", { property }, "};" ;

	enum-comment = "flags" ;
	enum = "enum", identifier, [ type-param ], [ enum-comment ], "{", { property }, "};" ;

	property-qualifier = identifier, [ type-param ] ;
	property-type = identifier ;
	property-name = identifier ;
	property-definition = [ property-qualifier ], [ property-type ], property-name ;
	property-value = identifier ;
	property-assignment = "=", property-value, { "|", property-value } ;
	property-comment = "obsolete" | "removed", [ string ] ;
	property = property-definition, [ property-assignment ], ";" , [ property-comment ];
*/
package parser
