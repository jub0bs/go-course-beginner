package p

// START OMIT
func (c *Checker) getUnresolvedSymbolForEntityName(name *ast.Node) *ast.Symbol {
	// ...
	result := c.unresolvedSymbols[path] // HL
	if result == nil {
		result := c.newSymbol(ast.SymbolFlagsTypeAlias, text) // HL
		c.unresolvedSymbols[path] = result
		result.Parent = parentSymbol
		c.declaredTypeLinks.Get(result).declaredType = c.unresolvedType
	}
	return result // HL
}

// END OMIT
