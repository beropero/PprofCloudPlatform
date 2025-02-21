export interface GetProjectByPageUserReq {
    page:               number;
    pageSize:           number;
    projectDescription: string;
    projectId:          number;
    projectName:        string;
}

export interface CreateProjectReq {
    projectDescription: string;
    projectName: string;
}
