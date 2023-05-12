import {request} from "umi";

export interface TimeBaseType {
  utime?: number;
  ctime?: number;
  dtime?: number;
}
export interface InstanceType extends TimeBaseType {
  mode: number;
  id?: number;
  datasource: string;
  dsn: string;
  name: string;
  clusterId?: number;
  configmap?: string;
  namespace?: string;
  prometheusTarget?: string;
  clusters?: string[];
  desc?: string;
}

export interface TestInstanceRequest {
  dsn: string;
  datasource: string;
}

export interface ClustersRequest {
  current?: number;
  pageSize?: number;
  query?: string;
}

export interface ClusterType extends TimeBaseType {
  id?: number;
  clusterName: string;
  apiServer: string;
  description: string;
  kubeConfig: string;
  status: number;
}

export interface CreatedDatabaseRequest {
  databaseName: string;
}

export default {
  async getAllInstances() {
    return request(process.env.PUBLIC_PATH + `api/v2/base/instances`, {
      method: "GET",
    });
  },

  // Getting a list of instances
  async getInstances() {
    return request<API.Res<InstanceType[]>>(
      process.env.PUBLIC_PATH + `api/v1/sys/instances`,
      {
        method: "GET",
      }
    );
  },

  // Test instance
  async testInstance(data: TestInstanceRequest) {
    return request<API.Res<any>>(
      process.env.PUBLIC_PATH + `api/v1/sys/instances/test`,
      {
        method: "POST",
        data,
      }
    );
  },

  // Get instance Info
  async getInstancesInfo(id: number) {
    return request<API.Res<any>>(
      process.env.PUBLIC_PATH + `api/v1/sys/instances/${id}`
    );
  },

  // Create an instance
  async createdInstance(data: InstanceType) {
    return request<API.Res<string>>(
      process.env.PUBLIC_PATH + `api/v1/sys/instances`,
      {
        method: "POST",
        data,
      }
    );
  },
  // Update instance
  async updatedInstance(id: number, data: InstanceType) {
    return request<API.Res<string>>(
      process.env.PUBLIC_PATH + `api/v1/sys/instances/${id}`,
      {
        method: "PATCH",
        data,
      }
    );
  },
  // Deleting an instance
  async deletedInstance(id: number) {
    return request<API.Res<string>>(
      process.env.PUBLIC_PATH + `api/v1/sys/instances/${id}`,
      {
        method: "DELETE",
      }
    );
  },

  // Obtaining the cluster List
  async getClusters(params?: ClustersRequest) {
    return request<API.ResPage<ClusterType>>(
      process.env.PUBLIC_PATH + `api/v1/sys/clusters`,
      {
        method: "GET",
        params,
      }
    );
  },
  // Creating a Cluster
  async createdCluster(data: ClusterType) {
    return request<API.Res<string>>(
      process.env.PUBLIC_PATH + `api/v1/sys/clusters`,
      {
        method: "POST",
        data,
      }
    );
  },
  // Updating a cluster
  async updatedCluster(id: number, data: ClusterType) {
    return request<API.Res<string>>(
      process.env.PUBLIC_PATH + `api/v1/sys/clusters/${id}`,
      {
        method: "PATCH",
        data,
      }
    );
  },
  // Deleting a Cluster
  async deletedCluster(id: number) {
    return request<API.Res<string>>(
      process.env.PUBLIC_PATH + `api/v1/sys/clusters/${id}`,
      {
        method: "DELETE",
      }
    );
  },

  // Creating a database
  async createdDatabase(iid: number, data: CreatedDatabaseRequest) {
    return request(
      process.env.PUBLIC_PATH + `api/v1/instances/${iid}/databases`,
      {
        method: "POST",
        data,
      }
    );
  },

  // Delete database
  async deletedDatabase(id: number) {
    return request(process.env.PUBLIC_PATH + `api/v1/databases/${id}`, {
      method: "DELETE",
    });
  },

  // Updated database
  async updatedDatabase(id: number, data: any) {
    return request(process.env.PUBLIC_PATH + `api/v1/databases/${id}`, {
      method: "PATCH",
      data,
    });
  },
};
