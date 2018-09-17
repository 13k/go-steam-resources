/*
Package parser implements a parser for SteamKit's SteamLanguage.

Grammar

The grammar (in EBNF) was inferred from SteamKit's source code and is not official.

Common definitions such as letter, digit or string use undefined special sequences but are common
enough in programming languages to be easily understandable.

White space outside rules is discarded.

	language = { preprocess-directive }, { comment | enum-definition | class-definition } ;

	letter = ? unicode-letter ? ;
	digit = ? unicode-digit ? ;
	character = ? unicode-character ? ;
	string = ? double-quote-delimited-string ? ;
	integer = ? integer ? | ? hex-integer ? | ? octal-integer ? ;
	float = ? float ? ;
	number = integer | float ;
	identifier = letter, { "_" | letter | digit | ":" | "." } ;
	terminator = ";" ;

	preprocess-directive = "#", string ;

	comment = "//", { character } ;

	type-ref = identifier | number ;
	type-param = "<", type-ref, ">" ;

	property-type-param = type-param ;
	property-qualifier = identifier, [ property-type-param ] ;
	property-type = identifier ;
	property-name = identifier ;
	property-value = identifier ;
	property-values = property-value, { "|", property-value } ;
	property-assignment = "=", property-values ;
	property-annotation-value = "obsolete" | "removed" ;
	property-annotation-comment = string ;
	property-annotation = property-annotation-value, [ property-annotation-comment ] ;
	property-definition = [ property-qualifier ], [ property-type ], property-name,
		                    [ property-assignment ], terminator, [ property-annotation ] ;

	scope-definition = "{", { property-definition }, "}" ;

	class-emsg = type-param ;
	class-annotation = "removed" ;
	class-definition = "class", identifier, [ class-emsg ], [ class-annotation ],
	                   scope-definition, terminator ;

	enum-type = type-param ;
	enum-annotation = "flags" ;
	enum-definition = "enum", identifier, [ enum-type ], [ enum-annotation ],
	                  scope-definition, terminator ;
*/
package parser
