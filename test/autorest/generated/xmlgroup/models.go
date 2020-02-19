// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package xmlgroup

import azinternal "generatortests/autorest/generated/xmlgroup/internal/xmlgroup"

type AccessTier = azinternal.AccessTier

const (
	AccessTierArchive = azinternal.AccessTierArchive
	AccessTierCool    = azinternal.AccessTierCool
	AccessTierHot     = azinternal.AccessTierHot
	AccessTierP10     = azinternal.AccessTierP10
	AccessTierP20     = azinternal.AccessTierP20
	AccessTierP30     = azinternal.AccessTierP30
	AccessTierP4      = azinternal.AccessTierP4
	AccessTierP40     = azinternal.AccessTierP40
	AccessTierP50     = azinternal.AccessTierP50
	AccessTierP6      = azinternal.AccessTierP6
)

func PossibleAccessTierValues() []AccessTier {
	return azinternal.PossibleAccessTierValues()
}

type ArchiveStatus = azinternal.ArchiveStatus

const (
	ArchiveStatusRehydratePendingToCool = azinternal.ArchiveStatusRehydratePendingToCool
	ArchiveStatusRehydratePendingToHot  = azinternal.ArchiveStatusRehydratePendingToHot
)

func PossibleArchiveStatusValues() []ArchiveStatus {
	return azinternal.PossibleArchiveStatusValues()
}

type BlobType = azinternal.BlobType

const (
	BlobTypeBlockblob  = azinternal.BlobTypeBlockblob
	BlobTypePageblob   = azinternal.BlobTypePageblob
	BlobTypeAppendblob = azinternal.BlobTypeAppendblob
)

func PossibleBlobTypeValues() []BlobType {
	return azinternal.PossibleBlobTypeValues()
}

type CopyStatusType = azinternal.CopyStatusType

const (
	CopyStatusTypePending = azinternal.CopyStatusTypePending
	CopyStatusTypeSuccess = azinternal.CopyStatusTypeSuccess
	CopyStatusTypeAborted = azinternal.CopyStatusTypeAborted
	CopyStatusTypeFailed  = azinternal.CopyStatusTypeFailed
)

func PossibleCopyStatusTypeValues() []CopyStatusType {
	return azinternal.PossibleCopyStatusTypeValues()
}

type LeaseDurationType = azinternal.LeaseDurationType

const (
	LeaseDurationTypeInfinite = azinternal.LeaseDurationTypeInfinite
	LeaseDurationTypeFixed    = azinternal.LeaseDurationTypeFixed
)

func PossibleLeaseDurationTypeValues() []LeaseDurationType {
	return azinternal.PossibleLeaseDurationTypeValues()
}

type LeaseStateType = azinternal.LeaseStateType

const (
	LeaseStateTypeAvailable = azinternal.LeaseStateTypeAvailable
	LeaseStateTypeLeased    = azinternal.LeaseStateTypeLeased
	LeaseStateTypeExpired   = azinternal.LeaseStateTypeExpired
	LeaseStateTypeBreaking  = azinternal.LeaseStateTypeBreaking
	LeaseStateTypeBroken    = azinternal.LeaseStateTypeBroken
)

func PossibleLeaseStateTypeValues() []LeaseStateType {
	return azinternal.PossibleLeaseStateTypeValues()
}

type LeaseStatusType = azinternal.LeaseStatusType

const (
	LeaseStatusTypeLocked   = azinternal.LeaseStatusTypeLocked
	LeaseStatusTypeUnlocked = azinternal.LeaseStatusTypeUnlocked
)

func PossibleLeaseStatusTypeValues() []LeaseStatusType {
	return azinternal.PossibleLeaseStatusTypeValues()
}

type PublicAccessType = azinternal.PublicAccessType

const (
	PublicAccessTypeBlob      = azinternal.PublicAccessTypeBlob
	PublicAccessTypeContainer = azinternal.PublicAccessTypeContainer
)

func PossiblePublicAccessTypeValues() []PublicAccessType {
	return azinternal.PossiblePublicAccessTypeValues()
}

// An Access policy
type AccessPolicy = azinternal.AccessPolicy

// A barrel of apples.
type AppleBarrel = azinternal.AppleBarrel

// A banana.
type Banana = azinternal.Banana

// An Azure Storage blob
type Blob = azinternal.Blob

type BlobPrefix = azinternal.BlobPrefix

// Properties of a blob
type BlobProperties = azinternal.BlobProperties

type Blobs = azinternal.Blobs

// I am a complex type with no XML node
type ComplexTypeNoMeta = azinternal.ComplexTypeNoMeta

// I am a complex type with XML node
type ComplexTypeWithMeta = azinternal.ComplexTypeWithMeta

// An Azure Storage container
type Container = azinternal.Container

// Properties of a container
type ContainerProperties = azinternal.ContainerProperties

// CORS is an HTTP feature that enables a web application running under one domain to access resources in another domain.
// Web browsers implement a security restriction known as same-origin policy that prevents a web page from calling APIs in
// a different domain; CORS provides a secure way to allow one domain (the origin domain) to call APIs in another domain
type CorsRule = azinternal.CorsRule

type Error = azinternal.Error

type JSONInput = azinternal.JSONInput

type JSONOutput = azinternal.JSONOutput

// An enumeration of blobs
type ListBlobsResponse = azinternal.ListBlobsResponse

// An enumeration of containers
type ListContainersResponse = azinternal.ListContainersResponse

// Azure Analytics Logging settings.
type Logging = azinternal.Logging

type Metrics = azinternal.Metrics

// the retention policy
type RetentionPolicy = azinternal.RetentionPolicy

// I am root, and I ref a model WITH meta
type RootWithRefAndMeta = azinternal.RootWithRefAndMeta

// I am root, and I ref a model with no meta
type RootWithRefAndNoMeta = azinternal.RootWithRefAndNoMeta

// signed identifier
type SignedIDentifier = azinternal.SignedIDentifier

// A slide in a slideshow
type Slide = azinternal.Slide

// Data about a slideshow
type Slideshow = azinternal.Slideshow

// Storage Service Properties.
type StorageServiceProperties = azinternal.StorageServiceProperties

// XMLGetACLsResponse contains the response from method XML.GetACLs.
type XMLGetACLsResponse = azinternal.XMLGetACLsResponse

// XMLGetComplexTypeRefNoMetaResponse contains the response from method XML.GetComplexTypeRefNoMeta.
type XMLGetComplexTypeRefNoMetaResponse = azinternal.XMLGetComplexTypeRefNoMetaResponse

// XMLGetComplexTypeRefWithMetaResponse contains the response from method XML.GetComplexTypeRefWithMeta.
type XMLGetComplexTypeRefWithMetaResponse = azinternal.XMLGetComplexTypeRefWithMetaResponse

// XMLGetEmptyChildElementResponse contains the response from method XML.GetEmptyChildElement.
type XMLGetEmptyChildElementResponse = azinternal.XMLGetEmptyChildElementResponse

// XMLGetEmptyListResponse contains the response from method XML.GetEmptyList.
type XMLGetEmptyListResponse = azinternal.XMLGetEmptyListResponse

// XMLGetEmptyRootListResponse contains the response from method XML.GetEmptyRootList.
type XMLGetEmptyRootListResponse = azinternal.XMLGetEmptyRootListResponse

// XMLGetEmptyWrappedListsResponse contains the response from method XML.GetEmptyWrappedLists.
type XMLGetEmptyWrappedListsResponse = azinternal.XMLGetEmptyWrappedListsResponse

// XMLGetHeadersResponse contains the response from method XML.GetHeaders.
type XMLGetHeadersResponse = azinternal.XMLGetHeadersResponse

// XMLGetRootListResponse contains the response from method XML.GetRootList.
type XMLGetRootListResponse = azinternal.XMLGetRootListResponse

// XMLGetRootListSingleItemResponse contains the response from method XML.GetRootListSingleItem.
type XMLGetRootListSingleItemResponse = azinternal.XMLGetRootListSingleItemResponse

// XMLGetServicePropertiesResponse contains the response from method XML.GetServiceProperties.
type XMLGetServicePropertiesResponse = azinternal.XMLGetServicePropertiesResponse

// XMLGetSimpleResponse contains the response from method XML.GetSimple.
type XMLGetSimpleResponse = azinternal.XMLGetSimpleResponse

// XMLGetWrappedListsResponse contains the response from method XML.GetWrappedLists.
type XMLGetWrappedListsResponse = azinternal.XMLGetWrappedListsResponse

// XMLJSONInputResponse contains the response from method XML.JSONInput.
type XMLJSONInputResponse = azinternal.XMLJSONInputResponse

// XMLJSONOutputResponse contains the response from method XML.JSONOutput.
type XMLJSONOutputResponse = azinternal.XMLJSONOutputResponse

// XMLListBlobsResponse contains the response from method XML.ListBlobs.
type XMLListBlobsResponse = azinternal.XMLListBlobsResponse

// XMLListContainersResponse contains the response from method XML.ListContainers.
type XMLListContainersResponse = azinternal.XMLListContainersResponse

// XMLPutACLsResponse contains the response from method XML.PutACLs.
type XMLPutACLsResponse = azinternal.XMLPutACLsResponse

// XMLPutComplexTypeRefNoMetaResponse contains the response from method XML.PutComplexTypeRefNoMeta.
type XMLPutComplexTypeRefNoMetaResponse = azinternal.XMLPutComplexTypeRefNoMetaResponse

// XMLPutComplexTypeRefWithMetaResponse contains the response from method XML.PutComplexTypeRefWithMeta.
type XMLPutComplexTypeRefWithMetaResponse = azinternal.XMLPutComplexTypeRefWithMetaResponse

// XMLPutEmptyChildElementResponse contains the response from method XML.PutEmptyChildElement.
type XMLPutEmptyChildElementResponse = azinternal.XMLPutEmptyChildElementResponse

// XMLPutEmptyListResponse contains the response from method XML.PutEmptyList.
type XMLPutEmptyListResponse = azinternal.XMLPutEmptyListResponse

// XMLPutEmptyRootListResponse contains the response from method XML.PutEmptyRootList.
type XMLPutEmptyRootListResponse = azinternal.XMLPutEmptyRootListResponse

// XMLPutEmptyWrappedListsResponse contains the response from method XML.PutEmptyWrappedLists.
type XMLPutEmptyWrappedListsResponse = azinternal.XMLPutEmptyWrappedListsResponse

// XMLPutRootListResponse contains the response from method XML.PutRootList.
type XMLPutRootListResponse = azinternal.XMLPutRootListResponse

// XMLPutRootListSingleItemResponse contains the response from method XML.PutRootListSingleItem.
type XMLPutRootListSingleItemResponse = azinternal.XMLPutRootListSingleItemResponse

// XMLPutServicePropertiesResponse contains the response from method XML.PutServiceProperties.
type XMLPutServicePropertiesResponse = azinternal.XMLPutServicePropertiesResponse

// XMLPutSimpleResponse contains the response from method XML.PutSimple.
type XMLPutSimpleResponse = azinternal.XMLPutSimpleResponse

// XMLPutWrappedListsResponse contains the response from method XML.PutWrappedLists.
type XMLPutWrappedListsResponse = azinternal.XMLPutWrappedListsResponse
