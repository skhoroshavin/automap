package provider

import (
	"github.com/skhoroshavin/automap/internal/mapper/node"
)

// Provider represents single provider
type Provider interface {
	Signature() *Signature

	Parent() Provider
	Children() []Provider

	Dependencies() []Request

	Map(node.Node, []node.Node) node.Node
}

// Nodes
// Node ID
// Node name
// Node Type
//  * Ident (name)
//  * Field (base.name)
//  * Struct (name{fields...})
//  * Method (base.name(args...))
//  * Func (name(args...))
//  * Addr (&base)
//  * Deref (*base)
// Node deps
//  * Base
//  * Args...

// Providers
//  * Ident - just a name of variable
//  * Field - access field
//  * Struct - build struct
//  * Method - get data by calling a method
//  * Func - transform data by calling a function
//  * Addr - get reference to underlying data
//  * Deref - dereference underlying pointer

// TypeConfig
//  * BlacklistMembers - forbid using some members
//  * WhitelistMembers - allow using only set of members
//  * IncludeMembers - allow member resolution without prefixing
//  * IgnoredPrefixes - list of prefixes to ignore
//  * Substitutions - list of substitutions
