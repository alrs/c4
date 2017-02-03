// Package c4 implements the C4 framework, consisting of C4 ID, C4lang, and C4 PKI.
//
// Creating IDs
//
// To create a C4 ID, use the IDEncoder type as a hash writer.
//
//     e := c4.NewIDEncoder()
//     io.Copy(e, src)
//     id := e.ID()
//
// Parsing IDs
//
// To parse an ID from a string, use the ParseID function.
//
//     c4.ParseID(src)
package c4
