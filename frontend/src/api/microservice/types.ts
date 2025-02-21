export interface GetMicroserviceByPageUserReq {
    Ip: string;
    MicroserviceDescription: string;
    MicroserviceId: number;
    MicroserviceName: string;
    page: number;
    pageSize: number;
    Port: number;
    ProjectId: number;
}

export interface CreateMicroserviceReq {
    Ip: string;
    MicroserviceDescription: string;
    MicroserviceName: string;
    Port: number;
    ProjectId: number;
}
