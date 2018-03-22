// Copyright 2018 NetApp, Inc. All Rights Reserved.

package docker

import (
	"fmt"

	hash "github.com/mitchellh/hashstructure"
	log "github.com/sirupsen/logrus"

	"github.com/netapp/trident/config"
	"github.com/netapp/trident/core"
	"github.com/netapp/trident/storage"
	"github.com/netapp/trident/storage_attribute"
	"github.com/netapp/trident/storage_class"
	"github.com/netapp/trident/utils"
)

// getStorageClass accepts a list of volume creation options and returns a
// matching storage class.  If the orchestrator already has a matching
// storage class, that is returned; otherwise a new one is created and
// registered with the orchestrator.
func getStorageClass(options map[string]string, o core.Orchestrator) (*storageclass.Config, error) {

	// Create a storage class based on available options
	newScConfig, err := makeStorageClass(options)
	if err != nil {
		return nil, err
	}

	// Check existing storage classes for a match based on the name
	sc, err := o.GetStorageClass(newScConfig.Name)
	if err != nil && !core.IsNotFoundError(err) {
		return nil, err
	}
	if sc != nil {
		log.WithField("storageClass", sc.Config.Name).Debug("Matched existing storage class.")
		return sc.Config, nil
	}

	// No match found, so register the new storage class
	addedSc, err := o.AddStorageClass(newScConfig)
	if err != nil {
		log.WithFields(log.Fields{
			"storageClass": newScConfig.Name,
		}).Error("Docker frontend couldn't add the storage class: ", err)
		return nil, err
	}

	return addedSc.Config, nil
}

// makeStorageClass accepts a list of volume creation options and creates a
// matching storage class.  The name of the new storage class contains a hash
// of the attributes it contains, thereby enabling comparison of storage
// classes generated by this method by simply comparing their names.
func makeStorageClass(options map[string]string) (*storageclass.Config, error) {

	scConfig := new(storageclass.Config)

	// Map options to storage class attributes
	scConfig.Attributes = make(map[string]storageattribute.Request)
	for k, v := range options {
		// format: attribute: "type:value"
		req, err := storageattribute.CreateAttributeRequestFromAttributeValue(k, v)
		if err != nil {
			log.WithFields(log.Fields{
				"storageClass":            scConfig.Name,
				"storageClass_parameters": options,
			}).Debug("Docker frontend ignoring storage class attribute: ", err)
			continue
		}
		scConfig.Attributes[k] = req
	}

	// Set name based on hash value
	scHash, err := hash.Hash(scConfig, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"storageClass":            scConfig.Name,
			"storageClass_parameters": options,
		}).Error("Docker frontend couldn't hash the storage class attributes: ", err)
		return nil, err
	}
	scConfig.Name = fmt.Sprintf(autoStorageClassPrefix, scHash)

	return scConfig, nil
}

// getVolumeConfig accepts a set of parameters describing a volume creation request
// and returns a volume config structure suitable for passing to the orchestrator core.
func getVolumeConfig(name, storageClass string, opts map[string]string) (*storage.VolumeConfig, error) {

	sizeBytes, err := utils.GetVolumeSizeBytes(opts, "0")
	if err != nil {
		return nil, fmt.Errorf("error creating volume: %v", err)
	}
	delete(opts, "size")

	return &storage.VolumeConfig{
		Name:                name,
		Size:                fmt.Sprintf("%d", sizeBytes),
		StorageClass:        storageClass,
		Protocol:            config.ProtocolAny,
		AccessMode:          config.ModeAny,
		SpaceReserve:        utils.GetV(opts, "spaceReserve", ""),
		SecurityStyle:       utils.GetV(opts, "securityStyle", ""),
		SplitOnClone:        utils.GetV(opts, "splitOnClone", ""),
		SnapshotPolicy:      utils.GetV(opts, "snapshotPolicy", ""),
		ExportPolicy:        utils.GetV(opts, "exportPolicy", ""),
		SnapshotDir:         utils.GetV(opts, "snapshotDir", ""),
		UnixPermissions:     utils.GetV(opts, "unixPermissions", ""),
		BlockSize:           utils.GetV(opts, "blocksize", ""),
		QoS:                 utils.GetV(opts, "qos", ""),
		QoSType:             utils.GetV(opts, "type", ""),
		FileSystem:          utils.GetV(opts, "fstype|fileSystemType", ""),
		Encryption:          utils.GetV(opts, "encryption", ""),
		CloneSourceVolume:   utils.GetV(opts, "from", ""),
		CloneSourceSnapshot: utils.GetV(opts, "fromSnapshot", ""),
	}, nil
}
