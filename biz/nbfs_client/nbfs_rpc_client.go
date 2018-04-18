/*
 *  Copyright (c) 2018, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package nbfs_client

import (
	"github.com/nebulaim/telegramd/grpc_util/service_discovery"
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/mtproto"
	"context"
)

type nbfsClient struct {
	client mtproto.RPCNbfsClient
}

var (
	nbfsInstance = &nbfsClient{}
)

func InstallNbfsClient(discovery *service_discovery.ServiceDiscoveryClientConfig) {
	conn, err := grpc_util.NewRPCClientByServiceDiscovery(discovery)

	if err != nil {
		glog.Error(err)
		panic(err)
	}

	nbfsInstance.client = mtproto.NewRPCNbfsClient(conn)
}

func UploadPhotoFile(ownerId int64, file *mtproto.InputFile) (*mtproto.PhotoDataRsp, error) {
	// TODO(@benqi): Check nbfsInstance.client inited

	request := &mtproto.UploadPhotoFileRequest{
		OwnerId: ownerId,
		File:    file,
	}
	reply, err := nbfsInstance.client.NbfsUploadPhotoFile(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func GetPhotoSizeList(photoId int64) ([]*mtproto.PhotoSize, error) {
	// TODO(@benqi): Check nbfsInstance.client inited

	request := &mtproto.GetPhotoFileDataRequest{
		PhotoId: photoId,
	}
	reply, err := nbfsInstance.client.NbfsGetPhotoFileData(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return reply.SizeList, nil
}
