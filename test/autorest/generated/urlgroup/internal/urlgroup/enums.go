// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package urlgroup

type UriColor string

const (
	UriColorRedColor UriColor = "red color"
	UriColorGreenColor UriColor = "green color"
	UriColorBlueColor UriColor = "blue color"
)

func PossibleUriColorValues() []UriColor {
	return []UriColor{UriColorRedColor, UriColorGreenColor, UriColorBlueColor}
}

func (c UriColor) ToPtr() *UriColor {
	return &c
}

