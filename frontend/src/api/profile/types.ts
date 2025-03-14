export interface GetProfileByPageUserReq {
    profileId: number;
    comment: string;
    microserviceId: number;
    page: number;
    pageSize: number;
    ptype: string;
}

export interface GetProfileContentReq {
    id?: string;
}
