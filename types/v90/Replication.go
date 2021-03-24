/*
 Copyright © 2020 Dell Inc. or its subsidiaries. All Rights Reserved.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
      http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package types

// RDFGroup contains information about an RDF group
type RDFGroup struct {
	RdfgNumber               int      `json:"rdfgNumber"`
	Label                    string   `json:"label"`
	RemoteRdfgNumber         int      `json:"remoteRdfgNumber"`
	RemoteSymmetrix          string   `json:"remoteSymmetrix"`
	NumDevices               int      `json:"numDevices"`
	TotalDeviceCapacity      float64  `json:"totalDeviceCapacity"`
	LocalPorts               []string `json:"localPorts"`
	RemotePorts              []string `json:"remotePorts"`
	Modes                    []string `json:"modes"`
	Type                     string   `json:"type"`
	Metro                    bool     `json:"metro"`
	Async                    bool     `json:"async"`
	Witness                  bool     `json:"witness"`
	WitnessName              string   `json:"witnessName"`
	WitnessProtectedPhysical bool     `json:"witnessProtectedPhysical"`
	WitnessProtectedVirtual  bool     `json:"witnessProtectedVirtual"`
	WitnessConfigured        bool     `json:"witnessConfigured"`
	WitnessEffective         bool     `json:"witnessEffective"`
	BiasConfigured           bool     `json:"biasConfigured"`
	BiasEffective            bool     `json:"biasEffective"`
	WitnessDegraded          bool     `json:"witnessDegraded"`
	LocalOnlinePorts         []string `json:"localOnlinePorts"`
	RemoteOnlinePorts        []string `json:"remoteOnlinePorts"`
	DevicePolarity           string   `json:"device_polarity"`
}

// Suspend action
type Suspend struct {
	Force      bool `json:"force"`
	SymForce   bool `json:"symForce"`
	Star       bool `json:"star"`
	Hop2       bool `json:"hop2"`
	Bypass     bool `json:"bypass"`
	Immediate  bool `json:"immediate"`
	ConsExempt bool `json:"consExempt"`
	MetroBias  bool `json:"metroBias"`
}

// Resume action
type Resume struct {
	Force        bool `json:"force"`
	SymForce     bool `json:"symForce"`
	Star         bool `json:"star"`
	Hop2         bool `json:"hop2"`
	Bypass       bool `json:"bypass"`
	Remote       bool `json:"remote"`
	RecoverPoint bool `json:"recoverPoint"`
}

// ModifySGRDFGroup holds parameters for rdf storage group updates
type ModifySGRDFGroup struct {
	Action          string   `json:"action"`
	Suspend         *Suspend `json:"suspend,omitempty"`
	Resume          *Resume  `json:"resume,omitempty"`
	ExecutionOption string   `json:"executionOption"`
}

// CreateSGSRDF contains parameters to create storage group replication {in u4p a.k.a "storageGroupSrdfCreate"}
type CreateSGSRDF struct {
	RemoteSymmID           string `json:"remoteSymmId"`
	ReplicationMode        string `json:"replicationMode"`
	RdfgNumber             int    `json:"rdfgNumber"`
	ForceNewRdfGroup       string `json:"forceNewRdfGroup"`
	Establish              bool   `json:"establish"`
	MetroBias              bool   `json:"metroBias"`
	RemoteStorageGroupName string `json:"remoteStorageGroupName"`
	ThinPool               string `json:"thinPool"`
	FastPolicy             string `json:"fastPolicy"`
	RemoteSLO              string `json:"remoteSLO"`
	NoCompression          bool   `json:"noCompression"`
	ExecutionOption        string `json:"executionOption"`
}

//SGRDFInfo contains parameters to hold srdf information of a storage group {in u4p a.k.a "storageGroupRDFg"}
type SGRDFInfo struct {
	SymmetrixID               string   `json:"symmetrixId"`
	StorageGroupName          string   `json:"storageGroupName"`
	RdfGroupNumber            int      `json:"rdfGroupNumber"`
	VolumeRdfTypes            []string `json:"volumeRdfTypes"`
	States                    []string `json:"states"`
	Modes                     []string `json:"modes"`
	Hop2Rdfgs                 []int    `json:"hop2Rdfgs"`
	Hop2States                []string `json:"hop2States"`
	Hop2Modes                 []string `json:"hop2Modes"`
	LargerRdfSides            []string `json:"largerRdfSides"`
	TotalTracks               int      `json:"totalTracks"`
	LocalR1InvalidTracksHop1  int      `json:"localR1InvalidTracksHop1"`
	LocalR2InvalidTracksHop1  int      `json:"localR2InvalidTracksHop1"`
	RemoteR1InvalidTracksHop1 int      `json:"remoteR1InvalidTracksHop1"`
	RemoteR2InvalidTracksHop1 int      `json:"remoteR2InvalidTracksHop1"`
	SrcR1InvalidTracksHop2    int      `json:"srcR1InvalidTracksHop2"`
	SrcR2InvalidTracksHop2    int      `json:"srcR2InvalidTracksHop2"`
	TgtR1InvalidTracksHop2    int      `json:"tgtR1InvalidTracksHop2"`
	TgtR2InvalidTracksHop2    int      `json:"tgtR2InvalidTracksHop2"`
}

//SGRDFGList contains list of all RDF enabled storage groups {in u4p a.k.a "storageGroupRDFg"}
type SGRDFGList struct {
	RDFGList []string `json:"rdfgs"`
}

//RDFStorageGroup contains information about protected SG {in u4p a.k.a "StorageGroup"}
type RDFStorageGroup struct {
	Name               string   `json:"name"`
	SymmetrixID        string   `json:"symmetrixId"`
	ParentName         string   `json:"parentName"`
	ChildNames         []string `json:"childNames"`
	NumDevicesNonGk    int      `json:"numDevicesNonGk"`
	CapacityGB         float64  `json:"capacityGB"`
	NumSnapVXSnapshots int      `json:"numSnapVXSnapshots"`
	SnapVXSnapshots    []string `json:"snapVXSnapshots"`
	Rdf                bool     `json:"rdf"`
	IsLinkTarget       bool     `json:"isLinkTarget"`
}

// LocalDeviceAutoCriteria holds parameters for auto selecting local device parameters
type LocalDeviceAutoCriteria struct {
	PairCount          int    `json:"pairCount"`
	Emulation          string `json:"emulation"`
	Capacity           int64  `json:"capacity"`
	CapacityUnit       string `json:"capacityUnit"`
	LocalThinPoolName  string `json:"localThinPoolName"`
	RemoteThinPoolName string `json:"remoteThinPoolName"`
}

// LocalDeviceListCriteria holds parameters for local device lis
type LocalDeviceListCriteria struct {
	LocalDeviceList    []string `json:"localDeviceList"`
	RemoteThinPoolName string   `json:"remoteThinPoolName"`
}

// CreateRDFPair holds SG create replica pair parameters
type CreateRDFPair struct {
	RdfMode                 string                   `json:"rdfMode"`
	RdfType                 string                   `json:"rdfType"`
	InvalidateR1            bool                     `json:"invalidateR1"`
	InvalidateR2            bool                     `json:"invalidateR2"`
	Establish               bool                     `json:"establish"`
	Restore                 bool                     `json:"restore"`
	Format                  bool                     `json:"format"`
	Exempt                  bool                     `json:"exempt"`
	NoWD                    bool                     `json:"noWD"`
	Remote                  bool                     `json:"remote"`
	Bias                    bool                     `json:"bias"`
	RecoverPoint            bool                     `json:"recoverPoint"`
	LocalDeviceAutoCriteria *LocalDeviceAutoCriteria `json:"localDeviceAutoCriteriaParam"`
	LocalDeviceListCriteria *LocalDeviceListCriteria `json:"localDeviceListCriteriaParam"`
	ExecutionOption         string                   `json:"executionOption"`
}

// RDFDevicePair holds RDF volume pair information
type RDFDevicePair struct {
	LocalSymmID          string `json:"localSymmetrixId"`
	RemoteSymmID         string `json:"remoteSymmetrixId"`
	LocalRdfGroupNumber  int    `json:"localRdfGroupNumber"`
	RemoteRdfGroupNumber int    `json:"remoteRdfGroupNumber"`
	LocalVolumeName      string `json:"localVolumeName"`
	RemoteVolumeName     string `json:"remoteVolumeName"`
	LocalVolumeState     string `json:"localVolumeState"`
	RemoteVolumeState    string `json:"remoteVolumeState"`
	VolumeConfig         string `json:"volumeConfig"`
	RdfMode              string `json:"rdfMode"`
	RdfpairState         string `json:"rdfpairState"`
	LargerRdfSide        string `json:"largerRdfSide"`
}

// RDFDevicePairList holds list of newly created RDF volume pair information
type RDFDevicePairList struct {
	RDFDevicePair []RDFDevicePair `json:"devicePair"`
}

// StorageGroupRDFG holds information about protected storage group
type StorageGroupRDFG struct {
	SymmetrixID      string   `json:"symmetrixId"`
	StorageGroupName string   `json:"storageGroupName"`
	RdfGroupNumber   int      `json:"rdfGroupNumber"`
	VolumeRdfTypes   []string `json:"volumeRdfTypes"`
	States           []string `json:"states"`
	Modes            []string `json:"modes"`
}
